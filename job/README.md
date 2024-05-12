# Jsonl MongoDB Job


The job's responsibility is to fetch several JSON Lines (jsonl) files from AWS S3, process them, and write the data to a MongoDB database.

## Requirements

- Golang
- MongoDB
- Docker
- Git

## Setup

1. Clone the repository:

    ```bash
    git clone <repository_url>
    ```

2. Install the required dependencies:

    ```bash
    go mod download
    ```

3. Set up your AWS credentials by either exporting them as environment variables or using a configuration file. Refer to the `aws-sdk-go-v2` documentation for more information.

4. Build the Docker image:

    ```bash
    docker build -t jsonl-db-job .
    ```

## Usage

1. Run the Docker container:

    ```bash
    docker run 
    -e AWS_ACCESS_KEY_ID=<access_key> 
    -e AWS_SECRET_ACCESS_KEY=<secret_key> 
    -e MONGODB_URI=<mongodb_uri> 
    jsonl-db-job
    ```

    Replace `<access_key>` and `<secret_key>` with your AWS credentials.

2. The job will fetch the JSON Lines files from the AWS S3 bucket, process them, and write the data to the MongoDB database.

## Performance Considerations

To handle potential performance issues, the job employs the following strategies:

- Concurrency: The job processes the files concurrently to improve overall throughput.
- Deduplication: Before writing data to the database, the job checks if the records already exist to avoid duplicates.

If the job runs a second time with the same product files or a new file contains nearly identical records to those already in the database, it could be a waste of resources.
