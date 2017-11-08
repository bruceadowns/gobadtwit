package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
	var count, criteria string
	var length int
	flag.StringVar(&count, "count", "15", "result count")
	flag.StringVar(&criteria, "criteria", "@JiveSoftware", "search criteria")
	flag.IntVar(&length, "length", 80, "maximum tweet length")
	flag.Parse()

	keys, err := ioutil.ReadFile("keys.json")
	if err != nil {
		log.Fatal("error reading keys file")
	}

	jKeys := struct {
		ConsumerKey       string
		ConsumerSecret    string
		AccessToken       string
		AccessTokenSecret string
	}{}
	if err = json.Unmarshal(keys, &jKeys); err != nil {
		log.Fatalf("error unmarshalling keys.json: %s", err)
	}

	anaconda.SetConsumerKey(jKeys.ConsumerKey)
	anaconda.SetConsumerSecret(jKeys.ConsumerSecret)
	api := anaconda.NewTwitterApi(jKeys.AccessToken, jKeys.AccessTokenSecret)

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

		log.Printf("%s: %s\n", tweet.User.ScreenName, s)
		//log.Printf("\n[%v]\n", tweet)
	}
}
