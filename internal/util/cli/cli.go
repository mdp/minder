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

// Package cli contains utility for the cli
package cli

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/erikgeiser/promptkit/confirmation"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	"github.com/stacklok/minder/internal/util"
	"github.com/stacklok/minder/internal/util/cli/useragent"
)

// ErrWrappedCLIError is an error that wraps another error and provides a message used from within the CLI
type ErrWrappedCLIError struct {
	Message string
	Err     error
}

func (e *ErrWrappedCLIError) Error() string {
	return e.Err.Error()
}

// PrintYesNoPrompt prints a yes/no prompt to the user and returns false if the user did not respond with yes or y
func PrintYesNoPrompt(cmd *cobra.Command, promptMsg, confirmMsg, fallbackMsg string, defaultYes bool) bool {
	// Print the warning banner with the prompt message
	cmd.Println(WarningBanner.Render(promptMsg))

	// Determine the default confirmation value
	defConf := confirmation.No
	if defaultYes {
		defConf = confirmation.Yes
	}

	// Prompt the user for confirmation
	input := confirmation.New(confirmMsg, defConf)
	ok, err := input.RunPrompt()
	if err != nil {
		cmd.Println(WarningBanner.Render(fmt.Sprintf("Error reading input: %v", err)))
		ok = false
	}

	// If the user did not confirm, print the fallback message
	if !ok {
		cmd.Println(Header.Render(fallbackMsg))
	}
	return ok
}

// GrpcForCommand is a helper for getting a testing connection from cobra flags
func GrpcForCommand(v *viper.Viper) (*grpc.ClientConn, error) {
	grpcHost := v.GetString("grpc_server.host")
	grpcPort := v.GetInt("grpc_server.port")
	insecureDefault := grpcHost == "localhost" || grpcHost == "127.0.0.1" || grpcHost == "::1"
	allowInsecure := v.GetBool("grpc_server.insecure") || insecureDefault

	issuerUrl := v.GetString("identity.cli.issuer_url")
	clientId := v.GetString("identity.cli.client_id")

	return util.GetGrpcConnection(
		grpcHost, grpcPort, allowInsecure, issuerUrl, clientId, grpc.WithUserAgent(useragent.GetUserAgent()))
}

// GetAppContext is a helper for getting the cmd app context
func GetAppContext(ctx context.Context, v *viper.Viper) (context.Context, context.CancelFunc) {
	return GetAppContextWithTimeoutDuration(ctx, v, 10)
}

// GetAppContextWithTimeoutDuration is a helper for getting the cmd app context with a custom timeout
func GetAppContextWithTimeoutDuration(ctx context.Context, v *viper.Viper, tout int) (context.Context, context.CancelFunc) {
	v.SetDefault("cli.context_timeout", tout)
	timeout := v.GetInt("cli.context_timeout")

	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
	return ctx, cancel
}

// GRPCClientWrapRunE is a wrapper for cobra commands that sets up the grpc client and context
func GRPCClientWrapRunE(
	runEFunc func(ctx context.Context, cmd *cobra.Command, c *grpc.ClientConn) error,
) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			return fmt.Errorf("error binding flags: %s", err)
		}

		ctx, cancel := GetAppContext(cmd.Context(), viper.GetViper())
		defer cancel()

		c, err := GrpcForCommand(viper.GetViper())
		if err != nil {
			return err
		}

		defer c.Close()

		return runEFunc(ctx, cmd, c)
	}
}

// MessageAndError prints a message and returns an error.
func MessageAndError(msg string, err error) error {
	return &ErrWrappedCLIError{Message: msg, Err: err}
}

// ExitNicelyOnError print a message and exit with the right code
func ExitNicelyOnError(err error, message string) {
	if err != nil {
		if message != "" {
			fmt.Fprintf(os.Stderr, "Message: %s\n", message)
		}
		// Check if the error is wrapped
		var wrappedErr *ErrWrappedCLIError
		if errors.As(err, &wrappedErr) {
			// Print the wrapped message
			fmt.Fprintf(os.Stderr, "Message: %s\n", wrappedErr.Message)
			// Continue processing the wrapped error
			err = wrappedErr.Err
		}
		// Check if the error is a grpc status
		if rpcStatus, ok := status.FromError(err); ok {
			nice := util.FromRpcError(rpcStatus)
			fmt.Fprintf(os.Stderr, "Details: %s\n", nice.Details)
			os.Exit(int(nice.Code))
		} else {
			fmt.Fprintf(os.Stderr, "Details: %s\n", err)
			os.Exit(1)
		}
	}
}
