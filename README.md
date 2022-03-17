# QuickBid

## Initialize
- Run postgres only with `docker-compose up -d postgres`
- Init DB versioning
```
    cd migrations/
    go run main.go init
    go run main.go up
```

## Run app
`docker-compose up quickbid -d --build`

## Stop app
`docker-compose stop`

## Generate mock
`go generate ./...`

## Code Arch
Based on Clean Architecture 
- Entity: Business objects
- UseCase: Business logics
- Repository: Data & resource management
- Bridge: Abstraction of libraries

- App: Connects layers
- Delivery: Delivery layer (eg HTTP/GRPC/Cron)

## Postman
https://www.getpostman.com/collections/02e2a5e7f2f4a44bb0e6
