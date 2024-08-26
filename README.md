# API Metering System

## Author
- Saifudin
- Email: qsaifudin.official@gmail.com
- LinkedIn: https://www.linkedin.com/in/qsaifudin/
- Personal Web: https://qsaifudin.site/

## Project Resources
- GitHub Repo : https://github.com/qsaifudin/golang-api-metering-system.git
- Video Demo : https://youtu.be/3VNMNxAbVBk

## Overview
The API Metering System is designed to monitor and control the usage of API endpoints and server storage. It tracks the number of requests made to different API endpoints and the amount of storage used by uploaded files. The system enforces usage limits, where API access is stopped after reaching a specified threshold:

- API Request Limit: Stops API access after 1,000 requests per endpoint.
- Storage Limit: Stops file uploads after the total storage used reaches 1 GB (1,073,741,824 bytes).

## Project Structure

```
├── main.go                 # Entry point for the application
├── metering
│   ├── api_metering.go     # Handles API request metering
│   ├── storage_metering.go # Handles storage metering
├── handlers
│   ├── api_handler.go      # HTTP handlers for API metering
│   └── storage_handler.go  # HTTP handlers for storage metering
├── tests
│   ├── api_metering_test.go    # Unit tests for API metering
│   ├── storage_metering_test.go# Unit tests for storage metering
│   └── handlers_test.go        # Integration tests for HTTP handlers
└── README.md              # Project documentation
```


## How to Run the Project
1. Clone the Repository:
    ```bash
    git clone https://github.com/qsaifudin/golang-api-metering-system.git
    cd api_metering_system
    ```

2. Start the Phoenix Server:
    ```bash
    go run main.go
    ```
    This will start the server on `http://localhost:8080`

## Available Endpoints
- Available Endpoints
- POST /api/endpoint1 : Track requests to endpoint1.
- GET /api/metrics : Retrieve request counts for all endpoints.
- POST /upload : Upload a file and track its size.
- GET /storage : Retrieve total storage used.

## How to Test

- Run All Unit Tests:
  ```bash
  go test ./tests -v
  ```
