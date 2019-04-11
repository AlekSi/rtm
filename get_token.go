// +build ignore

package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/AlekSi/rtm"
)

func main() {
	log.SetFlags(0)

	keyF := flag.String("key", "", "API key")
	secretF := flag.String("secret", "", "API secret")
	permsF := flag.String("perms", "read", "Permission: 'read', 'write', or 'delete'.")
	flag.Parse()
	if *keyF == "" || *secretF == "" {
		log.Fatal("Both -key and -secret flags should be used.")
	}

	client := &rtm.Client{
		APIKey:    *keyF,
		APISecret: *secretF,
	}
	ctx := context.Background()

	frob, err := client.Auth().GetFrob(ctx)
	if err != nil {
		log.Fatal(err)
	}

	u := client.AuthenticationURL(rtm.Perms(*permsF), frob)
	log.Printf("Visit this URL: %s", u)

	var token string
	for i := 0; i < 30; i++ {
		token, _ = client.Auth().GetToken(ctx, frob)
		if token != "" {
			break
		}
		time.Sleep(time.Second)
	}
	if token == "" {
		log.Fatal("Failed to get authentication token.")
	}
	log.Printf("Authentication token: %q.", token)
}
