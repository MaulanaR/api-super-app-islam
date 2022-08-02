# Zahir HRM API

## Getting Started
1. Make sure you have [Go](https://go.dev) installed.
2. Clone the repo
```bash
git clone https://bitbucket.org/zahironline/zahirhrm-api.git
```
3. Go to the directory and run go mod tidy to add missing requirements and to drop unused requirements
```bash
cd zahirhrm-api && go mod tidy
```
3. Setup your .env file
```bash
cp .env-example .env && vi .env
```
4. Start
```bash
go run main.go
```

## Generate / Update Swagger
1. Make sure you have added GOPATH/bin to the PATH environment.
2. Compile and install [Swag](https://github.com/swaggo/swag) packages and dependencies.
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```
3. Add comments to your API source code, [See Declarative Comments Format](https://github.com/swaggo/swag#declarative-comments-format).
4. Run the [Swag](https://github.com/swaggo/swag) in your Go project root folder which contains `main.go` file, [Swag](https://github.com/swaggo/swag) will parse comments and generate required files(`docs` folder and `docs/doc.go`).
```bash
swag init --outputTypes json
```
5. Start
```bash
go run main.go
```
6. Open http://localhost:4040/docs/ in your browser, you can see Swagger 2.0 Api documents.

## Test
1. Test all
```bash
go test ./...
```
2. Test all with verbose output that lists all of the tests and their results.
```bash
go test ./... -v
```
3. Test all with benchmark.
```bash
go test ./... -bench=.
```

## Build for production
1. Compile packages and dependencies
```bash
go build -o zahirhrm-api main.go
```
2. Setup .env file for production
```bash
cp .env-example .env && vi .env
```
3. Run executable file with systemd, supervisor, pm2 or other process manager
```bash
./zahirhrm-api
```