package public_api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"github.com/vds/gen"
)

var trainCollection *mgo.Collection
var routeCollection *mgo.Collection

type JsonError struct {
	error   string
	details string
}
type Stop struct {
	Name     string
	NextStop int32 `bson:"nextStop"`
}
type Line struct {
	LineName string `json:"lineName" bson:"lineName"`
	Stops    []Stop
}

func GetStationSchedule(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	query := bson.M{
		"nextstation": params["station"],
		"line":        params["line"],
		"direction":   params["direction"],
	}
	result := master_station.TrainPing{}
	err := trainCollection.Find(query).Sort("-timestamp").Limit(1).One(&result)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	dat, jsonErr := json.Marshal(result)
	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(dat)
}

func HandleRoutes(w http.ResponseWriter, r *http.Request) {
	result := []Line{}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	err := routeCollection.Find(bson.M{}).All(&result)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write(make([]byte, 0))
	}
	dat, jsonErr := json.Marshal(result)
	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(dat)
}



func ServeHTTP() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatalf("Unable to start API due to DB Error: %+v", err)
	}
	trainCollection = session.DB("test").C("trains")
	routeCollection = session.DB("test").C("routes")
	router := mux.NewRouter()
	router.HandleFunc("/station", GetStationSchedule).Methods("GET").
		Queries("station", "{station}", "line", "{line}", "direction", "{direction}")
	router.HandleFunc("/routes", HandleRoutes).Methods("GET")
	log.Println(http.ListenAndServe(":8888", router))
}
