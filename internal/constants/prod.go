//go:build !staging

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

package constants

const (
	// IdentitySeverURL is the URL of the identity server
	IdentitySeverURL = "https://auth.stacklok.com"
	// MinderGRPCHost is the host of the minder gRPC server
	MinderGRPCHost = "api.stacklok.com"
	// TrustyHttpURL is the URL of the trusty server
	TrustyHttpURL = "https://trustypkg.dev/"
)
