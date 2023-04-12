package rssfeed

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/YWJSonic/ycore/driver/writefile"
	"github.com/YWJSonic/ycore/module/mylog"

	"github.com/gorilla/feeds"
	"github.com/mmcdole/gofeed"
	"github.com/mmcdole/gofeed/rss"
)

func NewFeed(feedsDatas []*feeds.Item) {
	feed := &feeds.Feed{
		Title: "591 最新兩房資訊",
		Link:  &feeds.Link{Href: "http://34.81.111.226:8080/rss.xml"},
		// Description: "Description of your feeds",
		// Author:  &feeds.Author{Name: "author name"},
		Created: time.Now(),
	}

	var feedItems []*feeds.Item
	for _, feedData := range feedsDatas {
		feedItems = append(feedItems,
			&feeds.Item{
				Id:          feedData.Id,
				Title:       feedData.Title,
				Link:        feedData.Link,
				Description: feedData.Description,
				Created:     time.Now(),
			})
	}
	feed.Items = feedItems

	rssFeed := (&feeds.Rss{Feed: feed}).RssFeed()
	mylog.Info("writefile.ProtocolXml")
	err := writefile.ProtocolXml(`/home/bunkeryangtw/nginxhome/rss.xml`, rssFeed.FeedXml(), os.ModePerm)
	if err != nil {
		mylog.Errorf("[Rss] err: %v", err)
	}
}

func ReadRssByString(feedData string) {
	// feedData := `<rss version="2.0">
	// <channel>
	// <webMaster>example@site.com (Example Name)</webMaster>
	// </channel>
	// </rss>`
	fp := rss.Parser{}
	rssFeed, err := fp.Parse(strings.NewReader(feedData))
	fmt.Println(rssFeed, err)
}

func ReadRssByUrl(url string) {
	// feed, _ := fp.ParseURL("http://feeds.twit.tv/twit.xml")
	fp := gofeed.NewParser()
	rssFeed, err := fp.ParseURL(url)
	fmt.Println(rssFeed, err)
}

func ReadRssByFilePath(path string) {
	// file, _ := os.Open("/path/to/a/file.xml")
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	fp := gofeed.NewParser()
	rssFeed, err := fp.Parse(file)
	fmt.Println(rssFeed, err)
}

func ReadRssByUrlWithTimeout(ctx context.Context, url string, timeout time.Duration) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURLWithContext(url, ctx)
	fmt.Println(feed, err)
}
