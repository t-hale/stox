default: help

.PHONY: apply destroy gen help init login update

DEPLOYMENT ?= qa
GITHUB_REPO ?= https://github.com/t-hale/stox
PROJECT ?= subtle-canto-412404
REGION ?= us-central1
SERVICE_ACCOUNT ?= infrastructure@subtle-canto-412404.iam.gserviceaccount.com

apply: gen gazelle	## Deploy infrastructure
	@gcloud infra-manager deployments apply projects/$(PROJECT)/locations/$(REGION)/deployments/$(DEPLOYMENT) \
      --service-account projects/$(PROJECT)/serviceAccounts/$(SERVICE_ACCOUNT) \
      --git-source-repo=$(GITHUB_REPO) \
      --git-source-directory=terraform \
      --git-source-ref=$$(git rev-parse --abbrev-ref HEAD)

destroy:	## Run terraform destroy
	@gcloud infra-manager deployments delete projects/$(PROJECT)/locations/$(REGION)/deployments/$(DEPLOYMENT)

gazelle:	## Run gazelle
	@gazelle update

gen:		## Run the goa gen command
	@goa gen github.com/t-hale/stox/design

help:		## Show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

login:		## Login with gcloud
	@gcloud auth application-default login

update:		## Update gcloud components
	@gcloud components update