locals {
  project = "subtle-canto-412404"
  region = "us-east1"
}

provider "google-beta" {
  project     = local.project
  region      = local.region
}

resource "google_api_gateway_api" "api" {
  provider = google-beta
  api_id = "stox-api"
}

resource "google_api_gateway_api_config" "api_cfg" {
  provider = google-beta
  api = google_api_gateway_api.api.api_id
  api_config_id = "stox-api-config"

  openapi_documents {
    document {
      path = "spec.yaml"
      contents = filebase64("../gen/http/openapi.yaml")
    }
  }
  lifecycle {
    create_before_destroy = true
  }
}

resource "google_api_gateway_gateway" "api_gw" {
  provider = google-beta
  api_config = google_api_gateway_api_config.api_cfg.id
  gateway_id = "stox-gateway"
}

resource "google_cloud_run_v2_service" "default" {
  provider = google-beta
  name     = "stox"
  location = local.region
  ingress = "INGRESS_TRAFFIC_ALL"

  template {
    containers {
      image = "gcr.io/${local.project}/ubuntu"
    }
  }
}