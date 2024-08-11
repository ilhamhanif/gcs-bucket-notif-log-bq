# Create Pub/Sub Topic, and its Subscriber
# with Push method, authenticated with App Engine default service account.
resource "google_pubsub_topic" "pubsub_topic" {
  name = var.pubsub_topic_name
}

data "google_app_engine_default_service_account" "gcp_default_app_engine_service_account" {
  depends_on = [
    null_resource.resource_api_activation_complete
  ]
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
    oidc_token {
      service_account_email = data.google_app_engine_default_service_account.gcp_default_app_engine_service_account.email
    }
  }
}
