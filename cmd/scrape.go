package cmd

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"

	stream "github.com/FlavioAmurrioCS/lazystream"
	"github.com/spf13/cobra"
	"golang.org/x/net/html"
)

var lookUp = ""
var scrapeCmd = &cobra.Command{
	Use:   "scrape URL",
	Short: "Scrape a website",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("URL is required")
		}
		_, err := url.Parse(args[0])
		if err != nil {
			return fmt.Errorf("invalid URL: %w", err)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := url.Parse(args[0])
		client := http.Client{}
		resp, err := client.Get(url.String())
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		doc, _ := html.Parse(resp.Body)
		links := stream.FromSlice(extractLinks(doc))
		links = stream.Map(links, ensureFullUrl(url)) // Ensure full URLs
		if lookUp != "" {
			r, _ := regexp.Compile(lookUp)
			links = links.Filter(r.MatchString)
		}
		stream.
			Sort(links).
			Reversed().
			ForEach(func(link string) {
				fmt.Println(link)
			})
	},
}

func init() {
	scrapeCmd.Flags().StringVarP(&lookUp, "filter", "f", "", "Search for a specific string")
}

func extractLinks(n *html.Node) []string {
	var links []string
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			}
		}
	}
	// Recursively traverse the node tree.
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		links = append(links, extractLinks(child)...)
	}
	return links
}

func ensureFullUrl(url *url.URL) func(string) string {
	return func(link string) string {
		if len(link) > 0 && link[0] == '/' {
			nurl := *url
			nurl.Path = link
			return nurl.String()
		}
		return url.JoinPath(link).String()
	}
}
