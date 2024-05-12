package main

import (
	"fmt"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)
func updateRecord(collection *mongo.Collection, ctx context.Context, data map[string]interface{})(map[string]interface{}, error){
	PrintDict(data)
	fmt.Printf("Filter id: %s\n", data["id"])
	filter:= bson.M{"id": data["id"]}
	fields:= bson.M{"$set": data}

	_, err:= collection.UpdateOne(ctx , filter,  fields)

	if err!= nil{
		res:= map[string]interface{}{
			"data": "Document update failed."}

		return res, err
	}

	res:= map[string]interface{}{
		"data": "Document updated successfully.",
	}

	return res, nil
}
