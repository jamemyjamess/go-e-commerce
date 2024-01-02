# How to run postgresql with docker
## example for dev
```
docker-compose -f docker-compose.dev.yml --env-file .env.dev up --build -d
```
## Example for dev
docker-compose -f docker-compose.test.yml --env-file .env.test up --build -d
## Example for production
docker-compose -f docker-compose.prod.yml --env-file .env.prod up --build -d
## Example when current dir is ./ of project
docker-compose -f pkg/databases/postgres/docker-compose.dev.yml --env-file pkg/databases/postgres/.env.dev up --build -d
### Flags
| Flags      | Description |
| --------- | -----:|
| -f     |   path of docker-compose |
| --envfile  | environment variable on build runtime |
| --build |   build images and containers |
| -d      |    build as a detach mode or background mode |


# How to create a migration
## Create
```
# migrate create -ext sql -dir ./migrations -seq {name}
migrate create -ext sql -dir ./migrations -seq init-e-commerce-db
```
### Flags
| Flags      | Description |
| --------- | -----:|
| -dir     |   path to a migration directory |
