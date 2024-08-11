# Enable API Services
resource "google_project_service" "gcp_api_services" {
  count              = length(var.api_services)
  project            = var.project_id
  service            = var.api_services[count.index]
  disable_on_destroy = false
}

## Wait for All API successfully enabled
resource "null_resource" "resource_api_activation_complete" {
  depends_on = [
    google_project_service.gcp_api_services
  ]
}
