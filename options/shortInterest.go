package options

import (
  "fmt"
  "log"
  "net/http"
  "strings"
  "io/ioutil"
  "encoding/json"

  "github.com/joho/godotenv"
  "os"
)

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

//var cache map[string] string

func GetShortInterest(w http.ResponseWriter, ticker string){

  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  var yx,sq ShortInterestData

  var xchange = func (data *ShortInterestData, xchangeName, startDate, endDate string){
    reqURL := fmt.Sprintf("https://data.nasdaq.com/api/v3/datasets/FINRA/%s_%s/data.json?start_date=%s&end_date=%s&api_key=%s",

      strings.ToUpper(xchangeName),
      strings.ToUpper(ticker),
      startDate,
      endDate,
      os.Getenv("NASDAQ_KEY"))
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
        log.Println(err)
        return
      }

    }
  }

  //https://data.nasdaq.com/api/v3/datasets/FINRA/FNSQ_GME/data.json?start_date=2021-12-01&end_date=2021-12-01&api_key=
  var (
    nsdq = "FNSQ" //FINRA NASDAQ
    nyse = "FNYX" //FINRA NYSE
    start = "2021-12-01"
    end = "2021-12-03"
  )
  xchange(&yx,nyse, start, end)
  xchange(&sq,nsdq, start, end)

  fmt.Fprint(w, "\nNYSE")
  printFormattedSIData(w,&yx)
  fmt.Fprint(w, "\nNASDAQ")
  printFormattedSIData(w,&sq)
  calculateTotalVolume(w,&sq,&yx)
}


func printFormattedSIData(w http.ResponseWriter,s *ShortInterestData){
  fmt.Fprint(w, "\tDate\t\tShort Volume\t\tShort Exempt Volume\t\tTotal Volume\n")
  fmt.Fprint(w, fmt.Sprintf("\t%s", s.DatasetData.Data[0][0]))
  fmt.Fprint(w, fmt.Sprintf("\t%.1f", s.DatasetData.Data[0][1]))
  fmt.Fprint(w, fmt.Sprintf("\t\t%.1f", s.DatasetData.Data[0][2]))
  fmt.Fprint(w, fmt.Sprintf("\t\t\t%.1f\n", s.DatasetData.Data[0][3]))
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
