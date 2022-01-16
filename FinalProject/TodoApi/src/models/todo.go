package models

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"host.local/go/golang-todo-api/src/database"
)

var collection *mongo.Collection = database.GetCollectionPointer()

// Todo is the basic todo struct
type Todo struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task      string             `json:"task" bson:"task" validate:"required"`
	Completed bool               `json:"completed" bson:"copleted" validate:"required"`
}

// Validate runs govalidator for the struct
func (todo *Todo) Validate() error {
	validate := validator.New()
	return validate.Struct(todo)
}

// InsertOne insert a new entry
func (todo *Todo) InsertOne() error {
	_, err := database.GetCollectionPointer().InsertOne(context.Background(), todo)

	if err != nil {
		return err
	}

	return nil
}

// Update updates an element
func (todo *Todo) Update(id string) error {
	log.Println("Updating", todo)
	oid, err := primitive.ObjectIDFromHex(id)
	collection := database.GetCollectionPointer()
	doc, err := collection.UpdateOne(
		context.Background(),
		bson.M{"_id": bson.M{"$eq": oid}},
		bson.M{"$set": todo},
	)

	if err != nil {
		log.Error("Error updating", doc, err)
	}

	return err
}

// Delete removes an entry
func (todo *Todo) Delete(id string) error {
	log.Println("Deleting", todo)
	oid, err := primitive.ObjectIDFromHex(id)
	collection := database.GetCollectionPointer()
	doc, err := collection.DeleteOne(
		context.Background(),
		bson.M{"_id": bson.M{"$eq": oid}},
	)

	if err != nil {
		log.Error("Error deleting", doc, err)
	}

	return err
}

// GetAll searchs all todos in the database
func (todo *Todo) GetAll() ([]Todo, error) {
	todos := []Todo{}
	cursor, err := database.GetCollectionPointer().Find(context.TODO(), bson.D{})
	if err != nil {
		log.Error("Error with collection pointer", err)
		return todos, err
	}

	for cursor.Next(context.TODO()) {
		var elem Todo
		err := cursor.Decode(&elem)

		if err != nil {
			log.Error("Error decoding cursor", err)
		}

		todos = append(todos, elem)
	}

	cursor.Close(context.TODO())
	return todos, nil
}
