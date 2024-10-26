// Code generated by paramgen. DO NOT EDIT.
// Source: github.com/ConduitIO/conduit-commons/tree/main/paramgen

package destination

import (
	"github.com/conduitio/conduit-commons/config"
)

const (
	ConfigRouterEndpoints = "routerEndpoints"
	ConfigTopic           = "topic"
)

func (Config) Parameters() map[string]config.Parameter {
	return map[string]config.Parameter{
		ConfigRouterEndpoints: {
			Default:     "",
			Description: "RouterEndpoints is a comma separated list of socket endpoints that we wish to deal messages to",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{
				config.ValidationRequired{},
			},
		},
		ConfigTopic: {
			Default:     "",
			Description: "Topic is the topic to publish to when receiving a record to write",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{
				config.ValidationRequired{},
			},
		},
	}
}
