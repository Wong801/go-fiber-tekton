/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"

	"git.finsoft.id/finsoft.id/go-example/db"
	"git.finsoft.id/finsoft.id/go-example/handlers"
	"git.finsoft.id/finsoft.id/go-example/usecase"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/jackc/pgx/v5"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the API",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Flag("toggle").Value)

		ctx := context.Background()

		dbConn, err := pgx.Connect(ctx, viper.GetString("GOX_DB_DSN"))
		if err != nil {
			log.Fatal(err)
		}
		defer dbConn.Close(ctx)

		usecase.DbConn = dbConn
		usecase.Queries = db.New(dbConn)

		opts, err := redis.ParseURL(viper.GetString("GOX_REDIS_DSN"))
		if err != nil {
			panic(err)
		}

		rdb := redis.NewClient(opts)

		usecase.Redis = rdb

		rabbitMqConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		if err != nil {
			log.Panicf("%s: %s", "Failed to connect to rabbitMQ server", err)
		}
		defer rabbitMqConn.Close()

		ch, err := rabbitMqConn.Channel()
		if err != nil {
			log.Panicf("%s: %s", "Failed to open rabbitMQ Channel", err)
		}
		defer ch.Close()

		usecase.RabbitMQChannel = ch

		app := fiber.New(fiber.Config{
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				resp := handlers.Response{
					Success: false,
					Code:    "INTERNAL_SERVER_ERROR",
					Message: err.Error(),
					Data:    nil,
				}
				return c.Status(fiber.StatusInternalServerError).JSON(resp)
			},
		})

		app.Use(recover.New())
		app.Use(requestid.New(requestid.Config{
			Header: "X-Request-ID",
			// Generator: func() string {
			// 	return "static-id"
			// },
		}))

		// init swagger
		cfg := swagger.Config{
			BasePath: "/api/v1/",
			FilePath: "./docs/swagger.json",
			Path:     "docs",
			Title:    "Finsoft API Docs",
		}

		app.Use(swagger.New(cfg))

		app.Use(healthcheck.New(healthcheck.Config{
			LivenessProbe: func(c *fiber.Ctx) bool {
				return true
			},
			LivenessEndpoint: "/livez",
			ReadinessProbe: func(c *fiber.Ctx) bool {
				return true
			},
			ReadinessEndpoint: "/healthz",
		}))
		jwtMiddleware := jwtware.New(jwtware.Config{
			SigningKey: jwtware.SigningKey{Key: []byte("secret")},
		})

		app.Get("/users", handlers.GetUsers)

		// kalo ketemu case kayak gini, ada 2 endpoint yg patternnya tabrakan,
		// yg non param taroh di atas karena kalo yg param ditaroh di atas,
		// ntar yg non param langsung dianggap param.
		// Sebisa mungkin dibikin nggak tabrakan ya.
		app.Get("/user/info", jwtMiddleware, handlers.UserInfo)
		app.Get("/user/:user_id", handlers.GetUser)

		app.Post("/register", handlers.Register)
		app.Post("/login", handlers.Login)

		fmt.Println("server started")
		log.Fatal(app.Listen(":8888"))
		fmt.Println("server terminated")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
