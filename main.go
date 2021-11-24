package main

import (
  "fmt"
  "strings"
  "go-invest/options"
  "go-invest/quote"
  "os"
)


var (
  tickerName string
  c string
)


func main(){

  fmt.Println("Options: o | Trends: t | Quit: q | Recommendations: r")
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
        go options.GetYahooOptionsInfo(tickerName)
      case strings.ToUpper("t"):
        go options.GetTrends()
      case strings.ToUpper("r"):
        fmt.Print("\nEnter ticker to retrieve market data: ")
        fmt.Scanf("%s", &tickerName)
        fmt.Println()
        go quote.GetRecommendations(tickerName)
      default:

    }

  }



}
