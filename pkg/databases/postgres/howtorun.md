# how to run
# for dev
docker-compose -f docker-compose.dev.yml --env-file .env.dev up --build -d
<!-- --env-file is a environment variable on build runtime -->

# for test
docker-compose -f docker-compose.test.yml --env-file .env.test up --build -d

# for production
docker-compose -f docker-compose.prod.yml --env-file .env.prod up --build -d


# example when current dir is ./ of project
docker-compose -f pkg/databases/postgres/docker-compose.dev.yml --env-file pkg/databases/postgres/.env.dev up --build -d