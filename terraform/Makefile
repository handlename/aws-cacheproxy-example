PLAN_FILE := plan.out

init:
	terraform init

plan: init
	terraform plan -out=$(PLAN_FILE)

apply: init
	terraform apply $(PLAN_FILE)
	rm $(PLAN_FILE)


lock:
	terraform provider lock \
	   -platform=linux_amd64 \
	   -platform=linux_arm64 \
	   -platform=darwin_amd64 \
	   -platform=darwin_arm64
