// You can edit this code!
// Click here and start typing.
package main

import (
	"context"
	ad_listing "ct-backend-course-baonguyen/clients/ad-listing"
	"log"
	"os"
)

func main() {

	// TODO #5 setup output for logger to write it to a file
	file, err := os.Create("myLogFile.log")

	if err != nil {
		log.Fatal("Failed to create log file:", err)
	}

	defer file.Close()

	logger := log.New(file, "file log: ", 3)
	
	c := ad_listing.NewClient(
		ad_listing.WithBaseUrl(ad_listing.BaseUrl),
		ad_listing.WithRetryTimes(3),
		ad_listing.WithLogger(logger),
	)
	ads, err := c.GetAdByCate(context.TODO(), ad_listing.CatePty)
	if err != nil {
		panic("GetAdByCate " + err.Error())
	}

	logger.Printf("Number of Ads: %v", ads.Total)
}