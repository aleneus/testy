check:
	@go test

cover:
	@go test --coverprofile=cover.out ./...
	@go tool cover -func=cover.out | tail -n1 | tr -s "\t"
	@go tool cover -html="cover.out" -o "cover.html"
	@rm -f cover.out
