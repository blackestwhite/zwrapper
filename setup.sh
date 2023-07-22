#!/bin/bash

# Function to prompt the user for input with default values
function prompt_with_default() {
    read -p "$1 (default: $2): " response
    echo "${response:-$2}"
}

# Prompt for environment variables
MERCHANT_ID=$(prompt_with_default "MERCHANT_ID" "merchant id by zarinpal")
ADMIN_USERNAME=$(prompt_with_default "ADMIN_USERNAME" "admin_username")
ADMIN_PASSWORD=$(prompt_with_default "ADMIN_PASSWORD" "admin_password")
BASE_URL=$(prompt_with_default "BASE_URL" "https://example.com/")

# Save environment variables to .env
cat > .env <<EOF
MERCHANT_ID="$MERCHANT_ID"
ADMIN_USERNAME="$ADMIN_USERNAME"
ADMIN_PASSWORD="$ADMIN_PASSWORD"
BASE_URL="$BASE_URL"
EOF

# Prompt for external port for docker-compose.yml
EXTERNAL_PORT=$(prompt_with_default "External port for docker-compose.yml" "8080")

# Save docker-compose.yml
cat > docker-compose.yml <<EOF
version: '3'
services:
  mongodb:
    image: mongo
    volumes:
      - /vol/zwrapper/data/db:/data/db
    restart: always
  app:
    build: .
    ports:
      - "$EXTERNAL_PORT:8080"
    restart: always
    depends_on:
      - mongodb
    links:
      - mongodb
EOF
