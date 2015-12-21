package cligraphite

//Copyright 2015 MediaMath <http://www.mediamath.com>.  All rights reserved.
//Use of this source code is governed by a BSD-style
//license that can be found in the LICENSE file.

import (
	"fmt"

	"github.com/MediaMath/govent/graphite"
	"github.com/codegangsta/cli"
)

//UserFlag has the username for http authenticated graphite apis
var UserFlag = cli.StringFlag{
	Name:   "graphite_user",
	Usage:  "Authenticate to graphite endpoint with this user.",
	EnvVar: "GRAPHITE_USER",
}

//PasswordFlag has the password for http authenticated graphite apis
var PasswordFlag = cli.StringFlag{
	Name:   "graphite_password",
	Usage:  "Authenticate to graphite endpoint with this password.",
	EnvVar: "GRAPHITE_PASSWORD",
}

//URLFlag is the graphite events endpoint.
var URLFlag = cli.StringFlag{
	Name:   "graphite_url",
	Usage:  "Graphite endpoint to send graphite events to.",
	EnvVar: "GRAPHITE_URL",
}

//Flags is the common graphite cli flags
var Flags = []cli.Flag{UserFlag, PasswordFlag, URLFlag}

//NewClientFromContext creates a graphite client from the cli context assuming it is using
//the flags in this package
func NewClientFromContext(ctx *cli.Context) (*graphite.Graphite, error) {
	url := ctx.String(URLFlag.Name)
	if url == "" {
		return nil, fmt.Errorf("%s is required", URLFlag.Name)
	}

	return graphite.New(ctx.String(UserFlag.Name), ctx.String(PasswordFlag.Name), url), nil
}
