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

package zeromq

import (
	"testing"
	"time"

	sdk "github.com/conduitio/conduit-connector-sdk"
	"go.uber.org/goleak"
)

type zeroMQAcceptanceTestDriver struct {
	sdk.ConfigurableAcceptanceTestDriver
}

//nolint:paralleltest // we don't need the paralleltest here
func TestAcceptance(t *testing.T) {
	sdk.AcceptanceTest(t, zeroMQAcceptanceTestDriver{
		ConfigurableAcceptanceTestDriver: sdk.ConfigurableAcceptanceTestDriver{
			Config: sdk.ConfigurableAcceptanceTestDriverConfig{
				Connector:         Connector,
				GoleakOptions:     []goleak.Option{goleak.IgnoreCurrent()},
				SourceConfig:      map[string]string{"portBindings": "tcp://127.0.0.1:5555", "topic": "a"},
				DestinationConfig: map[string]string{"routerEndpoints": "tcp://127.0.0.1:5555", "topic": "a"},
				GenerateDataType:  sdk.GenerateRawData,
				Skip: []string{
					"TestSource_Configure_RequiredParams",
					"TestDestination_Configure_RequiredParams",
					"TestSource_Open_ResumeAtPositionCDC",
					"TestSource_Open_ResumeAtPositionSnapshot",
					"TestDestination_Write_Success",
					"TestSource_Read_Success",
					"TestSource_Read_SuccessSnapshot",
					"TestSource_Read_SuccessCDC",
				},
				WriteTimeout: 500 * time.Millisecond,
				ReadTimeout:  3000 * time.Millisecond,
			},
		},
	})
}
