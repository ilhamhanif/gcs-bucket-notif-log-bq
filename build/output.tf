# API
output "gcp_api_services_id" {
  description = "API ID"
  value       = google_project_service.gcp_api_services
}

# BigQuery
output "bq_dataset_id" {
  description = "BQ Dataset ID"
  value       = google_bigquery_dataset.bq_dataset.id
}

output "bq_table_id" {
  description = "BQ Table ID"
  value       = google_bigquery_table.bq_table.id
}

# Cloud Function
output "gcs_bucket_cf_zip_source_code_id" {
  description = "GCS Bucket ID"
  value       = google_storage_bucket.bucket_cf_zip_source_code.id
}

output "gcs_object_cf_zip_source_code_output_name" {
  description = "GCS Object CF ZIP source code output name"
  value       = google_storage_bucket_object.upload_to_bucket_cf_zip_source_code.output_name
}

output "cf_id" {
  description = "CF ID"
  value       = google_cloudfunctions_function.cloud_function.id
}

# GCS
output "gcs_bucket_notif_id" {
  description = "GCS Bucket (with Notif) ID"
  value       = google_storage_bucket.gcs_notif_bucket.id
}

output "gcs_sa_auth_binding_etag" {
  description = "GCS SA auth binding ETag"
  value       = google_pubsub_topic_iam_binding.binding.etag
}

output "gcs_notif_bucket_notification_id" {
  description = "GCS Bucket (with Notif) Notification ID"
  value       = google_storage_notification.gcs_notif_bucket_notification.id
}

# PubSub
output "pubsub_topic" {
  description = "PubSub Topic ID"
  value       = google_pubsub_topic.pubsub_topic.id
}

output "pubsub_subscriber" {
  description = "PubSub Subscriber ID"
  value       = google_pubsub_subscription.pubsub_subscriber.id
}
