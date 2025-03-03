package main

import (
	"fmt"
	"io"
	"net/url"
	"regexp"
	"strings"

	"github.com/charmbracelet/log"
)

type GithubSource struct {
	domain  string
	org     string
	repo    string
	version string
}

func NewGithubSourceFromUrl(link string) *GithubSource {
	// https://github.com/
	// junegunn/fzf/releases/tag/v0.60.0
	e, _ := url.Parse(link)
	domain := e.Scheme + "://" + e.Host
	path_parts := strings.Split(strings.TrimLeft(e.Path, "/"), "/")
	version := "latest"
	if len(path_parts) >= 5 && path_parts[2] == "releases" && path_parts[3] == "tag" {
		version = path_parts[4]
	}
	org, repo := path_parts[0], path_parts[1]
	return &GithubSource{domain, org, repo, version}
}

func (githubSource *GithubSource) links() []string {
	client := getGithubClient()
	version := githubSource.version
	if version == "latest" {
		// client.Get()
		version = "latest"
	}
	// https://github.com/junegunn/fzf/releases/expanded_assets/v0.59.0
	url := fmt.Sprintf("%s/%s/%s/releases/expanded_assets/%s", githubSource.domain, githubSource.org, githubSource.repo, version)
	log.Info(url)
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	// href="/junegunn/fzf/releases/download/v0.60.0/fzf-0.60.0-linux_loong64.tar.gz"
	link_pattern := fmt.Sprintf(`/%s/%s/releases/download/%s/[^"]+`, githubSource.org, githubSource.repo, version)
	re := regexp.MustCompile(link_pattern)

	finds := re.FindAllString(string(body), -1)
	ret := []string{}
	for _, find := range finds {
		ret = append(ret, githubSource.domain+find)
	}
	return ret
}
