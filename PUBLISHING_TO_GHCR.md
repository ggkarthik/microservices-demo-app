# Publishing Microservices to GitHub Container Registry (GHCR)

## Overview
This document describes the process of publishing all 12 microservices from the Online Boutique application to GitHub Container Registry.

## Changes Made

### 1. Updated Build-Publish Workflow
**File**: `.github/workflows/build-publish.yml`

Added the missing `loadgenerator` service to ensure all 12 microservices are built and published:

- ✅ Added `loadgenerator` to the service matrix (line 37)
- ✅ Added `loadgenerator` image reference update in manifest creation (line 211)
- ✅ Added `loadgenerator` to release notes Docker images list (line 259)

### 2. Services Being Published
All 12 microservices are now included in the build-publish workflow:

1. **frontend** - HTTP server serving the website (Go)
2. **cartservice** - Shopping cart management (C#)
3. **productcatalogservice** - Product listings and search (Go)
4. **currencyservice** - Currency conversion (Node.js)
5. **paymentservice** - Payment processing (Node.js)
6. **shippingservice** - Shipping cost calculation (Go)
7. **emailservice** - Order confirmation emails (Python)
8. **checkoutservice** - Checkout orchestration (Go)
9. **recommendationservice** - Product recommendations (Python)
10. **adservice** - Contextual advertisements (Java)
11. **loadgenerator** - Simulated user traffic (Python/Locust) ⭐ *newly added*
12. **shoppingassistantservice** - AI product suggestions (Python)

## Triggering the Workflow

The build-publish workflow can be triggered in two ways:

### Method 1: Version Tag (Recommended)
```bash
# Create and push a version tag
git tag -a v1.0.0 -m "Release v1.0.0"
git push origin v1.0.0
```

**Status**: ✅ Tag `v1.0.0` has been created and pushed

### Method 2: Manual Workflow Dispatch
Navigate to: `Actions` → `Build and Publish` → `Run workflow`

Set the version parameter (e.g., `v1.0.0` or `latest`)

## Workflow Process

When triggered, the workflow performs the following steps for each microservice:

### 1. Build Stage
- Checks out the code
- Sets up Docker Buildx
- Logs into GitHub Container Registry
- Builds Docker image for each service

### 2. Security Scanning
- Runs Trivy vulnerability scanner (fails on CRITICAL/HIGH vulnerabilities)
- Runs Trivy misconfiguration scanner
- Generates Software Bill of Materials (SBOM) using Syft

### 3. Publish Stage
- Pushes Docker images to GHCR with multiple tags:
  - Semantic version (e.g., `v1.0.0`)
  - Major.minor version (e.g., `1.0`)
  - Major version (e.g., `1`)
  - Custom version or `latest`

### 4. Manifest Creation
- Updates Kubernetes manifests with new image references
- Creates a single consolidated deployment manifest
- Uploads as workflow artifact

### 5. GitHub Release (for tag pushes only)
- Creates a GitHub release
- Attaches Kubernetes manifest
- Lists all published Docker images

## Published Image Names

All images will be available at:
```
ghcr.io/ggkarthik/microservices-demo-<service-name>:<version>
```

Examples:
- `ghcr.io/ggkarthik/microservices-demo-frontend:v1.0.0`
- `ghcr.io/ggkarthik/microservices-demo-cartservice:v1.0.0`
- `ghcr.io/ggkarthik/microservices-demo-loadgenerator:v1.0.0`

## Monitoring the Workflow

### Check Workflow Status
1. Go to: https://github.com/ggkarthik/microservices-demo-app/actions
2. Look for the "Build and Publish" workflow run
3. Click on the run to see detailed logs for each service

### Check Published Images
1. Go to: https://github.com/ggkarthik?tab=packages
2. Look for packages starting with `microservices-demo-`
3. Each service will have its own package

### Check GitHub Release
1. Go to: https://github.com/ggkarthik/microservices-demo-app/releases
2. Look for Release `v1.0.0`
3. Download the Kubernetes manifest if needed

## Troubleshooting

### Common Issues

#### 1. Security Scan Failures
**Problem**: Trivy finds CRITICAL or HIGH vulnerabilities
**Solution**: 
- Review the vulnerability report in the workflow logs
- Update base images or dependencies
- Consider using `ignore-unfixed: true` if vulnerabilities don't have fixes

#### 2. Build Failures
**Problem**: Docker build fails for a service
**Solution**:
- Check Dockerfile syntax
- Ensure all dependencies are available
- Verify context path is correct

#### 3. Authentication Issues
**Problem**: Cannot push to GHCR
**Solution**:
- Ensure workflow has `packages: write` permission
- Check that `GITHUB_TOKEN` has correct scopes
- Verify repository settings allow package publishing

#### 4. Manifest Update Issues
**Problem**: sed commands fail to update manifests
**Solution**:
- Check that kubernetes-manifests directory exists
- Verify image reference patterns match
- Ensure manifests are in correct YAML format

## Next Steps

### Using the Published Images

1. **Pull an image**:
   ```bash
   docker pull ghcr.io/ggkarthik/microservices-demo-frontend:v1.0.0
   ```

2. **Deploy to Kubernetes**:
   ```bash
   # Download the release manifest
   curl -L https://github.com/ggkarthik/microservices-demo-app/releases/download/v1.0.0/kubernetes-manifests.yaml -o deployment.yaml
   
   # Apply to cluster
   kubectl apply -f deployment.yaml
   ```

3. **Run locally with Docker Compose**:
   Update `docker-compose.yml` to use GHCR images:
   ```yaml
   services:
     frontend:
       image: ghcr.io/ggkarthik/microservices-demo-frontend:v1.0.0
   ```

### Creating New Releases

For subsequent releases:
```bash
# Create a new version
git tag -a v1.1.0 -m "Release v1.1.0 - Description of changes"
git push origin v1.1.0
```

## Security Best Practices

1. **Image Scanning**: All images are scanned with Trivy before publishing
2. **SBOM Generation**: Software Bill of Materials created for each service
3. **Least Privilege**: Workflow uses minimal required permissions
4. **Signed Commits**: Consider enabling commit signature verification
5. **Dependency Updates**: Keep dependencies up-to-date using Dependabot

## Additional Resources

- [GitHub Container Registry Documentation](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry)
- [GitHub Actions Workflow Syntax](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions)
- [Kubernetes Deployment Guide](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)
- [Online Boutique Documentation](https://github.com/GoogleCloudPlatform/microservices-demo)

## Commit History

- **2025-10-17**: Added loadgenerator service to build-publish workflow (PR #15)
- **2025-10-17**: Created and pushed v1.0.0 tag to trigger initial publication

---

**Current Status**: 
- ✅ Workflow updated with all 12 services
- ✅ PR #15 merged to main branch  
- ✅ Tag v1.0.0 created and pushed
- ⏳ Workflow running - publishing images to GHCR

Monitor progress at: https://github.com/ggkarthik/microservices-demo-app/actions
