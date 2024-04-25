package main

import (
	"fmt"
	"net/http"
	"html/template"

)

func main(){
	fmt.Println("hello");
	
	http.HandleFunc("/", getHome)

	http.ListenAndServe(":3333", nil)
}

func getHome(w http.ResponseWriter, r *http.Request){
var tmplFile = "templates/layout.html"
tmpl, err:= template.ParseFiles(tmplFile)

if err!=nil{
	panic(err)
}

err = tmpl.Execute(w, nil)
if err!=nil{
	panic(err)
}

}

