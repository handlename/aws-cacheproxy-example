init:
	terraform init

plan: init
	terraform plan -out=plan.out

apply: init
	terraform apply plan.out

lock:
	terraform provider lock \
	   -platform=linux_amd64 \
	   -platform=linux_arm64 \
	   -platform=darwin_amd64 \
	   -platform=darwin_arm64
