# Simple Meeting Room Booking App

![image](https://github.com/MrWildanMD/meeting-room-booking/assets/60354284/beed153d-de1e-4efc-b104-9f599581b962)

# Main Functionality
- Notification via Telegram bot if new Booking added, you can join here: https://t.me/+gL6b2AJeSrMyZGQ1
- Work under Golang-Gin for backend and Svelte for fronted
- Extendable as your need!
- Run ```docker-compose up``` to fire-up the backend
- Run ```cd frontend/``` and install dependencies with ```pnpm install``` after done installing all dependencies run ```pnpm dev``` to fire-up the frontend
- App run in port :8000 for backend and :5173 for frontend

# Configuration
- Change ```APP_MODE``` in ```.env``` to ```docker``` if you want to use docker, change to ```local``` if you want to use without docker
- Change Telegram Token if you want to use your own BOT API TOKEN
- To customize the bot target group you can edit it in the ```config/notif.go``` and change ChatID to your own group ChatID

# Additional Information
To log in/register you can make your own credential by hit ```/api/users/register``` and ```/api/users/login```
and what? nothing you are ready to go with your own front-end/ready made but yet still need improvement xD.

# Tech Stack
- Golang
- Postgresql
- Svelte

# About
This project is mainly for portofolio, but it can be extendable to use in production, bring your own frontend and integrate it with the Go backend, all endpoints are stored under ```routers/``` dir
