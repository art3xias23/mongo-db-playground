package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

func createRecord(collection *mongo.Collection, ctx context.Context, data []byte)(map[string]interface{}, error){
	req, err:= collection.InsertOne(ctx, data)

	if err != nil {
		fmt.Println("Insert one error: ", err)
		return nil, err
	}

	insertedId := req.InsertedID

	res := map[string]interface{}{
		"data": map[string]interface{}{
			"inserted": insertedId,
		},
	}

	return res, nil
}
