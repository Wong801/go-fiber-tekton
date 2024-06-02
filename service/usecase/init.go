package usecase

import (
	"git.finsoft.id/finsoft.id/go-example/service/db"
	"git.finsoft.id/finsoft.id/go-example/service/lib/slack"
	"github.com/jackc/pgx/v5"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var DbConn *pgx.Conn
var Redis *redis.Client
var Queries *db.Queries
var RabbitMQChannel *amqp.Channel
var Slack slack.SlackClient

func init() {
	Slack = slack.SlackClient{
		WebHookUrl: viper.GetString("GOX_SLACK_WEBHOOK_URL"),
		UserName:   "testing-alert-please-ignore",
		Channel:    "service-alert",
	}
}
