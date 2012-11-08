/*
	Written by José Carlos Nieto <xiam@menteslibres.org>
	(c) 2012 Astrata Software http://astrata.mx

	MIT License
*/

package twitter

import (
	"fmt"
	"github.com/garyburd/go-oauth/oauth"
	"github.com/gosexy/sugar"
	"github.com/gosexy/to"
	"github.com/gosexy/yaml"
	"testing"
	"time"
)

var SettingsFile = "settings.yaml"

var conf *yaml.Yaml

func TestSettings(t *testing.T) {
	var err error
	conf, err = yaml.Open(SettingsFile)

	if err != nil {
		panic(err.Error())
	}

}

func TestApi(t *testing.T) {

	var err error

	client := New(&oauth.Credentials{
		to.String(conf.Get("twitter/app/key")),
		to.String(conf.Get("twitter/app/secret")),
	})

	client.SetAuth(&oauth.Credentials{
		to.String(conf.Get("twitter/user/token")),
		to.String(conf.Get("twitter/user/secret")),
	})

	_, err = client.VerifyCredentials()

	if err != nil {
		t.Errorf("Test failed.\n")
	}

	_, err = client.HomeTimeline()

	if err != nil {
		t.Errorf("Test failed.\n")
	}

	_, err = client.MentionsTimeline()

	if err != nil {
		t.Errorf("Test failed.\n")
	}

	_, err = client.UserTimeline()

	if err != nil {
		t.Errorf("Test failed.\n")
	}

	_, err = client.RetweetsOfMe()

	if err != nil {
		t.Errorf("Test failed.\n")
	}

	_, err = client.Retweets(int64(21947795900469248))

	if err != nil {
		t.Errorf("Test failed.\n")
	}

	var status *sugar.Tuple
	status, err = client.Update(fmt.Sprintf("Test message @ %s", time.Now()), nil)

	if err != nil {
		t.Errorf("Test failed.\n")
	}

	tweetId := to.Int64(status.Get("id_str"))

	_, err = client.Destroy(tweetId)

	if err != nil {
		t.Errorf("Test failed.\n")
	}

	_, err = client.Retweet(int64(21947795900469248))

	if err != nil {
		t.Errorf("Test failed.\n")
	}

}
