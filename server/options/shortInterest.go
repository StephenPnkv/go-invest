package options

import (
  "fmt"
  "log"
  "net/http"
  "net/url"
  "strings"
  "io/ioutil"
  "encoding/json"
  "github.com/joho/godotenv"
  "os"
  "github.com/fxtlabs/date"
  "github.com/patrickmn/go-cache"
  "time"
  "math"
)

type Exchanges struct{
  NASDAQ *ShortInterestData `json:"nsdq"`
  NYSE *ShortInterestData `json:"nyse"`
  PercentDifference []float64 `json:"pDifference"`
}

type ShortInterestData struct {
	DatasetData struct {
		Limit       interface{}     `json:"limit"`
		Transform   interface{}     `json:"transform"`
		ColumnIndex interface{}     `json:"column_index"`
		ColumnNames []string        `json:"column_names"`
		StartDate   string          `json:"start_date"`
		EndDate     string          `json:"end_date"`
		Frequency   string          `json:"frequency"`
		Data        [][]interface{} `json:"data"`
		Collapse    interface{}     `json:"collapse"`
		Order       interface{}     `json:"order"`
	} `json:"dataset_data"`
}

var(
  cachedData = cache.New(60*time.Minute, 60*time.Minute)
)

func GetShortInterest(w http.ResponseWriter, r *http.Request){

  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  u, err := url.Parse(r.URL.RequestURI())
  if err != nil{
    log.Fatal(err)
  }

  m,_ := url.ParseQuery(u.RawQuery)
  ticker := strings.ToUpper(m["symbol"][0])

  val, found := cachedData.Get(ticker)
  if found{
    log.Println(ticker, "retrieved from cache.")
    data := val.(*OptionsData)
    json.NewEncoder(w).Encode(data)
    return
  }

   var yx,sq ShortInterestData
   var x Exchanges

  var xchange = func (data *ShortInterestData, xchangeName, startDate, endDate string){

    reqURL := fmt.Sprintf("https://data.nasdaq.com/api/v3/datasets/FINRA/%s_%s/data.json?start_date=%s&end_date=%s&api_key=%s",
      strings.ToUpper(xchangeName),
      strings.ToUpper(ticker),
      startDate,
      endDate,
      os.Getenv("NASDAQ_KEY"))
    log.Println(reqURL)
    res, err := http.Get(reqURL)
    if err != nil{
      log.Fatal(err)
    }
    defer res.Body.Close()

    if res.StatusCode == http.StatusOK{
      body, err := ioutil.ReadAll(res.Body)
      if err != nil{
        log.Fatal(err)
      }

      err = json.Unmarshal(body,&data)
      if err != nil{
        log.Fatal(err)
        return
      }
    }
  }

  nsdq := "FNSQ" //FINRA NASDAQ
  nyse := "FNYX" //FINRA NYSE

  d := date.Today()
  end := fmt.Sprintf("%d-%d-%d", d.Year(), d.Month(), d.Day())
  start := fmt.Sprintf("%d-%d-%d", d.Year(), d.Month(), d.Day()-3)

  xchange(&yx,nyse, start, end)
  xchange(&sq,nsdq, start, end)

  x.NASDAQ = &sq
  x.NYSE = &yx
  //x.PercentDifference[0] = x.NASDAQ
  json.NewEncoder(w).Encode(x)
//  log.Println(yx)
//  json.NewEncoder(w).Encode(x)
//  fmt.Fprint(w, "\nNYSE")
//  printFormattedSIData(w,&yx)

//  printFormattedSIData(w,&sq)

  //calculateTotalVolume(w,&sq,&yx)
}

func percentDifference(first,second int64)float64{
  dV := float64(second - first)
  return float64(100*(dV/math.Abs(float64(first))))
}


func calculateTotalVolume(w http.ResponseWriter,fnsq,fnyx *ShortInterestData){

  vals := make([]float64,len(fnsq.DatasetData.Data[0]))
  for i := 1; i < len(fnsq.DatasetData.Data[0]); i++{
    sq, ok := fnsq.DatasetData.Data[0][i].(float64)
    if !ok{
      log.Fatal(ok)
    }
    yx, ok := fnsq.DatasetData.Data[0][i].(float64)
    if !ok{
      log.Fatal(ok)
    }
    vals[i] = sq + yx
  }

  fmt.Fprintf(w, "\n\nTotal short volume: %.1f\t\tTotal exempt: %.1f\t\tTotal volume: %.1f", vals[1], vals[2], vals[3])
}
