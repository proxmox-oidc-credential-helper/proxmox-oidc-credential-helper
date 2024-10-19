all: clean test build

clean:
	rm -rf bin || true
build:
	go build -o bin/proxmox-oidc-credentials-helper github.com/camaeel/proxmox-oidc-credential-helper/cmd/proxmox-oidc-credential-helper

test:
	go test ./...

lint:
	golangci-lint run ./... -v

vet:
	$(gocmd)  vet ./...

trivy:
	trivy fs .
