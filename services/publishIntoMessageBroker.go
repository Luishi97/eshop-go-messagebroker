package services

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"

	"assuresfot/harmons/messagebrokerredis/dtos"
	"assuresfot/harmons/messagebrokerredis/messagebroker"
)

func PublishIntoMessageBroker(c *fiber.Ctx) error {
	newUrl := fmt.Sprintf("%v", c.Locals("newUrl"))
	if newUrl == "" {
		newUrl = c.Path()
	}

	pay := &dtos.Payload{
		Url:    newUrl,
		Method: c.Method(),
	}

	payloadJson, err := json.Marshal(pay)
	if err != nil {
		panic(err)
	}

	rdb := messagebroker.GetClient()
	err = rdb.Publish(ctx, "harmonschannel", payloadJson).Err()
	if err != nil {
		panic(err)
	}
	return c.Send([]byte(c.Path()))
}
