package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

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

	http.HandleFunc("/", getHome)
	http.HandleFunc("/update", getUpdate)

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

	// Print out all the form fields and their values.
	for key, values := range r.Form {
		// `values` is a slice of strings (because form fields can have multiple values)
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
	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{}

	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+dbUser+":"+dbPass+"@172.28.224.1:27017"))

	if err != nil {
		fmt.Println("Error in mongodb client")
		fmt.Println(err.Error())
	}

	collection := client.Database(dbName).Collection("animals")

	data := map[string]interface{}{}

	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		fmt.Println("Error in getting animal collection")
		fmt.Println(err.Error())
	}

	switch r.Method {
	case "POST":
		response, err = createRecord(collection, ctx, data)
	case "PUT":
		response, err = updateRecord(collection, ctx, data)
	case "GET":
		response, err = getRecords(collection, ctx)
	case "DELETE":
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
