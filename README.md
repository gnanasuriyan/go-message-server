# go-message-server
Simple message server implemented in golang (Fiber framework)

## Dependencies
1. Golang
2. Docker
3. Docker-compose

## How to run
1. Clone the repository
2. Run the following command to initialize the server

    ```cd docker && docker-compose up db -d```

    ```cd .. && do.sh init-db```

3. Run the following command to start the server

    ```go run main.go```

4. The server will be running on port 3000
