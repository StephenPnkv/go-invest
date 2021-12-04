
package options

import (
  "fmt"
  "log"
  "net/http"
  "strings"
  "io/ioutil"
  "encoding/json"
//  "github.com/jwalton/gchalk"
  "time"
  "github.com/joho/godotenv"
  "os"
  "errors"
)

type YahooQuote struct {
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
)

func GetStockInfo(w http.ResponseWriter,ticker string) (YahooQuote, error){

  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  var op YahooQuote
  client := &http.Client{}
  t = strings.ToUpper(ticker)
  URL := fmt.Sprintf("https://yfapi.net/v7/finance/options/%s",t)
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
    return op,nil
  }else{
    return op, errors.New(fmt.Sprintf("\nError: %s\n", http.StatusNoContent))
  }
}

func PrintFormattedQuoteData(w http.ResponseWriter, op YahooQuote){
  fmt.Fprint(w,"\tMarket cap.\tPrice\tHigh\tLow\tOpen\t52-week High\t52-week Low\tAvg. volume\tVolume\t\tP/E\t\n")
  for i := 0; i < len(op.OptionChain.Result);i++{
    fmt.Fprintf(w,"\t%d\t%.2f\t%.2f\t%.2f\t%.2f\t%.2f\t\t%.2f\t\t%d\t%d\t%.2f\n",
      op.OptionChain.Result[i].Quote.MarketCap,
      op.OptionChain.Result[i].Quote.RegularMarketPrice,
      op.OptionChain.Result[i].Quote.RegularMarketDayHigh,
      op.OptionChain.Result[i].Quote.RegularMarketDayLow,
      op.OptionChain.Result[i].Quote.RegularMarketOpen,
      op.OptionChain.Result[i].Quote.FiftyTwoWeekHigh,
      op.OptionChain.Result[i].Quote.FiftyTwoWeekLow,
      op.OptionChain.Result[i].Quote.AverageDailyVolume10Day,
      op.OptionChain.Result[i].Quote.RegularMarketVolume,
      op.OptionChain.Result[i].Quote.PriceEpsCurrentYear,
    )
  }
  fmt.Fprint(w,"\n")
}

func GetOpenInterestStats(op YahooQuote)(int, int){
  totalPutInterest, totalCallInterest := 0,0

  for res := 0; res < len(op.OptionChain.Result); res++{

    for i := 0; i < len(op.OptionChain.Result[res].Options); i++{

      for j := 0; j < len(op.OptionChain.Result[res].Options[i].Puts); j++{
        totalPutInterest += op.OptionChain.Result[res].Options[i].Puts[j].OpenInterest
      }

      for j := 0; j < len(op.OptionChain.Result[res].Options[i].Calls); j++{
        totalCallInterest += op.OptionChain.Result[res].Options[i].Calls[j].OpenInterest
      }

    }

  }
  return totalPutInterest,totalCallInterest
}

func PrintFormattedOptionsData(w http.ResponseWriter, op YahooQuote){
  //#FF3333 - red
  //#00FF80 - green
  for res := 0; res < len(op.OptionChain.Result); res++{

    for i := 0; i < len(op.OptionChain.Result[res].Options); i++{
      //Options expiration dates returns as seconds since Thursday, 1 January 1970 from the Yahoo API
      exp:= int64(op.OptionChain.Result[res].Options[i].ExpirationDate)
      expDate := time.Unix(exp,0)

      fmt.Fprintf(w,"\t$%s\t\t\t\t   %s\t\t\t\t\t\t\t\n",op.OptionChain.Result[res].Quote.Symbol,expDate)
      fmt.Fprint(w,"\t\t\t\t\t\t       Puts\t\t\t\t\t\t\t\t\t\n")
      fmt.Fprint(w,"\tStrike\t\tBid\t\tAsk\t\tVolume\t\tOpen Int.\tIV\t\t% Change\tITM\t\n")
      for j := 0; j < len(op.OptionChain.Result[res].Options[i].Puts); j++{
        pAsk := (fmt.Sprintf("%.2f", op.OptionChain.Result[res].Options[i].Puts[j].Ask))
        pBid := (fmt.Sprintf("%.2f", op.OptionChain.Result[res].Options[i].Puts[j].Bid))
        var pChange string
        if op.OptionChain.Result[res].Options[i].Puts[j].PercentChange < 0{
          pChange = (fmt.Sprintf("%.2f", op.OptionChain.Result[res].Options[i].Puts[j].PercentChange))
        }else{
          pChange = (fmt.Sprintf("%.2f", op.OptionChain.Result[res].Options[i].Puts[j].PercentChange))
        }
        var itm string
        if op.OptionChain.Result[res].Options[i].Puts[j].InTheMoney == true{
          itm = (fmt.Sprintf("%v", op.OptionChain.Result[res].Options[i].Puts[j].InTheMoney))
        }else{
          itm = (fmt.Sprintf("%v", op.OptionChain.Result[res].Options[i].Puts[j].InTheMoney))
        }

        coloredData := fmt.Sprintf("\t%.2f\t\t%s\t\t%s\t",op.OptionChain.Result[res].Options[i].Puts[j].Strike,pBid,pAsk)
        fmt.Fprint(w,coloredData)

        plainData := fmt.Sprintf("\t%d\t\t%d\t\t%.2f\t\t%s\t\t%s\t\n",
          op.OptionChain.Result[res].Options[i].Puts[j].Volume,
          op.OptionChain.Result[res].Options[i].Puts[j].OpenInterest,
          op.OptionChain.Result[res].Options[i].Puts[j].ImpliedVolatility,
          pChange,
          itm)
        fmt.Fprint(w,plainData)
      }

      fmt.Fprint(w,"\t\t\t\t\t\t       Calls\t\t\t\t\t\t\t\t\t\n")
      for j := 0; j < len(op.OptionChain.Result[res].Options[i].Calls); j++{
        pAsk := (fmt.Sprintf("%.2f", op.OptionChain.Result[res].Options[i].Calls[j].Ask))
        pBid := (fmt.Sprintf("%.2f", op.OptionChain.Result[res].Options[i].Calls[j].Bid))
        var pChange string
        if op.OptionChain.Result[res].Options[i].Calls[j].PercentChange < 0{
          pChange = (fmt.Sprintf("%.2f", op.OptionChain.Result[res].Options[i].Calls[j].PercentChange))
        }else{
          pChange = (fmt.Sprintf("%.2f", op.OptionChain.Result[res].Options[i].Calls[j].PercentChange))
        }
        var itm string
        if op.OptionChain.Result[res].Options[i].Calls[j].InTheMoney == true{
          itm = (fmt.Sprintf("%v", op.OptionChain.Result[res].Options[i].Calls[j].InTheMoney))
        }else{
          itm = (fmt.Sprintf("%v", op.OptionChain.Result[res].Options[i].Calls[j].InTheMoney))
        }

        coloredData := fmt.Sprintf("\t%.2f\t\t%s\t\t%s\t",op.OptionChain.Result[res].Options[i].Calls[j].Strike,pBid,pAsk)
        fmt.Fprint(w,coloredData)

        plainData := fmt.Sprintf("\t%d\t\t%d\t\t%.2f\t\t%s\t\t%s\t\n",
          op.OptionChain.Result[res].Options[i].Calls[j].Volume,
          op.OptionChain.Result[res].Options[i].Calls[j].OpenInterest,
          op.OptionChain.Result[res].Options[i].Calls[j].ImpliedVolatility,
          pChange,
          itm)
        fmt.Fprint(w,plainData)
      }
    }


  }
  fmt.Println("\n")
}
