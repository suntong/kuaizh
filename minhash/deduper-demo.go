////////////////////////////////////////////////////////////////////////////
// Porgram: deduper-demo.go
// Purpose: Deduper Minhash LSH Demo
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits:
//          https://github.com/mauidude/deduper
////////////////////////////////////////////////////////////////////////////

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"gopkg.in/suntong/deduper.v1/command"
	"gopkg.in/suntong/deduper.v1/minhash"
)

type CarAd struct {
	Price, Title, Description, Details string
}

type config struct {
	debug     bool
	bands     int
	rows      int
	shingles  int
	threshold float64
}

var cfg *config

func init() {
	cfg = &config{}

	flag.BoolVar(&cfg.debug, "debug", false, "Enable debug logging")
	flag.IntVar(&cfg.bands, "bands", 100, "Number of bands")
	flag.IntVar(&cfg.rows, "hashes", 2, "Number of hashes to use")
	flag.IntVar(&cfg.shingles, "shingles", 2, "Number of shingles")
	flag.Float64Var(&cfg.threshold, "threshold", 0.5, "Threshold")
}

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	if cfg.debug {
	}

	bytes := []byte(carAdsData)

	// Unmarshal string into structs.
	var carAds []CarAd
	json.Unmarshal(bytes, &carAds)
	//fmt.Printf("%+v\n", carAds)

	mh := minhash.New(cfg.bands, cfg.rows, cfg.shingles)

	for _, v := range carAds {
		//fmt.Printf("*%v*: %v, %v\n", v.Title, v.Details, v.Price)
		matches := mh.FindSimilar(
			strings.NewReader(v.Description), cfg.threshold)
		if len(matches) != 0 {
			fmt.Printf("\n[%v] matches\n ", v.Title)
			json.NewEncoder(os.Stdout).Encode(matches)
			//fmt.Printf(" [%v]\n", v.Description)
		} else {
			command.NewWriteCommand(v.Title, v.Description).Apply(mh)
		}
	}
}

