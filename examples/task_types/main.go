package main

import (
	"fmt"
	"log"

	"github.com/gomorpheus/morpheus-go-sdk"
)

func main() {
	client := morpheus.NewClient("https://yourmorpheus.com")
	client.SetUsernameAndPassword("username", "password")
	resp, err := client.Login()
	if err != nil {
		fmt.Println("LOGIN ERROR: ", err)
	}
	fmt.Println("LOGIN RESPONSE:", resp)

	// List task types
	req := &morpheus.Request{}
	response, err := client.ListTaskTypes(req)
	if err != nil {
		log.Fatal(err)
	}
	result := response.Result.(*morpheus.ListTaskTypesResult)
	log.Println(result.TaskTypes)
}
