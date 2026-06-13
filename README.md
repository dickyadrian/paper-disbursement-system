# Paper Disbursement System

A small REST API for creating users, recording their balances, and processing
disbursements (payouts) asynchronously via a background worker.

## Stack

- **HTTP**: [Echo](https://echo.labstack.com/) v4
- **Database**: PostgreSQL via [pgx/v5](https://github.com/jackc/pgx) (`pgxpool`)
- **Background jobs**: [asynq](https://github.com/hibiken/asynq) (Redis-backed)
- **Migrations**: [golang-migrate](https://github.com/golang-migrate/migrate)
