package aggregator

import (
	"gopkg.in/mgo.v2"
	"time"
	"fmt"
	pb "github.com/vds/gen"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type TrainSchedule struct {
	trainId    string
	minutesOut int32
}

type StationSchedule struct {
	StationName string
	Inbound     []TrainSchedule
	Outbound    []TrainSchedule
}

type Key struct {
	Line        string `bson:"line"`
	Direction   string `bson:"direction"`
	Trainid     string `bson:"trainId"`
	Nextstation string `bson:"nextstation"`
}

type GroupResult struct {
	Id    Key            `bson:"_id"`
	Pings []pb.TrainPing `bson:"pings"`
}

//func createSchedule(trains []GroupResult, schedule *map[string]*StationSchedule) {
//	for _, train := range trains {
//		if val, ok := schedule[train.Id.Nextstation]; ok {
//			if train.Id.Direction == "INBOUND" {
//				if len(train.Pings) > 0 {
//					val.Inbound = append(val.Inbound, TrainSchedule{train.Id.Trainid, train.Pings[0].MinutesToNextStop})
//				}
//			}
//			if train.Id.Direction == "OUTBOUND"{
//				if len(train.Pings) > 0 {
//					val.Outbound = append(val.Outbound, TrainSchedule{train.Id.Trainid, train.Pings[0].MinutesToNextStop})
//				}
//			}
//		} else {
//			stationSched := &StationSchedule{StationName:train.Id.Nextstation}
//			//schedule[train.Id.Nextstation] =
//			if train.Id.Direction == "INBOUND" {
//				if len(train.Pings) > 0 {
//					stationSched.Inbound = append(stationSched.Inbound, TrainSchedule{train.Id.Trainid, train.Pings[0].MinutesToNextStop})
//				}
//			}
//			if train.Id.Direction == "OUTBOUND"{
//				if len(train.Pings) > 0 {
//					stationSched.Outbound = append(stationSched.Outbound, TrainSchedule{train.Id.Trainid, train.Pings[0].MinutesToNextStop})
//				}
//			}
//			schedule[train.Id.Nextstation] = stationSched
//		}
//	}
//	fmt.Println(schedule)
//
//}

func Aggregate(s *mgo.Session) {
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		defer s.Close()
		for t := range ticker.C {
			fmt.Println(t)
			trains := make([]GroupResult, 0)
			trainCollection := s.DB("test").C("trains")
			err := trainCollection.Pipe([]bson.M{{"$sort": bson.M{"timestamp": -1}},
				{"$group": bson.M{"_id": bson.M{"line": "$line", "direction": "$direction", "trainid": "$trainid", "nextstation": "$nextstation"}, "pings": bson.M{"$push": "$$ROOT"}}}}).
				All(&trains)
			if err != nil {
				log.Fatalln("Error: ", err)
			}
			//var schedule = make(map[string]*StationSchedule)
			//createSchedule(trains, &schedule)
			fmt.Printf("Result: %+v", trains)
		}
	}()
}
