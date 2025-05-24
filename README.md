# **Swift Codes API**

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
1. Clone repository:
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
*Note: Occasionally, a database connection error may occur on startup. If this happens, simply stop the process (Ctrl+C) and rerun the command.*  
  
*If you are on windows: a window will pop up with permission etc, allow it and then run the command again*  
  
The API will be available at: http://localhost:8080  


## **Endpoints**
### **GET:**  
```bash
/v1/swift-codes/:swift_code
```  
Description:  
Retrieves a JSON object containing the details of a specific bank by its ID.  
If the bank is a headquarter, the response also includes the parameters of all its branches.  

Response Example:
```bash
{
  "address": "123 Main Street",
  "bankName": "National Bank",
  "countryISO2": "PL",
  "countryName": "Poland",
  "isHeadquarter": true,
  "swiftCode": "NATBPLPWXXX",
  "branches": [
    {
      "address": "Branch Street 1",
      "bankName": "National Bank Branch Krakow",
      "countryISO2": "PL",
      "isHeadquarter": false,
      "swiftCode": "NATBPLPWKRK"
    },
    {
      "address": "Branch Avenue 2",
      "bankName": "National Bank Branch Gdansk",
      "countryISO2": "PL",
      "isHeadquarter": false,
      "swiftCode": "NATBPLPWDAN"
    }
  ]
}

```  

### **GET:**  
```bash
/v1/swift-codes/country/:countryISO2code
```  
Description:  
Returns a list of all banks located in the country specified by the ISO 3166-1 alpha-2 code (ISO2).  
  
Response Example:
```bash
{
  "countryISO2": "PL",
  "countryName": "Poland",
  "swiftCodes": [
    {
      "address": "123 Main Street",
      "bankName": "National Bank",
      "countryISO2": "PL",
      "isHeadquarter": true,
      "swiftCode": "NATBPLPWXXX"
    },
    {
      "address": "Branch Street 1",
      "bankName": "National Bank Branch Krakow",
      "countryISO2": "PL",
      "isHeadquarter": false,
      "swiftCode": "NATBPLPWKRK"
    },
    {
      "address": "456 Central Ave",
      "bankName": "Polish Savings Bank",
      "countryISO2": "PL",
      "isHeadquarter": true,
      "swiftCode": "POLBPLPWXXX"
    }
  ]
}
```  


### **POST:**  
```bash
/v1/swift-codes
```
Description:  
Adds a new bank to the database.  
If a bank with the same unique identifier already exists, the operation is ignored or rejected.  

Rexuest Body Example:
```bash
{
  "address": "123 Main Street",
  "bankName": "National Bank",
  "countryISO2": "PL",
  "countryName": "Poland",
  "isHeadquarter": true,
  "swiftCode": "NATBPLPWXXX"
}

```  
Success Response Example:  
```bash
{
  "message": "Bank successfully added."
}
```    

### **DELETE:**  
```bash
/v1/swift-codes/:swiftCode
```  
Description:  
Deletes a bank from the database by its ID, but only if it exists.  

Success Response Example:  
```bash
{
  "message": "Bank successfully added."
}
```    
  
  
  
## **Running tests**
In repository location
```bash
go test ./internal/parser
```
```bash
go test ./internal/responeses
```
```bash
go test ./internal/integration_test
```
*Known issue: The final test fails from the terminal but passes in the VS Code UI. The endpoint itself returns the expected result when tested manually.*  

I have implemented a few integration tests for the GET endpoints. These tests check basic functionality, ensuring that the API returns the expected HTTP status codes and responses.  
  
However, I have not yet written tests for all endpoints, as I am still learning about integration testing. I plan to improve my knowledge and add more tests in the future.


## Stopping and removing containers:
```bash
docker-compose down
```

## What I Learned
This project helped me improve in several key areas:
- Building REST APIs in Go using Gin and GORM
- Dockerizing applications and using Docker Compose
- Working with PostgreSQL 
- Writing basic unit and integration tests
- Structuring a clean Go project with packages  
  
## Feedback
I'm open to feedback and suggestions! Feel free to reach out via issues or contact me directly.  
