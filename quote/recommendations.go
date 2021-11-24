package quote

import (
  "fmt"
  "log"
  "net/http"
  //"strings"
  "io/ioutil"
  "encoding/json"
//  "github.com/jwalton/gchalk"
  //"time"
  "github.com/joho/godotenv"
  "os"
)
type Recommendations struct {
	Finance struct {
		Error  interface{} `json:"error"`
		Result []struct {
			RecommendedSymbols []struct {
				Score  float64 `json:"score"`
				Symbol string  `json:"symbol"`
			} `json:"recommendedSymbols"`
			Symbol string `json:"symbol"`
		} `json:"result"`
	} `json:"finance"`
}


func GetRecommendations(ticker string){

  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  var tr Recommendations
  client := &http.Client{}
  url := "https://yfapi.net/v6/finance/recommendationsbysymbol/" + ticker

  req, err := http.NewRequest("GET",url,nil)
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
    err = json.Unmarshal(body,&tr)
    if err != nil{
      log.Println("No ticker found.")
      return
    }
    printFormattedRecommendations(tr)
  }
}

func printFormattedRecommendations(tr Recommendations){
  fmt.Println("Recommended US stocks: \n")
  for i := 0; i < len(tr.Finance.Result[0].RecommendedSymbols); i++{
    fmt.Printf("\t%s %.2f\n", tr.Finance.Result[0].RecommendedSymbols[i].Symbol, tr.Finance.Result[0].RecommendedSymbols[i].Score)
  }
}
