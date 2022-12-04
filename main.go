package main

import (
	"brazuerabot/brasao"
	"brazuerabot/brazuera"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

type brazueraData struct {
	addressToConselho    map[string][]int
	conselhoToAddress    map[int]string
	addressToBrasaoCount map[string]int
	hasBothCount         int
	openseaData
	marketData
	*sync.RWMutex
}

type openseaData struct {
	conselhoOPStats *openseaStats
	brasaoOPStats   *openseaStats
	*sync.RWMutex
	// monstros --coming soon--
}

type marketData struct {
	conselhoStats *marketStats
	brasaoStats   *marketStats
	*sync.RWMutex
	// monstros --coming soon--
}

type openseaStats struct {
	OneDayDiff           float64 `json:"one_day_diff"`
	OneDaySales          float64 `json:"one_day_sales"`
	OneDayVolume         float64 `json:"one_day_volume"`
	OneDayAveragePrice   float64 `json:"one_day_average_price"`
	SevenDaysDiff        float64 `json:"seven_days_diff"`
	SevenDaySales        float64 `json:"seven_day_sales"`
	SevenDayVolume       float64 `json:"seven_day_volume"`
	SevenDayAveragePrice float64 `json:"seven_day_average_price"`
	TotalVolume          float64 `json:"total_volume"`
	FloorPrice           float64 `json:"floor_price"`
	NumOwners            int64   `json:"num_owners"`
	TotalSales           float64 `json:"total_sales"`
	AveragePrice         float64 `json:"average_price"`
}

type marketStats struct {
	sales             []sale
	lastSale          float64
	totalSales        int64
	thirtyDaysPercent float64
	thirtyDaysSales   int64
	sevenDaysPercent  float64
	sevenDaysSales    int64
	oneDayPercent     float64
	oneDaySales       int64
}

type sale struct {
	timestamp time.Time
	price     float64
}

type Bot struct {
	*discordgo.Session
	*ethclient.Client
	*brasao.Brasao
	*brazuera.Brazuera
	*brazueraData
	*sync.WaitGroup
}

func newBot(s *discordgo.Session, e *ethclient.Client) Bot {
	instance, err := brazuera.NewBrazuera(common.HexToAddress("0xC6A633cE248ed3D30317464343B33B4C647BE5Ac"), e)
	if err != nil {
		log.Fatal().Msgf("error while instantiating Brazuera contract : %s ", err.Error())
	}
	instc, err := brasao.NewBrasao(common.HexToAddress("0x969C2E2021bBc5987f67CfDb34877fe53f932AB8"), e)

	if err != nil {
		log.Fatal().Msgf("error while instantiating Brasao contract : %s ", err.Error())
	}

	return Bot{
		Session:   s,
		Client:    e,
		Brazuera:  instance,
		Brasao:    instc,
		WaitGroup: &sync.WaitGroup{},
		brazueraData: &brazueraData{
			conselhoToAddress:    make(map[int]string),
			addressToBrasaoCount: make(map[string]int),
			RWMutex:              &sync.RWMutex{},
			marketData: marketData{
				conselhoStats: &marketStats{
					sales: make([]sale, 0),
				},
				brasaoStats: &marketStats{
					sales: make([]sale, 0),
				},
				RWMutex: &sync.RWMutex{},
			},
			openseaData: openseaData{
				RWMutex: &sync.RWMutex{},
			},
		},
	}
}

func init() {
	_ = godotenv.Load(".env")
}

func main() {
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatal().Msgf("error while instantiating Discord bot : %s ", err.Error())
	}

	client, err := ethclient.Dial(os.Getenv("ETH_RPC"))

	bot := newBot(dg, client)

	err = bot.Open()
	if err != nil {
		log.Fatal().Msgf("error while opening connection with Discord : %s ", err.Error())
	}

	go bot.updateData()
	go bot.getOpenseaStats()
	go bot.updateBotInfo()
	go bot.getContractSales()

	bot.AddHandler(bot.messageHandler)

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	_ = bot.Session.Close()

}

