package main

import (
	"log"

	_ "github.com/ilhamhanif/gcs-bucket-notif-log-bq"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func main() {

	/*
		A main function to start function-framework
	*/

	port := "8080"
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}

}
