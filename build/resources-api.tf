# Enable API Services
# Cloud Build API
resource "google_project_service" "api_cloud_build" {
  project            = var.project_id
  service            = "cloudbuild.googleapis.com"
  disable_on_destroy = false
}

# Artifact Registry API
resource "google_project_service" "api_artifact_registry" {
  project            = var.project_id
  service            = "artifactregistry.googleapis.com"
  disable_on_destroy = false
}

# Cloud Function API
resource "google_project_service" "api_cloud_functions" {
  project            = var.project_id
  service            = "cloudfunctions.googleapis.com"
  disable_on_destroy = false
}

# BigQuery API
resource "google_project_service" "api_bigquery" {
  project            = var.project_id
  service            = "bigquery.googleapis.com"
  disable_on_destroy = false
}

# Give Delay
resource "null_resource" "resource_api_activation_complete" {
  depends_on = [
    google_project_service.api_cloud_build,
    google_project_service.api_artifact_registry,
    google_project_service.api_cloud_functions,
    google_project_service.api_bigquery
  ]
}
