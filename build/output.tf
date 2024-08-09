# API
output "api_cloud_build_id" {
  description = "API Cloud Build ID"
  value       = google_project_service.api_cloud_build.id
}

output "api_artifact_registry_id" {
  description = "API Artifact Registry ID"
  value       = google_project_service.api_artifact_registry.id
}

output "api_cloud_functions_id" {
  description = "API Cloud Function ID"
  value       = google_project_service.api_cloud_functions.id
}

output "api_bigquery_id" {
  description = "API BigQuery ID"
  value       = google_project_service.api_bigquery.id
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

output "cf_auth_member_etag" {
  description = "CF auth member ETag"
  value       = google_cloudfunctions_function_iam_member.cloud_function_auth_member.etag
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
