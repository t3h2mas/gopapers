# gopapers
Download wallpapers from Reddit using Go.
Save images from a subreddit to a folder named after the image posters.

## usage

### import the Reddit api wrapper
Gopaper uses the [Geddit](https://github.com/jzelinskie/geddit) Go wrapper for the Reddit API. You can get the library by running

 `go get github.com/jzelinskie/geddit`

## build it

 `go build wallpapers.go`

## run it

 `./wallpapers`

### optional flag

**set the subreddit**

 `./wallpapers -s=earthporn`

 s: subreddit

 `./wallpapers -a=100`

a: amount

 ## future features?
 * save to a directory to keep the app directory clean [ ]
 * implement go routines [ ]
 * handle different sorting options [ ]
