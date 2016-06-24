// Copyright 2016 yati authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"

	"github.com/fsouza/go-dockerclient"
)

func createContainer(address, image string, env []string) error {
	client, err := docker.NewClient(address)
	if err != nil {
		return err
	}
	pullOpts := docker.PullImageOptions{
		Repository:   image,
		OutputStream: os.Stdout,
		Tag:          "latest",
	}
	err = client.PullImage(pullOpts, docker.AuthConfiguration{})
	if err != nil {
		return err
	}
	imageInspect, err := client.InspectImage(image)
	if err != nil {
		return err
	}
	hostConfig := &docker.HostConfig{}
	if len(imageInspect.Config.ExposedPorts) > 0 {
		hostConfig.PortBindings = make(map[docker.Port][]docker.PortBinding)
		for k := range imageInspect.Config.ExposedPorts {
			hostConfig.PortBindings[k] = []docker.PortBinding{{HostIP: "0.0.0.0", HostPort: k.Port()}}
		}
	}
	config := docker.Config{Image: image, Env: env}
	opts := docker.CreateContainerOptions{Config: &config}
	container, err := client.CreateContainer(opts)
	if err != nil {
		return err
	}
	return client.StartContainer(container.ID, hostConfig)
}
