package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	proto "github.com/idoqo/adder"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("add-service:5000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	addClient := proto.NewAddServiceClient(conn)

	routes := mux.NewRouter()
	routes.HandleFunc("/", indexHandler).Methods("GET")
	routes.HandleFunc("/add/{a}/{b}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json;charset=UTF-8")

		ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)
		defer cancel()

		req := &proto.AddRequest{A: 5, B: 5}
		if res, err := addClient.Compute(ctx, req); err != nil {
			msg := fmt.Sprintf("Internal server error: %v", err)
			json.NewEncoder(w).Encode(msg)
		} else {
			msg := fmt.Sprintf("%d", res.Result)
			json.NewEncoder(w).Encode(msg)
		}

	}).Methods("GET")

	fmt.Println("application is running on :8080")
	http.ListenAndServe(":8080", routes)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	json.NewEncoder(w).Encode("server is up...")
}
