package main

import (
	"context"
	"testing"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

type bqRecords struct {
	EventType string
	Cnt       int
}

var rowRecord bqRecords
var projectId = "sb-gcs-bucket-notif-log-bq"
var jobId = "20240815075826"
var expectedBqValue = 10000

func TestBqRowRecords(t *testing.T) {

	ctx := context.Background()
	bqClient, err := bigquery.NewClient(ctx, projectId)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	defer bqClient.Close()

	q := bqClient.Query(`
	SELECT event_type eventType, count(*) cnt
	FROM sb-gcs-bucket-notif-log-bq.ops.gcs_bucketnotif_log
	WHERE 1=1
	AND REPLACE(SPLIT(name, "_")[SAFE_ORDINAL(1)], "target/file", "") = @jobId
	GROUP BY 1;
	`)
	q.Parameters = []bigquery.QueryParameter{
		{
			Name:  "jobId",
			Value: jobId,
		},
	}

	rowIterator, err := q.Read(ctx)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	for {

		err := rowIterator.Next(&rowRecord)
		if err != nil {
			if err == iterator.Done {
				break
			} else {
				t.Errorf("Error: %s", err)
			}
		}

		if rowRecord.Cnt != expectedBqValue {
			t.Errorf("Error: Row Mismatch")
		}
	}

}
