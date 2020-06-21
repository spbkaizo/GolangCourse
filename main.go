package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func LookupLabel(w http.ResponseWriter, r *http.Request) {
	log.Printf("Starting up....\n")
	var label Label

	client := &http.Client{}
	var uid string
	vars := mux.Vars(r)
	uid = vars["uid"]

	url := "https://musicbrainz.org/ws/2/label/" + uid + "?inc=aliases&fmt=json"
	resp, err := client.Get(url)
	if err != nil {
		log.Printf("ERROR: %v")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	//log.Printf("Target URL: %v", url)
	//log.Printf(string(body))
	err = json.Unmarshal(body, &label)
	if err != nil {
		log.Printf("ERROR: %v", err)
	}
	log.Printf("%v of %v", label.Name, label.Country)
	io.WriteString(w, "Hello, world!\n")
	return
}

func LookupArtist(w http.ResponseWriter, r *http.Request) {
	var release Release
	// https://musicbrainz.org/ws/2/release/59211ea4-ffd2-4ad9-9a4e-941d3148024a?inc=artist-credits+labels+discids+recordings&fmt=json
	client := &http.Client{}
	var uid string

	vars := mux.Vars(r)
	uid = vars["uid"]
	log.Printf("Artist var was : %v", uid)

	url := "https://musicbrainz.org/ws/2/release/" + uid + "?inc=artist-credits+labels+discids+recordings&fmt=json"
	resp, err := client.Get(url)
	if err != nil {
		log.Printf("ERROR: %v")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	err = json.Unmarshal(body, &release)
	if err != nil {
		log.Printf("ERROR: %v", err)
	}
	log.Printf("%v", release)
	io.WriteString(w, "Hello, world!\n")
	return
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//releaselabel := LookupReleaseLabel("1673b8f5-31ed-4bba-ba71-24bf95183a6a")
	//log.Printf("Release Label is %v", releaselabel)
	//label := LookupLabel(releaselabel)
	//log.Printf("%v", label)

	r := mux.NewRouter()
	r.HandleFunc("/artist/{uid}", LookupArtist)
	r.HandleFunc("/label/{uid}", LookupLabel)
	log.Fatal(http.ListenAndServe(":8081", r))
}
