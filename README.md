# Welcome

## docker

`docker run --name transactionRoutine-mysql-docker -e MYSQL_ROOT_PASSWORD=password -d mysql:8`

## run app

`go run main.go`

## Curl for testing endpoints

POST /accounts

```sh
curl -X POST  \
-H 'Content-Type: application/json' \
-d '{"document_number":"123456"}' \
localhost:8080/account
```

POST /transactions

```sh
curl -X POST  \
-H 'Content-Type: application/json' \
-d '{"account_id":"5","operation_type_id":"1","amount":"-1223.23"}' \
localhost:8080/transaction
```
