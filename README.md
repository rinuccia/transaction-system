# HTTP Transaction system server.

The application was written for the purpose of practicing Go programming ([task](TASK.md))

## Run Project

* Configure `.env` and `config.yml` files
* Run `docker-compose up -d`

## API Endpoints

### POST `"/user/new"`
Create new user

Request
```json
{
"first_name": "John",
"last_name": "Doe",
"email": "johndoe@gmail.com"
}
```

Response
HTTP Status Code 200
```json
{
"userID": "1"
}
```

### GET `"/user/:id"`
Get user by ID

Response
HTTP Status Code 200
```json
{
"id": "1",
"first_name": "John",
"last_name": "Doe",
"email": "johndoe@gmail.com",
"balance": 20.30
}
```

### PUT `"/user/request"`
Create a user request to replenish or withdraw.
Operation Code:
* `-` - withdraw
* `+` - replenish

Request
```json
{
"id": "1",
"amount_of_money": 2.33,
"operation_code": "+"
}
```

Response
HTTP Status Code: 200
