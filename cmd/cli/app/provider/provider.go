//
// Copyright 2023 Stacklok, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// NOTE: This file is for stubbing out client code for proof of concept
// purposes. It will / should be removed in the future.
// Until then, it is not covered by unit tests and should not be used
// It does make a good example of how to use the generated client code
// for others to use as a reference.

// Package provider is the root command for the provider subcommands
package provider

import (
	"github.com/spf13/cobra"

	"github.com/stacklok/mediator/cmd/cli/app"
)

// ProviderCmd is the root command for the provider subcommands
var ProviderCmd = &cobra.Command{
	Use:   "provider",
	Short: "Manage providers within a mediator control plane",
	Long:  `The medic provider commands manage providers within a mediator control plane.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Usage()
	},
}

func init() {
	app.RootCmd.AddCommand(ProviderCmd)
}