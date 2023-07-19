# Welcome

## Starting app

To start the application, run `make up` on the root of the project.

## Verify

After the app is up, you will receieve a message like this:

```sh
app    | db connected
app    |
app    |    ____    __
app    |   / __/___/ /  ___
app    |  / _// __/ _ \/ _ \
app    | /___/\__/_//_/\___/ v4.10.2
app    | High performance, minimalist Go web framework
app    | https://echo.labstack.com
app    | ____________________________________O/_______
app    |                                     O\
app    | ⇨ http server started on 192.168.208.3:8080
```

The IP showed above is the IP from the app.
The IP for the DB is the same as the app, but the last digit is `2` instead of `3`. On the example above, the IP for the DB would be `192.168.208.2`.

## Endpoints

### POST /accounts

Request Body:

```json
{
  "document_number": "12345678900"
}
```

Response:

```json
{
  "4"
}
```

### POST /transaction

Request Body:

```json
{
  "account_id": 1,
  "operation_type_id": 4,
  "amount": 123.45
}
```

Response Body:

```json
{
  "1"
}
```

### GET /accounts/:accountId

Response Body:

```json
{
  "account_id": 1,
  "document_number": "12345678900"
}
```

## Curl examples for testing endpoints

POST /accounts

```sh
curl -X POST  \
-H 'Content-Type: application/json' \
-d '{"document_number":123456}' \
localhost:8080/account
```

POST /transaction

```sh
curl -X POST  \
-H 'Content-Type: application/json' \
-d '{"account_id":1,"operation_type_id":1,"amount":-1223.23}' \
localhost:8080/transaction
```

GET /accounts/:accountId

```sh
curl -X GET localhost:8080/account/1
```

## Exiting app

For shutting down the application, just run `make down` on the root of the project.
