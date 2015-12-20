package cligraphite

//Copyright 2015 MediaMath <http://www.mediamath.com>.  All rights reserved.
//Use of this source code is governed by a BSD-style
//license that can be found in the LICENSE file.

import (
	"flag"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MediaMath/gevent/graphite"
	"github.com/codegangsta/cli"
)

func TestURLRequired(t *testing.T) {
	set := flag.NewFlagSet("test", 0)
	_, err := NewClientFromContext(cli.NewContext(nil, set, nil))

	if err == nil {
		t.Fatal(err)
	}

	set.String("graphite_url", "http://example.com", "")
	_, err = NewClientFromContext(cli.NewContext(nil, set, nil))

	if err != nil {
		t.Fatal(err)
	}
}

func TestClientFromContext(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			http.Error(w, "Auth Req", 401)
		}
	}))
	defer ts.Close()

	set := flag.NewFlagSet("test", 0)
	set.String("graphite_url", ts.URL, "")
	set.String("graphite_user", "foo", "")
	set.String("graphite_password", "bar", "")

	g, err := NewClientFromContext(cli.NewContext(nil, set, nil))
	if err != nil {
		t.Fatal(err)
	}

	err = g.Publish(graphite.NewEvent("boo", "yeah", "boy"))
	if err != nil {
		t.Fatal(err)
	}
}
