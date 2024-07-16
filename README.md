# go-simple-server
Simple go web server to be used with some terraform projects

## Endpoints

- `GET /hello_world`: 1. Returns a 200 status code with `{ "message": "Hello World!" }` JSON response
- `GET /current_time?name=same_name`: Returns a 200 status code with `{ "timestamp": 1700000000, "message": "Hello some_name" }`
- `GET /healthcheck`: Returns a 200 status code to indicate that the service is healthy

## Building the Docker Image

1. Clone the repository.
2. Navigate to the project directory.
3. Run the following command to build the Docker image:

```sh
docker build -t go-simple-server .
```

## Run container locally and expose port for the service (optional)
```sh
docker run -p 8080:8080 go-simple-server
```

## Authenticate against AWS ECR
https://docs.aws.amazon.com/AmazonECR/latest/userguide/registry_auth.html#:~:text=To%20authenticate%20Docker%20to%20an%20Amazon%20ECR%20registry%20with%20get,you%20want%20to%20authenticate%20to.

## How to tag the image
```sh
docker tag go-simple-server:latest <account_id>.dkr.ecr.us-east-1.amazonaws.com/go-simple-server:latest
```

## How to push the image to registry
```sh
docker push <account_id>.dkr.ecr.us-east-1.amazonaws.com/go-simple-server:latest
```

