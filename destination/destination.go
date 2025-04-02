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

package destination

import (
	"context"
	"fmt"

	"github.com/conduitio/conduit-commons/config"
	"github.com/conduitio/conduit-commons/opencdc"
	sdk "github.com/conduitio/conduit-connector-sdk"
	"github.com/zeromq/goczmq"
)

type Destination struct {
	sdk.UnimplementedDestination

	config     Config
	pubChannel *goczmq.Channeler
}

func New() sdk.Destination {
	// Create Destination and wrap it in the default middleware.
	return sdk.DestinationWithMiddleware(&Destination{}, sdk.DefaultDestinationMiddleware()...)
}

func (d *Destination) Parameters() config.Parameters {
	// Parameters is a map of named Parameters that describe how to configure
	// the Destination. Parameters can be generated from DestinationConfig with
	// paramgen.
	return d.config.Parameters()
}

func (d *Destination) Configure(ctx context.Context, cfg config.Config) error {
	sdk.Logger(ctx).Info().Msg("Configuring Destination...")
	err := sdk.Util.ParseConfig(ctx, cfg, &d.config, New().Parameters())
	if err != nil {
		return fmt.Errorf("invalid config: %w", err)
	}

	return nil
}

func (d *Destination) Open(ctx context.Context) error {
	sdk.Logger(ctx).Info().Msg("Opening Destination...")
	d.pubChannel = goczmq.NewPubChanneler(d.config.RouterEndpoints)

	return nil
}

func (d *Destination) Write(_ context.Context, records []opencdc.Record) (int, error) {
	for _, rec := range records {
		d.pubChannel.SendChan <- [][]byte{[]byte(d.config.Topic), rec.Bytes()}
	}

	return len(records), nil
}

func (d *Destination) Teardown(ctx context.Context) error {
	sdk.Logger(ctx).Info().Msg("Teardown Destination...")
	if d.pubChannel != nil {
		d.pubChannel.Destroy()
	}

	return nil
}
