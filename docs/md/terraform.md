<!-- BEGIN_TF_DOCS -->
## Requirements

No requirements.

## Providers

| Name | Version |
|------|---------|
| <a name="provider_archive"></a> [archive](#provider\_archive) | 2.5.0 |
| <a name="provider_google"></a> [google](#provider\_google) | 5.40.0 |
| <a name="provider_null"></a> [null](#provider\_null) | 3.2.2 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [google_bigquery_dataset.bq_dataset](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/bigquery_dataset) | resource |
| [google_bigquery_table.bq_table](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/bigquery_table) | resource |
| [google_cloudfunctions_function.cloud_function](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/cloudfunctions_function) | resource |
| [google_project_service.gcp_api_services](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/project_service) | resource |
| [google_pubsub_subscription.pubsub_subscriber](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/pubsub_subscription) | resource |
| [google_pubsub_topic.pubsub_topic](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/pubsub_topic) | resource |
| [google_pubsub_topic_iam_binding.binding](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/pubsub_topic_iam_binding) | resource |
| [google_storage_bucket.bucket_cf_zip_source_code](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/storage_bucket) | resource |
| [google_storage_bucket.gcs_notif_bucket](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/storage_bucket) | resource |
| [google_storage_bucket_object.upload_to_bucket_cf_zip_source_code](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/storage_bucket_object) | resource |
| [google_storage_notification.gcs_notif_bucket_notification](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/storage_notification) | resource |
| [null_resource.resource_api_activation_complete](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) | resource |
| [archive_file.source](https://registry.terraform.io/providers/hashicorp/archive/latest/docs/data-sources/file) | data source |
| [google_app_engine_default_service_account.gcp_default_app_engine_service_account](https://registry.terraform.io/providers/hashicorp/google/latest/docs/data-sources/app_engine_default_service_account) | data source |
| [google_storage_project_service_account.gcs_service_account](https://registry.terraform.io/providers/hashicorp/google/latest/docs/data-sources/storage_project_service_account) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_api_services"></a> [api\_services](#input\_api\_services) | GCP API Services List | `list(string)` | n/a | yes |
| <a name="input_bq_dataset_name"></a> [bq\_dataset\_name](#input\_bq\_dataset\_name) | GCP BigQuery Dataset Name | `string` | n/a | yes |
| <a name="input_bq_table_conf_deletion_protection"></a> [bq\_table\_conf\_deletion\_protection](#input\_bq\_table\_conf\_deletion\_protection) | GCP BigQuery Config: Delete Protection | `bool` | n/a | yes |
| <a name="input_bq_table_name"></a> [bq\_table\_name](#input\_bq\_table\_name) | GCP BigQuery Table Name | `string` | n/a | yes |
| <a name="input_cf_conf_entry_point"></a> [cf\_conf\_entry\_point](#input\_cf\_conf\_entry\_point) | GCP Cloud Function Config: Function Entry Point | `string` | n/a | yes |
| <a name="input_cf_conf_memory_mib"></a> [cf\_conf\_memory\_mib](#input\_cf\_conf\_memory\_mib) | GCP Cloud Function Config: Memory (in MiB) | `number` | n/a | yes |
| <a name="input_cf_conf_trigger_http"></a> [cf\_conf\_trigger\_http](#input\_cf\_conf\_trigger\_http) | GCP Cloud Function Config: Trigger HTTP | `bool` | n/a | yes |
| <a name="input_cf_gcs_bucket"></a> [cf\_gcs\_bucket](#input\_cf\_gcs\_bucket) | GCP Cloud Function Cloud Storage Bucket Name | `string` | n/a | yes |
| <a name="input_cf_name"></a> [cf\_name](#input\_cf\_name) | GCP Cloud Function Name | `string` | n/a | yes |
| <a name="input_cf_runtime"></a> [cf\_runtime](#input\_cf\_runtime) | GCP Cloud Function Runtime | `string` | n/a | yes |
| <a name="input_cf_zip_fileloc"></a> [cf\_zip\_fileloc](#input\_cf\_zip\_fileloc) | GCP Cloud Function ZIP File Location | `string` | n/a | yes |
| <a name="input_cf_zip_filename"></a> [cf\_zip\_filename](#input\_cf\_zip\_filename) | GCP Cloud Function ZIP Filename | `string` | n/a | yes |
| <a name="input_gcs_bucket_with_notification"></a> [gcs\_bucket\_with\_notification](#input\_gcs\_bucket\_with\_notification) | GCP Cloud Storage Bucket Name (contains bucket notif) | `string` | n/a | yes |
| <a name="input_project_id"></a> [project\_id](#input\_project\_id) | GCP Project ID Name | `string` | n/a | yes |
| <a name="input_pubsub_ack_deadline_timeout_seconds"></a> [pubsub\_ack\_deadline\_timeout\_seconds](#input\_pubsub\_ack\_deadline\_timeout\_seconds) | GCP PubSub ACK Deadline Timeout (seconds) | `number` | n/a | yes |
| <a name="input_pubsub_subscriber_name"></a> [pubsub\_subscriber\_name](#input\_pubsub\_subscriber\_name) | GCP PubSub Subscriber Name | `string` | n/a | yes |
| <a name="input_pubsub_topic_name"></a> [pubsub\_topic\_name](#input\_pubsub\_topic\_name) | GCP PubSub Topic Name | `string` | n/a | yes |
| <a name="input_region"></a> [region](#input\_region) | GCP Region Name | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_bq_dataset_id"></a> [bq\_dataset\_id](#output\_bq\_dataset\_id) | BQ Dataset ID |
| <a name="output_bq_table_id"></a> [bq\_table\_id](#output\_bq\_table\_id) | BQ Table ID |
| <a name="output_cf_id"></a> [cf\_id](#output\_cf\_id) | CF ID |
| <a name="output_gcp_api_services_id"></a> [gcp\_api\_services\_id](#output\_gcp\_api\_services\_id) | API ID |
| <a name="output_gcs_bucket_cf_zip_source_code_id"></a> [gcs\_bucket\_cf\_zip\_source\_code\_id](#output\_gcs\_bucket\_cf\_zip\_source\_code\_id) | GCS Bucket ID |
| <a name="output_gcs_bucket_notif_id"></a> [gcs\_bucket\_notif\_id](#output\_gcs\_bucket\_notif\_id) | GCS Bucket (with Notif) ID |
| <a name="output_gcs_notif_bucket_notification_id"></a> [gcs\_notif\_bucket\_notification\_id](#output\_gcs\_notif\_bucket\_notification\_id) | GCS Bucket (with Notif) Notification ID |
| <a name="output_gcs_object_cf_zip_source_code_output_name"></a> [gcs\_object\_cf\_zip\_source\_code\_output\_name](#output\_gcs\_object\_cf\_zip\_source\_code\_output\_name) | GCS Object CF ZIP source code output name |
| <a name="output_gcs_sa_auth_binding_etag"></a> [gcs\_sa\_auth\_binding\_etag](#output\_gcs\_sa\_auth\_binding\_etag) | GCS SA auth binding ETag |
| <a name="output_pubsub_subscriber"></a> [pubsub\_subscriber](#output\_pubsub\_subscriber) | PubSub Subscriber ID |
| <a name="output_pubsub_topic"></a> [pubsub\_topic](#output\_pubsub\_topic) | PubSub Topic ID |
<!-- END_TF_DOCS -->