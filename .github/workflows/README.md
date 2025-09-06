# GitHub Actions Workflows

This directory contains GitHub Actions workflow files for CI/CD pipelines.

## Main Workflows

- `ci-pipeline.yml`: Main CI pipeline that runs on every push and PR
- `build-publish.yml`: Workflow for building and publishing versioned images
- `security-scan.yml`: Weekly security scanning workflow

## Infrastructure Workflows

- `helm-chart-ci.yaml`: CI for Helm charts
- `kubevious-manifests-ci.yaml`: CI for Kubernetes manifests
- `kustomize-build-ci.yaml`: CI for Kustomize configurations
- `terraform-validate-ci.yaml`: CI for Terraform configurations
