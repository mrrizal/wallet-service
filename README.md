# wallet-service
Wallet Service

## Installation
1. copy env.example to .env
2. build docker `docker-compose build`
3. migrate database, run the database first using `docker-compose up wallet-service-db -d` then run `make migrate-up`
4. udpate `POSTGRES_HOST` at .env file to `POSTGRES_HOST=wallet-service-db`

## Run
`docker-compose up wallet-service-api -d`
