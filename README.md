# Requester

Do a GET request to a given URL and parse the result into a given target.

## How to install

    go get github.com/gepser/requester

## How to use

``` go
package main

import (
    "fmt"

    "github.com/gepser/requester"
)

//GithubRootStruct is an example struct to be used to parse the response of the root github api
type GithubRootStruct struct {
    CurrentUserURL                   string `json:"current_user_url"`
    CurrentUserAuthorizationsHTMLURL string `json:"current_user_authorizations_html_url"`
    AuthorizationsURL                string `json:"authorizations_url"`
    CodeSearchURL                    string `json:"code_search_url"`
    EmailsURL                        string `json:"emails_url"`
    EmojisURL                        string `json:"emojis_url"`
    EventsURL                        string `json:"events_url"`
    FeedsURL                         string `json:"feeds_url"`
    FollowersURL                     string `json:"followers_url"`
    FollowingURL                     string `json:"following_url"`
    GistsURL                         string `json:"gists_url"`
    HubURL                           string `json:"hub_url"`
    IssueSearchURL                   string `json:"issue_search_url"`
    IssuesURL                        string `json:"issues_url"`
    KeysURL                          string `json:"keys_url"`
    NotificationsURL                 string `json:"notifications_url"`
    OrganizationRepositoriesURL      string `json:"organization_repositories_url"`
    OrganizationURL                  string `json:"organization_url"`
    PublicGistsURL                   string `json:"public_gists_url"`
    RateLimitURL                     string `json:"rate_limit_url"`
    RepositoryURL                    string `json:"repository_url"`
    RepositorySearchURL              string `json:"repository_search_url"`
    CurrentUserRepositoriesURL       string `json:"current_user_repositories_url"`
    StarredURL                       string `json:"starred_url"`
    StarredGistsURL                  string `json:"starred_gists_url"`
    TeamURL                          string `json:"team_url"`
    UserURL                          string `json:"user_url"`
    UserOrganizationsURL             string `json:"user_organizations_url"`
    UserRepositoriesURL              string `json:"user_repositories_url"`
    UserSearchURL                    string `json:"user_search_url"`
}

func main() {
    github := &GithubRootStruct{}
    req := requester.NewRequester("https://api.github.com/")
    err := req.Get("", github)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(github.GistsURL)
}
```
