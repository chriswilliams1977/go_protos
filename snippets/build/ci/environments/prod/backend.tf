terraform {
  backend "gcs" {
    bucket = "williamscj-demos-tfstate"
    prefix = "env/prod"
  }
}