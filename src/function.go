package function

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("GCSBucketNotifBQLog", GCSBucketNotifBQLog)
}

func GCSBucketNotifBQLog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}