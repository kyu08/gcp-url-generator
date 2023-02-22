# gcp-url-generator
GCP URL generator for [fzf-bookmark-opener](https://github.com/kyu08/fzf-bookmark-opener)

# how to use
1. install [fzf-bookmark-opener](https://github.com/kyu08/fzf-bookmark-opener)
1. edit `config/variables.go` as you like(enter following items)
  - GCP Project Names
  - Cloud Run Service Names
1. execute `make generate`
1. append generated bookmarks to your `~/.config/fzf-bookmark-opener/config.yaml`

- for now, you can generate URLs of following GCP services
  - Dashboard
  - Datastore
  - App Engine
  - Artifact Registry
  - Cloud Storage
  - Spanner
  - BigQuery
  - Cloud Scheduler
  - PubSub
  - Dataflow
  - Cloud Build
  - Cloud Tasks
  - Cloud Run
  - Cloud SQL
  - Cloud Functions
  - IAM
  - Secret Manager
  - AI Platform
  - Compute Engine
  - Logging
  - Monitoring
  - Workflows
  - Firebase Realtime Database (DB)
  - Firebase Hosting
  - Firebase Remote Config

