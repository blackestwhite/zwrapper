# Zwrapper: Zarinpal API wrapper
to run the app, fill the entries in `.env.dev` and rename it to `.env`, also check ports in `docker-compose.yml` then use the following command to run the app:
```console
docker compose up -d --build
```

## Features
- automatically refresh authority before end-user interaction
- support for webhooks
- multiple consumers

## TODO
- [ ] check if vpn/proxy is being used