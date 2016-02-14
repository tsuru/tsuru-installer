// Copyright 2016 yati authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/fsouza/go-dockerclient"
)

func createContainer(address, image string) error {
	endpoint := fmt.Sprintf("tcp://%s", address)
	client, err := docker.NewClient(endpoint)
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
