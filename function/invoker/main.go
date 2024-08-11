package main

import (
	"log"

	_ "github.com/ilhamhanif/gcs-bucket-notif-log-bq/src"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func main() {

	port := "8080"
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}

}
