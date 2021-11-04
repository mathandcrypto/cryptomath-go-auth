# cryptomath-go-auth
**Auth** Go microservice for managing authentication sessions.

## Install dependencies

```bash
$ make deps
```

## Build
```bash
$ make clean
$ make vendor
$ make copy-configs
$ make build
```

## Database migrations
### Apply all migrations

```bash
$ make migrate-up
```

### Down all migrations
```bash
$ make migrate-down
```

## Running the app

```bash
$ make run
```

## Local launch of docker containers application services
### PostgreSQL
```bash
# start service
$ make start-database

# init database
$ make init-database

# stop service
$ make stop-database
```

### Redis
```bash
# start service
$ make start-redis

# stop service
$ make stop-redis
```