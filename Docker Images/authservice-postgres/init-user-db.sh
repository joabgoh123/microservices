#!/bin/bash

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
        CREATE TABLE Users (
            user_id serial PRIMARY KEY,
            username VARCHAR ( 50 ) UNIQUE NOT NULL,
            password VARCHAR ( 50 ) NOT NULL,
            created_on TIMESTAMP NOT NULL,
            last_login TIMESTAMP
        )
EOSQL