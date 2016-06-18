////////////////////////////////////////////////////////////////////////////
// Program: scrape_kijiji_car-r
// Purpose: Kijiji web scrapping using the Raw interface
// Authors: Tong Sun (c) 2016, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-shaper/shaper"
	"github.com/suntong/goscrape"
	"github.com/suntong/goscrape/extract"
	"github.com/suntong/goscrape/paginate"
)

func main() {
	config := &scrape.ScrapeConfig{
		DividePage: scrape.DividePageBySelector("div.info-container"),

		Pieces: []scrape.Piece{
			{Name: "Raw", Selector: "&", Extractor: extract.OuterHtml{}},
			{Name: "Price", Selector: "div.price", Extractor: extract.Text{}},
			// {Name: "Id", Selector: "div.title>a.attr('href')",
			// 	Extractor: extract.Regex{Regex: regexp.MustCompile(`.*/(\d+)$`)}},
			{Name: "Title", Selector: "div.title", Extractor: extract.Text{}},
			{Name: "Link", Selector: "div.title > a", Extractor: extract.Attr{Attr: "href"}},
			{Name: "Location", Selector: "div.location", Extractor: extract.Text{}},
			{Name: "Description", Selector: "div.description", Extractor: extract.Text{}},
			{Name: "Details", Selector: "div.details", Extractor: extract.Text{}},
			//{Name: "", Selector: "div.", Extractor: extract.Text{}},
		},

		Paginator:   paginate.BySelector("a[title='Next']", "href"),
		Opts:        scrape.ScrapeOptions{MaxPages: 1},
		PieceShaper: shaper.NewFilter().ApplyTrim().ApplyRegSpaces(),
	}

	scraper, err := scrape.New(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating scraper: %s\n", err)
		os.Exit(1)
	}

	results, err := scraper.ScrapeHTML(initHTML)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error scraping: %s\n", err)
		os.Exit(1)
	}

	for i, r := range results.Results[0] {
		// X: fmt.Printf("%d: %v\n", i, json.NewEncoder(os.Stdout).Encode(r))
		fmt.Printf("%d: %+v\n", i, r)
		s := r["Raw"].(*goquery.Selection)
		location := s.Find("div.location")
		location.Find("span").Remove()
		fmt.Printf("   %v\n", config.PieceShaper.Process(location.Text()))
		fmt.Printf("   %v\n", config.PieceShaper.Process(
			s.Find("div.price").Text()))
		fmt.Printf("   %v\n\n", s.Find("div.details").Text())
	}
}

