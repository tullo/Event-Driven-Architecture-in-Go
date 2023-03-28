#!/bin/sh
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  CREATE DATABASE pactdb;

  CREATE USER pactuser WITH ENCRYPTED PASSWORD 'pactpass';

  GRANT CREATE, CONNECT ON DATABASE pactdb TO pactuser;
EOSQL

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "pactdb" <<-EOSQL

  GRANT CREATE ON SCHEMA public TO pactuser;
EOSQL
