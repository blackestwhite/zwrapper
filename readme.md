# Zwrapper: Zarinpal API wrapper
to run the app, make sure you have docker, docker compose and curl installed. then use the following commands:
```terminal
git clone git@github.com:blackestwhite/zwrapper.git
cd zwrapper
chmod +x ./setup.sh
./setup.sh
docker compose up -d --build
```

## Features
- automatically refresh authority before end-user interaction
- support for webhooks
- multiple consumers

## Endpoints
### Admin
#### creating new consumers

`POST /api/v1/admin/newConsumer?username=admin_username&password=admin_password`

body:
```json
{
    "consumer": "consumer name"
}
```

on success returns:
```json
{
    "ok": true,
    "result": {
        "message": "token generated successfully",
        "id": "insertion ID",
        "consumer": "consumer name",
        "token": "uuid access token"
    }
}
```

### Payments

#### creating new payments

`POST /api/v1/payment/new`

access token fetched in consumer creation should be included as a header which name is `x-zwrapper-access-token`

body:
```json
{
    "amount": 10000,
    "next": "https://example.com/afterPaymentFailOrSuccess",
    "webhook": "https://example.com/afterPaymentSuccessWebhook",
    "description": "payment description"
}
```

on success returns:
```json
{
    "ok": true,
    "result": {
        "id": "payment id"
    }
}
```
#### going to payment page(for end-user)

`GET /api/v1/payment/pay/:id`

`:id` is the id fetched in payment creation step

### verify payments(wheter a payment is paid or not)

`POST /api/v1/payment/verify/:id`

`:id` is the id fetched in payment creation step

access token fetched in consumer creation should be included as a header which name is `x-zwrapper-access-token`

if paid retuns:
```json
{
    "ok": true,
    "result": {
        "paid": true,
        "ref": "ref"
    }
}
```

if not paid returns:
```json
{
    "ok": true,
    "result": {
        "paid": false,
        "verified": false,
        "ref": "ref",
        "status": 123,
        "error": "error message"
    }
}
```


## TODO

- [ ] check if vpn/proxy is being used