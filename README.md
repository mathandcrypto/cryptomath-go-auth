# cryptomath-go-auth
**Auth** microservice for managing authentication sessions.

## Microservice structure
The microservice has the following applications:
- **Auth** - the main application of the microservice that start up a [gRPC](https://grpc.io/) server.
- **Clear** - application for clearing expired refresh sessions from the database. It can usually be used in cron jobs.
- **Migrate** - application that apply all up database migrations.

## Install dependencies

```bash
$ make install-deps
```

## Build applications
Compiled application binaries saved to the local `/out/bin` folder.
Before build, it is necessary to do [vendoring](https://go.dev/ref/mod#vendoring):

```bash
$ make vendor
````

### Build `auth` application
```bash
$ make build-auth
```

### Build `clear` application
```bash
$ make build-clear
```

### Build `migrate` application
```bash
$ make build-migrate
```

## Database migrations
Replace `{DATABASE_URL}` to PostgreSQL database url, should look like: `postgres://username:password@localhost:5432/dbname`.

### Apply all migrations

```bash
$ DATABASE_URL="{DATABASE_URL}" make migrate-up
```

### Down all migrations
```bash
$ DATABASE_URL="{DATABASE_URL}" make migrate-down
```

## Database model generation
This service uses the [SQLBoiler](https://github.com/volatiletech/sqlboiler) package to generate models based on a database schema.
To generate, you must create a SQLBoiler configuration file `sqlboiler.toml` at the root of the project.
You can see an example of this configuration in the `sqlboiler.toml.example` file at the root of the project.
To start generating models, run the command
```bash
$ make boil-generate
```

## Project configuration files

### Application configuration
These are basically gRPC server settings.
You can see an example of this configuration in the `/configs/app/config.env.sample` file.

### Auth configuration
These are authorization parameters settings.
You can see an example of this configuration in the `/configs/auth/config.env.sample` file.

### Database configuration
These are the settings for connecting to the PostgreSQL database, which stores refresh session records.
You can see an example of this configuration in the `/configs/db/config.env.sample` file.

### Redis configuration
These are the settings for connecting to the Redis database, where access session data stored.
You can see an example of this configuration in the `/configs/redis/config.env.sample` file.