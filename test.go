package main

import (
  "fmt"
  "strings"
  "go-invest/options"
  "go-invest/quote"
  "os"
  "github.com/patrickmn/go-cache"
	"time"
  "log"
)


var (
  tickerName string
  c string
)

func init(){
}


func main(){

  quotes := cache.New(60*time.Minute, 60*time.Minute)

  fmt.Println("Options menu: \nMarket data: o\nStock Recommendations: r\nTrending stocks: t \nQuit: q\n")
  for{
    fmt.Scanf("%s", &c)
    c = strings.ToUpper(c)
    switch c{
      case strings.ToUpper("q"):
        os.Exit(0)
      case strings.ToUpper("o"):
        fmt.Print("\nEnter ticker to retrieve market data: ")
        fmt.Scanf("%s", &tickerName)
        fmt.Println()

        val, found := quotes.Get(tickerName)
        if found{
          data := val.(*options.YahooQuote)
          options.PrintFormattedQuoteData(*data)
          options.PrintFormattedOptionsData(*data)
          break
        }

        quote, err := options.GetYahooOptionsInfo(tickerName)
        if err != nil{
          log.Println(err)
        }
        quotes.Set(tickerName,&quote, 10*time.Minute)
        options.PrintFormattedQuoteData(quote)
        options.PrintFormattedOptionsData(quote)

      case strings.ToUpper("t"):
        //o options.GetTrends()

      case strings.ToUpper("r"):
        fmt.Print("\nEnter ticker to retrieve market data: ")
        fmt.Scanf("%s", &tickerName)
        fmt.Println()
        go quote.GetRecommendations(tickerName)

      default:

    }

  }



}
