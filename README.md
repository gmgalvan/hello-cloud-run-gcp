# Hello Cloud Run GCP

### GCP cli  commands

Get the accounts to check if you are logged in:
```bash
gcloud auth list
```

Authenticate with your Google Cloud account:
```bash
gcloud auth login
```

List the available projects:
```bash
gcloud projects list
```

Set the project you want to use (replace `your-project-id` with your actual project ID):
```bash
gcloud config set project your-project-id
```


Enable Artifact Registry:
```bash
gcloud services enable artifactregistry.googleapis.com
```

### Build image and push to Artifact Registry

Build  Docker image
```bash
docker build -t gcr.io/your-project-id/hello-world-app .
```

Create a new Docker repository in Artifact Registry
```bash
gcloud artifacts repositories create hello-world-app-repo \
--repository-format=docker \
--location=us-central1
```

Configure Docker to authenticate with Artifact Registry
```bash
gcloud auth configure-docker us-central1-docker.pkg.dev
```

Tag your Docker image with the Artifact Registry repository URL:
```bash
docker tag gcr.io/your-project-id/hello-world-app us-central1-docker.pkg.dev/your-project-id/hello-world-app-repo/hello-world-app
```


Push the tagged Docker image to your Artifact Registry repository:
```bash
docker push us-central1-docker.pkg.dev/your-project-id/hello-world-app-repo/hello-world-app
```

### Deploy to cloud run
```bash
gcloud run deploy hello-world-app \
--image us-central1-docker.pkg.dev/your-project-id/hello-world-repo/hello-world-app \
--platform managed \
--region us-central1 \
--allow-unauthenticated
```


### Setup gcp for automatic deployment from GitHub Actions
To set up automatic deployment from your GitHub develop branch to Google Cloud Run, you'll need to create a GitHub Actions workflow. 

Create service account
```bash
gcloud iam service-accounts create github-actions \
    --description="Service account for GitHub Actions" \
    --display-name="GitHub Actions"
```

Grant necessary roles
```bash
gcloud projects add-iam-policy-binding your-project-id \
    --member="serviceAccount:github-actions@your-project-id.iam.gserviceaccount.com" \
    --role="roles/run.admin"

gcloud projects add-iam-policy-binding your-project-id \
    --member="serviceAccount:github-actions@your-project-id.iam.gserviceaccount.com" \
    --role="roles/storage.admin"

gcloud projects add-iam-policy-binding your-project-id \
    --member="serviceAccount:github-actions@your-project-id.iam.gserviceaccount.com" \
    --role="roles/artifactregistry.admin"
```

Grant the serviceAccountUser role
```bash
gcloud projects add-iam-policy-binding your-project-id \
    --member="serviceAccount:github-actions@your-project-id.iam.gserviceaccount.com" \
    --role="roles/iam.serviceAccountUser"

gcloud projects add-iam-policy-binding your-project-id \
    --member="serviceAccount:github-actions@your-project-id.iam.gserviceaccount.com" \
    --role="roles/run.serviceAgent"
```

Create and download service account key
```bash
gcloud iam service-accounts keys create github-actions-key.json \
    --iam-account=github-actions@your-project-id.iam.gserviceaccount.com
```


### References:
https://www.datacamp.com/tutorial/cloud-run
