package main

import (
	"log"
	"os"

	"github.com/MediaMath/gevent/cligraphite"
	"github.com/MediaMath/gevent/graphite"
	"github.com/codegangsta/cli"
)

var (
	whatFlag = cli.StringFlag{
		Name:   "what",
		Usage:  "The 'What' field in the event.",
		EnvVar: "GEVENT_WHAT",
	}

	dataFlag = cli.StringFlag{
		Name:  "data",
		Usage: "The 'Data' field in the event.",
	}

	tagsFlag = cli.StringSliceFlag{
		Name:   "tag",
		Usage:  "The 'Tag' field in the event.",
		EnvVar: "GEVENT_TAGS",
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "gevent"
	app.Usage = "send events to the graphite api"
	app.Flags = append(cligraphite.Flags, whatFlag, dataFlag, tagsFlag)

	app.Action = func(ctx *cli.Context) {
		what := ctx.String(whatFlag.Name)
		data := ctx.String(dataFlag.Name)
		tags := ctx.StringSlice(tagsFlag.Name)

		if what == "" || data == "" {
			log.Fatalf("%s and %s are required.", whatFlag.Name, dataFlag.Name)
		}

		client, err := cligraphite.NewClientFromContext(ctx)
		if err != nil {
			log.Fatal(err)
		}

		err = client.Publish(graphite.NewEvent(what, data, tags...))
		if err != nil {
			log.Fatal(err)
		}
	}

	app.Run(os.Args)
}
