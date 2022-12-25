## **Normal Use**
1. run with DATABASE_URL=postgres://... PORT=:2565 go run server.go 
2. You can use REST API
- POST /expenses
- GET /expenses/:id
- PUT /expenses/:id
- GET /expenses
3. test view expenses GET /expenses with basic auth 
- username : apidesign 
- password : 123456

## ** Unit Test **
Use `go test --tags=unit -v ./...`

## **Docker Unit Test **
Use `docker build -t kbtg-assessment-peach:1.0.0 .`

## **Docker Integration Test **
1. config environment.env file
2. Use `docker-compose -f docker-compose.yml --env-file=environment.env up --build --abort-on-container-exit --exit-code-from it_tests`

