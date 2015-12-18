package graphite

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

//Event to record in graphite
type Event struct {
	What string `json:"what"`
	Tags string `json:"tags"`
	Data string `json:"data"`
}

//NewEvent creates an event with the provided data
func NewEvent(what string, data string, tags ...string) *Event {
	return &Event{
		What: what,
		Tags: strings.Join(tags, ","),
		Data: data,
	}
}

//New creates a new graphite client
func New(username, password, addr string) *Graphite {
	return &Graphite{username, password, addr, &http.Client{}}
}

//Graphite is a wrapper around the graphite events API
type Graphite struct {
	username string
	password string
	addr     string
	client   *http.Client
}

//Publish sends the event to the graphite API
func (g *Graphite) Publish(event *Event) error {
	b, err := json.Marshal(event)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", g.addr, bytes.NewBuffer(b))

	if err != nil {
		return err
	}

	if g.username != "" && g.password != "" {
		req.SetBasicAuth(g.username, g.password)
	}

	resp, err := g.client.Do(req)

	if err != nil {
		return err
	}

	if resp.Body != nil {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("error response code:%v", resp.StatusCode)
	}
	return nil
}
