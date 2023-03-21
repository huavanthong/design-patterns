package repository

import (
	"context"

	"github.com/huavanthong/design-patterns/APP/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBCircleRepository struct {
	collection *mongo.Collection
}

func NewMongoDBCircleRepository(collection *mongo.Collection) *MongoDBCircleRepository {
	return &MongoDBCircleRepository{
		collection: collection,
	}
}

func (repo *MongoDBCircleRepository) Create(circle *entity.Circle) error {
	_, err := repo.collection.InsertOne(context.Background(), circle)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MongoDBCircleRepository) GetByID(id string) (*entity.Circle, error) {
	var circle entity.Circle
	err := repo.collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&circle)
	if err != nil {
		return nil, err
	}
	return &circle, nil
}

func (repo *MongoDBCircleRepository) Update(circle *entity.Circle) error {
	_, err := repo.collection.ReplaceOne(context.Background(), bson.M{"id": circle.ID}, circle)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MongoDBCircleRepository) Delete(id string) error {
	_, err := repo.collection.DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil {
		return err
	}
	return nil
}
