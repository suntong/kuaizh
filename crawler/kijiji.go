package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

/*

<div class="info-container">
  <div class="price">$14,900.00</div>

  <div class="title">
    <a href="/v-cars-trucks/mississauga-peel-region/2014-dodge-journey-sxt-bluetooth-gps-nav-heated-backup-sensor/1173567915?src=topAdSearch" class="title enable-search-navigation-flag ">
    2014 Dodge Journey SXT Bluetooth GPS Nav Heated Backup Sensor</a>
  </div>

  <div class="distance">
  </div>

  <div class="location">
    Mississauga / Peel Region<span class="date-posted">13/06/2016</span>
  </div>

  <div class="description">
    I am selling my 2014 Dodge Journey SXT... <div class="details">
    47000km | Automatic</div>
  </div>
</div>

*/

func main() {
	doc, _ := goquery.NewDocument("http://www.kijiji.ca/b-cars-trucks/gta-greater-toronto-area/gps/k0c174l1700272?ad=offering&price=__18000&kilometers=__90000&transmission=2&for-sale-by=ownr")
	doc.Find("div.info-container").Each(func(_ int, s *goquery.Selection) {
		title := s.Find("div.title")
		location := s.Find("div.location")
		description := s.Find("div.description")
		titleHtml, _ := title.Html()
		locationHtml, _ := location.Html()
		descriptionHtml, _ := description.Html()
		fmt.Printf("%v\n%v\n%v\n%v\n\n=>\n",
			titleHtml, locationHtml, location.Text(), descriptionHtml)
		ta := title.Find("a").Remove()
		taHtml, _ := ta.Html()
		titleHtml, _ = title.Html()
		location.Find("span").Remove()
		locationHtml, _ = location.Html()
		description.Find("div").Remove()
		descriptionHtml, _ = description.Html()
		fmt.Printf("%v\n%v\n%v\n%v\n\n---\n",
			taHtml, titleHtml, locationHtml, descriptionHtml)
	})
}
