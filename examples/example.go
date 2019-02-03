package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	fsm "github.com/verseatile/conway"
)

const API_KEY = ""

func main() {
	// create a new state machine
	machine := fsm.NewMachine()
	// create a state instance, a default one is already initialized if not
	s := &fsm.State{
		State: make(map[string]interface{}, 0)}

	// Set the state machine's state to the one you just created
	machine.SetCurrent(s)

	// we'll use a channel to get the data back from our fetch function
	ch := make(chan []byte)
	go grabData(ch, "https://newsapi.org/v2/top-headlines?sources=techradar&apiKey="+API_KEY)

	// set the state of a specfic property of the state currently in the machine
	machine.SetState("techradar", string(<-ch))

	// a sleep timer in case to handle routines that dont get the chance to run
	time.Sleep(500 * time.Millisecond)

	fmt.Println(machine.GetState("techradar"))
}

// Basic fetch function to grab some api data
func grabData(ch chan []byte, url string) {
	client := &http.Client{
		// Timeout:
	}
	resp, err := client.Get(url)
	if err != nil {
		log.Println("ERROR: ", err)
	}
	data, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	ch <- data
}
