# Generates an archive of the source code compressed as a .zip file.
data "archive_file" "source" {
  type        = "zip"
  source_dir  = "../function"
  output_path = "${var.cf_zip_fileloc}/${var.cf_zip_filename}"
}

# Create GCS Bucket to Store Cloud Function ZIP file.
resource "google_storage_bucket" "bucket_cf_zip_source_code" {
  name     = "${var.project_id}-${var.gcs_bucket}"
  location = var.region
}

resource "google_storage_bucket_object" "upload_to_bucket_cf_zip_source_code" {
  depends_on = [
    data.archive_file.source,
    google_storage_bucket.bucket_cf_zip_source_code
  ]

  name   = var.cf_zip_filename
  bucket = google_storage_bucket.bucket_cf_zip_source_code.name
  source = "${var.cf_zip_fileloc}/${var.cf_zip_filename}"
}

# Create Cloud Function from stored ZIP file.
# And setup its authentication
resource "google_cloudfunctions_function" "cloud_function" {
  depends_on = [
    time_sleep.resource_api_cloud_function_sleep
  ]

  name    = var.cf_name
  runtime = var.cf_runtime

  available_memory_mb   = var.cf_conf_memory_mib
  source_archive_bucket = google_storage_bucket.bucket_cf_zip_source_code.name
  source_archive_object = google_storage_bucket_object.upload_to_bucket_cf_zip_source_code.name
  trigger_http          = var.cf_conf_trigger_http
  entry_point           = var.cf_conf_entry_point
}

resource "google_cloudfunctions_function_iam_member" "cloud_function_auth_member" {
  cloud_function = google_cloudfunctions_function.cloud_function.name
  region         = google_cloudfunctions_function.cloud_function.region

  role   = "roles/cloudfunctions.invoker"
  member = var.cf_auth_member
}
