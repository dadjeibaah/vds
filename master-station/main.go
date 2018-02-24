package main

import (
	pb "github.com/vds/gen"
	"io"
	"fmt"
	"net"
	"log"
	"google.golang.org/grpc"
	"gopkg.in/mgo.v2"
	"time"
	"github.com/vds/master-station/public-api"
	"github.com/vds/master-station/aggregator"
	"gopkg.in/mgo.v2/bson"
)

type masterServiceServer struct {
	session *mgo.Session
	trains  *mgo.Collection
}

func (m *masterServiceServer) ReceiveTrainPing(stream pb.MasterService_ReceiveTrainPingServer) error {
	for {
		ping, err := stream.Recv()
		if err == io.EOF {
			stream.SendAndClose(&pb.Empty{})
			log.Println("Lost connection to train")
		}
		if err != nil {
			log.Printf("stream closed %+v", err)
			return err
		}
		log.Printf("Pinged! %+v", ping)
		go func() {
			ping.Timestamp = time.Now().UTC().Unix()
			err = m.trains.Insert(ping)
			if err != nil {
				log.Printf("DB operation failed %+v", err)
			}
			if ping.Status == "DISEMBARK"{
				info, err := m.trains.RemoveAll(bson.M{
					"trainid": ping.TrainId,
					"timestamp":
					bson.M{
						"$lte": ping.Timestamp,
					}})
				if err != nil {
					log.Printf("DB operation failed %+v", err)
				}
				log.Printf("DB Info: %+v", info)
			}
		}()
	}
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8881))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	session, err := mgo.Dial("localhost:27017")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := &masterServiceServer{
		session: session,
		trains:  session.DB("test").C("trains"),
	}
	go func() {
		public_api.ServeHTTP()
	}()
	go func(){
		aggregator.Aggregate(session.Copy())
	}()
	grpcServer := grpc.NewServer()
	pb.RegisterMasterServiceServer(grpcServer, server)
	grpcServer.Serve(lis)
}
