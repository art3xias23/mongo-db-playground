package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/yaml.v2"
)

var config DbConfig
var client MongoDbClient

func main() {
	fmt.Println("hello")

	configContent, err := os.ReadFile("dbinfo.yaml")

	if err != nil {
		fmt.Println("Could not read dbinfo.yaml")
	}

	err = yaml.Unmarshal(configContent, &config)
	if err != nil {
		fmt.Println("Could not unmarshal dbinfo.yaml", err)
	}

	fmt.Printf("password %s, username %s", config.DbPass, config.DbUser)

	http.HandleFunc("/endpoint", requestHandler)
	http.HandleFunc("/", getHome)

	http.ListenAndServe(":3000", nil)
}

func getUpdate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got update")
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	for key, values := range r.Form {
		fmt.Printf("%s: %s\n", key, values[0])
	}
}

func getHome(w http.ResponseWriter, r *http.Request) {
	var tmplFile = "templates/layout.html"
	var rowFile = "templates/row.html"
	var data []Animal

	ctx := context.Background()
	client, err := GetMongoDbClient(config.DbPass, config.DbUser, ctx)
	if err != nil {
		fmt.Print("\nCould not get mongo db client: ", err)
		return
	}

	filter := bson.M{}
	cursor, err := client.Collection.Find(ctx, filter)
	if err != nil {
		fmt.Print("\nCould not get mongo db cursor: ", err)
		return
	}
	var animals struct {
		Items []Animal
	} = struct {
		Items []Animal
	}{}
	for cursor.Next(ctx) {
		var animal Animal

		err = cursor.Decode(&animal)
		if err != nil {
			fmt.Print("\nCould not decode animal: ", err)
			return
		}
		animals.Items = append(animals.Items, animal)
	}

	fmt.Printf("Data: ", data)

	tmpl, err := template.ParseFiles(tmplFile, rowFile)

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, animals)
	if err != nil {
		panic(err)
	}

}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entered request handler")

	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{}

	filter := bson.M{}
	ctx := context.Background()
	client, err := GetMongoDbClient(config.DbPass, config.DbUser, ctx)
	if err != nil {
		fmt.Print("\nCould not get mongo db client: ", err)
		return
	}
	cursor, err := client.Collection.Find(ctx, filter)

	if err != nil {
		fmt.Print("\nCould not get cursor: ", err)
		return
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var animal Animal

		if err := cursor.Decode(&animal); err != nil {
			fmt.Print("\nCould not decode animal: ", err)
			return
		}
		fmt.Printf("Animal: %+v\n", animal)
	}

	data := map[string]interface{}{}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	id := r.FormValue("Id")
	tpe := r.FormValue("Type")
	name := r.FormValue("Name")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Could not convert id to int: ", id)
		return
	}

	animal := Animal{
		Id:   idInt,
		Type: tpe,
		Name: name,
	}

	bsonData, err := bson.Marshal(animal)
	if err != nil {
		fmt.Println("Could not marshal animal: ", id)
		return
	}

	switch r.Method {
	case "POST":
		fmt.Println("POST")
		response, err = createRecord(client.Collection, ctx, bsonData)
		var data, ok = response["data"].(map[string]interface{})
		if ok == false {
			fmt.Println("Could not convert response: ", id)
			return

		}
		objId, ok := data["inserted"].(primitive.ObjectID)
		if ok == false {
			fmt.Println("Could not convert data: ", id)
			return

		}
		fmt.Println("objId: ", objId)
		filter := bson.M{"_id": objId}
		var animal Animal
		err := client.Collection.FindOne(ctx, filter).Decode(&animal)
		if err != nil {
			fmt.Println("Could not get new animal: ")
			return

		}

		fmt.Println("Animal: ", animal)
		tmpl, err := template.ParseFiles("templates/row.html")

		if err != nil {
			panic(err)
		}

		err = tmpl.Execute(w, animal)
		if err != nil {
			panic(err)
		}

	case "PUT":
		fmt.Println("PUT")
		response, err = updateRecord(client.Collection, ctx, data)
	case "GET":
		fmt.Println("GET")
		response, err = getRecords(client.Collection, ctx)
	case "DELETE":
		fmt.Println("DELETE")
		response, err = deleteRecord(client.Collection, ctx, data)
	}

	if err != nil {

		fmt.Println("Error in obtaining response")
		response = map[string]interface{}{"error": err.Error()}
	}

}
