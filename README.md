# Welcome

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
-d '{"account_id":"1","operation_type_id":"2","amount":"3"}' \
localhost:8080/transaction
```
