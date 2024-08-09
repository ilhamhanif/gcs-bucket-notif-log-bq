package function

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

type HttpBody struct {
	Message struct {
		Attribute struct {
			BucketId           string
			EventTime          string
			EventType          string
			NotificationConfig string
			ObjectGeneration   string
			ObjectId           string
			PayloadFormat      string
		}
		Data              string
		MessageIdPascal   string
		MessageId         string
		PublishTimePascal string
		PublishTime       string
	}
	Subscription string
}

type DecodedData struct {
	Kind                    string
	Id                      string
	SelfLink                string
	Name                    string
	Bucket                  string
	Generation              string
	MetaGeneration          string
	ContentType             string
	TimeCreated             string
	Updated                 string
	StorageClass            string
	TimeStorageClassUpdated string
	Size                    string
	Md5Hash                 string
	MediaLink               string
	Metadata                struct {
		GoogReservedFileMtime string
	}
	Crc32c string
	Etag   string
}

func init() {
	functions.HTTP("GCSBucketNotifBQLog", GCSBucketNotifBQLog)
}

func GCSBucketNotifBQLog(w http.ResponseWriter, r *http.Request) {

	var httpBody HttpBody
	var decodedData DecodedData

	if err := json.NewDecoder(r.Body).Decode(&httpBody); err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	dataEncoded := httpBody.Message.Data
	dataDecoded, _ := base64.StdEncoding.DecodeString(dataEncoded)
	if err := json.Unmarshal(dataDecoded, &decodedData); err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	rows := []*BqRow{
		{
			BucketId:                httpBody.Message.Attribute.BucketId,
			EventTime:               httpBody.Message.Attribute.EventTime,
			EventType:               httpBody.Message.Attribute.EventType,
			NotificationConfig:      httpBody.Message.Attribute.NotificationConfig,
			ObjectGeneration:        httpBody.Message.Attribute.ObjectGeneration,
			PayloadFormat:           httpBody.Message.Attribute.PayloadFormat,
			ObjectId:                httpBody.Message.Attribute.ObjectId,
			Kind:                    decodedData.Kind,
			Id:                      decodedData.Id,
			SelfLink:                decodedData.SelfLink,
			Name:                    decodedData.Name,
			MetaGeneration:          decodedData.MetaGeneration,
			ContentType:             decodedData.ContentType,
			TimeCreated:             decodedData.TimeCreated,
			Updated:                 decodedData.Updated,
			StorageClass:            decodedData.StorageClass,
			TimeStorageClassUpdated: decodedData.TimeStorageClassUpdated,
			Size:                    decodedData.Size,
			MediaLink:               decodedData.MediaLink,
		},
	}

	if err := insertBqRows(rows); err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	fmt.Fprint(w, "ok")
}
