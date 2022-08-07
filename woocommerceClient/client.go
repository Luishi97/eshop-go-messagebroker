package woocommerceClient

import (
	"github.com/tgglv/wc-api-go/client"
	"github.com/tgglv/wc-api-go/options"
)

// Client is an instance to woocommerce client
var Client client.Client

// Connect create a connection to woocommerce and set the client
func Connect() {
	factory := client.Factory{}
	Client = factory.NewClient(options.Basic{
		URL:    "http://localhost:8000/",
		Key:    "ck_849d43c07e450eddea005a0f163fcecb1160e5f2",
		Secret: "cs_944119c8ddf08f6ef9981a16dabe791e5accfd7d",
		Options: options.Advanced{
			WPAPI:           true,
			WPAPIPrefix:     "/wp-json/",
			Version:         "wc/v3",
			QueryStringAuth: true,
			VerifySsl:       true,
		},
	})
}
