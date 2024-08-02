# Create Pub/Sub Topic, and Subscriber
resource "google_pubsub_topic" "pubsub_topic" {
  name = var.pubsub_topic_name
}

resource "google_pubsub_subscription" "pubsub_subscriber" {
  name  = var.pubsub_subscriber_name
  topic = google_pubsub_topic.pubsub_topic.id

  ack_deadline_seconds = var.pubsub_ack_deadline_timeout_seconds

  expiration_policy {
    ttl = ""
  }

  push_config {
    push_endpoint = google_cloudfunctions_function.cloud_function.https_trigger_url
  }
}
