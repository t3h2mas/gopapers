# gopapers
Download wallpapers from Reddit using Go.
D

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
 * implement go routines [ ]
 * set the limit via a flag [ ]
 * handle different sorting options [ ]
