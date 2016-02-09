package main

import "net/http"
import "io/ioutil"
import "os"
import "encoding/xml"
import "log"
import "fmt"
import (
	gj "github.com/kpawlik/geojson"
)

//Record is a single update of a single bus location
type Record struct {
	BusID            string  `xml:"Record>BusID"`
	BusName          string  `xml:"Record>BusName"`
	Latitude         float64 `xml:"Record>Latitude"`
	Longitude        float64 `xml:"Record>Longitude"`
	RouteID          string  `xml:"Record>RouteID"`
	TripID           string  `xml:"Record>TripID"`
	Direction        string  `xml:"Record>Direction"`
	ServiceDirection string  `xml:"Record>ServiceDirection"`
	Service          string  `xml:"Record>Service"`
	ServiceName      string  `xml:"Record>ServiceName"`
	TripHeadsign     string  `xml:"Record>TripHeadsign"`
	LocationUpdated  string  `xml:"Record>LocationUpdated"`
}

func parserecord(r **Record, body []byte) {
	err := xml.Unmarshal(body, r)
	if err != nil {
		log.Fatal("Unmarshall failed", err)
	}

}

func makeFeaturefromRecord(r *Record) string {
	var f *gj.Feature
	lon := gj.Coord(r.Longitude)
	lat :=  gj.Coord(r.Latitude)
	p := gj.NewPoint(gj.Coordinate{lon, lat})
	id := r.BusID
	props := make(map[string]interface{})
	props["BusName"] = r.BusName
	props["RouteID"] = r.RouteID
	props["TripID"] = r.TripID
	props["TripHeadsign"] = r.TripHeadsign
	props["Service"] = r.Service
	props["LocationUpdated"] = r.LocationUpdated
	f = gj.NewFeature(p, props, id)
	gjstr, err := gj.Marshal(f)
	if err != nil {
		log.Fatal("geojson marshall failed", err)
	}
	return gjstr
}

func init() {
	log.SetOutput(os.Stderr)
}

func main() {
	log.Println("starting")
	resp, err := http.Get("http://www.miamidade.gov/transit/WebServices/Buses/?RouteID=7")
	if err != nil {
		log.Fatal("error fetching url", err)
	}
	log.Println(resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("could not read Response Body")
	}
	resp.Body.Close()
	r := new(Record)
	parserecord(&r, body)
	log.Println(r.BusName)
	log.Printf("POINT(%f %f)", r.Longitude, r.Latitude)
	j := makeFeaturefromRecord(r)
	fmt.Println(j)
}
