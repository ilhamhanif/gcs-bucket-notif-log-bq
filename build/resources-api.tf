# Enable API Services
# Get service name from `gcloud services list --available | grep [SERVICE NAME]`
resource "google_project_service" "api_artifact_registry" {
  project            = var.project_id
  service            = "artifactregistry.googleapis.com"
  disable_on_destroy = false
}

resource "google_project_service" "api_cloud_build" {
  project            = var.project_id
  service            = "cloudbuild.googleapis.com"
  disable_on_destroy = false
}

# Give Delay
resource "time_sleep" "resource_api_cloud_function_sleep" {
  depends_on = [
    google_project_service.api_artifact_registry,
    google_project_service.api_cloud_build
  ]
  create_duration = "60s"
}
