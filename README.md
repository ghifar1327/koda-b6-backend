# Backend Coffee Shop

REST API backend for a coffee shop application, built with Go, Gin, PostgreSQL, and Redis.

## Features

- Authentication: register, login, forgot password, reset password
- Profile update and picture upload
- Products, variants, sizes, and categories
- Cart and transactions
- Product reviews
- Master data endpoints
- Swagger documentation

## Tech Stack

- Go 1.25
- Gin
- PostgreSQL
- Redis
- Swagger

## Project Structure

```text
cmd/                   application entrypoint
internal/handlers/     HTTP handler layer
internal/service/      business logic
internal/repository/   database and Redis access
internal/routes/       API routing
db/migrations/         SQL schema/migrations
docs/                  generated Swagger docs
```

## Environment Variables

Copy `.env.example` to `.env`, then adjust the values as needed.

```bash
cp .env.example .env
```

Environment variables used by the application:

- `PORT` application port
- `DATABASE_URL` PostgreSQL connection string for `pgxpool`
- `REDIS_HOST` Redis host
- `REDIS_PORT` Redis port
- `SECRET_KEY` secret key used for JWT signing
- `FRONTEND_URL` frontend URL
- `BACKEND_URL` backend URL
- `CORS_ORIGINS` allowed origins in JSON array format

Example:

```env
PORT=8888
DATABASE_URL=postgres://postgres:postgres@localhost:5432/coffee_shop?sslmode=disable
REDIS_HOST=localhost
REDIS_PORT=6379
SECRET_KEY=super-secret-key
FRONTEND_URL=http://localhost:3000
BACKEND_URL=http://localhost:8888
CORS_ORIGINS=["http://localhost:3000","http://localhost:5173"]
```

## Local Setup

Make sure you have:

- Go 1.25+
- PostgreSQL
- Redis

### 1. Install dependencies

```bash
go mod download
```

### 2. Create the database

Create a PostgreSQL database, then make sure `DATABASE_URL` points to that database.

### 3. Run migrations

This project stores SQL schema files in [db/migrations](/home/ghifar/koda-b6/backend-coffee-shop/db/migrations). Run the `.up.sql` files in this order:

1. `000001_init_db.up.sql`
2. `000002_forgot_pwd.up.sql`
3. `000003_product_variants.up.sql`
4. `000004_product_sizes.up.sql`
5. `000005_cart.up.sql`

If you use a migration tool such as `golang-migrate`, point it to that folder. If not, you can run the SQL files manually in PostgreSQL using the order above.

### 4. Prepare the upload folder

The picture upload endpoint saves files to `./uploads`, so create this folder before running the app:

```bash
mkdir -p uploads
```

### 5. Run the server

```bash
go run ./cmd
```

The server will run at:

```text
http://localhost:8888
```

## Docker

Build the image:

```bash
docker build -t backend-coffee-shop -f dockerfile .
```

Run the container:

```bash
docker run --env-file .env -p 8888:8888 backend-coffee-shop
```

Note: PostgreSQL, Redis, and the upload folder still need to be prepared separately.

## API Documentation

Swagger UI is available at:

```text
http://localhost:8888/swagger/index.html
```

Generated Swagger files:

- [docs/swagger.yaml](/home/ghifar/koda-b6/backend-coffee-shop/docs/swagger.yaml)
- [docs/swagger.json](/home/ghifar/koda-b6/backend-coffee-shop/docs/swagger.json)

## Endpoint Overview

Public endpoints:

- `GET /`
- `GET /products`
- `GET /products/:id`
- `GET /products/:id/variants`
- `GET /products/:id/sizes`
- `POST /auth/register`
- `POST /auth/login`
- `POST /auth/forgot-password`
- `PATCH /auth/reset-password`
- `GET /review-product`
- `GET /review-product/:id`
- `GET /landing/reviews`
- `GET /landing/reviews/:id`
- `GET /landing/recommended-product`
- `GET /landing/recommended-product/:id`
- `GET /master/:table`
- `GET /master/:table/:id`

Authenticated endpoints:

- `POST /transactions`
- `GET /transactions`
- `GET /transactions/:id`
- `GET /transactions/user/:id`
- `GET /cart/:user_id`
- `POST /cart`
- `DELETE /cart/:id`
- `POST /review-product`
- `PATCH /review-product/:id`
- `DELETE /review-product/:id`
- `PATCH /auth/:id/update`
- `PATCH /auth/:id/picture`

Admin endpoints:

- `GET /admin/users`
- `GET /admin/users/:id`
- `PATCH /admin/users/:id`
- `DELETE /admin/users/:id`
- `GET /admin/transaction`
- `PATCH /admin/transaction/:id`
- `DELETE /admin/transaction/:id`
- `POST /admin/master/:table`
- `PATCH /admin/master/:table/:id`
- `DELETE /admin/master/:table/:id`

## Notes

- JWT uses `SECRET_KEY` and the current token lifetime is 15 minutes.
- CORS reads `CORS_ORIGINS` in JSON array format, not as a plain string.
- Picture upload is limited to 1 MB.
