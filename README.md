 # **Remitly Internship 2025 Homework Task**

## **Swift Codes API**

A simple RESTful API for managing SWIFT (BIC) code data. Built with Go and PostgreSQL, containerized using Docker.

## **Technologies Used**
- **Go** (1.24+)
- **PostgreSQL**
- **Gin** – HTTP web framework
- **GORM** – ORM for Go
- **Docker** and **Docker Compose**
- **dotenv** – environment variable management

## **Prerequisites**
- Docker
- Git

## **Installation and setup**  
*if you are on Windows, use Powershell*  
In your terminal:
1. Clone my repository in your chosen place:
```bash
git clone http://github.com/bpietrzakk/swift_codes
```
```bash
cd swift_codes
```
2. Build the docker image:
```bash
docker build -t swift_codes .
```
4. Run the app with Docker Compose:  
```bash
docker-compose up --build
```
*If you are on windows: a window will pop up with permission etc, allow it and then run the command again*  
  
The API will be available at: http://localhost:8080  

## **Running tests**
In repository location
```bash
go test ./internal/parser
```
```bash
go test ./internal/responeses
```


## Stopping and removing containers:
```bash
docker-compose down
```


