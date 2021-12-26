
package options

import (
  "fmt"
  "log"
  "net/http"
  "strings"
  "io/ioutil"
  "encoding/json"
  "github.com/patrickmn/go-cache"
  "time"
  "github.com/joho/godotenv"
  "os"
//  "errors"
  "net/url"
)

type OptionsData struct {
	OptionChain struct {
		Result []struct {
			UnderlyingSymbol string    `json:"underlyingSymbol"`
			ExpirationDates  []int     `json:"expirationDates"`
			Strikes          []float64 `json:"strikes"`
			HasMiniOptions   bool      `json:"hasMiniOptions"`
			Quote            struct {
				Language                          string  `json:"language"`
				Region                            string  `json:"region"`
				QuoteType                         string  `json:"quoteType"`
				QuoteSourceName                   string  `json:"quoteSourceName"`
				Triggerable                       bool    `json:"triggerable"`
				ShortName                         string  `json:"shortName"`
				Currency                          string  `json:"currency"`
				FirstTradeDateMilliseconds        int64   `json:"firstTradeDateMilliseconds"`
				PriceHint                         int     `json:"priceHint"`
				PostMarketChangePercent           float64 `json:"postMarketChangePercent"`
				PostMarketTime                    int     `json:"postMarketTime"`
				PostMarketPrice                   float64 `json:"postMarketPrice"`
				PostMarketChange                  float64 `json:"postMarketChange"`
				RegularMarketChange               float64 `json:"regularMarketChange"`
				RegularMarketChangePercent        float64 `json:"regularMarketChangePercent"`
				RegularMarketTime                 int     `json:"regularMarketTime"`
				MarketState                       string  `json:"marketState"`
				RegularMarketPrice                float64 `json:"regularMarketPrice"`
				RegularMarketDayHigh              float64 `json:"regularMarketDayHigh"`
				RegularMarketDayRange             string  `json:"regularMarketDayRange"`
				RegularMarketDayLow               float64 `json:"regularMarketDayLow"`
				RegularMarketVolume               int     `json:"regularMarketVolume"`
				RegularMarketPreviousClose        float64 `json:"regularMarketPreviousClose"`
				Bid                               float64 `json:"bid"`
				Ask                               float64 `json:"ask"`
				BidSize                           int     `json:"bidSize"`
				AskSize                           int     `json:"askSize"`
				FullExchangeName                  string  `json:"fullExchangeName"`
				FinancialCurrency                 string  `json:"financialCurrency"`
				RegularMarketOpen                 float64 `json:"regularMarketOpen"`
				AverageDailyVolume3Month          int     `json:"averageDailyVolume3Month"`
				AverageDailyVolume10Day           int     `json:"averageDailyVolume10Day"`
				FiftyTwoWeekLowChange             float64 `json:"fiftyTwoWeekLowChange"`
				FiftyTwoWeekLowChangePercent      float64 `json:"fiftyTwoWeekLowChangePercent"`
				FiftyTwoWeekRange                 string  `json:"fiftyTwoWeekRange"`
				FiftyTwoWeekHighChange            float64 `json:"fiftyTwoWeekHighChange"`
				FiftyTwoWeekHighChangePercent     float64 `json:"fiftyTwoWeekHighChangePercent"`
				FiftyTwoWeekLow                   float64 `json:"fiftyTwoWeekLow"`
				FiftyTwoWeekHigh                  float64 `json:"fiftyTwoWeekHigh"`
				EarningsTimestamp                 int     `json:"earningsTimestamp"`
				EarningsTimestampStart            int     `json:"earningsTimestampStart"`
				EarningsTimestampEnd              int     `json:"earningsTimestampEnd"`
				EpsTrailingTwelveMonths           float64 `json:"epsTrailingTwelveMonths"`
				EpsForward                        float64 `json:"epsForward"`
				EpsCurrentYear                    float64 `json:"epsCurrentYear"`
				PriceEpsCurrentYear               float64 `json:"priceEpsCurrentYear"`
				SharesOutstanding                 int     `json:"sharesOutstanding"`
				BookValue                         float64 `json:"bookValue"`
				FiftyDayAverage                   float64 `json:"fiftyDayAverage"`
				FiftyDayAverageChange             float64 `json:"fiftyDayAverageChange"`
				FiftyDayAverageChangePercent      float64 `json:"fiftyDayAverageChangePercent"`
				TwoHundredDayAverage              float64 `json:"twoHundredDayAverage"`
				TwoHundredDayAverageChange        float64 `json:"twoHundredDayAverageChange"`
				TwoHundredDayAverageChangePercent float64 `json:"twoHundredDayAverageChangePercent"`
				MarketCap                         int     `json:"marketCap"`
				ForwardPE                         float64 `json:"forwardPE"`
				PriceToBook                       float64 `json:"priceToBook"`
				SourceInterval                    int     `json:"sourceInterval"`
				ExchangeDataDelayedBy             int     `json:"exchangeDataDelayedBy"`
				PageViewGrowthWeekly              float64 `json:"pageViewGrowthWeekly"`
				AverageAnalystRating              string  `json:"averageAnalystRating"`
				Tradeable                         bool    `json:"tradeable"`
				Exchange                          string  `json:"exchange"`
				LongName                          string  `json:"longName"`
				MessageBoardID                    string  `json:"messageBoardId"`
				ExchangeTimezoneName              string  `json:"exchangeTimezoneName"`
				ExchangeTimezoneShortName         string  `json:"exchangeTimezoneShortName"`
				GmtOffSetMilliseconds             int     `json:"gmtOffSetMilliseconds"`
				Market                            string  `json:"market"`
				EsgPopulated                      bool    `json:"esgPopulated"`
				DisplayName                       string  `json:"displayName"`
				Symbol                            string  `json:"symbol"`
			} `json:"quote"`
			Options []struct {
				ExpirationDate int  `json:"expirationDate"`
				HasMiniOptions bool `json:"hasMiniOptions"`
				Calls          []struct {
					ContractSymbol    string  `json:"contractSymbol"`
					Strike            float64 `json:"strike"`
					Currency          string  `json:"currency"`
					LastPrice         float64 `json:"lastPrice"`
					Change            float64 `json:"change"`
					PercentChange     float64 `json:"percentChange,omitempty"`
					Volume            int     `json:"volume,omitempty"`
					OpenInterest      int     `json:"openInterest"`
					Bid               float64 `json:"bid"`
					Ask               float64 `json:"ask"`
					ContractSize      string  `json:"contractSize"`
					Expiration        int     `json:"expiration"`
					LastTradeDate     int     `json:"lastTradeDate"`
					ImpliedVolatility float64 `json:"impliedVolatility"`
					InTheMoney        bool    `json:"inTheMoney"`
				} `json:"calls"`
				Puts []struct {
					ContractSymbol    string  `json:"contractSymbol"`
					Strike            float64 `json:"strike"`
					Currency          string  `json:"currency"`
					LastPrice         float64 `json:"lastPrice"`
					Change            float64 `json:"change"`
					PercentChange     float64 `json:"percentChange,omitempty"`
					Volume            int     `json:"volume,omitempty"`
					OpenInterest      int     `json:"openInterest"`
					Bid               float64 `json:"bid"`
					Ask               float64 `json:"ask"`
					ContractSize      string  `json:"contractSize"`
					Expiration        int     `json:"expiration"`
					LastTradeDate     int     `json:"lastTradeDate"`
					ImpliedVolatility float64 `json:"impliedVolatility"`
					InTheMoney        bool    `json:"inTheMoney"`
				} `json:"puts"`
			} `json:"options"`
		} `json:"result"`
		Error interface{} `json:"error"`
	} `json:"optionChain"`
}

