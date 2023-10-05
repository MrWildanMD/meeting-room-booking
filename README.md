# Simple Meeting Room Booking App

![image](https://github.com/MrWildanMD/meeting-room-booking/assets/60354284/beed153d-de1e-4efc-b104-9f599581b962)

# Main Functionality
- Notification via Telegram bot if new Booking added, you can join here: https://t.me/+gL6b2AJeSrMyZGQ1
- Work under Golang-Gin for backend and Svelte for fronted
- Extendable as your need!
- Run ```run-docker.sh``` if you are using linux or ```run-docker.bat``` if you are using windows to fire-up the backend using docker
- Run ```cd frontend/``` and install dependencies with ```pnpm install``` after done installing all dependencies run ```pnpm dev``` to fire-up the frontend
- App run in port :8000 for backend and :5173 for frontend

# Configuration
- Change Telegram Token if you want to use your own BOT API TOKEN
- To customize the bot target group you can edit it in the ```config/notif.go``` and change ChatID to your own group ChatID

# Project Structure Explanation
- ```config``` = Storing configuration that used by app
- ```controllers``` = Storing the api handler
- ```database``` = Storing the database things
- ```frontend``` = Storing the frontend
- ```middlewares``` = Storing the gin middleware
- ```models``` = Storing the models
- ```routers``` = Storing the endpoints
- ```utils``` = Storing the utilities that used by app
- ```.env``` = App ENV
- ```Dockerfile``` = Dockerize the app
- ```main.go``` = Main driver

# Additional Information
Default dummy user information
Admin: admin@test.com | 12345
Guest: guest@test.com | 12345

# Tech Stack
- Golang
- Postgresql
- Svelte

# About
This project is mainly for portofolio, but it can be extendable to use in production, bring your own frontend and integrate it with the Go backend, all endpoints are stored under ```routers/``` dir
