package main

import "net/http"
import "io/ioutil"
import "os"
import "encoding/xml"
import "encoding/json"
import "log"
import "fmt"

//Record is a single update of a single bus location
type Record struct {
	BusID            string `xml:"Record>BusID"`
	BusName          string `xml:"Record>BusName"`
	Latitude         string `xml:"Record>Latitude"`
	Longitude        string `xml:"Record>Longitude"`
	RouteID          string `xml:"Record>RouteID"`
	TripID           string `xml:"Record>TripID"`
	Direction        string `xml:"Record>Direction"`
	ServiceDirection string `xml:"Record>ServiceDirection"`
	Service          string `xml:"Record>Service"`
	ServiceName      string `xml:"Record>ServiceName"`
	TripHeadsign     string `xml:"Record>TripHeadsign"`
	LocationUpdated  string `xml:"Record>LocationUpdated"`
}

func parserecord(r **Record, body []byte) {
	err := xml.Unmarshal(body, r)
	if err != nil {
		log.Fatal("Unmarshall failed", err)
	}

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
	log.Printf("POINT(%s %s)", r.Longitude, r.Latitude)
	j, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(j))
}
