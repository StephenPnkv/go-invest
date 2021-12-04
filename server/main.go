package main

import (
 "fmt"
 "strings"
  "go-invest/options"
//  "go-invest/quote"
//  "os"
  "github.com/patrickmn/go-cache"
	"time"
  "log"
  "net/url"
  "net/http"
  "github.com/gorilla/mux"
)

var (
  quotes = cache.New(60*time.Minute, 60*time.Minute) //Update cache every hr
  //siData = cache.New(720*time.Minute, 720*time.Minute)//update cache every 12 hrs
)

func trends(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  options.GetTrends(w,r)
}


func stocks(w http.ResponseWriter, r *http.Request){
//  w.Header().Set("Content-Type", "application/json")

  u, err := url.Parse(r.URL.RequestURI())
  if err != nil{
    log.Fatal(err)
  }

  m,_ := url.ParseQuery(u.RawQuery)
  ticker := m["symbol"][0]

  val, found := quotes.Get(ticker)
  if found{
    log.Println(strings.ToUpper(ticker), "retrieved from cache.")
    data := val.(*options.YahooQuote)
    options.PrintFormattedQuoteData(w,*data)
    options.PrintFormattedOptionsData(w,*data)
    options.GetShortInterest(w,ticker)
    return
  }

  quote, err := options.GetStockInfo(w,ticker)
  if err != nil{
    log.Fatal(err)
  }
  quotes.Set(ticker,&quote, 10*time.Minute)
  options.PrintFormattedQuoteData(w,quote)
  options.PrintFormattedOptionsData(w,quote)
  options.GetShortInterest(w,ticker)

  fmt.Println()
}


func init(){
  //  t := cache.New(30*time.Minute, 30*time.Minute)
}


func main(){

  port := ":8080"
  r := mux.NewRouter()
  //r.HandleFunc("/", home)
  r.HandleFunc("/stocks", stocks)
  r.HandleFunc("/trends", trends)
  http.ListenAndServe(port,r)

}
