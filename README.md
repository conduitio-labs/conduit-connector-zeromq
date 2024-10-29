# Conduit Connector for ZeroMQ

ZeroMQ connector is one of [Conduit](https://github.com/ConduitIO/conduit) plugins. It provides both, a Source and a Destination ZeroMQ connectors.

## How to build?
Run `make build` to build the connector.

## Testing
Run `make test` to run all the unit tests. Run `make test-integration` to run the integration tests.

The Docker compose file at `test/docker-compose.yml` can be used to run the required resource locally.

## Source
A source is a subscriber that listens to a topic.

### Configuration

| name                  | description                           | required | default value |
|-----------------------|---------------------------------------|----------|---------------|
| `topic` | Topic is the topic to publish to when receiving a record to write. | true     |           |
| `portBindings` | PortBindings is a comma separated list of ports that we wish to bind to. | false     |           |

## Destination
A destination is a publisher that writes data to socket.

### Configuration

| name                       | description                                | required | default value |
|----------------------------|--------------------------------------------|----------|---------------|
| `topic` | Topic is the topic to publish to when receiving a record to write. | true     |           |
| `routerEndpoints` | RouterEndpoints is a comma separated list of socket endpoints that we wish to deal messages to. | true     |           |


## Useful resources
* [ZeroMQ Documentation](https://zeromq.org/get-started)
* [Go Library](https://github.com/zeromq/goczmq)

![scarf pixel](https://static.scarf.sh/a.png?x-pxid=c2ab29c2-181f-49b8-abcc-ac21897048d2)
