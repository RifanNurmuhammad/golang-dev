

test:
	@if [ -f coverage.txt ]; then rm coverage.txt; fi;
	@echo ">> running unit test and calculate coverage"
	@go test -timeout 1200s -p 1 ./... -cover -coverprofile=coverage.txt -covermode=count -coverpkg=$(PACKAGES)
	@go tool cover -func=coverage.txt