var initHTML = `
<!DOCTYPE html>
<html id="reset">
<head>
    <title>Gps | Find Great Deals on Used and New Cars &amp; Trucks in Toronto (GTA) | Kijiji Classifieds</title>
</head>

<body id="PageSRP"
      class="en">
<div class="container-results large-images">
  <div class="extra-bar">
    <div class="float-right">
      Register for&nbsp;<strong>Kijiji Alerts&nbsp;</strong>
      <span class="tooltip-trigger" data-content="KijijiAlertTooltip" data-position-my="bottom right" data-hasqtip="5">[?]</span>
      <div class="button-open with-arrow alert-modal-trigger short">
        Sign Up<span class="arrow-next"></span>
      </div>
    </div>
    <div class="float-left">
      <div id="ForSaleByPageTipTrigger" class="page-tip-trigger" data-content="KijijiAlertTooltip" data-position-my="bottom center" data-position-at="top center">
        <div class="message">For Sale By:</div>
      </div>
    </div>
  </div>
  <div class="adsense-top-bar">Sponsored Links:</div>
  <hr class="adsense-bottom-bar">
  <div data-vip-url="/v-cars-trucks/city-of-toronto/2014-fiat-500-l-pickup-truck/1174926196" class="search-item regular-ad">
    <div class="left-col">
      <div class="watch watchlist-star p-vap-lnk-actn-addwtch" title="Click to add to My Watchlist" data-action="add" data-adid="1174926196"></div>
      <div class="image">
        <img src="http://i.ebayimg.com/00/s/NjAwWDgwMA==/z/YIYAAOSw-4BXZIoh/$_35.JPG" alt="2014 Fiat 500 L Pickup Truck">
      </div>
    </div>

    <div class="info">
      <div class="info-container">
        <div class="price">
          $14,999.00</div>

        <div class="title">
          <a href="/v-cars-trucks/city-of-toronto/2014-fiat-500-l-pickup-truck/1174926196" class="title enable-search-navigation-flag ">
            2014 Fiat 500 L Pickup Truck</a>
        </div>

        <div class="distance">
        </div>

        <div class="location">
          City of Toronto<span class="date-posted">&lt; 18 hours ago</span>
        </div>

        <div class="description">
          Looking to let go of this 2014 Fiat 500L. This truck is front wheel drive, 1.4 L turbo engine, fully loaded with navigation GPS Brand new tires. Comes certified with safety and emissions Private <div class="details">
            2200km | Automatic</div>
        </div>
      </div>
    </div>
  </div>
  <div data-vip-url="/v-cars-trucks/markham-york-region/cadillac-cts-2003-fully-loaded-gps-nav-best-offer-moving-sale/1174859573" class="search-item regular-ad">
    <div class="left-col">
      <div class="watch watchlist-star p-vap-lnk-actn-addwtch" title="Click to add to My Watchlist" data-action="add" data-adid="1174859573"></div>
      <div class="image">
        <img src="http://i.ebayimg.com/00/s/NjAwWDgwMA==/z/xIoAAOSw3YNXZEch/$_35.JPG" alt="Cadillac cts 2003 fully loaded gps nav (best offer moving sale)">
      </div>
    </div>

    <div class="info">
      <div class="info-container">
        <div class="price">
          $3,000.00</div>

        <div class="title">
          <a href="/v-cars-trucks/markham-york-region/cadillac-cts-2003-fully-loaded-gps-nav-best-offer-moving-sale/1174859573" class="title enable-search-navigation-flag ">
            Cadillac cts 2003 fully loaded gps nav (best offer moving sale)</a>
        </div>

        <div class="distance">
        </div>

        <div class="location">
          Markham / York Region<span class="date-posted">&lt; 22 hours ago</span>
        </div>

        <div class="description">
          Moving overseas fire sale. Check out my other ads! Best offer/ first real person with cash, not PayPal. All reasonable offers considered, low balls ignored. Luxury without breaking the bank. Fully <div class="details">
            177km | Automatic</div>
        </div>
      </div>
    </div>
  </div>
  <div data-vip-url="/v-cars-trucks/city-of-toronto/2006-great-condition-honda-odyssey-minivan-fully-loaded-gps-dv/1174714358" class="search-item regular-ad">
    <div class="left-col">
      <div class="watch watchlist-star p-vap-lnk-actn-addwtch" title="Click to add to My Watchlist" data-action="add" data-adid="1174714358"></div>
      <div class="image">
        <img src="http://i.ebayimg.com/00/s/ODAwWDQ1MA==/z/2i8AAOSwNuxXY3Ne/$_35.JPG" alt="2006 Great condition Honda Odyssey Minivan, fully loaded gps, dv">
      </div>
    </div>

    <div class="info">
      <div class="info-container">
        <div class="price">
          $10,500.00</div>

        <div class="title">
          <a href="/v-cars-trucks/city-of-toronto/2006-great-condition-honda-odyssey-minivan-fully-loaded-gps-dv/1174714358" class="title enable-search-navigation-flag ">
            2006 Great condition Honda Odyssey Minivan, fully loaded gps, dv</a>
        </div>

        <div class="distance">
        </div>

        <div class="location">
          City of Toronto<span class="date-posted">Yesterday</span>
        </div>

        <div class="description">
          A mint condition 2006 fully loaded honda odyssey for sale. Leather inside, heated seats, navi, dvd.<div class="details">
            20700km | Automatic</div>
        </div>
      </div>
    </div>
  </div>
  <div data-vip-url="/v-cars-trucks/mississauga-peel-region/2004-mazda-rx-8-for-sale-below-value/1173813154" class="search-item regular-ad">
    <div class="left-col">
      <div class="watch watchlist-star p-vap-lnk-actn-addwtch" title="Click to add to My Watchlist" data-action="add" data-adid="1173813154"></div>
      <div class="image">
        <img src="http://i.ebayimg.com/00/s/NjAwWDgwMA==/z/viMAAOSwnNBXX3My/$_35.JPG" alt="2004 Mazda RX-8 For Sale Below Value!!!!">
      </div>
    </div>

    <div class="info">
      <div class="info-container">
        <div class="price">
          $4,650.00</div>

        <div class="title">
          <a href="/v-cars-trucks/mississauga-peel-region/2004-mazda-rx-8-for-sale-below-value/1173813154" class="title enable-search-navigation-flag ">
            2004 Mazda RX-8 For Sale Below Value!!!!</a>
        </div>

        <div class="distance">
        </div>

        <div class="location">
          Mississauga / Peel Region<span class="date-posted">14/06/2016</span>
        </div>

        <div class="description">
          For Sale 2004 Mazda RX8 Only 89 900Km Automatic Transmission Paddle Shifters, Leather Seats, GPS , Sun Roof 6-CD Changer Bose Speakers, Fully Loaded, New Battery, Well Maintained. Call me for more <div class="details">
            89941km | Automatic</div>
        </div>
      </div>
    </div>
  </div>
  <div data-vip-url="/v-cars-trucks/mississauga-peel-region/hyundai-genesis-sedan-2012-3-8l-v6/1173790662" class="search-item regular-ad">
    <div class="left-col">
      <div class="watch watchlist-star p-vap-lnk-actn-addwtch" title="Click to add to My Watchlist" data-action="add" data-adid="1173790662"></div>
      <div class="image">
        <img src="http://i.ebayimg.com/00/s/NDUwWDgwMA==/z/6FYAAOSwjXRXX1sX/$_35.JPG" alt="Hyundai Genesis sedan 2012 !! 3.8L v6">
      </div>
    </div>

    <div class="info">
      <div class="info-container">
        <div class="price">
          $16,500.00</div>

        <div class="title">
          <a href="/v-cars-trucks/mississauga-peel-region/hyundai-genesis-sedan-2012-3-8l-v6/1173790662" class="title enable-search-navigation-flag ">
            Hyundai Genesis sedan 2012 !! 3.8L v6</a>
        </div>

        <div class="distance">
        </div>

        <div class="location">
          Mississauga / Peel Region<span class="date-posted">13/06/2016</span>
        </div>

        <div class="description">
          I have for sale Genesis 2012 very clean car and premium car all automatic heated and cold seat rear and front all features GPS Rear view camera Tilt steering Tilt mirror on reverse Brembo breaks Hid <div class="details">
            76500km | Automatic</div>
        </div>
      </div>
    </div>
  </div>
  <div data-vip-url="/v-cars-trucks/oshawa-durham-region/mazda-3-gx-sport-2015-lease-assignment/1173446434" class="search-item regular-ad">
    <div class="left-col">
      <div class="watch watchlist-star p-vap-lnk-actn-addwtch" title="Click to add to My Watchlist" data-action="add" data-adid="1173446434"></div>
      <div class="image">
        <img src="http://i.ebayimg.com/00/s/ODAwWDQ1MQ==/z/kPUAAOSwepJXXeqo/$_35.JPG" alt="Mazda 3 GX Sport 2015 Lease Assignment">
      </div>
    </div>

    <div class="info">
      <div class="info-container">
        <div class="price">
          $420.00</div>

        <div class="title">
          <a href="/v-cars-trucks/oshawa-durham-region/mazda-3-gx-sport-2015-lease-assignment/1173446434" class="title enable-search-navigation-flag ">
            Mazda 3 GX Sport 2015 Lease Assignment</a>
        </div>

        <div class="distance">
        </div>

        <div class="location">
          Oshawa / Durham Region<span class="date-posted">12/06/2016</span>
        </div>

        <div class="description">
          Auto 6-gear 52KM Tinted power windows Heated front seats, mirrors Free 4 x snow tires and Mazda carpets (brand new) Back up camera Gps navigation (upgraded) Buy back $10000 approx. Hatchback Accident <div class="details">
            52000km | Automatic</div>
        </div>
      </div>
    </div>
  </div>
  <div data-vip-url="/v-cars-trucks/markham-york-region/2013-nissan-altima-sv-37000km-navi-sunroof-owner/1173443054" class="search-item regular-ad">
    <div class="left-col">
      <div class="watch watchlist-star p-vap-lnk-actn-addwtch" title="Click to add to My Watchlist" data-action="add" data-adid="1173443054"></div>
      <div class="image">
        <img src="http://i.ebayimg.com/00/s/NjAwWDgwMA==/z/OsQAAOSwbYZXXejk/$_35.JPG" alt="2013 Nissan Altima SV 37000km Navi Sunroof (Owner)">
      </div>
    </div>

    <div class="info">
      <div class="info-container">
        <div class="price">
          $16,900.00</div>

        <div class="title">
          <a href="/v-cars-trucks/markham-york-region/2013-nissan-altima-sv-37000km-navi-sunroof-owner/1173443054" class="title enable-search-navigation-flag ">
            2013 Nissan Altima SV 37000km Navi Sunroof (Owner)</a>
        </div>

        <div class="distance">
        </div>

        <div class="location">
          Markham / York Region<span class="date-posted">12/06/2016</span>
        </div>

        <div class="description">
          Selling our 2013 Nissan Altima SV The car is fully loaded in very good condition, only 37000km mileage. * 37000km mileage * 7'' in dash GPS navigation and back up camera * remote start/keyless <div class="details">
            37000km | Automatic</div>
        </div>
      </div>
    </div>
  </div>
  <div data-vip-url="/v-cars-trucks/city-of-toronto/2001-toyota-camry-other/1173211799" class="search-item regular-ad">
    <div class="left-col">
      <div class="watch watchlist-star p-vap-lnk-actn-addwtch" title="Click to add to My Watchlist" data-action="add" data-adid="1173211799"></div>
      <div class="image">
        <img src="http://i.ebayimg.com/00/s/NjAwWDgwMA==/z/BjUAAOSwtJZXXMxR/$_35.JPG" alt="2001 Toyota Camry Other">
      </div>
    </div>

    <div class="info">
      <div class="info-container">
        <div class="price">
          $1,900.00</div>

        <div class="title">
          <a href="/v-cars-trucks/city-of-toronto/2001-toyota-camry-other/1173211799" class="title enable-search-navigation-flag ">
            2001 Toyota Camry Other</a>
        </div>

        <div class="distance">
        </div>

        <div class="location">
          City of Toronto<span class="date-posted">12/06/2016</span>
        </div>

        <div class="description">
          Fully Loaded. Call me @ 4168786775 New engine / Transmission @ 170KM Auto Start remote Alarm Sun Roof New Tires New Battery Leather Seats A/C GPS CD player Lady Driven Much more........Mint <div class="details">
            31000km | Automatic</div>
        </div>
      </div>
    </div>
  </div>
  <div data-vip-url="/v-cars-trucks/city-of-toronto/2013-ford-fusion-se-2-0t/1136516476" class="search-item regular-ad">
    <div class="left-col">
      <div class="watch watchlist-star p-vap-lnk-actn-addwtch" title="Click to add to My Watchlist" data-action="add" data-adid="1136516476"></div>
      <div class="image">
        <img src="http://i.ebayimg.com/00/s/NjAwWDgwMA==/z/K78AAOSwX~dWrCxY/$_35.JPG" alt="2013 Ford Fusion SE 2.0T">
      </div>
    </div>

    <div class="info">
      <div class="info-container">
        <div class="price">
          $15,500.00</div>

        <div class="title">
          <a href="/v-cars-trucks/city-of-toronto/2013-ford-fusion-se-2-0t/1136516476" class="title enable-search-navigation-flag ">
            2013 Ford Fusion SE 2.0T</a>
        </div>

        <div class="distance">
        </div>

        <div class="location">
          City of Toronto<span class="date-posted">11/06/2016</span>
        </div>

        <div class="description">
          2013 FORD Fusion SE 2.0T GPS and 18' alloy OEM wheels, leather seats Well equipped. 37000km very low mileage.<div class="details">
            37000km | Automatic</div>
        </div>
      </div>
    </div>
  </div>
  <div data-vip-url="/v-cars-trucks/city-of-toronto/2004-mercedes-benz-e500-4matic/1172928088" class="search-item regular-ad">
    <div class="left-col">
      <div class="watch watchlist-star p-vap-lnk-actn-addwtch" title="Click to add to My Watchlist" data-action="add" data-adid="1172928088"></div>
      <div class="image">
        <img src="http://i.ebayimg.com/00/s/NDUwWDgwMA==/z/0D8AAOSwepJXW2sy/$_35.JPG" alt="2004 Mercedes-Benz E500 4matic">
      </div>
    </div>

    <div class="info">
      <div class="info-container">
        <div class="price">
          $7,000.00</div>

        <div class="title">
          <a href="/v-cars-trucks/city-of-toronto/2004-mercedes-benz-e500-4matic/1172928088" class="title enable-search-navigation-flag ">
            2004 Mercedes-Benz E500 4matic</a>
        </div>

        <div class="distance">
        </div>

        <div class="location">
          City of Toronto<span class="date-posted">10/06/2016</span>
        </div>

        <div class="description">
          2004 Mercedes Benz E500 AWD Mercedes-Benz 4matic fully fully loaded Very clean car Lady driven only With RECENT UPDATED GPS 170km original Must see Call Mike (416) 433-5508 As is $7000<div class="details">
            170km | Automatic</div>
        </div>
      </div>
    </div>
  </div>
  <div data-vip-url="/v-cars-trucks/city-of-toronto/2015-toyota-corolla-eco-premium-sedan/1172812523" class="search-item regular-ad">
    <div class="left-col">
      <div class="watch watchlist-star p-vap-lnk-actn-addwtch" title="Click to add to My Watchlist" data-action="add" data-adid="1172812523"></div>
      <div class="image">
        <img src="http://i.ebayimg.com/00/s/NjAwWDgwMA==/z/y0sAAOSwzJ5XWwHB/$_35.JPG" alt="2015 Toyota Corolla eco premium Sedan">
      </div>
    </div>

    <div class="info">
      <div class="info-container">
        <div class="price">
          $17,500.00</div>

        <div class="title">
          <a href="/v-cars-trucks/city-of-toronto/2015-toyota-corolla-eco-premium-sedan/1172812523" class="title enable-search-navigation-flag ">
            2015 Toyota Corolla eco premium Sedan</a>
        </div>

        <div class="distance">
        </div>

        <div class="location">
          City of Toronto<span class="date-posted">10/06/2016</span>
        </div>

        <div class="description">
          hightest top model corolla LE2015. 4 brand new michelin winter tires. bluetooth,GPS,real leather seat,all glass cover protection,sunroof auto, the best clean and protect body and botton anti <div class="details">
            30000km | Automatic</div>
        </div>
      </div>
    </div>
  </div>
  <div data-vip-url="/v-cars-trucks/mississauga-peel-region/2010-dodge-charger-sxt-sedan/1172154078" class="search-item regular-ad">
    <div class="left-col">
      <div class="watch watchlist-star p-vap-lnk-actn-addwtch" title="Click to add to My Watchlist" data-action="add" data-adid="1172154078"></div>
      <div class="image">
        <img src="http://i.ebayimg.com/00/s/NDUwWDgwMA==/z/If0AAOSwuhhXWCcF/$_35.JPG" alt="2010 Dodge Charger sxt Sedan">
      </div>
    </div>

    <div class="info">
      <div class="info-container">
        <div class="price">
          $12,999.00</div>

        <div class="title">
          <a href="/v-cars-trucks/mississauga-peel-region/2010-dodge-charger-sxt-sedan/1172154078" class="title enable-search-navigation-flag ">
            2010 Dodge Charger sxt Sedan</a>
        </div>

        <div class="distance">
        </div>

        <div class="location">
          Mississauga / Peel Region<span class="date-posted">08/06/2016</span>
        </div>

        <div class="description">
          VEHICLE OPTIONS: Air Conditioning Navigation system (GPS) Power locks Power trunk Remote keyless entry Power windows CD player Leather seats Power seats Alloy wheels LOW KMS... 3.5 litre V-6 <div class="details">
            86500km | Automatic</div>
        </div>
      </div>
    </div>
  </div>
  <div data-vip-url="/v-cars-trucks/city-of-toronto/2013-mazda-mazda3-sedan/1172081358" class="search-item regular-ad">
    <div class="left-col">
      <div class="watch watchlist-star p-vap-lnk-actn-addwtch" title="Click to add to My Watchlist" data-action="add" data-adid="1172081358"></div>
      <div class="image">
        <img src="http://i.ebayimg.com/00/s/ODAwWDYwMA==/z/bTQAAOSwOVpXV4jJ/$_35.JPG" alt="2013 Mazda Mazda3 Sedan">
      </div>
    </div>

    <div class="info">
      <div class="info-container">
        <div class="price">
          $8,400.00</div>

        <div class="title">
          <a href="/v-cars-trucks/city-of-toronto/2013-mazda-mazda3-sedan/1172081358" class="title enable-search-navigation-flag ">
            2013 Mazda Mazda3 Sedan</a>
        </div>

        <div class="distance">
        </div>

        <div class="location">
          City of Toronto<span class="date-posted">07/06/2016</span>
        </div>

        <div class="description">
          6472730214 nice car coming with gps back up camera<div class="details">
            51000km | Automatic</div>
        </div>
      </div>
    </div>
  </div>
  <div data-vip-url="/v-cars-trucks/city-of-toronto/2015-toyota-corolla-evo-premium-sedan/1171977906" class="search-item regular-ad">
    <div class="left-col">
      <div class="watch watchlist-star p-vap-lnk-actn-addwtch" title="Click to add to My Watchlist" data-action="add" data-adid="1171977906"></div>
      <div class="image">
        <img src="http://i.ebayimg.com/00/s/ODAwWDYwMA==/z/1wsAAOSwqBJXVy2y/$_35.JPG" alt="2015 Toyota Corolla Evo premium Sedan">
      </div>
    </div>

    <div class="info">
      <div class="info-container">
        <div class="price">
          $17,900.00</div>

        <div class="title">
          <a href="/v-cars-trucks/city-of-toronto/2015-toyota-corolla-evo-premium-sedan/1171977906" class="title enable-search-navigation-flag ">
            2015 Toyota Corolla Evo premium Sedan</a>
        </div>

        <div class="distance">
        </div>

        <div class="location">
          City of Toronto<span class="date-posted">07/06/2016</span>
        </div>

        <div class="description">
          highest top model with 4 brand new michelin winter tires coming with bluetooth,GPS,,real leather seat,sterling auto up and down adjustable, auto sunroof,all the glasses coveing protection dark dime. <div class="details">
            30000km | Automatic</div>
        </div>
      </div>
    </div>
  </div>
  <div data-vip-url="/v-cars-trucks/markham-york-region/2-year-lease-takeover-2014-chevrolet-cruze-2lt-rally-sport-rs/1170412657" class="search-item regular-ad">
    <div class="left-col">
      <div class="watch watchlist-star p-vap-lnk-actn-addwtch" title="Click to add to My Watchlist" data-action="add" data-adid="1170412657"></div>
      <div class="image">
        <img src="http://i.ebayimg.com/00/s/NjQwWDgwMA==/z/x9QAAOSw9eVXUDJG/$_35.JPG" alt="2 Year Lease Takeover - 2014 Chevrolet Cruze 2LT Rally Sport, RS">
      </div>
    </div>

    <div class="info">
      <div class="info-container">
        <div class="price">
          $182.00</div>

        <div class="title">
          <a href="/v-cars-trucks/markham-york-region/2-year-lease-takeover-2014-chevrolet-cruze-2lt-rally-sport-rs/1170412657" class="title enable-search-navigation-flag ">
            2 Year Lease Takeover - 2014 Chevrolet Cruze 2LT Rally Sport, RS</a>
        </div>

        <div class="distance">
        </div>

        <div class="location">
          Markham / York Region<span class="date-posted">02/06/2016</span>
        </div>

        <div class="description">
          2014 Chevy Cruze RS Package -Fully Loaded INCENTIVE: Buyer will receive a mini Ipad and $500.00 for gas! Body Kit, Fog Lights, Sun-roof, GPS, Back-up Camera, Heated Leather Seats, Navigation, XM <div class="details">
            80000km | Automatic</div>
        </div>
      </div>
    </div>
  </div>
  <div data-vip-url="/v-cars-trucks/oshawa-durham-region/2010-nissan-sentra-se-r-sedan/1170021437" class="search-item regular-ad">
    <div class="left-col">
      <div class="watch watchlist-star p-vap-lnk-actn-addwtch" title="Click to add to My Watchlist" data-action="add" data-adid="1170021437"></div>
      <div class="image">
        <img src="http://i.ebayimg.com/00/s/NTkyWDgwMA==/z/Xj0AAOSwOVpXTiDg/$_35.JPG" alt="2010 Nissan Sentra SE-R Sedan">
      </div>
    </div>

    <div class="info">
      <div class="info-container">
        <div class="price">
          $11,400.00</div>

        <div class="title">
          <a href="/v-cars-trucks/oshawa-durham-region/2010-nissan-sentra-se-r-sedan/1170021437" class="title enable-search-navigation-flag ">
            2010 Nissan Sentra SE-R Sedan</a>
        </div>

        <div class="distance">
        </div>

        <div class="location">
          Oshawa / Durham Region<span class="date-posted">31/05/2016</span>
        </div>

        <div class="description">
          Super clean and well kept 177hp 4cyl, 2.5L CVT SE-R. Second owner, no accidents. Fully loaded with GPS navigation, paddle shifters, backup camera, sunroof, auto dimming rear view mirror and after <div class="details">
            65800km | Automatic</div>
        </div>
      </div>
    </div>
  </div>
  <div data-vip-url="/v-cars-trucks/city-of-toronto/2015-chevrolet-cruze-ltz-sedan/1169439739" class="search-item regular-ad">
    <div class="left-col">
      <div class="watch watchlist-star p-vap-lnk-actn-addwtch" title="Click to add to My Watchlist" data-action="add" data-adid="1169439739"></div>
      <div class="image">
        <img src="http://i.ebayimg.com/00/s/ODAwWDgwMA==/z/RcIAAOSwR5dXS4nP/$_35.JPG" alt="2015 Chevrolet Cruze LTZ Sedan">
      </div>
    </div>

    <div class="info">
      <div class="info-container">
        <div class="price">
          $13,300.00</div>

        <div class="title">
          <a href="/v-cars-trucks/city-of-toronto/2015-chevrolet-cruze-ltz-sedan/1169439739" class="title enable-search-navigation-flag ">
            2015 Chevrolet Cruze LTZ Sedan</a>
        </div>

        <div class="distance">
        </div>

        <div class="location">
          City of Toronto<span class="date-posted">29/05/2016</span>
        </div>

        <div class="description">
          I would like to sell my Chevy Cruze LTZ RS package with limited options: GPS Backup Camera Sunroof Keyless entry Push to start Leather sets Heated sets Tinted windows Fully loaded Top of the line Car <div class="details">
            11000km | Automatic</div>
        </div>
      </div>
    </div>
  </div>
  <div data-vip-url="/v-cars-trucks/city-of-toronto/audi-q7-4-2-for-11-200-firm/1168333341" class="search-item regular-ad">
    <div class="left-col">
      <div class="watch watchlist-star p-vap-lnk-actn-addwtch" title="Click to add to My Watchlist" data-action="add" data-adid="1168333341"></div>
      <div class="image">
        <img src="http://i.ebayimg.com/00/s/NjAwWDYwMA==/z/uogAAOSwa81XRk-G/$_35.JPG" alt="Audi Q7 4.2 for $11.200 firm ">
      </div>
    </div>

    <div class="info">
      <div class="info-container">
        <div class="price">
          $11,200.00</div>

        <div class="title">
          <a href="/v-cars-trucks/city-of-toronto/audi-q7-4-2-for-11-200-firm/1168333341" class="title enable-search-navigation-flag ">
            Audi Q7 4.2 for $11.200 firm </a>
        </div>

        <div class="distance">
        </div>

        <div class="location">
          City of Toronto<span class="date-posted">25/05/2016</span>
        </div>

        <div class="description">
          Hello there I am selling Audi Q7 it has everything that a car needs. It has new brakes and tires. The car has GPS and backup camera.<div class="details">
            200km | Automatic</div>
        </div>
      </div>
    </div>
  </div>
  <div data-vip-url="/v-cars-trucks/oakville-halton-region/2014-jeep-compass/1145692677" class="search-item regular-ad">
    <div class="left-col">
      <div class="watch watchlist-star p-vap-lnk-actn-addwtch" title="Click to add to My Watchlist" data-action="add" data-adid="1145692677"></div>
      <div class="image">
        <img src="http://i.ebayimg.com/00/s/NzU0WDQwMQ==/z/RowAAOSwh-1W3Z-T/$_35.JPG" alt="2014 Jeep Compass">
      </div>
    </div>

    <div class="info">
      <div class="info-container">
        <div class="price">
          $15,800.00</div>

        <div class="title">
          <a href="/v-cars-trucks/oakville-halton-region/2014-jeep-compass/1145692677" class="title enable-search-navigation-flag ">
            2014 Jeep Compass</a>
        </div>

        <div class="distance">
        </div>

        <div class="location">
          Oakville / Halton Region<span class="date-posted">24/05/2016</span>
        </div>

        <div class="description">
          Only 15,800, great Offer!!! 57000 Kilometers, only sell for $16000! Lots to offer! GPS, Reverse video camera, Bluetooth, CD, SD card, pad, front seat cover!!<div class="details">
            57000km | Automatic</div>
        </div>
      </div>
    </div>
  </div>
  <div data-vip-url="/v-cars-trucks/city-of-toronto/2013-honda-civic-touring-sedan/1167900714" class="search-item regular-ad">
    <div class="left-col">
      <div class="watch watchlist-star p-vap-lnk-actn-addwtch" title="Click to add to My Watchlist" data-action="add" data-adid="1167900714"></div>
      <div class="image">
        <img src="http://i.ebayimg.com/00/s/NTMzWDgwMA==/z/2P0AAOSw7ehXRINl/$_35.JPG" alt="2013 Honda Civic TOURING Sedan">
      </div>
    </div>

    <div class="info">
      <div class="info-container">
        <div class="price">
          $17,999.00</div>

        <div class="title">
          <a href="/v-cars-trucks/city-of-toronto/2013-honda-civic-touring-sedan/1167900714" class="title enable-search-navigation-flag ">
            2013 Honda Civic TOURING Sedan</a>
        </div>

        <div class="distance">
        </div>

        <div class="location">
          City of Toronto<span class="date-posted">24/05/2016</span>
        </div>

        <div class="description">
          Hello, Am selling 2013 Honda Civic Touring fully loaded (GPS Navigartion,sunroof,back up camera, heated seats leather interior, and much more, no accsidents and will provide car proof. Looks brand <div class="details">
            42000km | Automatic</div>
        </div>
      </div>
    </div>
  </div>
  <div class="adsense-top-bar">Sponsored Links:</div>
</div>
</body>
</html>
`

