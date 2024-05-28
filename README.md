# Golang S3 Processing Microservice

## Overview
You are tasked with developing two applications: a job and a microservice as product_service. The job's responsibility is to fetch several JSON Lines (jsonl) files from AWS S3, process them, and write the data to a database. On the other hand, the microservice should provide one endpoint which gets a record by id.

## Job Application

### Dockerization
To dockerize the job application, you have a few options:

1. **AWS Lambda**: You can create a Lambda function that is triggered by an S3 event. The Lambda function can fetch the JSONL files from S3, process them, and write the data to the database.

2. **Azure Functions**: Similar to AWS Lambda, you can create an Azure Function that is triggered by an S3 event. The function can perform the required processing and database operations.

3. **Google Cloud Functions**: Google Cloud Functions also provide a way to trigger functions based on S3 events. You can write a function that fetches the files, processes them, and stores the data in the database.

4. **Cron Job**: Alternatively, you can schedule a cron job on a server or container that runs periodically and fetches the JSONL files, processes them, and writes the data to the database.

Choose the option that best fits your requirements and infrastructure.

## Microservice Application

### Hexagonal Architecture
The microservice application follows the hexagonal architecture pattern. This pattern separates the application into different layers, including the domain (core) layer, which contains the business logic and interfaces.

### README.md
For detailed instructions on how to run and use the job and microservice applications, please refer to the respective README.md files in their respective directories.

- [Job Application README.md](/path/to/job/README.md)
- [Microservice Application README.md](/path/to/product_microservice/README.md)

## Best Practices
Both the job and microservice applications have been developed following best practices in software development. These practices include:

- Modular and maintainable code structure
- Proper error handling and logging
- Unit testing and test coverage
- Documentation for functions and APIs
- Version control using Git

Please refer to the README.md files in the respective application directories for more information on the best practices followed.
