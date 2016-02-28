// Copyright 2016 yati authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/tsuru/gnuflag"
	"github.com/tsuru/tsuru-installer/tsuru-installer/iaas"
	_ "github.com/tsuru/tsuru-installer/tsuru-installer/iaas/dockermachine"
	"github.com/tsuru/tsuru/cmd"
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
	fmt.Println("Creating machine")
	i := iaas.Get("docker-machine")
	m, err := i.CreateMachine(nil)
	if err != nil {
		fmt.Println("Error creating machine")
		return err
	}
	fmt.Printf("Machine %s successfully created!\n", m.Address)
	fmt.Println("Installing MongoDB")
	err = createContainer(m.Address, m.Config["ca"], m.Config["cert"], m.Config["key"], "mongo")
	if err != nil {
		fmt.Println("Error installing MongoDB!")
		return err
	}
	fmt.Println("MongoDB successfully installed!")
	fmt.Println("Installing Redis")
	err = createContainer(m.Address, m.Config["ca"], m.Config["cert"], m.Config["key"], "redis")
	if err != nil {
		fmt.Println("Error installing Redis!")
		return err
	}
	fmt.Println("Redis successfully installed!")
	fmt.Println("Installing Docker Registry")
	err = createContainer(m.Address, m.Config["ca"], m.Config["cert"], m.Config["key"], "registry")
	if err != nil {
		fmt.Println("Error installing Docker Registry!")
		return err
	}
	fmt.Println("Docker Registry successfully installed!")
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