func (b *Bot) updateData() {
	firstExec := true
	ticker := *time.NewTicker(time.Second)
	for {
		<-ticker.C

		if firstExec {
			ticker.Stop()
			ticker = *time.NewTicker(time.Minute * 15)
			firstExec = false
		}
		b.addressToConselho = make(map[string][]int)

		for i := 0; i < 500; i = i + 50 {
			b.Add(1)

			go func(i int, b *Bot) {
				defer b.Done()
				initial := i
				for ; i <= initial+50; i++ {
					if i == 0 {
						continue
					}

					var in big.Int
					in.SetInt64(int64(i))

					add, err := b.Brazuera.OwnerOf(nil, &in)

					if err != nil {
						fmt.Println(err)
					}
					b.brazueraData.RWMutex.Lock()
					b.addressToConselho[add.Hex()] = append(b.addressToConselho[add.Hex()], i)
					b.conselhoToAddress[i] = add.Hex()
					b.brazueraData.RWMutex.Unlock()
				}
			}(i, b)

		}

		b.Wait()
		b.getBrasaoCounts()

		b.brazueraData.RWMutex.Lock()
		b.hasBothCount = 0
		for add := range b.addressToConselho {

			if b.addressToBrasaoCount[add] != 0 {
				b.hasBothCount++
			}

		}
		b.brazueraData.RWMutex.Unlock()
	}
}

func (b *Bot) getBrasaoCounts() {
	var cursor string

	for i := 1; i <= 5; i++ {
		httpClient := http.Client{}

		req, err := http.NewRequest("GET", os.Getenv("MORALIS_ENDPOINT"), nil)

		if err != nil {
			log.Fatal().Msgf("error while instantiating new http request for moralis: %s ", err.Error())
		}

		req.Header.Set("X-API-Key", os.Getenv("MORALIS_APIKEY"))
		req.Header.Set("Content-Type", "application/json")

		q := req.URL.Query()
		q.Add("chain", "eth")
		q.Add("format", "decimal")
		q.Add("normalizeMetadata", "false")

		if cursor != "" {
			q.Add("cursor", cursor)
		}

		req.URL.RawQuery = q.Encode()

		resp, err := httpClient.Do(req)
		if err != nil {
			log.Fatal().Msgf("error while http requesting moralis: %s ", err.Error())
		}

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatal().Msgf("error while reading body on moralis: %s ", err.Error())
		}

		if resp.StatusCode != 200 {

			log.Fatal().Msgf("moralis response code: %s , body: ", resp.StatusCode, string(body))
		}

		var rawData map[string]interface{}
		err = json.Unmarshal(body, &rawData)

		if err != nil {
			log.Fatal().Msgf("error while unmarshalling moralis body: %s ", err.Error())
		}

		if rawData["cursor"] != nil {
			cursor = rawData["cursor"].(string)
		}
		rs := rawData["result"]

		for _, r := range rs.([]interface{}) {
			add := r.(map[string]interface{})["owner_of"]
			amount := r.(map[string]interface{})["amount"]
			i, _ := strconv.Atoi(amount.(string))
			mixedAdd := common.NewMixedcaseAddress(common.HexToAddress(add.(string)))

			b.addressToBrasaoCount[mixedAdd.Original()] = i
		}

		err = resp.Body.Close()
		if err != nil {
			log.Fatal().Msgf("error while closing moralis body: %s ", err.Error())
		}

	}

}

func (b *Bot) getOpenseaStats() {
	firstExec := true
	ticker := *time.NewTicker(time.Second * 10)
	for {
		<-ticker.C

		if firstExec {
			ticker.Stop()
			ticker = *time.NewTicker(time.Second * 180)
			firstExec = false
		}

		collections := [2]string{"brazuera-o-conselho", "brasao-conselheiro"}

		for i, c := range collections {
			httpClient := http.Client{}

			req, err := http.NewRequest("GET", os.Getenv("OPENSEA_ENDPOINT")+c+"/stats", nil)

			if err != nil {
				log.Fatal().Msgf("error while instantiating new http request for opensea: %s ", err.Error())
			}

			req.Header.Set("accept", "application/json")

			resp, err := httpClient.Do(req)
			if err != nil {
				log.Fatal().Msgf("error while http requesting opensea: %s ", err.Error())
			}

			body, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				log.Fatal().Msgf("error while reading opensea body: %s ", err.Error())
			}

			if resp.StatusCode != 200 {

				log.Fatal().Msgf("opensea response code: %s , body: ", resp.StatusCode, string(body))
			}

			var stats = struct {
				Stats openseaStats `json:"stats"`
			}{}

			err = json.Unmarshal(body, &stats)

			if err != nil {
				log.Fatal().Msgf("error while unmarshalling opensea body: %s ", err.Error())
			}
			b.openseaData.RWMutex.Lock()
			if (i+1)%2 == 1 {
				b.conselhoOPStats = &stats.Stats

			} else {
				b.brasaoOPStats = &stats.Stats
			}
			b.openseaData.RWMutex.Unlock()

			err = resp.Body.Close()
			if err != nil {
				log.Fatal().Msgf("error closing opensea body: %s ", err.Error())
			}

		}

	}
}

