package options

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
type Trends struct {
	Finance struct {
		Error  interface{} `json:"error"`
		Result []struct {
			Count        int   `json:"count"`
			JobTimestamp int64 `json:"jobTimestamp"`
			Quotes       []struct {
				Symbol string `json:"symbol"`
			} `json:"quotes"`
			StartInterval int64 `json:"startInterval"`
		} `json:"result"`
	} `json:"finance"`
}

var cache map[string] string

func GetTrends(w http.ResponseWriter, r *http.Request){

  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  var tr Trends
  client := &http.Client{}

  req, err := http.NewRequest("GET","https://yfapi.net/v1/finance/trending/US",nil)
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
    updateCache(tr)
    fmt.Fprint(w, "\nTrending US stocks: \n ")
    for i := range cache{

      fmt.Fprint(w,fmt.Sprintf("\t\t%s\n",cache[i]))
    }
  }
}

func updateCache(tr Trends){
  if cache == nil{
    cache = make(map[string]string)
  }
  for i := 0; i < len(tr.Finance.Result[0].Quotes); i++{
    _, exists := cache[tr.Finance.Result[0].Quotes[i].Symbol]
    if !exists{
      log.Println("Adding ", tr.Finance.Result[0].Quotes[i].Symbol)
      cache[tr.Finance.Result[0].Quotes[i].Symbol] = tr.Finance.Result[0].Quotes[i].Symbol
    }
  }
}

func printFormattedTrends(tr Trends){
  fmt.Println("\nTrending US stocks: ")
  for i := range cache{
    fmt.Println("\t",cache[i])
  }
}
