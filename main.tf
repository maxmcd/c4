# terraform {
#   backend "gcs" {
#     bucket  = "c4"
#     prefix  = ""
#   }
# }

# data "terraform_remote_state" "c4" {
#   backend = "gcs"
#   config {
#     bucket  = "c4"
#     key    = "secrets.tfstate"
#   }
# }

provider "google" {
  project     = "media-9x16"
  region      = "us-central1"
}

resource "google_storage_bucket" "c4" {
  name     = "c4-data"
  location = "US"
}

resource "google_sql_database_instance" "c4" {
  name = "master-instance"
  database_version = "POSTGRES_9_6"
  region = "us-central1"
  
  settings {
    availability_type = "ZONAL"
    # Second-generation instance tiers are based on the machine
    # type. See argument reference below.
    tier = "db-f1-micro"
    location_preference {
      zone = "us-central1-a"
    }
  }
}

resource "google_sql_database" "c4" {
  instance  = "${google_sql_database_instance.c4.name}"
  name      = "c4"
}

# resource "google_sql_user" "users" {
#   instance = "${google_sql_database_instance.c4_db_instance.name}"
#   name     = "c4"
#   password = "changeme"
# }

resource "google_compute_address" "c4" {
  name = "c4"
}

resource "google_compute_instance" "c4" {
  name         = "c4"
  allow_stopping_for_update = true
  machine_type = "n1-standard-1"
  zone         = "us-central1-a"
  tags = ["c4"]

  network_interface {
    network = "default"
    access_config = {
      nat_ip = "${google_compute_address.c4.address}"
    }
  }
  boot_disk {
    initialize_params {
      image = "cos-cloud/cos-stable"
    }
  }
  metadata {
      gce-container-declaration = <<CONTAINER
spec:
  containers:
  - name: c4
    image: 'maxmcd/c4:latest'
    stdin: false
    tty: false
    env:
    - name: HOSTNAME
      value: c4.api.rivet.app
  restartPolicy: Always
CONTAINER
    }

    service_account {
      scopes = [
        "cloud-platform"
      ]
    }
}

resource "google_compute_firewall" "c4" {
  name    = "c4"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["80"]
  }

  target_tags = ["c4"]
}

