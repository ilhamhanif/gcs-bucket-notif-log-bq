# Create BigQuery Dataset, and Table
resource "google_bigquery_dataset" "bq_dataset" {
  dataset_id = var.bq_dataset_name
  location   = var.region
}

resource "google_bigquery_table" "bq_table" {
  dataset_id          = google_bigquery_dataset.bq_dataset.dataset_id
  table_id            = var.bq_table_name
  deletion_protection = var.bq_table_conf_deletion_protection

  time_partitioning {
    type  = "DAY"
    field = "event_date"
  }
  clustering = ["object_generation"]

  schema = <<EOF
[
  {
    "name": "bucket_id",
    "type": "STRING",
    "mode": "NULLABLE"
  },
  {
    "name": "event_date",
    "type": "DATE",
    "mode": "NULLABLE"
  },
  {
    "name": "event_time",
    "type": "DATETIME",
    "mode": "NULLABLE"
  },
  {
    "name": "event_type",
    "type": "STRING",
    "mode": "NULLABLE"
  },
  {
    "name": "notification_config",
    "type": "STRING",
    "mode": "NULLABLE"
  },
  {
    "name": "object_generation",
    "type": "STRING",
    "mode": "NULLABLE"
  },
  {
    "name": "payload_format",
    "type": "STRING",
    "mode": "NULLABLE"
  },
  {
    "name": "object_id",
    "type": "STRING",
    "mode": "NULLABLE"
  },
  {
    "name": "kind",
    "type": "STRING",
    "mode": "NULLABLE"
  },
  {
    "name": "id",
    "type": "STRING",
    "mode": "NULLABLE"
  },
  {
    "name": "self_link",
    "type": "STRING",
    "mode": "NULLABLE"
  },
  {
    "name": "name",
    "type": "STRING",
    "mode": "NULLABLE"
  },
  {
    "name": "metageneration",
    "type": "STRING",
    "mode": "NULLABLE"
  },
  {
    "name": "content_type",
    "type": "STRING",
    "mode": "NULLABLE"
  },
  {
    "name": "time_created",
    "type": "DATETIME",
    "mode": "NULLABLE"
  },
  {
    "name": "updated",
    "type": "DATETIME",
    "mode": "NULLABLE"
  },
  {
    "name": "storage_class",
    "type": "STRING",
    "mode": "NULLABLE"
  },
  {
    "name": "time_storage_class_updated",
    "type": "DATETIME",
    "mode": "NULLABLE"
  },
  {
    "name": "size",
    "type": "STRING",
    "mode": "NULLABLE"
  }
]
EOF

}
