package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"cloud.google.com/go/storage"
)

var projectId = "sb-gcs-bucket-notif-log-bq"
var targetFileTotal = 10000
var targetGCSBucket = projectId + "-" + "bucket-with-notif"

func orchestrateFileInGCS(bucket string, object_src string, object_gcs string) error {

	/*
		A function to orchestrate (create and delete) an object in Google GCS
	*/

	// Setting up GCS client
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("[Error] storage.NewClient: %w", err)
	}
	defer client.Close()

	// Get local object
	f, err := os.Open(object_src)
	if err != nil {
		return fmt.Errorf("[Error] os.Open: %w", err)
	}
	defer f.Close()

	// Setup automatic cancel after 50 seconds
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	o := client.Bucket(bucket).Object(object_gcs)
	wc := o.NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		return fmt.Errorf("[Error] io.Copy: %w", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("[Error] Writer.Close: %w", err)
	}

	// Delete an object with storage.Delete.
	if err := o.Delete(ctx); err != nil {
		return fmt.Errorf("[Error] Object(%q).Delete: %w", object_gcs, err)
	}

	return nil
}

func main() {

	/*
		A main function
	*/

	// Setup current time as jobId
	currTime := time.Now()
	currTimeFmt := currTime.Format("20060102150405")

	// Setup current directory and temp directory
	currDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get the current working directory.")
		return
	}

	tempDir := filepath.Join(currDir, "temp")
	if _, err = os.Stat(tempDir); os.IsNotExist(err) {
		err = os.Mkdir(tempDir, 0777)
		if err != nil {
			fmt.Println("Failed to create a temp directory.")
			return
		}
	}
	defer os.RemoveAll(tempDir)

	// Create file in local temp folder
	// and orchestrate it in GCS
	for i := 0; i < targetFileTotal; i++ {

		fileName := "file" + currTimeFmt + "_" + strconv.Itoa(i) + ".txt"
		filePath := filepath.Join(tempDir, fileName)
		filePathRel, err := filepath.Rel(currDir, filePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		// Upload File to GCS
		err = orchestrateFileInGCS(targetGCSBucket, filePath, filePathRel)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("File %s is succesfully created, uploaded, and deleted from GCS.\n", filePathRel)
	}

	fmt.Printf("Job ID: %s\n", currTimeFmt)
}
