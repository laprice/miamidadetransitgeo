package main

import "testing"

func Testparserecord(t *testing.T) {
	testbody := []byte(`<?xml version="1.0" encoding="UTF-8"?><RecordSet><Record><BusID>1118</BusID><BusName>02182</BusName><Latitude>25.77862</Latitude><Longitude>-80.19792</Longitude><RouteID>7</RouteID><TripID>3581207</TripID><Direction>S</Direction><ServiceDirection>Eastbound</ServiceDirection><Service>EB</Service><ServiceName>SUNDAY</ServiceName><TripHeadsign>7 - DOWNTOWN MIAMI</TripHeadsign><LocationUpdated>7:42:48 PM</LocationUpdated></Record></RecordSet>`)
	r := new(Record)
	parserecord(&r, testbody)
	if r.TripID != "32581207" {
		t.Error("parsercord failed")
	}
}

