package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nautilus/gateway"
	"github.com/nautilus/graphql"
)

const defaultPort = "8000"

//start server for gateway and introspect graphql schemas
//To run instrospection using localhost, use the -u flag at runtime
// and pass "http://localhost:9000/query" in as the parameter.

func main() {
	//instropect the apis
	schemas, err := graphql.IntrospectRemoteSchemas(
		"http://localhost:8080/query",
		"http://localhost:8081/query",
	)
	fmt.Println(schemas)
	if err != nil {
		log.Println("An Error occurred when instropecting the remote schema:", err)
	}

	//create the gateway instance
	gw, err := gateway.New(schemas)
	if err != nil {
		log.Println("Error occured creating gateway instance:", err)
	}
	http.HandleFunc("/graphql", gw.PlaygroundHandler)

	// start the server
	fmt.Println("Starting server")
	err = http.ListenAndServe(":3001", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
