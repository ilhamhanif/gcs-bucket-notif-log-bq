variable "project_id" {
  description = "GCP Project ID Name"
  type        = string
}

variable "region" {
  description = "GCP Region Name"
  type        = string
}

variable "gcs_bucket" {
  description = "GCP Cloud Storage Bucket Name"
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

variable "cf_auth_member" {
  description = "GCP Cloud Function Auth Member"
  type        = string
}

variable "pubsub_topic_name" {
  description = "GCP PubSub Topic Name"
  type        = string
}

variable "pubsub_subscriber_name" {
  description = "GCP PubSub Subscriber Name"
  type        = string
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
