package services

import (
	"api/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// AddContato insere um novo contato no MongoDB
func AddContato(client *mongo.Client, contato models.Contato) models.Contato {
	collection := client.Database("riwlp").Collection("contatos")

	contato.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(context.TODO(), contato)
	if err != nil {
		log.Fatal(err)
	}

	return contato
}
