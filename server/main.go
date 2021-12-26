package main

import (
  "go-invest/server/options"
  "go-invest/server/charts"
  //"context"
  "go-invest/server/quote"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
  "log"
//  "go.mongodb.org/mongo-driver/mongo"
//  mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
//  "os"
)


func trendHandler(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  quote.GetTrends(w,r)
}

func siHandler(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  options.GetShortInterest(w,r)
}

func chartHandler(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
  charts.GetChart(w,r)
}


func optionsHandler(w http.ResponseWriter, r *http.Request){
 w.Header().Set("Content-Type", "application/json")
 options.GetOptionsData(w,r)
}

func quoteHandler(w http.ResponseWriter, r *http.Request){
 w.Header().Set("Content-Type", "application/json")
 quote.GetQuote(w,r)
}

//var collection *mongo.Collection
//var ctx = context.TODO()

func logError(err error){
  if err != nil{
    log.Fatal(err)
  }
}

func init(){
//  clientOptions := mongoOptions.Client().ApplyURI(`mongodb://localhost:27017`)
//  client, err := mongo.Connect(ctx, clientOptions)
//  logError(err)

//  err = client.Ping(ctx, nil)
//  logError(err)

//  collection = client.Database("precision").Collection("users")

}


func main(){

  port := ":8080"
  r := mux.NewRouter()
  r.HandleFunc("/api/options", optionsHandler).Methods("GET")
  r.HandleFunc("/api/si",siHandler).Methods("GET")
  r.HandleFunc("/api/chart",chartHandler).Methods("GET")
  r.HandleFunc("/api/quote", quoteHandler).Methods("GET")
  r.HandleFunc("/api/trends", trendHandler).Methods("GET")
  log.Fatal(http.ListenAndServe(port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)))

}
