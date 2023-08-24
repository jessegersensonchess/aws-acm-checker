# ACM Verifier

This application scans specified AWS regions to list ACM (AWS Certificate Manager) certificates and their validation methods.

## Requirements

- Docker

## Building the Docker Image

1. First, clone the repository to your local machine.

   ```bash
   git clone [your-repository-url]
   cd [repository-directory]
   docker build -t acmverifier .
   ```

## Running the Application

After building the Docker image, you can run the application using:
```
docker run -profile your-aws-profile -regions us-west-1,us-west-2,us-east-1
```



