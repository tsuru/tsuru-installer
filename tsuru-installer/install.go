// Copyright 2016 yati authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"

	"github.com/fsouza/go-dockerclient"
)

func createContainer(address, ca, cert, key, image string) error {
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
	config := docker.Config{Image: image}
	opts := docker.CreateContainerOptions{Config: &config}
	container, err := client.CreateContainer(opts)
	if err != nil {
		return err
	}
	return client.StartContainer(container.ID, nil)
}
