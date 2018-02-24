package main

import (
	"log"
	"google.golang.org/grpc"
	pb "github.com/vds/gen"
	"context"
	"os"
	"time"
	"flag"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"errors"
	"math/rand"
)

var trainId string
var line string
var direction bool
var inService bool

type Stop struct {
	Name     string
	NextStop int
}
type Line struct {
	LineName string
	Stops    []Stop
}

type Data struct {
	Routes []Line
}

type Train struct {
	status             string
	trainId            string
	line               Line
	direction          bool
	stopNum            int
	holdCount          int
	speed              float64
	seconds            float64
	distanceToNextStop int
	minutesToNextStop  int32
}

func (t *Train) printRoute() string {
	return fmt.Sprintf("Line %s. Heading %s. %s => %s station",
		t.line.LineName,
		determineDirection(t.direction),
		t.status,
		t.line.Stops[t.stopNum].Name)
}
func (t *Train) ping() *pb.TrainPing {

	return &pb.TrainPing{
		TrainId:           t.trainId,
		Route:             t.printRoute(),
		Status:            t.status,
		Speed:             t.speed,
		NextStop:          float64(t.distanceToNextStop),
		MinutesToNextStop: t.minutesToNextStop,
		Line:              t.line.LineName,
		Direction:         determineDirection(t.direction),
		NextStation:       t.line.Stops[t.stopNum].Name,
		Timestamp:         0,
	}
}

func (t *Train) accelerate() {
	max := 12.0
	a := 2.0
	distanceTraveled := (t.speed * t.seconds) + (0.5 * a * (t.seconds + 1*t.seconds + 1))
	t.distanceToNextStop -= int(distanceTraveled)
	t.speed += distanceTraveled / float64(1.0)
	if t.speed > max {
		t.speed = max
	}
	t.minutesToNextStop = int32(t.distanceToNextStop / int(t.speed))
	t.seconds += 1

}

func (t *Train) moveTrain() {
	if t.status == "NOT_IN_SERVICE" {
		t.status = "IN_TRANSIT"
		t.distanceToNextStop = t.line.Stops[t.stopNum].NextStop
	} else if t.status == "IN_TRANSIT" {
		t.accelerate()
		if t.distanceToNextStop <= 0 {
			t.status = "HOLDING"
			t.distanceToNextStop = 0
			t.speed = 0
			t.seconds = 0
			t.minutesToNextStop = 0
			t.holdCount = rand.Intn(3)
		}
	} else if t.status == "HOLDING" {
		if t.holdCount != 0 {
			t.holdCount -= 1
		} else {
			t.status = "DISEMBARK"
		}

	} else if t.status == "DISEMBARK" {
		t.status = "IN_TRANSIT"
		if t.stopNum < len(t.line.Stops)-1 && t.direction {
			t.stopNum += 1
		} else if t.stopNum > 0 && !t.direction {
			t.stopNum -= 1
		} else {
			if t.direction {
				t.stopNum = len(t.line.Stops) - 1
			} else {
				t.stopNum = 0
			}
			t.direction = !t.direction
			time.Sleep(10 * time.Second)
		}
		t.distanceToNextStop = t.line.Stops[t.stopNum].NextStop
	}
	time.Sleep(5 * time.Second)
}

func determineDirection(direction bool) string {
	if direction {
		return "INBOUND"
	} else {
		return "OUTBOUND"
	}
}

func findLine(routes []Line, lineName string) (Line, error) {
	for i, r := range routes {
		if r.LineName == lineName {
			return routes[i], nil
		}
	}
	return Line{}, errors.New("line not found")
}

func runTrain(client pb.MasterServiceClient) error {
	var data Data
	dat, err := ioutil.ReadFile("./car/routes.json")
	if err != nil {
		log.Fatalf("Error reading file %+v", err)
		return err
	}
	err = json.Unmarshal(dat, &data)
	if err != nil {
		log.Fatalf("Error parsing json %+v", err)
	}
	selectedLine, err := findLine(data.Routes, line)
	if err != nil {
		log.Fatalf("Unable to find route information")
		return err
	}

	train := &Train{
		trainId:   trainId,
		line:      selectedLine,
		status:    "NOT_IN_SERVICE",
		direction: direction,
	}

	stream, err := client.ReceiveTrainPing(context.Background())
	if err != nil {
		log.Printf("failed with %+v", err)
		return err
	}
	for {
		train.moveTrain()
		ping := train.ping()
		log.Println("Ping sent: ", ping.TrainId, ping.Route, ping.Speed, ping.NextStop, ping.MinutesToNextStop)
		err := stream.Send(ping)
		if err != nil {
			log.Fatalf("failed with %+v", err)
			stream.CloseSend()
			break
		}
	}
	return err
}

func main() {
	flag.StringVar(&trainId, "trainId", "1593", "The Id for the train")
	flag.StringVar(&line, "line", "K", "The line the train is on")
	flag.BoolVar(&direction, "direction", true, "Direction true for inbound false for outbound")
	flag.BoolVar(&inService, "inService", false, "Train is in or out of service")
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:8881", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewMasterServiceClient(conn)
	err = runTrain(client)
	if err != nil {
		os.Exit(1)
	}
}
