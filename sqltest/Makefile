TARG:=testdbconn
ZIP:=$(TARG).zip

.PHONY: createfunc build clean

vendor:
	dep ensure -v

build: $(ZIP)

clean:
	rm -f $(ZIP) $(TARG)

$(TARG):
	GOOS=linux go build -o $@ ./handler/lambda_handler.go
	
$(ZIP): $(TARG)
	zip $@ ./$(TARG)

createfunc: $(ZIP)
	aws lambda create-function \
	--region ap-southeast-2 \
	--function-name $(TARG) \
	--memory 256 \
	--role arn:aws:iam::142833200213:role/Lambda_BI_Server_Access \
	--runtime go1.x \
	--vpc-config SubnetIds=$(SUBNETIDS),SecurityGroupIds=$(SECURITYGROUPIDS) \
	--zip-file fileb://$(ZIP) \
	--handler $(TARG) \
	--publish \
	--timeout 30 \
	--environment Variables={USERNAME=$(USERNAME),PASSWORD=$(PASSWORD),HOST=$(HOST),PORT=$(PORT),DATABASE=$(DATABASE),INSTANCE=$(INSTANCE)} 

delfunc:
	aws lambda delete-function \
	--function-name $(TARG)

invokefunc:
	aws lambda invoke --function-name=$(TARG) out.log

testLocal:
	go run ./cli/main.go