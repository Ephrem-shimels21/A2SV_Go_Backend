package infrastructure

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection interface {
	FindOne(context.Context, interface{}) SingleResult
	InsertOne(context.Context, interface{}) (interface{}, error)
	DeleteOne(context.Context, interface{}) (int64, error)
	Find(context.Context, interface{}, ...*options.FindOptions) (Cursor, error)
	CountDocuments(context.Context, interface{}, ...*options.CountOptions) (int64, error)
	UpdateOne(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	// UpdateMany(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (*mongo.UpdateResult, error)
}

type SingleResult interface {
	Decode(interface{}) error
}

type Cursor interface {
	Close(context.Context) error
	Next(context.Context) bool
	Decode(interface{}) error
	All(context.Context, interface{}) error
}

type Database interface {
	Collection(name string) Collection
}

type mongoDatabase struct {
	db *mongo.Database
}

type mongoCollection struct {
	collection *mongo.Collection
}

func (mongoDb *mongoDatabase) Collection(name string) Collection {
	return &mongoCollection{
		collection: mongoDb.db.Collection(name),
	}
}

func (mongoColl *mongoCollection) FindOne(cxt context.Context, filter interface{}) SingleResult {
	SingleResult := mongoColl.collection.FindOne(cxt, filter)
	return SingleResult
}

func (mongoColl *mongoCollection) InsertOne(cxt context.Context, document interface{}) (interface{}, error) {
	id, err := mongoColl.collection.InsertOne(cxt, document)
	return id.InsertedID, err
}

func (mongoColl *mongoCollection) DeleteOne(cxt context.Context, filter interface{}) (int64, error) {
	result, err := mongoColl.collection.DeleteOne(cxt, filter)
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}

func (mongoColl *mongoCollection) Find(cxt context.Context, filter interface{}, opts ...*options.FindOptions) (Cursor, error) {
	cursor, err := mongoColl.collection.Find(cxt, filter, opts...)
	if err != nil {
		return nil, err
	}
	return cursor, nil
}

func (mongoColl *mongoCollection) CountDocuments(cxt context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	count, err := mongoColl.collection.CountDocuments(cxt, filter, opts...)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (mongoColl *mongoCollection) UpdateOne(cxt context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	result, err := mongoColl.collection.UpdateOne(cxt, filter, update, opts...)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func ConnectDb() (Database, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	db := client.Database("TaskManager")
	return &mongoDatabase{db: db}, nil
}
