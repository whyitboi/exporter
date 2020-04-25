package main

import (
	"fmt"
	interfaces "week3"
	"week3/exporter"
	"week3/facebook"
	"week3/linkedin"
	"week3/twitter"
)

func main() {

	fb := new(facebook.Facebook)
	twtr := new(twitter.Twitter)
	lnkdin := new(linkedin.Linkedin)
	//b := new(exporter.Message)

	exporter.Export(fb, "fbdata")
	exporter.Export(twtr, "twtrdata")
	exporter.Export(lnkdin, "lnkdindata")

	fmt.Println("Data Exported")

}

//ScrollFeeds prints the social media feeds
func ScrollFeeds(platforms ...interfaces.SocialMedia) {
	for _, sm := range platforms {
		for _, fd := range sm.Feed() {
			fmt.Println(fd)
		}
		fmt.Println("============================================")
	}

}
