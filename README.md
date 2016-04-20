# gopapers
Download wallpapers from Reddit using Go.

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

 `./wallpapers -sub=earthporn`


 ## future features?
 * save to a directory to keep the app directory clean [ ]
 * handle different sorting options [ ]
 * implement go routines [ ]
