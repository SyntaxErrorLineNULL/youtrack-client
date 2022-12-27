package yootrack

import (
	"context"
	"fmt"
	"net/url"
)

type Photos struct {
	Albums []Album `json:"enabled"`
}

type Album struct {
	AlbumID      string `json:"albumId"`
	ID           string `json:"id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	ThumbnailUrl string `json:"thumbnailUrl"`
}

func main() {
	ctx := context.Background()
	opts := RequestOptions{
		Headers: []RequestHeader{
			AddHeader("Content-Type", "application/json"),
		},
	}

	link, err := url.Parse("https://jsonplaceholder.typicode.com/photos")
	if err != nil {
		fmt.Print("ERR: ", err)
	}

	d, err := Get[Photos](ctx, link, opts)
	if err != nil {
		fmt.Print("ERR: ", err)
	}

	fmt.Print(d)
}
