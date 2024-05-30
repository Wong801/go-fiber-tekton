package usecase

import (
	"git.finsoft.id/finsoft.id/go-example/db"
	"git.finsoft.id/finsoft.id/go-example/lib/slack"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
	amqp "github.com/rabbitmq/amqp091-go"
)

var DbConn *pgx.Conn
var Redis *redis.Client
var Queries *db.Queries
var RabbitMQChannel *amqp.Channel
var Slack slack.SlackClient

func init() {
	Slack = slack.SlackClient{
		WebHookUrl: "https://hooks.slack.com/services/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",
		UserName:   "testing-alert-please-ignore",
		Channel:    "service-alert",
	}
}