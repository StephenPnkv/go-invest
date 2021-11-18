package main

import (
  "fmt"
  "math"
  "io/ioutil"
  "net/http"
  "log"
  "encoding/json"
  "strings"
  "go-invest/options"
  "os"
)


var (
  startDate = "2020-01-08"
  endDate = "2021-04-08"
  tickerName = ""
)

type QuoteData struct {
	QuoteResponse struct {
		Result []struct {
			Language                          string  `json:"language"`
			Region                            string  `json:"region"`
			QuoteType                         string  `json:"quoteType"`
			QuoteSourceName                   string  `json:"quoteSourceName"`
			Triggerable                       bool    `json:"triggerable"`
			Currency                          string  `json:"currency"`
			MarketState                       string  `json:"marketState"`
			EpsCurrentYear                    float64 `json:"epsCurrentYear"`
			PriceEpsCurrentYear               float64 `json:"priceEpsCurrentYear"`
			SharesOutstanding                 int64   `json:"sharesOutstanding"`
			BookValue                         float64 `json:"bookValue"`
			FiftyDayAverage                   float64 `json:"fiftyDayAverage"`
			FiftyDayAverageChange             float64 `json:"fiftyDayAverageChange"`
			FiftyDayAverageChangePercent      float64 `json:"fiftyDayAverageChangePercent"`
			TwoHundredDayAverage              float64 `json:"twoHundredDayAverage"`
			TwoHundredDayAverageChange        float64 `json:"twoHundredDayAverageChange"`
			TwoHundredDayAverageChangePercent float64 `json:"twoHundredDayAverageChangePercent"`
			MarketCap                         int64   `json:"marketCap"`
			ForwardPE                         float64 `json:"forwardPE"`
			PriceToBook                       float64 `json:"priceToBook"`
			SourceInterval                    int     `json:"sourceInterval"`
			ExchangeDataDelayedBy             int     `json:"exchangeDataDelayedBy"`
			AverageAnalystRating              string  `json:"averageAnalystRating"`
			Tradeable                         bool    `json:"tradeable"`
			ShortName                         string  `json:"shortName"`
			Exchange                          string  `json:"exchange"`
			LongName                          string  `json:"longName"`
			MessageBoardID                    string  `json:"messageBoardId"`
			ExchangeTimezoneName              string  `json:"exchangeTimezoneName"`
			ExchangeTimezoneShortName         string  `json:"exchangeTimezoneShortName"`
			GmtOffSetMilliseconds             int     `json:"gmtOffSetMilliseconds"`
			Market                            string  `json:"market"`
			EsgPopulated                      bool    `json:"esgPopulated"`
			FirstTradeDateMilliseconds        int64   `json:"firstTradeDateMilliseconds"`
			PriceHint                         int     `json:"priceHint"`
			PostMarketChangePercent           float64 `json:"postMarketChangePercent"`
			PostMarketTime                    int     `json:"postMarketTime"`
			PostMarketPrice                   float64 `json:"postMarketPrice"`
			PostMarketChange                  float64 `json:"postMarketChange"`
			RegularMarketChange               float64 `json:"regularMarketChange"`
			RegularMarketChangePercent        float64 `json:"regularMarketChangePercent"`
			RegularMarketTime                 int     `json:"regularMarketTime"`
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
			DividendDate                      int     `json:"dividendDate"`
			EarningsTimestamp                 int     `json:"earningsTimestamp"`
			EarningsTimestampStart            int     `json:"earningsTimestampStart"`
			EarningsTimestampEnd              int     `json:"earningsTimestampEnd"`
			TrailingAnnualDividendRate        float64 `json:"trailingAnnualDividendRate"`
			TrailingPE                        float64 `json:"trailingPE"`
			TrailingAnnualDividendYield       float64 `json:"trailingAnnualDividendYield"`
			EpsTrailingTwelveMonths           float64 `json:"epsTrailingTwelveMonths"`
			EpsForward                        float64 `json:"epsForward"`
			DisplayName                       string  `json:"displayName"`
			Symbol                            string  `json:"symbol"`
		} `json:"result"`
		Error interface{} `json:"error"`
	} `json:"quoteResponse"`
}



func GetTickerInfo(ticker string){
  var td QuoteData
  client := &http.Client{}
  t := strings.ToUpper(ticker)
  URL := fmt.Sprintf("https://yfapi.net/v6/finance/quote?region=US&lang=en&symbols=%s",t)
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

    err = json.Unmarshal(body,&td)
    if err != nil{
      log.Println("No ticker found.")
      return
    }
    fmt.Printf("\n| Ticker: $%s\n", t)
    fmt.Printf("| Market cap: %d \n| Price-Earnings ratio: %.2f \n| Dividend yield: %.2f \n| Average Volume: %d \n",td.QuoteResponse.Result[0].MarketCap, td.QuoteResponse.Result[0].PriceEpsCurrentYear, td.QuoteResponse.Result[0].TrailingAnnualDividendYield, td.QuoteResponse.Result[0].AverageDailyVolume10Day)
    fmt.Printf("| High today: %.2f \n| Low today: %.2f \n| Open price: %.2f  \n| Volume: %d \n",td.QuoteResponse.Result[0].RegularMarketDayHigh, td.QuoteResponse.Result[0].RegularMarketDayLow, td.QuoteResponse.Result[0].RegularMarketOpen, td.QuoteResponse.Result[0].RegularMarketVolume)
    fmt.Printf("| 52 Week high: %.2f \n| 52 Week low: %.2f \n| Post-market price: %.2f \n",td.QuoteResponse.Result[0].FiftyTwoWeekHigh, td.QuoteResponse.Result[0].FiftyTwoWeekLow,td.QuoteResponse.Result[0].PostMarketPrice)
  }

}


func percentDiff(current, previous float64) float64{
  var res float64
  res = math.Abs(current - previous) / ( math.Abs(current + previous) / 2 )
  return res * 100
}


func main(){

  for{
    fmt.Print("\nEnter ticker to retrieve market data: ")
    fmt.Scanf("%s", &tickerName)
    fmt.Println()
    options.GetYahooOptionsInfo(tickerName)
  }



}
