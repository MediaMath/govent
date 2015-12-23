package main

//Copyright 2015 MediaMath <http://www.mediamath.com>.  All rights reserved.
//Use of this source code is governed by a BSD-style
//license that can be found in the LICENSE file.

import (
	"flag"
	"testing"

	"github.com/codegangsta/cli"
)

func TestWhatRequired(t *testing.T) {
	set := flag.NewFlagSet("test", 0)
	s := cli.StringSlice([]string{"fo", "sho"})
	set.Var(&s, "tag", "boom")
	ctx := cli.NewContext(nil, set, nil)
	set.Parse([]string{"data"})

	_, err := eventFromCtx(ctx)
	if err == nil {
		t.Fatal("Should have failed")
	}

	set.String("what", "it", "is")

	_, err = eventFromCtx(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDataRequired(t *testing.T) {
	set := flag.NewFlagSet("test", 0)
	s := cli.StringSlice([]string{"fo", "sho"})
	set.Var(&s, "tag", "boom")
	set.String("what", "it", "is")

	ctx := cli.NewContext(nil, set, nil)

	_, err := eventFromCtx(ctx)
	if err == nil {
		t.Fatal("Should have failed")
	}

	set.Parse([]string{"data"})

	_, err = eventFromCtx(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestEventUsesTagIfOnlyOneTagDefinedAndNoWhat(t *testing.T) {
	set := flag.NewFlagSet("test", 0)
	s := cli.StringSlice([]string{"fo"})
	set.Var(&s, "tag", "boom")
	set.Parse([]string{"data"})

	ctx := cli.NewContext(nil, set, nil)

	event, err := eventFromCtx(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if event.What != "fo" && event.Tags != "fo" && event.Data != "data" {
		t.Fatalf("%v", event)
	}

}

func TestEventFromCtx(t *testing.T) {
	set := flag.NewFlagSet("test", 0)
	s := cli.StringSlice([]string{"fo", "sho"})
	set.Var(&s, "tag", "boom")
	set.String("what", "it", "is")
	set.Parse([]string{"data"})

	ctx := cli.NewContext(nil, set, nil)

	event, err := eventFromCtx(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if event.What != "it" && event.Tags != "fo" && event.Data != "data" {
		t.Fatalf("%v", event)
	}

}
