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

#### Start server

```bash
go run cmd/payments/main.go
```