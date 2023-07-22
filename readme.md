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

## TODO
- [ ] check if vpn/proxy is being used