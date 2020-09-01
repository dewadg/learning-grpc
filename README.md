# Learning gRPC

This repo demonstrates aggregation process between 2 services (new service and user service). Both these services exposes only gRPC endpoints and will be aggregated by gateway to merge the responses

## Modules

### Gateway

Gateway exposes HTTP endpoint `GET /news` to display news data. This endpoint will aggregate news data from news service with user data from user service.

### News service

News service exposes gRPC endpoint that proxifies call to jsonplaceholder.typicode.com to get news (posts) data.

### User service

User service exposes gRPC endpoint that proxifies call to jsonplaceholder.typicode.com to get users data.

## How to run

Docker compose is mandatory, just run this command:

```
docker-compose up --build -d
```

Then you can hit `GET http://localhost:8000/news`.
