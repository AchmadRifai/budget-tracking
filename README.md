# BUDGET TRACKING

## Requirements
- Docker:latest
- Golang:^1.19.0
- nodejs:^19
- nginx:latest

## Pre Running
### Running DB
1. Enter command `cd tech`
2. Enter command `bash start.sh`
3. Wait for DB to completely started

### Install Dependencies / Libraries
#### On Back End
1. Enter command `cd be`
2. Copy file .env.example to .env to allow set env variables
3. Edit .env file
4. Enter command `go mod download` to get all dependencies

#### On Front End
1. Enter command `cd fe`
2. Enter command `npm i`

## Running Applications
### On Back End
1. Enter command `cd be`
2. Enter command `go build main.go`
3. If use Windows, Enter command `main.exe`. If not, Enter command `./main`
### On Front End
1. Enter command `cd fe`
2. Enter command `npm run dev`
### On Nginx
1. Set PATH nginx in environtment
2. Enter command `nginx -c $(pwd)/nginx.conf`