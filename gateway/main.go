package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/nautilus/gateway"
	"github.com/nautilus/graphql"
	slog "github.com/sirupsen/logrus"
)

const defaultPort = "8000"

//start server for gateway and introspect graphql schemas
//To run instrospection using localhost, use the -u flag at runtime
// and pass "http://localhost:9000/query" in as the parameter.

func main() {
	slog.SetLevel(slog.DebugLevel)

	//instropect the apis
	schemas, err := graphql.IntrospectRemoteSchemas(
		"http://localhost:8080/query",
		"http://localhost:8081/query",
	)
	fmt.Println(schemas)
	if err != nil {
		log.Println("An Error occurred when instropecting the remote schema:", err)
	}

	// create queryer that can batch requests whenever we query a service
	factory := gateway.QueryerFactory(func(ctx *gateway.PlanningContext, url string) graphql.Queryer {
		return graphql.NewMultiOpQueryer(url, 10*time.Millisecond, 1000)
	})

	//create the gateway instance
	gw, err := gateway.New(schemas, gateway.WithMiddlewares(addHeader, logResponse), gateway.WithQueryerFactory(&factory))
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

var logResponse = gateway.ResponseMiddleware(func(ctx *gateway.ExecutionContext, response map[string]interface{}) error {
	// you can also modify the response directly if you wanted
	fmt.Println("done", response)

	return nil
})

var addHeader = gateway.RequestMiddleware(func(r *http.Request) error {
	r.Header.Add("X-Array-Batching", "true")
	return nil
})
