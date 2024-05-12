package main

import ("fmt")

func PrintDict(data map[string]interface{}){
	fmt.Println("Start Printing dict")
	for key, value := range data{
		fmt.Printf("Key: %v Value: %v\n", key, value)	
	}
	fmt.Println("End Printing dict")
}
