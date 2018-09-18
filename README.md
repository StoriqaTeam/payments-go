# Payments
Payments backend for wallet apps

## Install

#### Linter

```bash
brew tap alecthomas/homebrew-tap
brew install gometalinter
```


For vs code
```json
"go.lintTool": "gometalinter",
"go.lintFlags": [
    "--config",
    "${workspaceRoot}/.gometalinter"
]
```

#### Dependency management

```bash
# Install dep management tool
brew install dep
# Fetch required sources
dep ensure
```


#### Database
```bash
cd docker && docker-compose run --service-ports db

```

#### Migration tool
Install

```bash
go get -u -d github.com/golang-migrate/migrate/cli
cd $GOPATH/src/github.com/golang-migrate/migrate/cli
# This one could run up to 10 minutes
dep ensure
go build -tags 'postgres' -o /usr/local/bin/migrate github.com/golang-migrate/migrate/cli
```

Run migrations:
```bash
migrate -path ./migrations -database postgres://payments:payments@localhost:5432/payments?sslmode=disable up
```

Create migration:
```bash
migrate create -ext sql migration_name
```

#### Start server

```bash
go run cmd/payments/main.go
```