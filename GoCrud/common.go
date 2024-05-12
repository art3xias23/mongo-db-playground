package main

import ("fmt")

func PrintDict(data map[string]interface{}){
	for key, value := range data{
		fmt.Printf("Key: %v Value: %v\n", key, value)	
	}
}
