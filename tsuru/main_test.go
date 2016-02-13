// Copyright 2015 yati authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/tsuru/tsuru/cmd"
	"gopkg.in/check.v1"
)

func (s *S) TestCommandsFromBaseManagerAreRegistered(c *check.C) {
	baseManager := cmd.BuildBaseManager("tsuru", version, "", nil)
	manager := buildManager("tsuru")
	for name, instance := range baseManager.Commands {
		command, ok := manager.Commands[name]
		c.Assert(ok, check.Equals, true)
		c.Assert(command, check.FitsTypeOf, instance)
	}
}
func (s *S) TestInstallIsRegistered(c *check.C) {
	manager := buildManager("tsuru")
	install, ok := manager.Commands["install"]
	c.Assert(ok, check.Equals, true)
	c.Assert(install, check.FitsTypeOf, &install{})
}
