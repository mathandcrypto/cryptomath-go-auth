source ./scripts/database/get-env.sh

docker exec cryptomath-auth-postgres psql --username="${POSTGRES_USER}" <<-EOSQL
  CREATE DATABASE $POSTGRES_DB;
EOSQL