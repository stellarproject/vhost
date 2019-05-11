/*
	Copyright (c) 2019 Stellar Project

	Permission is hereby granted, free of charge, to any person
	obtaining a copy of this software and associated documentation
	files (the "Software"), to deal in the Software without
	restriction, including without limitation the rights to use, copy,
	modify, merge, publish, distribute, sublicense, and/or sell copies
	of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:

	The above copyright notice and this permission notice shall be
	included in all copies or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
	EXPRESS OR IMPLIED,
	INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
	IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
	HOLDERS BE LIABLE FOR ANY CLAIM,
	DAMAGES OR OTHER LIABILITY,
	WHETHER IN AN ACTION OF CONTRACT,
	TORT OR OTHERWISE,
	ARISING FROM, OUT OF OR IN CONNECTION WITH
	THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package main

import (
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/stellarproject/terraos/cmd"
	"github.com/urfave/cli"
)

var updateCommand = cli.Command{
	Name:  "update",
	Usage: "update terra components",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "http",
			Usage: "download via http",
		},
		cli.StringFlag{
			Name:  "dest",
			Usage: "destination directory for component",
			Value: "/",
		},
	},
	Action: func(clix *cli.Context) error {
		component := clix.Args().First()
		r := "docker.io/stellarproject/" + component + ":latest"
		logger := logrus.WithField("repo", r)
		repo := cmd.Repo(r)

		logger.Info("creating new content store")
		csDir, err := ioutil.TempDir("", "terra-")
		if err != nil {
			return err
		}
		defer os.RemoveAll(csDir)
		store, err := cmd.NewContentStore(csDir)
		if err != nil {
			return err
		}

		ctx := cmd.CancelContext()
		desc, err := cmd.Fetch(ctx, clix.Bool("http"), store, string(repo))
		if err != nil {
			return err
		}
		return cmd.Unpack(ctx, store, desc, clix.String("dest"))
	},
}
