# Sardine Take Home Exercise
#### Candidate

- Miguel Benedit

### Observations

The screen recording has been uploaded to YouTube, and you can find it at the following link: https://youtu.be/IkjYhUkCOf4. Please note that it may still be in the process of uploading. If you have any questions or need further clarification, feel free to reach out to me. Additionally, I want to mention that the only addition made after the recording was the inclusion of the Dockerfile.

### Run the App

1- Build the Dockerfile

```
docker build -t sardine-api .
```

2- Run the app

```
docker run -p 8080:8080 sardine-api
```


# Examples

## 1-  Valid request
The assigment example data request

```
curl --location 'http://localhost:8080/transactions/risk-evaluations' \
--header 'Content-Type: application/json' \
--data '{
  "transactions": [
    {"id": 1, "user_id": 1, "amount_us_cents": 200000, "card_id": 1},
    {"id": 2, "user_id": 1, "amount_us_cents": 600000, "card_id": 1},
    {"id": 3, "user_id": 1, "amount_us_cents": 1100000, "card_id": 1},
    {"id": 4, "user_id": 2, "amount_us_cents": 100000, "card_id": 2},
    {"id": 5, "user_id": 2, "amount_us_cents": 100000, "card_id": 3},
    {"id": 6, "user_id": 2, "amount_us_cents": 100000, "card_id": 4}
  ]
}'
```
Response
```
{
    "risk_ratings": [
        "low",
        "medium",
        "high",
        "low",
        "medium",
        "high"
    ]
}
```

## 2-  Invalid Empty request
The assigment example data request

```
curl --location 'http://localhost:8080/transactions/risk-evaluations' \
--header 'Content-Type: application/json' \
--data '{
  "transactions": [
  ]
}'
```
Response Bad request with 400 error code
```
empty transactions not allowed
```

## 3-  Invalid duplicated ID request
The assigment example data request

```
curl --location 'http://localhost:8080/transactions/risk-evaluations' \
--header 'Content-Type: application/json' \
--data '{
  "transactions": [
    {"id": 1, "user_id": 1, "amount_us_cents": 200000, "card_id": 1},
    {"id": 1, "user_id": 1, "amount_us_cents": 600000, "card_id": 1}
  ]
}'
```
Response Bad request with 400 error code
```
duplicated transactions ids not allowed
```