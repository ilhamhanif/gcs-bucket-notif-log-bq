package function

import (
	"context"
	"time"

	"cloud.google.com/go/bigquery"
)

type BqRow struct {
	BucketId                string
	EventTime               string
	EventType               string
	NotificationConfig      string
	ObjectGeneration        string
	PayloadFormat           string
	ObjectId                string
	Kind                    string
	Id                      string
	SelfLink                string
	Name                    string
	MetaGeneration          string
	ContentType             string
	TimeCreated             string
	Updated                 string
	StorageClass            string
	TimeStorageClassUpdated string
	Size                    string
	MediaLink               string
}

func (row *BqRow) Save() (map[string]bigquery.Value, string, error) {

	eventTime, _ := time.Parse("2006-01-02T15:04:05.999999Z0700", row.EventTime)
	eventTimeFormat := eventTime.Format("2006-01-02T15:04:05")
	eventDate := eventTime.Format("2006-01-02")
	timeCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z0700", row.TimeCreated)
	timeCreatedFormat := timeCreated.Format("2006-01-02T15:04:05")
	updated, _ := time.Parse("2006-01-02T15:04:05.999999Z0700", row.Updated)
	updatedFormat := updated.Format("2006-01-02T15:04:05")
	timeStorageClassUpdated, _ := time.Parse("2006-01-02T15:04:05.999999Z0700", row.TimeStorageClassUpdated)
	timeStorageClassUpdatedFormat := timeStorageClassUpdated.Format("2006-01-02T15:04:05")

	return map[string]bigquery.Value{
		"bucket_id":                  row.BucketId,
		"event_date":                 eventDate,
		"event_time":                 eventTimeFormat,
		"event_type":                 row.EventType,
		"notification_config":        row.NotificationConfig,
		"object_generation":          row.ObjectGeneration,
		"payload_format":             row.PayloadFormat,
		"object_id":                  row.ObjectId,
		"kind":                       row.Kind,
		"id":                         row.Id,
		"self_link":                  row.SelfLink,
		"name":                       row.Name,
		"metageneration":             row.MetaGeneration,
		"content_type":               row.ContentType,
		"time_created":               timeCreatedFormat,
		"updated":                    updatedFormat,
		"storage_class":              row.StorageClass,
		"time_storage_class_updated": timeStorageClassUpdatedFormat,
		"size":                       row.Size,
		"media_link":                 row.MediaLink,
	}, bigquery.NoDedupeID, nil
}

func insertBqRows(rows []*BqRow) error {

	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, projectId)
	if err != nil {
		return err
	}
	defer client.Close()

	inserter := client.Dataset(bqDatasetName).Table(bqTableName).Inserter()
	if err := inserter.Put(ctx, rows); err != nil {
		return err
	}

	return nil
}
