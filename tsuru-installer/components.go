// Copyright 2016 yati authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/fsouza/go-dockerclient"
	"github.com/tsuru/tsuru-installer/tsuru-installer/iaas"
)

var TsuruComponents = []TsuruComponent{
	&MongoDB{},
	&Redis{},
	&PlanB{},
	&Registry{},
	&TsuruAPI{},
}

type TsuruComponent interface {
	Name() string
	Install(*iaas.Machine) error
}

type MongoDB struct{}

func (c *MongoDB) Name() string {
	return "MongoDB"
}

func (c *MongoDB) Install(machine *iaas.Machine) error {
	return createContainer(machine.Address, "mongo", &docker.Config{Image: "mongo"})
}

type PlanB struct{}

func (c *PlanB) Name() string {
	return "PlanB"
}

func (c *PlanB) Install(machine *iaas.Machine) error {
	config := &docker.Config{
		Image: "tsuru/planb",
		Cmd:   []string{"--listen", ":80", "--read-redis-host", machine.IP, "--write-redis-host", machine.IP},
	}
	return createContainer(machine.Address, "planb", config)
}

type Redis struct{}

func (c *Redis) Name() string {
	return "Redis"
}

func (c *Redis) Install(machine *iaas.Machine) error {
	return createContainer(machine.Address, "redis", &docker.Config{Image: "redis"})
}

type Registry struct{}

func (c *Registry) Name() string {
	return "Docker Registry"
}

func (c *Registry) Install(machine *iaas.Machine) error {
	return createContainer(machine.Address, "registry", &docker.Config{Image: "registry"})
}

type TsuruAPI struct{}

func (c *TsuruAPI) Name() string {
	return "Tsuru API"
}

func (c *TsuruAPI) Install(machine *iaas.Machine) error {
	env := []string{fmt.Sprintf("MONGODB_ADDR=%s", machine.IP),
		"MONGODB_PORT=27017",
		fmt.Sprintf("REDIS_ADDR=%s", machine.IP),
		"REDIS_PORT=6379",
	}
	config := &docker.Config{
		Image: "tsuru/api",
		Env:   env,
	}
	return createContainer(machine.Address, "tsuru", config)
}