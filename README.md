# ACM Verifier

This application scans specified AWS regions to list ACM (AWS Certificate Manager) certificates and their validation methods.

## Requirements

- Docker

## Building the Docker Image

1. First, clone the repository to your local machine.

   ```bash
   git clone https://github.com/jessegersensonchess/aws-acm-checker
   cd aws-acm-checker
   docker build -t acmverifier .
   ```

## Running the Application

After building the Docker image, you can run the application using:
```
docker run -v ${HOME}/.aws:/root/.aws -profile your-aws-profile -regions us-west-1,us-west-2,us-east-1
```



