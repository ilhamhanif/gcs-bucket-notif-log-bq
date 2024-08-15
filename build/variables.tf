variable "project_id" {
  description = "GCP Project ID Name"
  type        = string
}

variable "region" {
  description = "GCP Region Name"
  type        = string
}

variable "api_services" {
  description = "GCP API Services List"
  type        = list(string)
}

variable "cf_gcs_bucket" {
  description = "GCP Cloud Function Cloud Storage Bucket Name"
  type        = string
}

variable "cf_zip_filename" {
  description = "GCP Cloud Function ZIP Filename"
  type        = string
}

variable "cf_zip_fileloc" {
  description = "GCP Cloud Function ZIP File Location"
  type        = string
}

variable "cf_name" {
  description = "GCP Cloud Function Name"
  type        = string
}

variable "cf_runtime" {
  description = "GCP Cloud Function Runtime"
  type        = string
}

variable "cf_conf_memory_mib" {
  description = "GCP Cloud Function Config: Memory (in MiB)"
  type        = number
}

variable "cf_conf_trigger_http" {
  description = "GCP Cloud Function Config: Trigger HTTP"
  type        = bool
}

variable "cf_conf_entry_point" {
  description = "GCP Cloud Function Config: Function Entry Point"
  type        = string
}

variable "cf_conf_min_instance" {
  description = "GCP Cloud Function Config: Minimum Instance"
  type        = number
}

variable "cf_conf_max_instance" {
  description = "GCP Cloud Function Config: Maximum Instance"
  type        = number
}

variable "pubsub_topic_name" {
  description = "GCP PubSub Topic Name"
  type        = string
}

variable "pubsub_subscriber_name" {
  description = "GCP PubSub Subscriber Name"
  type        = string
}

variable "pubsub_ack_deadline_timeout_seconds" {
  description = "GCP PubSub ACK Deadline Timeout (seconds)"
  type        = number
}

variable "bq_dataset_name" {
  description = "GCP BigQuery Dataset Name"
  type        = string
}

variable "bq_table_name" {
  description = "GCP BigQuery Table Name"
  type        = string
}

variable "bq_table_conf_deletion_protection" {
  description = "GCP BigQuery Config: Delete Protection"
  type        = bool
}

variable "gcs_bucket_with_notification" {
  description = "GCP Cloud Storage Bucket Name (contains bucket notif)"
  type        = string
}
