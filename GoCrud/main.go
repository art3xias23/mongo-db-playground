package main

import (
	"fmt"
	"net/http"
	"html/template"

)

func main(){
	fmt.Println("hello");
	
	http.HandleFunc("/", getHome)

	http.ListenAndServe(":3000", nil)
}

type Animal struct{
	Id int
	Type string
	Name string
}

func getHome(w http.ResponseWriter, r *http.Request){
var tmplFile = "templates/layout.html"
var rowFile = "templates/row.html"
items:= []Animal{
	{Type: "Dog", Name: "Rob"},
	{Type: "Cat", Name: "Mandy"},
}
 data := struct{
	 Items []Animal
 }{
	 Items: items}

tmpl, err:= template.ParseFiles(tmplFile, rowFile)

if err!=nil{
	panic(err)
}

err = tmpl.Execute(w, data)
if err!=nil{
	panic(err)
}

}

