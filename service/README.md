# Golang Project Example

## Tech Stack
* [Golang](https://go.dev/)
* [Fiber](https://gofiber.io/)
* [Sqlc](https://sqlc.dev/)
* [Viper](https://github.com/spf13/viper)
* [Cobra](https://github.com/spf13/cobra)
* [Postgresql](https://www.postgresql.org/)
    * [Posgresql Driver](https://github.com/jackc/pgx)
* [Redis](https://redis.io/)
    * [Redis Driver](https://github.com/redis/go-redis)
* [RabbitMQ](https://www.rabbitmq.com/)
    * [RabbitMQ Driver](https://github.com/rabbitmq/amqp091-go)
* [Database Migration](https://github.com/golang-migrate/migrate)

## Environment Variables
```bash
export GOX_DB_DSN="postgresql://postgres:12345@localhost:5432/go-example-db?sslmode=disable"
export GOX_REDIS_DSN="redis://localhost:6379/0"
```

Please modify yourself to match your environment.

## How To Run The Service
```bash
go build
./go-example serve
```
or 
```bash
go run main.go serve
```

## About SQLC In Brief
SQLC is a type-safe code generator from SQL. As developer, we only need define the raw SQL query and SQLC will generate functions based on our queries.
This way, we can automatically document all of our query and we can make sure what query that will be executed in the code because of the SQL-first approach.

## About Cobra In Brief
Cobra is tools to create modern cli applications. In this example the cli application has been initialized and the required commands may be already created. If you want to create a new command, first of course you must [install cobra](https://github.com/spf13/cobra?tab=readme-ov-file#installing) (please make sure you have install golang before), and run this command:
```bash
cobra-cli add <commandName>
```
or if you want to add child command, you can run this command:
```bash
cobra-cli add <commandName> -p '<parentCommandName>Cmd'
```

## About Viper In Brief
Viper is a tools for pasring and loading configuration from toml, yaml, env, json and remote sources. In our scenario, we will use env as our application configuration.
Viper can be integrated to cobra directly when initializing cli application with it.

## Application Layer
Handler --> [Usecase] --> Repository

* In our case, the repository layer is the code generated by SQLC.
* Use case layer is optional. Not all scenario need a use case, for example if you just want to get data from master table. In this case, it's okay from handler layer directly access the repository layer. This way, our application still clean and more straightforward.

## Database Migration
To rollup migration, run:
```
go-example migrate up
```
It will reapply migration from beginning to end.

To rolldown migration, run:
```
go-example migrate down
```
It will migrate database 1 version before.

## Other Note
* Hashing Password using [Argon2](https://en.wikipedia.org/wiki/Argon2), precisely Argon2id. In short, Argon2id is powerful enough defending password crack and side-channel attacks using GPU.
* RabbitMQ example still work in progress since it's need some modification related to the consumer deployment. Hence, the code is not committed yet. If you interested to further explanation, just slack me and I will explain through huddle.
