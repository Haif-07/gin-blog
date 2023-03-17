package config

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type GithubUserInfo struct {
	AvatarURL         string      `json:"avatar_url"`
	Bio               interface{} `json:"bio"`
	Blog              string      `json:"blog"`
	Company           interface{} `json:"company"`
	CreatedAt         string      `json:"created_at"`
	Email             string      `json:"email"`
	EventsURL         string      `json:"events_url"`
	Followers         int         `json:"followers"`
	FollowersURL      string      `json:"followers_url"`
	Following         int         `json:"following"`
	FollowingURL      string      `json:"following_url"`
	GistsURL          string      `json:"gists_url"`
	GravatarID        string      `json:"gravatar_id"`
	Hireable          interface{} `json:"hireable"`
	HTMLURL           string      `json:"html_url"`
	ID                int         `json:"id"`
	Location          interface{} `json:"location"`
	Login             string      `json:"login"`
	Name              interface{} `json:"name"`
	OrganizationsURL  string      `json:"organizations_url"`
	PublicGists       int         `json:"public_gists"`
	PublicRepos       int         `json:"public_repos"`
	ReceivedEventsURL string      `json:"received_events_url"`
	ReposURL          string      `json:"repos_url"`
	SiteAdmin         bool        `json:"site_admin"`
	StarredURL        string      `json:"starred_url"`
	SubscriptionsURL  string      `json:"subscriptions_url"`
	Type              string      `json:"type"`
	UpdatedAt         string      `json:"updated_at"`
	URL               string      `json:"url"`
}

func Setup() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     "f0363cf539cb420d3e2f",
		ClientSecret: "622e1faf2e26eb29e06262ed5db0f6e99cff669c",
		Scopes:       []string{"repo", "user"},
		Endpoint:     github.Endpoint,
	}
}

func GetUrl(conf *oauth2.Config) string {
	url := conf.AuthCodeURL("state")
	return url
}

// GetToken 检索github oauth2令牌
func GetToken(ctx context.Context, conf *oauth2.Config, code string) (*oauth2.Token, error) {
	// url := conf.AuthCodeURL("state")
	// fmt.Printf("Type the following url into your browser and follow the directions on screen: %v\n", url)
	// fmt.Println("Paste the code returned in the redirect URL and hit Enter:")

	return conf.Exchange(ctx, code)
}

// GetUsers 使用oauth2获取用户信息
func GetUsers(client *http.Client) (user GithubUserInfo, err error) {
	var githubuserInfo GithubUserInfo
	url := fmt.Sprintf("https://api.github.com/user")

	resp, err := client.Get(url)

	if err != nil {
		return githubuserInfo, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return githubuserInfo, err
	}

	err = json.Unmarshal(body, &githubuserInfo)

	// fmt.Println("Status Code from", url, ":", resp.StatusCode)
	return githubuserInfo, err
}
