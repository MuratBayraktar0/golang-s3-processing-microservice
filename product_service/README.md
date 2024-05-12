# Jsonl MongoDB Microservice

This microservice is designed to provide a RESTful API for fetching product data from a MongoDB database.

## Technologies Used

- Golang
- Go Convey (for unit testing)
- MongoDB
- Docker

## Prerequisites

Before running the microservice, make sure you have the following:

- Golang installed
- MongoDB installed and running
- Docker installed

## Setup

1. Clone the repository:

    ```bash
    git clone <repository_url>
    ```

2. Install the required dependencies:

    ```bash
    go mod download
    ```

3. Create the necessary environment files:

    - Create a `.env.prod` file with the production environment variables.
    - Create a `.env.dev` file with the development environment variables.

4. Build the Docker image:

    ```bash
    docker build -t jsonl-db-microservice .
    ```

## Usage

1. Run the Docker container:

    ```bash
    docker run 
    -e ENV=dev 
    -e MONGODB_URI=<mongodb_uri>
    jsonl-db-microservice
    ```

    Replace `<mongodb_uri>` with the URI of your MongoDB database.

2. The microservice will start and listen for incoming requests.

## Endpoints

The microservice provides the following endpoints:

- `GET /product/:id`: Fetches a single product by its ID from the database.
- `GET /products`: Fetches all products from the database.

## Scalability Considerations

To design the microservice with scalability in mind, consider the following:

- Use a load balancer to distribute incoming requests across multiple instances of the microservice.
- Implement caching mechanisms to reduce the load on the database.
- Optimize database queries and indexes for efficient retrieval of data.