/*

Output:

0: map[Location:City of Toronto< 18 hours ago Description:Looking to let go of this 2014 Fiat 500L. This truck is front wheel drive, 1.4 L turbo engine, fully loaded with navigation GPS Brand new tires. Comes certified with safety and emissions Private 2200km | Automatic Details:2200km | Automatic Raw:0xc8201478c0 Price:$14,999.00 Title:2014 Fiat 500 L Pickup Truck Link:/v-cars-trucks/city-of-toronto/2014-fiat-500-l-pickup-truck/1174926196]
   City of Toronto
   $14,999.00

            2200km | Automatic

1: map[Price:$3,000.00 Title:Cadillac cts 2003 fully loaded gps nav (best offer moving sale) Link:/v-cars-trucks/markham-york-region/cadillac-cts-2003-fully-loaded-gps-nav-best-offer-moving-sale/1174859573 Location:Markham / York Region< 22 hours ago Description:Moving overseas fire sale. Check out my other ads! Best offer/ first real person with cash, not PayPal. All reasonable offers considered, low balls ignored. Luxury without breaking the bank. Fully 177km | Automatic Details:177km | Automatic Raw:0xc8201478f0]
   Markham / York Region
   $3,000.00

            177km | Automatic

...

19: map[Location:City of Toronto24/05/2016 Description:Hello, Am selling 2013 Honda Civic Touring fully loaded (GPS Navigartion,sunroof,back up camera, heated seats leather interior, and much more, no accsidents and will provide car proof. Looks brand 42000km | Automatic Details:42000km | Automatic Raw:0xc820147c50 Price:$17,999.00 Title:2013 Honda Civic TOURING Sedan Link:/v-cars-trucks/city-of-toronto/2013-honda-civic-touring-sedan/1167900714]
   City of Toronto
   $17,999.00

            42000km | Automatic

*/
