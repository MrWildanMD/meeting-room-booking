@echo off
setlocal enabledelayedexpansion

REM Read the environment variables from .env file
for /f "tokens=*" %%a in ('type .env') do (
  set "%%a"
)

REM Determine the database engine
if "%DB_ENGINE%" == "postgresql" (
  REM PostgreSQL configuration
  (
    echo version: "3.9"
    echo services:
    echo   app:
    echo     container_name: golang_container
    echo     environment:
    echo       - POSTGRES_USER=!DB_USERNAME!
    echo       - POSTGRES_PASSWORD=!DB_PASSWORD!
    echo       - POSTGRES_DB=!DB_NAME!
    echo       - DATABASE_HOST=!DB_HOST!
    echo       - DATABASE_PORT=!DB_PORT!
    echo     tty: true
    echo     build: .
    echo     ports:
    echo       - !PORT!:!PORT!
    echo     restart: on-failure
    echo     volumes:
    echo       - .:/app
    echo     networks:
    echo       - localnet
    echo   postgresdb:
    echo     image: postgres:latest
    echo     container_name: postgres_container
    echo     environment:
    echo       - POSTGRES_USER=!DB_USERNAME!
    echo       - POSTGRES_PASSWORD=!DB_PASSWORD!
    echo       - POSTGRES_DB=!DB_NAME!
    echo       - DATABASE_HOST=!DB_HOST!
    echo     ports:
    echo       - !DB_PORT!:!DB_PORT!
    echo     volumes:
    echo       - ./pg_data:/var/lib/postgresql/data
    echo     restart: always
    echo     networks:
    echo       - localnet
    echo volumes:
    echo   pg_data:
    echo networks:
    echo   localnet:
    echo      driver: bridge
  ) >docker-compose.yml
) elif "%DB_ENGINE%" == "mysql" (
  REM MySQL configuration
  (
    echo version: "3.9"
    echo services:
    echo   app:
    echo     container_name: golang_container
    echo     environment:
    echo       - MYSQL_ROOT_USER=!DB_USERNAME!
    echo       - MYSQL_ROOT_PASSWORD=!DB_PASSWORD!
    echo       - MYSQL_DATABASE=!DB_NAME!
    echo       - MYSQL_USER=!DB_USERNAME!
    echo       - MYSQL_PASSWORD=!DB_PASSWORD!
    echo       - DATABASE_HOST=!DB_HOST!
    echo       - DATABASE_PORT=!DB_PORT!
    echo     tty: true
    echo     build: .
    echo     ports:
    echo       - !PORT!:!PORT!
    echo     restart: on-failure
    echo     volumes:
    echo       - .:/app
    echo     networks:
    echo       - localnet
    echo   mysql:
    echo     image: mysql:latest
    echo     container_name: mysql_container
    echo     environment:
    echo       - MYSQL_ROOT_USER=!DB_USERNAME!
    echo       - MYSQL_ROOT_PASSWORD=!DB_PASSWORD!
    echo       - MYSQL_DATABASE=!DB_NAME!
    echo       - MYSQL_USER=!DB_USERNAME!
    echo       - MYSQL_PASSWORD=!DB_PASSWORD!
    echo     restart: always
    echo     ports:
    echo       - !DB_PORT!:!DB_PORT!
    echo     volumes:
    echo       - ./mysql_data:/var/lib/mysql
    echo     networks:
    echo       - localnet
    echo volumes:
    echo   mysql_data:
    echo networks:
    echo   localnet:
    echo      driver: bridge
  ) >docker-compose.yml
) else (
  echo Unknown or unsupported database engine specified in DB_ENGINE.
  exit /b 1
)

REM Run Docker Compose
docker-compose up -d