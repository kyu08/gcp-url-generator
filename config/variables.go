package config

type project struct {
	name string
	key  string
}

type service struct {
	name string
	url  string
}

var (
	// Cloud Run
	// replace to your cloud run services
	runServices = []string{
		"service-a",
		"service-b",
		"service-c",
	}

	// replace to your region
	region = "asia-northeast1"

	// GCP project
	// replace to your gcp projects
	projects = []project{
		{name: "stg", key: "stg"},
		{name: "dev", key: "dev"},
		{name: "prod", key: "prod"},
	}

	// NOTE: Implementation of the function converts gaeServices to String does not exist for now.
	// projectName and serviceName is needed
	gae = []service{
		{
			name: "App Engine (Service)",
			url:  "https://console.cloud.google.com/appengine/versions?project=%s&serviceId={service}",
		},
	}
)

// GCP services(needs only a project name)
var services = []service{
	{
		name: "Home",
		url:  "https://console.cloud.google.com/home/dashboard?project=%s",
	},
	{
		name: "Datastore",
		url:  "https://console.cloud.google.com/datastore/entities?project=%s",
	},
	{
		name: "App Engine",
		url:  "https://console.cloud.google.com/appengine/services?project=%s",
	},
	{
		name: "Artifact Registry",
		url:  "https://console.cloud.google.com/artifacts?project=%s",
	},
	{
		name: "Cloud Storage",
		url:  "https://console.cloud.google.com/storage/browser?project=%s",
	},
	{
		name: "Spanner",
		url:  "https://console.cloud.google.com/spanner/instances?project=%s",
	},
	{
		name: "BigQuery",
		url:  "https://console.cloud.google.com/bigquery?project=%s",
	},
	{
		name: "Cloud Scheduler",
		url:  "https://console.cloud.google.com/cloudscheduler?project=%s",
	},
	{
		name: "PubSub - topic",
		url:  "https://console.cloud.google.com/cloudpubsub/topic/list?project=%s",
	},
	{
		name: "PubSub - subscription",
		url:  "https://console.cloud.google.com/cloudpubsub/subscription/list?project=%s",
	},
	{
		name: "Dataflow",
		url:  "https://console.cloud.google.com/dataflow/jobs?project=%s",
	},
	{
		name: "Cloud Build",
		url:  "https://console.cloud.google.com/cloud-build/builds?project=%s",
	},
	{
		name: "Cloud Tasks",
		url:  "https://console.cloud.google.com/cloudtasks?project=%s",
	},
	{
		name: "Cloud Run",
		url:  "https://console.cloud.google.com/run?project=%s",
	},
	{
		name: "Cloud Functions",
		url:  "https://console.cloud.google.com/functions/list?referrer=search&project=%s",
	},
	{
		name: "IAM",
		url:  "https://console.cloud.google.com/iam-admin/iam?project=%s",
	},
	{
		name: "Secret Manager",
		url:  "https://console.cloud.google.com/security/secret-manager?referrer=search&project=%s",
	},
	{
		name: "AI Platform",
		url:  "https://console.cloud.google.com/ai-platform?project=%s",
	},
	{
		name: "Compute Engine",
		url:  "https://console.cloud.google.com/compute/instances?project=%s",
	},
	{
		name: "Logging",
		url:  "https://console.cloud.google.com/logs/query?project=%s",
	},
	{
		name: "Monitoring",
		url:  "https://console.cloud.google.com/monitoring?referrer=search&project=%s",
	},
	{
		name: "Workflows",
		url:  "https://console.cloud.google.com/workflows?project=%s",
	},
	{
		name: "Firebase Realtime Database (DB)",
		url:  "https://console.firebase.google.com/project/%s/database",
	},
	{
		name: "Firebase Hosting",
		url:  "https://console.firebase.google.com/project/%s/hosting",
	},
	{
		name: "Firebase Remote Config",
		url:  "https://console.firebase.google.com/project/%s/config",
	},
}
