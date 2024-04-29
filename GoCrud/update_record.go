package main

import (
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/tools/go/analysis/passes/ifaceassert"
)
func updateRecord(collection *mongo.Collection, ctx context.Context, map[string]interface{})(map[string]interface{}, error){
	filter:= bson.M("Id": data["Id"])
	fields:= bson.M{"$set": data}

	_, err:= collection.UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions)

	if err!= nil{
		return nil, err
	}

	res:= map[string]interface{}{
		"data": "Document updated successfully.",
	}

	return res, nil
}
