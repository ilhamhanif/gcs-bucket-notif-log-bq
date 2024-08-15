package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Payload struct {
	Message      Message
	Subscription string
}

type Message struct {
	Attributes  map[string]string
	Data        string
	MessageId   string
	PublishTime string
}

func main() {

	/*
		A function to publish dummy data to local environment endpoint
	*/

	// Setup local endpoint and published message
	url := "http://localhost:8080/GCSBucketNotifBQLog"
	message := Message{
		Attributes: map[string]string{
			"bucketId":           "test",
			"eventTime":          "2022-08-12T23:22:36.901891Z",
			"eventType":          "OBJECT_FINALIZE",
			"notificationConfig": "projects/_/buckets/ldg_iap_mars_onl_dev/notificationConfigs/5",
			"objectGeneration":   "1660346556827567",
			"objectId":           "itm/20220813062019/520052095209003220220805_itm_20220806110121.csv",
			"payloadFormat":      "JSON_API_V1",
		},
		Data:        "ewogICJraW5kIjogInN0b3JhZ2Ujb2JqZWN0IiwKICAiaWQiOiAibGRnX2lhcF9tYXJzX29ubF9kZXYvaXRtLzIwMjIwODEzMDYyMDE5LzUyMDA1MjA5NTIwOTAwMzIyMDIyMDgwNV9pdG1fMjAyMjA4MDYxMTAxMjEuY3N2LzE2NjAzNDY1NTY4Mjc1NjciLAogICJzZWxmTGluayI6ICJodHRwczovL3d3dy5nb29nbGVhcGlzLmNvbS9zdG9yYWdlL3YxL2IvbGRnX2lhcF9tYXJzX29ubF9kZXYvby9pdG0lMkYyMDIyMDgxMzA2MjAxOSUyRjUyMDA1MjA5NTIwOTAwMzIyMDIyMDgwNV9pdG1fMjAyMjA4MDYxMTAxMjEuY3N2IiwKICAibmFtZSI6ICJpdG0vMjAyMjA4MTMwNjIwMTkvNTIwMDUyMDk1MjA5MDAzMjIwMjIwODA1X2l0bV8yMDIyMDgwNjExMDEyMS5jc3YiLAogICJidWNrZXQiOiAibGRnX2lhcF9tYXJzX29ubF9kZXYiLAogICJnZW5lcmF0aW9uIjogIjE2NjAzNDY1NTY4Mjc1NjciLAogICJtZXRhZ2VuZXJhdGlvbiI6ICIxIiwKICAiY29udGVudFR5cGUiOiAidGV4dC9jc3YiLAogICJ0aW1lQ3JlYXRlZCI6ICIyMDIyLTA4LTEyVDIzOjIyOjM2LjkwMVoiLAogICJ1cGRhdGVkIjogIjIwMjItMDgtMTJUMjM6MjI6MzYuOTAxWiIsCiAgInN0b3JhZ2VDbGFzcyI6ICJNVUxUSV9SRUdJT05BTCIsCiAgInRpbWVTdG9yYWdlQ2xhc3NVcGRhdGVkIjogIjIwMjItMDgtMTJUMjM6MjI6MzYuOTAxWiIsCiAgInNpemUiOiAiMjIwIiwKICAibWQ1SGFzaCI6ICJ1bUVkYW5NdERSbEpVZkdBWGNYemp3PT0iLAogICJtZWRpYUxpbmsiOiAiaHR0cHM6Ly93d3cuZ29vZ2xlYXBpcy5jb20vZG93bmxvYWQvc3RvcmFnZS92MS9iL2xkZ19pYXBfbWFyc19vbmxfZGV2L28vaXRtJTJGMjAyMjA4MTMwNjIwMTklMkY1MjAwNTIwOTUyMDkwMDMyMjAyMjA4MDVfaXRtXzIwMjIwODA2MTEwMTIxLmNzdj9nZW5lcmF0aW9uPTE2NjAzNDY1NTY4Mjc1NjcmYWx0PW1lZGlhIiwKICAibWV0YWRhdGEiOiB7CiAgICAiZ29vZy1yZXNlcnZlZC1maWxlLW10aW1lIjogIjE2NTk4NzU0NjIiCiAgfSwKICAiY3JjMzJjIjogIlBHQk5nUT09IiwKICAiZXRhZyI6ICJDSytQNWZXNHd2a0NFQUU9Igp9Cg==",
		MessageId:   "5333919906745759",
		PublishTime: "2022-08-12T23:22:36.971Z",
	}
	payload := Payload{
		Message:      message,
		Subscription: "projects/idf-corp-dev/subscriptions/subscription_bucketnotif_idf",
	}

	// Convert the data as JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		panic("Error")
	}

	// Sent the data to local endpoint using HTTP call method POST
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		panic("Error")
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic("Error")
	}
	defer resp.Body.Close()

	// Print response and status code
	fmt.Println(resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
