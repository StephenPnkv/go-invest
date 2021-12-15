package quote

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
  "net/url"
)

type YahooQuote struct {
	QuoteResponse struct {
		Error  interface{} `json:"error"`
		Result []struct {
			Ask                               float64 `json:"ask"`
			AskSize                           int     `json:"askSize"`
			AverageDailyVolume10Day           int     `json:"averageDailyVolume10Day"`
			AverageDailyVolume3Month          int     `json:"averageDailyVolume3Month"`
			Bid                               float64 `json:"bid"`
			BidSize                           int     `json:"bidSize"`
			BookValue                         float64 `json:"bookValue"`
			Currency                          string  `json:"currency"`
			DisplayName                       string  `json:"displayName"`
			DividendDate                      int     `json:"dividendDate"`
			EarningsTimestamp                 int     `json:"earningsTimestamp"`
			EarningsTimestampEnd              int     `json:"earningsTimestampEnd"`
			EarningsTimestampStart            int     `json:"earningsTimestampStart"`
			EpsCurrentYear                    float64 `json:"epsCurrentYear"`
			EpsForward                        float64 `json:"epsForward"`
			EpsTrailingTwelveMonths           float64 `json:"epsTrailingTwelveMonths"`
			EsgPopulated                      bool    `json:"esgPopulated"`
			Exchange                          string  `json:"exchange"`
			ExchangeDataDelayedBy             int     `json:"exchangeDataDelayedBy"`
			ExchangeTimezoneName              string  `json:"exchangeTimezoneName"`
			ExchangeTimezoneShortName         string  `json:"exchangeTimezoneShortName"`
			FiftyDayAverage                   float64 `json:"fiftyDayAverage"`
			FiftyDayAverageChange             float64 `json:"fiftyDayAverageChange"`
			FiftyDayAverageChangePercent      float64 `json:"fiftyDayAverageChangePercent"`
			FiftyTwoWeekHigh                  float64 `json:"fiftyTwoWeekHigh"`
			FiftyTwoWeekHighChange            float64 `json:"fiftyTwoWeekHighChange"`
			FiftyTwoWeekHighChangePercent     float64 `json:"fiftyTwoWeekHighChangePercent"`
			FiftyTwoWeekLow                   float64 `json:"fiftyTwoWeekLow"`
			FiftyTwoWeekLowChange             float64 `json:"fiftyTwoWeekLowChange"`
			FiftyTwoWeekLowChangePercent      float64 `json:"fiftyTwoWeekLowChangePercent"`
			FiftyTwoWeekRange                 string  `json:"fiftyTwoWeekRange"`
			FinancialCurrency                 string  `json:"financialCurrency"`
			FirstTradeDateMilliseconds        int64   `json:"firstTradeDateMilliseconds"`
			ForwardPE                         float64 `json:"forwardPE"`
			FullExchangeName                  string  `json:"fullExchangeName"`
			GmtOffSetMilliseconds             int     `json:"gmtOffSetMilliseconds"`
			Language                          string  `json:"language"`
			LongName                          string  `json:"longName"`
			Market                            string  `json:"market"`
			MarketCap                         int64   `json:"marketCap"`
			MarketState                       string  `json:"marketState"`
			MessageBoardID                    string  `json:"messageBoardId"`
			PostMarketChange                  float64 `json:"postMarketChange"`
			PostMarketChangePercent           float64 `json:"postMarketChangePercent"`
			PostMarketPrice                   float64 `json:"postMarketPrice"`
			PostMarketTime                    int     `json:"postMarketTime"`
			PriceEpsCurrentYear               float64 `json:"priceEpsCurrentYear"`
			PriceHint                         int     `json:"priceHint"`
			PriceToBook                       float64 `json:"priceToBook"`
			QuoteSourceName                   string  `json:"quoteSourceName"`
			QuoteType                         string  `json:"quoteType"`
			Region                            string  `json:"region"`
			RegularMarketChange               float64 `json:"regularMarketChange"`
			RegularMarketChangePercent        float64 `json:"regularMarketChangePercent"`
			RegularMarketDayHigh              float64 `json:"regularMarketDayHigh"`
			RegularMarketDayLow               float64 `json:"regularMarketDayLow"`
			RegularMarketDayRange             string  `json:"regularMarketDayRange"`
			RegularMarketOpen                 float64 `json:"regularMarketOpen"`
			RegularMarketPreviousClose        float64 `json:"regularMarketPreviousClose"`
			RegularMarketPrice                float64 `json:"regularMarketPrice"`
			RegularMarketTime                 int     `json:"regularMarketTime"`
			RegularMarketVolume               int     `json:"regularMarketVolume"`
			SharesOutstanding                 int64   `json:"sharesOutstanding"`
			ShortName                         string  `json:"shortName"`
			SourceInterval                    int     `json:"sourceInterval"`
			Symbol                            string  `json:"symbol"`
			Tradeable                         bool    `json:"tradeable"`
			TrailingAnnualDividendRate        float64 `json:"trailingAnnualDividendRate"`
			TrailingAnnualDividendYield       float64 `json:"trailingAnnualDividendYield"`
			TrailingPE                        float64 `json:"trailingPE"`
			Triggerable                       bool    `json:"triggerable"`
			TwoHundredDayAverage              float64 `json:"twoHundredDayAverage"`
			TwoHundredDayAverageChange        float64 `json:"twoHundredDayAverageChange"`
			TwoHundredDayAverageChangePercent float64 `json:"twoHundredDayAverageChangePercent"`
		} `json:"result"`
	} `json:"quoteResponse"`
}

var(
  quoteCache = cache.New(1*time.Minute, 1*time.Minute)//Update cache every minute 
)

func GetQuote(w http.ResponseWriter, r *http.Request){

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

  val, found := quoteCache.Get(ticker)
  if found{
    log.Println(ticker, "quote retrieved from cache.")
    data := val.(*YahooQuote)
    json.NewEncoder(w).Encode(data)
    return
  }

  var quote YahooQuote
  client := &http.Client{}
  URL := fmt.Sprintf("https://yfapi.net/v6/finance/quote?region=US&lang=en&symbols=%s",ticker)

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
    err = json.Unmarshal(body,&quote)
    if err != nil{
      log.Println(err)
    }
    json.NewEncoder(w).Encode(quote)
  }
}
