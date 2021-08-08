test:
    go test ./...

coverage:
    go test ./... -coverprofile cover.out
    go tool cover -html=cover.out -o coverage.html
    chrome ${PWD}/coverage.html

lint:
    golangci-lint run