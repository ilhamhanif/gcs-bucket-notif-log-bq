package main

import (
	"log"

	// Blank-import the function package so the init() runs
	_ "function/gcs-bucket-notif-bq-log"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func main() {
  // Use PORT environment variable, or default to 8080.
  port := "8080"
  err := funcframework.Start(port)
  if err != nil {
    log.Fatalf("funcframework.Start: %v\n", err)
  }
}