package main

import (
	"flag"
	"fmt"
	"log"
	"reflect"

	"github.com/joho/godotenv"
	"github.com/rizefinance/rize-go-sdk"
	"github.com/rizefinance/rize-go-sdk/examples"
	"github.com/rizefinance/rize-go-sdk/internal"
)

func init() {
	// Load local env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	v := rize.Version()
	log.Printf("Loading Rize SDK version: %s\n", v)
}

func main() {
	var (
		showHelp = flag.Bool("h", false, "Show help menu")
		service  = flag.String("s", "CustomerService", "Service Name")
		method   = flag.String("m", "List", "Method Name")
		e        examples.Example
	)
	flag.Parse()

	help := "--- Rize Go SDK Example Runner --- \n" +
		"main.go [-s ServiceName] [-m MethodName] \n" +
		"ServiceName: SDK service name i.e. CustomerService, DebitCardService, ProductService \n" +
		"MethodName: Service method to execute i.e. List, Get, Create \n" +
		"Example: \n" +
		"main.go -s CustomerService -m List"

	if *showHelp {
		log.Println(help)
		return
	} else if service == nil {
		log.Println("No service provided. Using CustomerService")
	} else if method == nil {
		log.Println("No method provided. Using List()")
	}

	// Create new Rize client for examples
	config := rize.Config{
		ProgramUID:  internal.CheckEnvVariable("program_uid"),
		HMACKey:     internal.CheckEnvVariable("hmac_key"),
		Environment: internal.CheckEnvVariable("environment"),
		Debug:       true,
	}
	rc, err := rize.NewClient(&config)
	if err != nil {
		log.Fatal("Error building RizeClient\n", err)
	}

	// Dynamically call example method based on input flags
	methodName := fmt.Sprintf("Example%s_%s", *service, *method)
	v := reflect.ValueOf(e).MethodByName(methodName)
	if !v.IsValid() {
		log.Fatalf("Method %s does not exist", methodName)
	}

	log.Println("Calling", methodName)
	v.Call([]reflect.Value{reflect.ValueOf(rc)})
}
