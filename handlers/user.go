package handlers

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"git.finsoft.id/finsoft.id/go-example/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Login(c *fiber.Ctx) error {
	loginReq := new(LoginRequest)
	if err := c.BodyParser(loginReq); err != nil {
		return err
	}

	err := Validate.Struct(loginReq)
	if err != nil {
		return err
	}

	err = usecase.Login(c.Context(), loginReq.Email, loginReq.Password)
	if err != nil {
		return err
	}

	resp := Response{
		Success: true,
		Code:    "LOGIN.SUCCESS",
		Message: "login success",
		Data:    nil,
	}

	return c.JSON(resp)
}

func Register(c *fiber.Ctx) error {
	registerReq := new(RegisterRequest)
	if err := c.BodyParser(registerReq); err != nil {
		return err
	}

	err := Validate.Struct(registerReq)
	if err != nil {
		return err
	}

	err = usecase.Register(c.Context(), registerReq.Email, registerReq.Password, registerReq.RoleIds)
	if err != nil {
		return err
	}

	usecase.Slack.SendInfo("Registration success")

	return c.JSON(map[string]string{"message": "register success"})
}

func GetUsers(c *fiber.Ctx) error {
	redisKey := c.OriginalURL()

	userCache, err := usecase.Redis.Get(c.Context(), redisKey).Bytes()
	if err == nil {
		return c.Send(userCache)
	}

	users, err := usecase.Queries.GetUsers(c.Context())
	if err != nil {
		return err
	}

	userJson, _ := json.Marshal(users)
	usecase.Redis.SetEx(c.Context(), redisKey, string(userJson), 5*time.Minute)

	// var response []map[string]any
	// for _, user := range users {
	// 	response = append(response, map[string]any{
	// 		"id":    user.ID,
	// 		"name": user.Name,
	// 		"email": user.Email,
	// 	})
	// }
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	user, err := usecase.Queries.GetUserById(c.Context(), uuid.MustParse(userID))
	if err != nil {
		return err
	}

	// var response []map[string]any
	// for _, user := range users {
	// 	response = append(response, map[string]any{
	// 		"id":    user.ID,
	// 		"name": user.Name,
	// 		"email": user.Email,
	// 	})
	// }
	return c.JSON(user)
}

func SendRabbitMQMessage(c *fiber.Ctx) error {
	q, err := usecase.RabbitMQChannel.QueueDeclare(
		"greetings", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	body := `{"greeting": "hello world!"}`
	err = usecase.RabbitMQChannel.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})
	if err != nil {
		return err
	}
	log.Printf(" [x] Sent %s\n", body)

	resp := Response{
		Success: true,
		Code:    "RABBITMQ.SEND",
		Message: "Send message success",
		Data:    body,
	}

	return c.JSON(resp)
}
