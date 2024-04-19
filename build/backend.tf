terraform {
  backend "gcs" {
    prefix = "cloud_function/gcs_bucket_notif_log_bq"
  }
}
