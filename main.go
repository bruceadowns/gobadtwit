package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
	var count, criteria string
	var length int
	flag.StringVar(&count, "count", "15", "result count")
	flag.StringVar(&criteria, "criteria", "@JiveSoftware", "search criteria")
	flag.IntVar(&length, "length", 80, "maximum tweet length")
	flag.Parse()

	keys, err := os.Open("keys.json")
	if err != nil {
		log.Fatal(err)
	}
	defer keys.Close()

	j := struct {
		ConsumerKey       string
		ConsumerSecret    string
		AccessToken       string
		AccessTokenSecret string
	}{}
	if err := json.NewDecoder(keys).Decode(&j); err != nil {
		log.Fatal(err)
	}

	anaconda.SetConsumerKey(j.ConsumerKey)
	anaconda.SetConsumerSecret(j.ConsumerSecret)
	api := anaconda.NewTwitterApi(j.AccessToken, j.AccessTokenSecret)

	v := url.Values{}
	v.Set("count", count)
	res, err := api.GetSearch(criteria, v)
	if err != nil {
		log.Fatal(err)
	}

	for _, tweet := range res.Statuses {
		s := tweet.Text
		if len(s) > length {
			s = s[:length] + "..."
		}

		fmt.Printf("%s: %s\n", tweet.User.ScreenName, s)
		//fmt.Printf("\n[%v]\n", tweet)
	}
}
