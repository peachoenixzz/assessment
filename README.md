## **Normal Use**
1. Run with `DATABASE_URL=postgres://... PORT=:2565 go run server.go` 
2. You can use REST API
- POST /expenses
- GET /expenses/:id
- PUT /expenses/:id
- GET /expenses
3. If test with basic auth use `expenses.postman_collection_with_auth.json` or Use username , password
- Username : apidesign
- Password : 123456

## **Unit Test**
Use `go test --tags=unit -v ./...`

## **Docker Build With Unit Test**
1. Config environment.env file 
```
DATABASE_URL=YOUR URL DATABASE
PORT=:2565
```
2. Use `docker build -t kbtg-assessment-peach:1.0.0 .`
3. Use `docker run --env-file=environment.env -p 2565:2565 kbtg-assessment-peach:1.0.0`

## **Docker Integration Test**
1. Config environment.env file
```
DATABASE_URL=YOUR URL DATABASE
PORT=:2565
```
2. Use `docker-compose -f docker-compose.yml --env-file=environment.env up --build --abort-on-container-exit --exit-code-from it_tests`

