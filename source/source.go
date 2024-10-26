package source

import (
	"context"
	"fmt"

	"github.com/conduitio/conduit-commons/config"
	"github.com/conduitio/conduit-commons/opencdc"
	sdk "github.com/conduitio/conduit-connector-sdk"
	"github.com/oklog/ulid/v2"
	"github.com/zeromq/goczmq"
)

type Source struct {
	sdk.UnimplementedSource

	config        SourceConfig
	routerChannel *goczmq.Channeler
	readBuffer    chan opencdc.Record
}

type SourceConfig struct {
	Config
}

func New() sdk.Source {
	return sdk.SourceWithMiddleware(&Source{
		readBuffer: make(chan opencdc.Record, 1),
	}, sdk.DefaultSourceMiddleware()...)
}

func (s *Source) Parameters() config.Parameters {
	return s.config.Parameters()
}

func (s *Source) Configure(ctx context.Context, cfg config.Config) error {
	sdk.Logger(ctx).Info().Msg("Configuring Source...")

	err := sdk.Util.ParseConfig(ctx, cfg, &s.config, New().Parameters())
	if err != nil {
		return err
	}

	return nil
}

func (s *Source) Open(ctx context.Context, pos opencdc.Position) error {
	s.routerChannel = goczmq.NewSubChanneler(s.config.PortBindings, s.config.Topic)
	go s.listen(ctx)

	return nil
}

func (s *Source) Read(ctx context.Context) (opencdc.Record, error) {
	select {
	case rec := <-s.readBuffer:
		return rec, nil
	case <-ctx.Done():
		return opencdc.Record{}, ctx.Err()
	}
}

func (s *Source) Ack(ctx context.Context, position opencdc.Position) error {
	return nil
}

func (s *Source) Teardown(ctx context.Context) error {
	if s.routerChannel != nil {
		s.routerChannel.Destroy()
	}

	return nil
}

func (s *Source) listen(ctx context.Context) {
	for {
		select {
		case msg := <-s.routerChannel.RecvChan:
			if msg != nil {
				recFrame := msg[0]
				fmt.Println(len(msg))

				for _, frame := range msg[1:] {
					recBytes := frame

					recUlid := ulid.Make()

					rec := sdk.Util.Source.NewRecordCreate(
						opencdc.Position(string(recFrame)+"_"+recUlid.String()),
						opencdc.Metadata{
							"frame": string(recFrame),
						},
						nil,
						opencdc.RawData(recBytes))

					s.readBuffer <- rec
				}
			}
		case <-ctx.Done():
			return
		}
	}
}
