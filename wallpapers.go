package main

import (
	"flag"
	"fmt"
	"github.com/jzelinskie/geddit"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// author == the save to directory
// the path is added to the app directory
func grabImg(url, author string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	res.Body.Close()

	// filename / path setup
	parts := strings.Split(url, "/")
	filename := parts[len(parts)-1]

	err = os.MkdirAll(author, 0777)
	if err != nil {
		return err
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	fullpath := filepath.Join(wd, author, filename)

	fmt.Println("[!] PATH: ", fullpath)

	err = ioutil.WriteFile(fullpath, data, 0666)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// this returns a str pointer
	subreddit := flag.String("s", "wallpaper", "subreddit name")
	limit := flag.Int("a", 25, "amount of posts to download")
	flag.Parse()

	fmt.Println("Subreddit: ", *subreddit)

	session := geddit.NewSession("wallpapers.go v1")
	opts := geddit.ListingOptions{
		Limit: *limit,
	}

	// TODO: add a flag for sorting options
	submissions, err := session.SubredditSubmissions(*subreddit, geddit.HotSubmissions, opts)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, s := range submissions {
		// TODO: test s.URL for nil
		if strings.HasSuffix(s.URL, ".jpg") || strings.HasSuffix(s.URL, ".png") {
			fmt.Printf("Title: %s\nAuthor: %s\nUrl: %s\n\n", s.Title, s.Author, s.URL)
			err := grabImg(s.URL, s.Author)
			if err != nil {
				fmt.Println("ERROR: ", err)
			}
		}
	}
}