func (b *Bot) getContractSales() {
	firstExec := true
	ticker := *time.NewTicker(time.Second * 10)
	for {
		<-ticker.C

		if firstExec {
			ticker.Stop()
			ticker = *time.NewTicker(time.Second * 720)
			firstExec = false
		}
		b.marketData.RWMutex.Lock()
		b.conselhoStats.sales = make([]sale, 0)
		b.brasaoStats.sales = make([]sale, 0)

		collections := [2]string{"b3c28cf9d336330bcc22ede2506c4308", "b711dad1795ea8fc9ea0158756a7ecb4"}

		for i, c := range collections {
			var next string
			for {
				httpClient := http.Client{}
				var req *http.Request
				var err error
				if next == "" {
					req, err = http.NewRequest("GET", os.Getenv("SIMPLEHASH_ENDPOINT")+c+"?order_by=timestamp_desc&limit=50", nil)
				} else {
					req, err = http.NewRequest("GET", next, nil)
				}
				if err != nil {
					log.Fatal().Msgf("error while instantiating new http request for simplehash: %s ", err.Error())
				}

				req.Header.Set("accept", "application/json")
				req.Header.Set("X-API-KEY", os.Getenv("SIMPLEHASH_APIKEY"))

				resp, err := httpClient.Do(req)
				if err != nil {
					log.Fatal().Msgf("error while http requesting simplehash: %s ", err.Error())
				}

				body, err := ioutil.ReadAll(resp.Body)

				if err != nil {
					log.Fatal().Msgf("error while reading simplehash body: %s ", err.Error())
				}

				var rawData map[string]interface{}
				err = json.Unmarshal(body, &rawData)

				if err != nil {
					log.Fatal().Msgf("error while unmarshalling simplehash body: %s ", err.Error())
				}

				if resp.StatusCode != 200 {

					log.Fatal().Msgf("simplehash response code: %s , body: ", resp.StatusCode, string(body))
				}

				if rawData["next"] != nil {
					next = rawData["next"].(string)
				} else {
					next = ""
				}
				trs := rawData["transfers"].([]interface{})

				for _, t := range trs {
					if t.(map[string]interface{})["sale_details"] == nil {
						continue
					}

					ti, err := time.Parse(time.RFC3339, t.(map[string]interface{})["timestamp"].(string))

					if err != nil {
						log.Fatal().Msgf("could not parse timestamp: %s ", err.Error())
					}

					dts := t.(map[string]interface{})["sale_details"]

					var s = sale{
						timestamp: ti,
						price:     dts.(map[string]interface{})["unit_price"].(float64) / 1000000000000000000,
					}

					if (i+1)%2 == 1 {
						b.conselhoStats.sales = append(b.conselhoStats.sales, s)

					} else {
						b.brasaoStats.sales = append(b.brasaoStats.sales, s)
					}
				}

				err = resp.Body.Close()
				if err != nil {
					log.Fatal().Msgf("error while closing simplehash body: %s ", err.Error())
				}

				if next == "" {
					break
				}

			}
		}

		for _, stat := range []*marketStats{b.conselhoStats, b.brasaoStats} {
			stat.lastSale = stat.sales[0].price
			stat.oneDayPercent, stat.sevenDaysPercent, stat.thirtyDaysPercent = 0, 0, 0
			stat.oneDaySales, stat.sevenDaysSales, stat.thirtyDaysSales = 0, 0, 0
			one, seven, thirty := false, false, false

			for _, s := range stat.sales {

				if time.Since(s.timestamp).Hours() >= 24 && one == false {
					stat.oneDayPercent = ((stat.lastSale - s.price) / s.price) * 100
					one = true
				} else if time.Since(s.timestamp).Hours() >= 168 && seven == false {
					stat.sevenDaysPercent = ((stat.lastSale - s.price) / s.price) * 100
					seven = true
				} else if time.Since(s.timestamp).Hours() >= 720 && thirty == false {
					stat.thirtyDaysPercent = ((stat.lastSale - s.price) / s.price) * 100
					thirty = true
				}

				if time.Since(s.timestamp).Hours() <= 24 {
					stat.oneDaySales++
				}
				if time.Since(s.timestamp).Hours() <= 168 {
					stat.sevenDaysSales++
				}
				if time.Since(s.timestamp).Hours() <= 720 {
					stat.thirtyDaysSales++
				}

				stat.totalSales++

			}

		}
		b.marketData.RWMutex.Unlock()
	}
}
