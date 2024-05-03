package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbUser = "tino"
	dbPass = "password"
	dbName = "gocrud"
)

func main() {
	fmt.Println("hello")

	http.HandleFunc("/endpoint", requestHandler)
	http.HandleFunc("/", getHome)

	http.ListenAndServe(":3000", nil)
}

type Animal struct {
	Id   int
	Type string
	Name string
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
	items := []Animal{
		{Type: "Dog", Name: "Rob"},
		{Type: "Cat", Name: "Mandy"},
	}
	data := struct {
		Items []Animal
	}{
		Items: items}

	tmpl, err := template.ParseFiles(tmplFile, rowFile)

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		panic(err)
	}

}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entered request handler")

	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{}

	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+dbUser+":"+dbPass+"@172.28.224.1:27017/gocrud?authSource=gocrud"))

	if err != nil {
		fmt.Println("Error in mongodb client")
		fmt.Println(err.Error())
	}

	fmt.Println("connection established")

	

	collection := client.Database(dbName).Collection("animals")
	filter := bson.M{}

	cursor, err:= collection.Find(ctx, filter)
	if err != nil {
		fmt.Print("\nCould not get cursor: ", err)
		return
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx){
		var animal Animal

		if err:= cursor.Decode(&animal); err!=nil {
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

	animal:= Animal{
		Id: idInt,
		Type: tpe,
		Name:name,
	}

	bsonData, err:= bson.Marshal(animal)
	if err != nil {
		fmt.Println("Could not marshal animal: ", id)
		return
	}

	switch r.Method {
	case "POST":
		fmt.Println("POST")
		response, err = createRecord(collection, ctx, bsonData)
	case "PUT":
		fmt.Println("PUT")
		response, err = updateRecord(collection, ctx, data)
	case "GET":
		fmt.Println("GET")
		response, err = getRecords(collection, ctx)
	case "DELETE":
		fmt.Println("DELETE")
		response, err = deleteRecord(collection, ctx, data)
	}


	   if err != nil { 

		fmt.Println("Error in obtaining response") 
         response = map[string]interface{}{"error": err.Error(),}  
     }

	enc:= json.NewEncoder(w)

	enc.SetIndent("", "  ")

	  if err := enc.Encode(response); err != nil {
         fmt.Println(err.Error())
     }
}
