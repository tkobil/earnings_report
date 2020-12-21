package internal

import (
	"fmt"
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func getEnvVar(environVarKey string) string {
	// Get's an environment variable and raises Error if does not exist
	envVal := os.Getenv(environVarKey)
	if envVal == "" {
		log.Fatal(fmt.Sprintf("ERROR: you must set the environment variable %v before running this program", environVarKey))
	}
	return envVal
}

func getClient() *twitter.Client {
	ConsumerKey := getEnvVar("CONSUMERKEY")
	ConsumerSecret := getEnvVar("CONSUMERSECRET")
	AccessToken := getEnvVar("ACCESSTOKEN")
	AccessTokenSecret := getEnvVar("ACCESSTOKENSECRET")

	config := oauth1.NewConfig(ConsumerKey, ConsumerSecret)
	token := oauth1.NewToken(AccessToken, AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	// we can retrieve the user and verify if the credentials
	// we have used successfully allow us to log in!
	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		log.Fatal(err) //Change to logging
	}

	log.Printf("User's ACCOUNT:\n%+v\n", user)
	return client
}

// SendTweets will synchronously tweet out every tweet string in the tweet passed in
func SendTweets(tweets []string) {
	/*
		:params tweet ([]string) - A slice of strings to be tweeted out sequentially
	*/

	client := getClient()
	for _, tweet := range tweets {
		fmt.Printf("TWEETING %v", tweet)
		tweet, resp, err := client.Statuses.Update(tweet, nil)
		if err != nil {
			log.Println(err)
		}
		log.Printf("%+v\n", resp)
		log.Printf("%+v\n", tweet)
	}
}
