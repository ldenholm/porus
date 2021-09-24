package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/ldenholm/porus/pb"
	"google.golang.org/grpc"
)

// Http Service exposed to Front-End

const (
	event     = "balance-requested"
	aggregate = "balance"
	grpcUri   = "localhost:7777"
)

func main() {
	// Create the Server
	server := &http.Server{
		Addr:    ":9090",
		Handler: initRoutes(),
	}
	log.Println("HTTP Sever listening...")
	// Running the HTTP Server
	server.ListenAndServe()
}

func initRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/orders", getBalance).Methods("POST")
	return router
}

func getBalance(rw http.ResponseWriter, req *http.Request) {
	balanceReq := &pb.BalanceRequest{}
	err := json.NewDecoder(req.Body).Decode(balanceReq)
	if err != nil {
		http.Error(rw, "Invalid User Id", 500)
		return
	}

	err = getBalanceRPC(balanceReq)
}

// Call gRPC service and create a getBalance event
func getBalanceRPC(balanceReq pb.BalanceRequest) error {
	conn, err := grpc.Dial(grpcUri, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewAccountServiceClient(conn)
	balanceJSON, _ := json.Marshal(balanceReq)

	event := &pb.Event{
		EventId:       uuid.NewUUID(),
		EventType:     event,
		AggregateId:   fmt.Sprint(balanceReq.UserId),
		AggregateType: aggregate,
		EventData:     string(balanceJSON),
		Channel:       event,
	}

	response, err := client.CreateEvent(context.Background(), event)
	if err != nil {
		return errors.Wrap(err, "Error from RPC server")
	}
	if resp.IsSuccess {
		return nil
	} else {
		return errors.Wrap(err, "Error from RPC server")
	}
}
