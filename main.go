package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {

	var brief bool
	flag.BoolVar(&brief, "brief", false, "specify brief output. defaults to false.")
	flag.Parse()

	var raw rawData

	stdin := bufio.NewReader(os.Stdin)
	decoder := json.NewDecoder(stdin)
	err := decoder.Decode(&raw)
	if err != nil {
		fmt.Errorf("error decoding raw json input")
		os.Exit(1)
	}

	places := raw.Places
	if brief {
		Brief(places)
		return
	}
	Verbose(places)
}

func Verbose(places Places) {
	fmt.Println("Name,Address,Lat,Lng,Icon,OpenNow,Types")
	for _, p := range places {
		fmt.Printf("%q,%q,%f,%f,%q,%t,%v\n",
			p.Name,
			p.Vicinity,
			p.Geometry.Lat,
			p.Geometry.Lng,
			p.Icon,
			p.Opening_hours.Open_now,
			p.Types,
		)
	}
}

func Brief(places Places) {
	fmt.Println("Name,Address,Lat,Lng")
	for _, p := range places {
		fmt.Printf("%q,%q,%f,%f",
			p.Name,
			p.Vicinity,
			p.Geometry.Lat,
			p.Geometry.Lng,
		)
	}
}
