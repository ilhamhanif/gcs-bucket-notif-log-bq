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

// Function: To upload local object to GCS
func orchestrateFileInGCS(bucket string, object_src string, object_gcs string) error {

	// Setting up GCS client
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("[Error] storage.NewClient: %w", err)
	}
	defer client.Close()

	// Get by opening local object
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

	currTime := time.Now()
	currTimeFmt := currTime.Format("20060102150405")
	curr_dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get the current working directory.")
		return
	}

	target_dir := filepath.Join(curr_dir, "target")
	if _, err = os.Stat(target_dir); os.IsNotExist(err) {
		err = os.Mkdir(target_dir, 0777)
		if err != nil {
			fmt.Println("Failed to create a target directory.")
			return
		}
	}
	defer os.RemoveAll(target_dir)

	for i := 0; i < targetFileTotal; i++ {

		fileName := "file" + currTimeFmt + "_" + strconv.Itoa(i) + ".txt"
		filePath := filepath.Join(target_dir, fileName)
		filePathRel, err := filepath.Rel(curr_dir, filePath)
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

}
