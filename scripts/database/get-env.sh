source ./scripts/env.sh

ENV_FILE=./configs/database/config.env

DATABASE_SCHEMA=$(read_var DATABASE_SCHEMA ${ENV_FILE})
DATABASE_HOST=$(read_var DATABASE_HOST ${ENV_FILE})
DATABASE_PORT=$(read_var DATABASE_PORT ${ENV_FILE})
POSTGRES_USER=$(read_var POSTGRES_USER ${ENV_FILE})
POSTGRES_PASSWORD=$(read_var POSTGRES_PASSWORD ${ENV_FILE})
POSTGRES_DB=$(read_var POSTGRES_DB ${ENV_FILE})

export POSTGRES_USER
export POSTGRES_DB
export DATABASE_URL="postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${POSTGRES_DB}?&sslmode=disable"