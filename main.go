package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/fitbit"
)

func main() {
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     os.Getenv("FITBIT_CLIENT_ID"),
		ClientSecret: os.Getenv("FITBIT_CLIENT_SECRET"),
		Scopes:       []string{"activity"},
		Endpoint:     fitbit.Endpoint,
	}

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	day := time.Now()
	const layout = "2006-01-02"
	date := day.Format(layout)

	client := conf.Client(ctx, tok)
	res, _ := client.Get(fmt.Sprintf("https://api.fitbit.com/1/user/-/activities/date/%s.json", date))
	b, _ := ioutil.ReadAll(res.Body)

	fmt.Printf(string(b))
}
