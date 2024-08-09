# Setup Backend Configuration
terraform {
  backend "local" {
    path = "state-gcs-bucket-notif-log-bq.tfstate"
  }
}
