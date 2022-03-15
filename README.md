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