// Copyright © 2024 Meroxa, Inc.
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

package source

import "github.com/mer-oscar/conduit-connector-zeromq/common"

//go:generate paramgen -output=config_paramgen.go Config
type Config struct {
	common.Config
	// PortBindings is a comma separated list of ports that we wish to bind to
	PortBindings string `json:"portBindings"`
}
