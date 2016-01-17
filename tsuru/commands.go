// Copyright 2016 yati authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/tsuru/tsuru/cmd"
	"launchpad.net/gnuflag"
)

type install struct {
	fs *gnuflag.FlagSet
}

func (c *install) Info() *cmd.Info {
	return &cmd.Info{
		Name:    "install",
		Usage:   "install",
		Desc:    "",
		MinArgs: 0,
	}
}

func (c *install) Run(context *cmd.Context, client *cmd.Client) error {
	return nil
}
