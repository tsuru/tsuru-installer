// Copyright 2016 yati authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/andrewsmedina/yati/tsuru/iaas"
	_ "github.com/andrewsmedina/yati/tsuru/iaas/dockermachine"
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
	i := iaas.Get("docker-machine")
	m, err := i.CreateMachine(nil)
	if err != nil {
		return err
	}
	fmt.Printf("Machine %s successfully created!", m.Address)
	return nil
}

type uninstall struct {
	fs *gnuflag.FlagSet
}

func (c *uninstall) Info() *cmd.Info {
	return &cmd.Info{
		Name:    "uninstall",
		Usage:   "uninstall",
		Desc:    "",
		MinArgs: 0,
	}
}

func (c *uninstall) Run(context *cmd.Context, client *cmd.Client) error {
	i := iaas.Get("docker-machine")
	err := i.DeleteMachine(&iaas.Machine{})
	if err != nil {
		return err
	}
	fmt.Println("Machine successfully removed!")
	return nil
}
