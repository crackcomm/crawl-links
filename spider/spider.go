// TODO: depth

package spider

import (
	"net/url"

	"golang.org/x/net/context"

	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/crackcomm/crawl"
)

// Spider - Links spider.
type Spider struct {
	crawl.Crawler
	Output func(*Result) error
}

// Result - Links crawl result.
type Result struct {
	// Results - List of resulted links.
	Results []string `json:"results,omitempty"`
	// Source - URL of the source.
	Source string `json:"source,omitempty"`
}

// Links - Gather links from page.
var Links = "github.com/crackcomm/crawl-links/spider.Links"

// Register - Registers spider.
func (spider *Spider) Register() {
	spider.Crawler.Register(Links, spider.Links)
}

// Links - Crawls Links.
func (spider *Spider) Links(ctx context.Context, resp *crawl.Response) (err error) {
	res := &Result{
		Source: resp.URL().String(),
	}

	resMap := make(map[string]bool)

	// Crawl all links
	resp.Query().Find("a").Each(func(_ int, link *goquery.Selection) {
		href, _ := link.Attr("href")
		if len(href) == 0 {
			return
		}

		for _, prefix := range []string{"javascript:", "mailto:"} {
			if strings.HasPrefix(href, prefix) {
				return
			}
		}

		u, err := url.Parse(href)
		if err != nil {
			return
		}

		// Append resolved URL to results map
		uri := resp.URL().ResolveReference(u).String()
		resMap[uri] = true
	})

	for uri := range resMap {
		res.Results = append(res.Results, uri)
	}

	return spider.Output(res)
}
