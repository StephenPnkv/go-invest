
package charts

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

type ChartData struct {
	Chart struct {
		Result []struct {
			Meta struct {
				Currency             string  `json:"currency"`
				Symbol               string  `json:"symbol"`
				ExchangeName         string  `json:"exchangeName"`
				InstrumentType       string  `json:"instrumentType"`
				FirstTradeDate       int     `json:"firstTradeDate"`
				RegularMarketTime    int     `json:"regularMarketTime"`
				Gmtoffset            int     `json:"gmtoffset"`
				Timezone             string  `json:"timezone"`
				ExchangeTimezoneName string  `json:"exchangeTimezoneName"`
				RegularMarketPrice   float64 `json:"regularMarketPrice"`
				ChartPreviousClose   float64 `json:"chartPreviousClose"`
				PriceHint            int     `json:"priceHint"`
				CurrentTradingPeriod struct {
					Pre struct {
						Timezone  string `json:"timezone"`
						Start     int    `json:"start"`
						End       int    `json:"end"`
						Gmtoffset int    `json:"gmtoffset"`
					} `json:"pre"`
					Regular struct {
						Timezone  string `json:"timezone"`
						Start     int    `json:"start"`
						End       int    `json:"end"`
						Gmtoffset int    `json:"gmtoffset"`
					} `json:"regular"`
					Post struct {
						Timezone  string `json:"timezone"`
						Start     int    `json:"start"`
						End       int    `json:"end"`
						Gmtoffset int    `json:"gmtoffset"`
					} `json:"post"`
				} `json:"currentTradingPeriod"`
				DataGranularity string   `json:"dataGranularity"`
				Range           string   `json:"range"`
				ValidRanges     []string `json:"validRanges"`
			} `json:"meta"`
			Timestamp   []int `json:"timestamp"`
			Comparisons []struct {
				Symbol             string    `json:"symbol"`
				High               []float64 `json:"high"`
				Low                []float64 `json:"low"`
				ChartPreviousClose float64   `json:"chartPreviousClose"`
				Close              []float64 `json:"close"`
				Open               []float64 `json:"open"`
			} `json:"comparisons"`
			Indicators struct {
				Quote []struct {
					Low    []float64 `json:"low"`
					High   []float64 `json:"high"`
					Open   []float64 `json:"open"`
					Close  []float64 `json:"close"`
					Volume []int     `json:"volume"`
				} `json:"quote"`
				Adjclose []struct {
					Adjclose []float64 `json:"adjclose"`
				} `json:"adjclose"`
			} `json:"indicators"`
		} `json:"result"`
		Error interface{} `json:"error"`
	} `json:"chart"`
}

var (
  chartCache = cache.New(60*time.Minute, 60*time.Minute)
)

func GetChart(w http.ResponseWriter, r *http.Request){

  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  log.Println("GET /api/chart => IP addr: ", r.RemoteAddr)
  u, err := url.Parse(r.URL.RequestURI())
  if err != nil{
    log.Fatal(err)
  }

  m,_ := url.ParseQuery(u.RawQuery)
  ticker := strings.ToUpper(m["symbol"][0])

  val, found := chartCache.Get(ticker)
  if found{
    log.Println(ticker, " retrieved from chart cache.")
    data := val.(*ChartData)
    json.NewEncoder(w).Encode(data)
    return
  }

  var ch ChartData
  client := &http.Client{}

  URL := fmt.Sprintf("https://yfapi.net/v8/finance/chart/%s?range=1mo&region=US&interval=1d&lang=en",ticker)

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
    err = json.Unmarshal(body,&ch)
    if err != nil{
      log.Println(err)
    }

    chartCache.Set(ch.Chart.Result[0].Meta.Symbol, &ch, cache.DefaultExpiration)
    log.Println(ch.Chart.Result[0].Meta.Symbol, " set in chart cache")
    json.NewEncoder(w).Encode(ch)
  }
}
