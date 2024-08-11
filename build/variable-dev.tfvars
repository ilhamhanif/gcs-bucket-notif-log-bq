project_id = "sb-gcs-bucket-notif-log-bq"
region     = "asia-southeast2"
api_services = [
  "compute.googleapis.com",
  "cloudbuild.googleapis.com",
  "artifactregistry.googleapis.com",
  "cloudfunctions.googleapis.com",
  "bigquery.googleapis.com"
]
cf_gcs_bucket                       = "cf-zip-source-code"
cf_zip_filename                     = "gcs-bucket-notif-bq-log.zip"
cf_zip_fileloc                      = "zip"
cf_name                             = "gcs-bucket-notif-bq-log"
cf_runtime                          = "go122"
cf_conf_memory_mib                  = 128
cf_conf_trigger_http                = true
cf_conf_entry_point                 = "GCSBucketNotifBQLog"
cf_auth_member                      = "allUsers"
pubsub_topic_name                   = "gcs-bucket-notif-bq-log"
pubsub_subscriber_name              = "gcs-bucket-notif-bq-log"
pubsub_ack_deadline_timeout_seconds = 60
bq_dataset_name                     = "ops"
bq_table_name                       = "gcs_bucketnotif_log"
bq_table_conf_deletion_protection   = false
gcs_bucket_with_notification        = "bucket-with-notif"