var carAdsData = `
[{"Description":"Take over my lease due to moving countries for work. 2016 Honda Civic Touring Fully loaded with leather, heated front and rear seats, GPS, wireless cell phone charging, back up camera, automatic? 12,000km | Automatic | CarProof","Details":"12,000km | Automatic | CarProof","Price":"$254.00","Title":"2016 Civic Touring - Fully Loaded"},{"Description":"Chrysler Pacifica Limited addition all wheel drive.. in great running condition runs and drive great..no rust no problem All power options are working Power windows power locks power seats heated? 236km | Automatic","Details":"236km | Automatic","Price":"$2,800.00","Title":"2005 Chrysler Pacifica Limited DVD-GPS"},{"Description":"A/C Heated front seats Driver power seat Power sunroof Power windows Alloy rims Automatic Head Lights Automatic winshield wiper Leather seats BOSE speakers Xenon Headlights Built-in GPS navigation? 65,400km | Automatic","Details":"65,400km | Automatic","Price":"$12,900.00","Title":"2012 Mazda3 GT (Full Option)"},{"Description":"A/C Heated front seats Driver power seat Power sunroof Power windows Alloy rims Automatic Head Lights Automatic winshield wiper Leather seats BOSE speakers Xenon Headlights Built-in GPS navigation? 65,400km | Automatic","Details":"65,400km | Automatic","Price":"$12,900.00","Title":"2012 Mazda3 GT (Full Option)"},{"Description":"2007 PONTIAC GRAND PRIX GP 4DOOR SEDAN FULLY LOADED INCLUDING ICE COLD AC/CAR DRIVES VERY SMOOTH MECHANICALLY A1,CLEAN IN SIDE AND OUT NO RIPS OR TARES NO RUST VERY WELL MAINTAINED REGULER OIL CHARGE? 200km | Automatic","Details":"200km | Automatic","Price":"$3,500.00","Title":"2007 Pontiac Grand Am cloth Sedan"},{"Description":"***Honda Civic Sedan SE 4 Sale*** -Great car -Great on gas -Currently Driving -No damage/No rust/No scraps -Great body -Automatic transmission/AC/CD/Heat -Including Garmin GPS/Seat covers \u0026 Mats? 200km | Automatic","Details":"200km | Automatic","Price":"$1,800.00","Title":"04 Honda Sedan 4 Sale"},{"Description":"Take over my lease due to moving countries for work. 2016 Honda Civic Touring Fully loaded with leather, heated front and rear seats, GPS, wireless cell phone charging, back up camera, automatic? 12,000km | Automatic | CarProof","Details":"12,000km | Automatic | CarProof","Price":"$254.00","Title":"2016 Civic Touring - Fully Loaded"},{"Description":"I'm selling my 2012 Volkswagen jetta sportline. This car is in fantastic condition! 88000km Beautiful exhaust, rims, tinted windows. Car is fully loaded. Only feature missing is gps! I can include? 87,500km | Automatic","Details":"87,500km | Automatic","Price":"$14,000.00","Title":"2012 Volkswagen jetta sportline"},{"Description":"Full Loaded BMW x5, 4.4 V8, accident free, very clean, mint condotion. Front and back camera, gps, dvd, cd changer, bluetooth,the car look like a new. 21 inch wheels. another set of wheels 19 inch? 240km | Automatic","Details":"240km | Automatic","Price":"$10,800.00","Title":"2006 BMW X5 SUV, Crossover"},{"Description":"LEASE TAKEOVER 2015 Nissan Murano SL with leather, GPS, interactive touchscreen, bluetooth, power liftgate, power locks, power moonroof, Lane switch assist, Front, Side and Rear cameras, tinted? 15,000km | Automatic","Details":"15,000km | Automatic","Price":"$477.01","Title":"INCREDIBLE LEASE DEAL!! 2015 Nissan Murano SL $477 FULLY LOADED"},{"Description":"Hi. I am reposting my 1995 Benz for sale. It's a fully loaded with all the options available from the year it was made. I have an extra GPS system that can be yours for the asking price. Power? 145km | Automatic","Details":"145km | Automatic","Price":"$1,900.00","Title":"1995 Mercedes-Benz 200-Series C220 Sedan"},{"Description":"Mazda 6 GT in excellent condition . Leather interior with sunroof and all options except GPS. 6 CD player with Bluetooth connectivity. 11 speaker Bose system. Car comes with a 3 year /60000 km? 74,000km | Automatic","Details":"74,000km | Automatic","Price":"$15,500.00","Title":"2011 Mazda Mazda6 GT 3.7 litre Sedan"},{"Description":"Only 15,800, great Offer!!! 57000 Kilometers, only sell for $16000! Lots to offer! GPS, Reverse video camera, Bluetooth, CD, SD card, pad, front seat cover!! 57,000km | Automatic","Details":"57,000km | Automatic","Price":"$15,800.00","Title":"2014 Jeep Compass"},{"Description":"Power steering, power brakes, captain seats, tinted windows, alarm, remote start, shelving, roof rack, running boards, cruise control, GPS, back up camera, $16,500. 72,000km | Automatic","Details":"72,000km | Automatic","Price":"$16,500.00","Title":"2007 Ford E-250 Minivan, Van"},{"Description":"Fully loaded, sunroof, heated and cooled seats,heated steering wheel, gps, back up camera, auto start, 2 key fobs, tinted windows.2l ecoboost. $398/m tax included. 0 down ,11 months left on the walk? 34,200km | Automatic","Details":"34,200km | Automatic","Price":"$398.00","Title":"2014 Ford Fusion Titanium Sedan"},{"Description":"Perfect time to enjoy this beautiful sports car for only $409 per month including Tax Get the lease of my BRZ Sport-Tech Car comes fully loaded with GPS navigation Dual climate control heated seats 6? 33,100km | Automatic","Details":"33,100km | Automatic","Price":"$409.00","Title":"2014 Subaru BRZ Technology Coupe (2 door)"},{"Description":"Cub van 2011 ford E 450 (16 feet)(cert\u0026emiss)available. ..6new tire's\u0026breaks\u0026rotors+back up camera .7\"cd player+gps in dash . extremely good on gas . heavy duty shocks and spring .......... no? 1km | Automatic","Details":"1km | Automatic","Price":"$17,900.00","Title":"2011 Cub van heavy duty ford E450 (cer\u0026emiss)available"},{"Description":"2011 Mitsubishi Lancer DE with under 52000 km, 8\" touch screen / Bluetooth gps nav system, Rear view camera. In excellent condition, no rust. Safetied and e tested as its my daily driving vehicle. 51,000km | Automatic","Details":"51,000km | Automatic","Price":"$10,000.00","Title":"2011 Mitsubishi Lancer DE 51000km"},{"Description":"2012 Hyundai Veloster TECH | NAVIGATION | PANORAMIC SUNROOF | REAR CAMERA | Women driven car. Car is NEAT in and out, you will be impressed! I am selling the car because I changed job - I no longer? 60,000km | Automatic | CarProof","Details":"60,000km | Automatic | CarProof","Price":"$13,500.00","Title":"2012 Hyundai Veloster Tech Package - GPS, BLUETOOTH AND MORE"},{"Description":"Fully Loaded!!!! Leather heated seats sunroof, tinted windows Sony Sync, 7 inch display with GPS Winter tires included, will certify for asking price, no accidents, one owner $10500 OBO, AS IS 82,000km | Automatic | CarProof","Details":"82,000km | Automatic | CarProof","Price":"$10,500.00","Title":"2012 Ford Focus TITANIUM Hatchback"},{"Description":"Looking for someone to take over Our lease for Toyota Camry Hybrid XLE (Fully loaded) , Clean and free from accident Car features: -Leather Seating -GPS -Heated front seats -Powerful Ecoboost engine,? 43,000km | Automatic","Details":"43,000km | Automatic","Price":"$470.00","Title":"2015 Toyota Camry HBRID XLE\""},{"Description":"Take over my lease due to moving countries for work. 2016 Honda Civic Touring Fully loaded with leather, heated front and rear seats, GPS, wireless cell phone charging, back up camera, automatic? 12,000km | Automatic | CarProof","Details":"12,000km | Automatic | CarProof","Price":"$254.00","Title":"2016 Civic Touring - Fully Loaded"},{"Description":"-Gently used and lady driven -Extremely clean 4 cylinder, save on gas -89000 miles -well maintained - comes with rain guards, side steps and a roof rack - it has a rear cam, GPS, dvd player and? 89,000km | Automatic","Details":"89,000km | Automatic","Price":"$6,000.00","Title":"2003 Toyota RAV4 SUV, Crossover"},{"Description":"Hi everyone, I'm selling my 2012 Honda Civic, had all power options and Bluetooth. GPS, SUNROOF,LEATHER SEATS, Hands free calling. Only has 63k, still in very good condition. Clean and ready for the? 63,505km | Automatic","Details":"63,505km | Automatic","Price":"$13,500.00","Title":"2012 HONDA CIVIC LX Sunroof - Bluetooth fully Loaded"},{"Description":"For Sale 06 F150 : extra leafs in rear :All new front end parts , tie rods,ball joints, control arms, :Four wheel alignment Has \"highway flashing lights front and rear :rear traffic LED light bar GPS? 1km | Automatic","Details":"1km | Automatic","Price":"$7,200.00","Title":"SOLD SOLD SOLD"},{"Description":"A mint condition 2006 fully loaded honda odyssey for sale. Leather inside, heated seats, navi, dvd. 20,700km | Automatic","Details":"20,700km | Automatic","Price":"$10,500.00","Title":"2006 Great condition Honda Odyssey Minivan, fully loaded gps, dv"},{"Description":"I have for sale Genesis 2012 very clean car and premium car all automatic heated and cold seat rear and front all features GPS Rear view camera Tilt steering Tilt mirror on reverse Brembo breaks Hid? 76,500km | Automatic","Details":"76,500km | Automatic","Price":"$16,500.00","Title":"Hyundai Genesis sedan 2012 !! 3.8L v6"},{"Description":"Fully Loaded. Call me @ 4168786775 New engine / Transmission @ 170KM Auto Start remote Alarm Sun Roof New Tires New Battery Leather Seats A/C GPS CD player Lady Driven Much more........Mint? 31,000km | Automatic","Details":"31,000km | Automatic","Price":"$1,900.00","Title":"2001 Toyota Camry Other"},{"Description":"2013 FORD Fusion SE 2.0T GPS and 18' alloy OEM wheels, leather seats Well equipped. 37000km very low mileage. 37,000km | Automatic","Details":"37,000km | Automatic","Price":"$15,000.00","Title":"2013 Ford Fusion SE 2.0T"},{"Description":"hightest top model corolla LE2015. 4 brand new michelin winter tires. bluetooth,GPS,real leather seat,all glass cover protection,sunroof auto, the best clean and protect body and botton anti? 30,000km | Automatic","Details":"30,000km | Automatic","Price":"$17,500.00","Title":"2015 Toyota Corolla eco premium Sedan"},{"Description":"VEHICLE OPTIONS: Air Conditioning Navigation system (GPS) Power locks Power trunk Remote keyless entry Power windows CD player Leather seats Power seats Alloy wheels LOW KMS... 3.5 litre V-6? 86,500km | Automatic","Details":"86,500km | Automatic","Price":"$12,999.00","Title":"2010 Dodge Charger sxt Sedan"},{"Description":"highest top model with 4 brand new michelin winter tires coming with bluetooth,GPS,,real leather seat,sterling auto up and down adjustable, auto sunroof,all the glasses coveing protection dark dime.? 30,000km | Automatic","Details":"30,000km | Automatic","Price":"$17,900.00","Title":"2015 Toyota Corolla Evo premium Sedan"}]
`

/*

$ go run minhash-test.go

[2012 Mazda3 GT (Full Option)] matches
 [{"id":"2012 Mazda3 GT (Full Option)","similarity":1}]

[2016 Civic Touring - Fully Loaded] matches
 [{"id":"2016 Civic Touring - Fully Loaded","similarity":1}]

[2016 Civic Touring - Fully Loaded] matches
 [{"id":"2016 Civic Touring - Fully Loaded","similarity":1}]

*/
