# Paper Disbursement System

A small REST API for creating users, recording their balances, and processing
disbursements (payouts) asynchronously via a background worker.

## Stack

- **HTTP**: [Echo](https://echo.labstack.com/) v4
- **Database**: PostgreSQL via [pgx/v5](https://github.com/jackc/pgx) (`pgxpool`)
- **Background jobs**: [asynq](https://github.com/hibiken/asynq) (Redis-backed)
- **Migrations**: [golang-migrate](https://github.com/golang-migrate/migrate)

## Running with Docker

Start the API server, worker, Postgres, and Redis with Docker Compose:

```sh
docker compose up --build
```

This will:
- start Postgres (`localhost:5432`) and Redis (`localhost:6379`)
- run database migrations automatically via the `migrate` service
- build and start the app (server + worker) on `localhost:3000`

To stop and remove containers:

```sh
docker compose down
```

To also remove the Postgres data volume:

```sh
docker compose down -v
```
