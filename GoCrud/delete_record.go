package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func deleteRecord(collection *mongo.Collection, ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {
	fmt.Print()
	fmt.Println("Enter deleteRecord")
	PrintDict(data)
	fmt.Println("\nend Printing dict")
	fmt.Printf("Id is %v", data["id"])
	fmt.Println()
	id:=data["id"].(int32)
	fmt.Printf("Id is %d\n", id)

if collection == nil{
	fmt.Print("collection is nil\n")
}
	_, err := collection.DeleteOne(ctx, bson.M{"id": id})

	if err != nil {
		fmt.Println("\nError during deletion")
		return nil, err
	}

	res := map[string]interface{}{
		"data": "Document Deleted Successfully",
	}

	return res, nil
}
