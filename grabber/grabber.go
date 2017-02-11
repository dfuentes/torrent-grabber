package grabber

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"

	"github.com/dfuentes/torrent-grabber/config"
	"github.com/mmcdole/gofeed"
)

func Grab(config config.Config) {
	for _, feed := range config.Feeds {
		grabFeed(feed)
	}
}

func grabFeed(feedConfig config.Feed) {
	filters := compileFilters(feedConfig.Filters)

	feedParser := gofeed.NewParser()
	feed, err := feedParser.ParseURL(feedConfig.URL)
	if err != nil {
		log.Printf("Failed to fetch feed from url '%s': %s", feedConfig.URL, err)
		return
	}

	log.Printf("Grabbing from feed %s...", feed.Title)
	for _, item := range feed.Items {
		if !anyMatch(item.Title, filters) {
			continue
		}
		log.Printf("Downloading %s...", item.Title)
		downloadItem(item, feedConfig.OutputDir)
	}
}

func anyMatch(title string, filters []*regexp.Regexp) bool {
	for _, filter := range filters {
		if filter.MatchString(title) {
			return true
		}
	}
	return false
}

func compileFilters(filters []string) []*regexp.Regexp {
	compiled := []*regexp.Regexp{}

	for _, filter := range filters {
		re := regexp.MustCompile(filter)
		compiled = append(compiled, re)
	}
	return compiled
}

func downloadItem(item *gofeed.Item, path string) {
	if len(item.Enclosures) == 0 {
		log.Printf("item missing enclosure: %s", item.Title)
		return
	}

	enc := item.Enclosures[0]

	url, err := url.Parse(enc.URL)
	if err != nil {
		log.Printf("failed to parse url '%s': %s", enc.URL, err)
		return
	}

	filename := filepath.Base(url.Path)

	out, err := os.Create(filepath.Join(path, filename))
	if err != nil {
		log.Printf("failed to create file %s: %s", filename, err)
		return
	}
	defer out.Close()

	resp, err := http.Get(enc.URL)
	if err != nil {
		log.Printf("failed to fetch url '%s': %s", enc.URL, err)
		return
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
}
