package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"assuresfot/harmons/messagebrokerredis/dtos"
	"assuresfot/harmons/messagebrokerredis/messagebroker"
	"assuresfot/harmons/messagebrokerredis/woocommerceClient"
)

var ctx = context.Background()

func main() {
	rdb := messagebroker.GetClient()
	woocommerceClient.Connect()
	client := woocommerceClient.Client

	pubsub := rdb.Subscribe(ctx, "harmonschannel")
	ch := pubsub.Channel()

	for msg := range ch {
		var payload dtos.Payload
		json.Unmarshal([]byte(msg.Payload), &payload)

		var body url.Values
		json.Unmarshal(payload.Body, &body)

		var res *http.Response
		var err error
		switch payload.Method {
		case "GET":
			res, err = client.Get(payload.Url, body)
		case "POST":
			res, err = client.Post(payload.Url, body)
		case "PUT":
			res, err = client.Put(payload.Url, body)
		case "DELETE":
			res, err = client.Delete(payload.Url, body)
		}

		if err != nil {
			log.Fatal(err)
		} else if res.StatusCode != http.StatusOK {
			log.Fatal("Unexpected StatusCode:", res)
		} else {
			defer res.Body.Close()
			if bodyBytes, err := ioutil.ReadAll(res.Body); err != nil {
				log.Fatal(err)
			} else {
				fmt.Println(string(bodyBytes))
			}
		}
	}

	defer pubsub.Close()
}
