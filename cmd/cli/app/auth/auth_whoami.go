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

package auth

import (
	"context"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"github.com/stacklok/minder/internal/util/cli"
	minderv1 "github.com/stacklok/minder/pkg/api/protobuf/go/minder/v1"
)

// whoamiCmd represents the whoami command
var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "whoami for current user",
	Long:  `whoami gets information about the current user from the minder server`,
	RunE:  cli.GRPCClientWrapRunE(whoamiCommand),
}

// whoamiCommand is the whoami subcommand
func whoamiCommand(ctx context.Context, cmd *cobra.Command, conn *grpc.ClientConn) error {
	client := minderv1.NewUserServiceClient(conn)

	userInfo, err := client.GetUser(ctx, &minderv1.GetUserRequest{})
	if err != nil {
		return cli.MessageAndError("Error getting information for user", err)
	}

	cmd.Println(cli.Header.Render("Here are your details:"))
	renderUserInfoWhoami(conn.Target(), userInfo)
	return nil
}

func init() {
	AuthCmd.AddCommand(whoamiCmd)
}
