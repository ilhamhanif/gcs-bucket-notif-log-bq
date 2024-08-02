# Create GCS Bucket
# and attach object notification
resource "google_storage_bucket" "gcs_notif_bucket" {
  name     = "${var.project_id}-${var.gcs_bucket_with_notification}"
  location = var.region
}

# Allow GCS SA to publish bucket notification message to Pub/Sub topic.
data "google_storage_project_service_account" "gcs_service_account" {}

resource "google_pubsub_topic_iam_binding" "binding" {
  topic   = google_pubsub_topic.pubsub_topic.id
  role    = "roles/pubsub.publisher"
  members = ["serviceAccount:${data.google_storage_project_service_account.gcs_service_account.email_address}"]
}

# Create GCS bucket notification with allowed SA
resource "google_storage_notification" "gcs_notif_bucket_notification" {
  bucket         = google_storage_bucket.gcs_notif_bucket.name
  payload_format = "JSON_API_V1"
  topic          = google_pubsub_topic.pubsub_topic.id
  event_types    = ["OBJECT_FINALIZE", "OBJECT_METADATA_UPDATE", "OBJECT_ARCHIVE", "OBJECT_DELETE"]
}
