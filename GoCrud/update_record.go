package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)
func updateRecord(collection *mongo.Collection, ctx context.Context, data map[string]interface{})(map[string]interface{}, error){
	filter:= bson.M{"Id": data["Id"]}
	fields:= bson.M{"$set": data}

	_, err:= collection.UpdateOne(ctx , filter,  fields)

	if err!= nil{
		return nil, err
	}

	res:= map[string]interface{}{
		"data": "Document updated successfully.",
	}

	return res, nil
}
