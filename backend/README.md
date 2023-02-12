# Backend

### How to connect to PG Admin with this Postgres Container:
First, start docker compose: `docker-compose up -d`
Then, open PG Admin at the configured port: `localhost:${PGADMIN_PORT}`
At PG Admin panel, add a new server with the following configs:
Host: `postgres` (Postgres container name)
Port: `5432` (Or whatever posgtres port you configured in .env file)
User: `user` (Or whatever posgtres user you configured in .env file)
Password: `password` (Or whatever posgtres password you configured in .env file)
