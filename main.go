package main

import (
	"api/handlers"
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri_mongo = "mongodb://localhost:27017"

func main() {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri_mongo).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Passar o cliente MongoDB para os manipuladores
	handlerContext := handlers.NewHandlerContext(client)

	r := mux.NewRouter()

	r.HandleFunc("/contato", handlerContext.AddContato).Methods("POST")

	log.Println("Starting server on :3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
