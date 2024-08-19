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
			"notificationConfig": "test",
			"objectGeneration":   "16603467",
			"objectId":           "itm/v",
			"payloadFormat":      "JSON_API_V1",
		},
		Data:        "eyJraW5kIjogInN0b3JhZ2Ujb2JqZWN0IiwgImlkIjogInRlc3QuY3N2LzE2NjAzNDY1NTY4Mjc1NjciLCAic2VsZkxpbmsiOiAiaHR0cHM6Ly93d3cuZ29vZ2xlYXBpcy5jb20vc3RvcmFnZS92MS9iL3Rlc3QuY3N2IiwgIm5hbWUiOiAidGVzdC5jc3YiLCAiYnVja2V0IjogInRlc3QiLCAiZ2VuZXJhdGlvbiI6ICIxNjYwMzQ2NTU2ODI3NTY3IiwgIm1ldGFnZW5lcmF0aW9uIjogIjEiLCAiY29udGVudFR5cGUiOiAidGV4dC9jc3YiLCAidGltZUNyZWF0ZWQiOiAiMjAyMi0wOC0xMlQyMzoyMjozNi45MDFaIiwgInVwZGF0ZWQiOiAiMjAyMi0wOC0xMlQyMzoyMjozNi45MDFaIiwgInN0b3JhZ2VDbGFzcyI6ICJNVUxUSV9SRUdJT05BTCIsICJ0aW1lU3RvcmFnZUNsYXNzVXBkYXRlZCI6ICIyMDIyLTA4LTEyVDIzOjIyOjM2LjkwMVoiLCAic2l6ZSI6ICIyMjAiLCAibWQ1SGFzaCI6ICJ1bUVkYW5NdERSbEpVZkdBWGNYemp3PT0iLCAibWVkaWFMaW5rIjogImh0dHBzOi8vd3d3Lmdvb2dsZWFwaXMuY29tL2Rvd25sb2FkL3N0b3JhZ2UvdjEvYi90ZXN0LmNzdj9nZW5lcmF0aW9uPTE2NjAzNDY1NTY4Mjc1NjcmYWx0PW1lZGlhIiwgIm1ldGFkYXRhIjogeyJnb29nLXJlc2VydmVkLWZpbGUtbXRpbWUiOiAiMTY1OTg3NTQ2MiJ9LCAiY3JjMzJjIjogIlBHQk5nUT09IiwgImV0YWciOiAiQ0srUDVmVzR3dmtDRUFFPSJ9",
		MessageId:   "5333919906745759",
		PublishTime: "2022-08-12T23:22:36.971Z",
	}
	payload := Payload{
		Message:      message,
		Subscription: "projects/test/subscriptions/test",
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
