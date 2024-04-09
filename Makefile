default: help

.PHONY: apply destroy gen help init login updatem

apply:		## Run terraform apply
	@terraform -chdir=terraform apply

destroy:	## Run terraform destroy
	@terraform -chdir=terraform destroy

gen:		## Run the goa gen command
	@goa gen github.com/t-hale/stox/design

help:		## Show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

init:		## Run terraform init
	@terraform -chdir=terraform init

login:		## Login with gcloud
	@gcloud auth application-default login

update:		## Update gcloud components
	@gcloud components update