var (
  t string
  total float64
  optionsCache = cache.New(60*time.Minute, 60*time.Minute)
)

func GetOptionsData(w http.ResponseWriter, r *http.Request){

  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  log.Println("Incoming request: ", r.RemoteAddr)
  u, err := url.Parse(r.URL.RequestURI())
  if err != nil{
    log.Fatal(err)
  }

  m,_ := url.ParseQuery(u.RawQuery)
  ticker := strings.ToUpper(m["symbol"][0])

  val, found := optionsCache.Get(ticker)
  if found{
    log.Println(ticker, "retrieved from options cache")
    data := val.(*OptionsData)
    json.NewEncoder(w).Encode(data)
    return
  }

  var op OptionsData
  client := &http.Client{}
  URL := fmt.Sprintf("https://yfapi.net/v7/finance/options/%s",ticker)

  req, err := http.NewRequest("GET",URL,nil)
  req.Header.Add("X-API-KEY",os.Getenv("API_KEY"))
  req.Header.Add("Content-Type", "application/json")

  res, err := client.Do(req)
  if err != nil{
    log.Fatal(err)
  }

  if res.StatusCode == http.StatusOK{
    body, err := ioutil.ReadAll(res.Body)
    if err != nil{
      log.Fatal(err)
    }
    defer res.Body.Close()
    err = json.Unmarshal(body,&op)
    if err != nil{
      log.Println(err)
    }
    optionsCache.Set(ticker,&op,cache.DefaultExpiration)
    log.Println(ticker, " set to options cache")
    json.NewEncoder(w).Encode(op)
  }
}
