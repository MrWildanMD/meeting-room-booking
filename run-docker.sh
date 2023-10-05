#!/bin/bash

# Read the environment variables from .env file
source .env

# Determine the database engine
if [ "$DB_ENGINE" = "postgresql" ]; then
    # PostgreSQL configuration
    cat <<EOF >docker-compose.yml
version: "3.9"
services:
  app:
    container_name: golang_container
    environment:
      - POSTGRES_USER=${DB_USERNAME}
      - POSTGRES_PASSWORD=${DB_PASSWORD}
      - POSTGRES_DB=${DB_NAME}
      - DATABASE_HOST=${DB_HOST}
      - DATABASE_PORT=${DB_PORT}
    tty: true
    build: .
    ports:
      - ${PORT}:${PORT}
    restart: on-failure
    volumes:
      - .:/app

  postgresdb:
    image: postgres:latest
    container_name: postgres_container
    environment:
      - POSTGRES_USER=${DB_USERNAME}
      - POSTGRES_PASSWORD=${DB_PASSWORD}
      - POSTGRES_DB=${DB_NAME}
      - DATABASE_HOST=${DB_HOST}
    ports:
      - ${DB_PORT}:${DB_PORT}
    volumes:
      - ./pg_data:/var/lib/postgresql/data
    restart: always

volumes:
  pg_data:
EOF
elif [ "$DB_ENGINE" = "mysql" ]; then
    # MySQL configuration
    cat <<EOF >docker-compose.yml
version: "3.9"
services:
  app:
    container_name: golang_container
    environment:
      - MYSQL_ROOT_USER=${DB_USERNAME}
      - MYSQL_ROOT_PASSWORD=${DB_PASSWORD}
      - MYSQL_DATABASE=${DB_NAME}
      - MYSQL_USER=${DB_USERNAME}
      - MYSQL_PASSWORD=${DB_PASSWORD}
      - DATABASE_HOST=${DB_HOST}
      - DATABASE_PORT=${DB_PORT}
    tty: true
    build: .
    ports:
      - ${PORT}:${PORT}
    restart: on-failure
    volumes:
      - .:/app

  mysql:
    image: mysql:latest
    container_name: mysql_container
    environment:
      - MYSQL_ROOT_USER=${DB_USERNAME}
      - MYSQL_ROOT_PASSWORD=${DB_PASSWORD}
      - MYSQL_DATABASE=${DB_NAME}
      - MYSQL_USER=${DB_USERNAME}
      - MYSQL_PASSWORD=${DB_PASSWORD}
    restart: always
    ports:
      - ${DB_PORT}:${DB_PORT}
    volumes:
      - ./mysql_data:/var/lib/mysql

volumes:
  mysql_data:
EOF
else
    echo "Unknown or unsupported database engine specified in DB_ENGINE."
    exit 1
fi

# Run Docker Compose
docker-compose up -d