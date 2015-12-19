package main

import (
	"flag"
	"testing"

	"github.com/codegangsta/cli"
)

func TestWhatRequired(t *testing.T) {
	set := flag.NewFlagSet("test", 0)
	set.String("tags", "fo", "")
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
	set.String("tags", "fo", "")
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

func TestEventFromCtx(t *testing.T) {
	set := flag.NewFlagSet("test", 0)
	set.String("tags", "fo", "")
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
