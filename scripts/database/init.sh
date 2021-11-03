source ./scripts/database/get-env.sh

docker-compose exec postgres psql --username="${POSTGRES_USER}" --owner="${POSTGRES_USER}" <<-EOSQL
  CREATE DATABASE $POSTGRES_DB;
EOSQL