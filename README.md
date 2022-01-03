# BDJuno
[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/pylons-tech/bdjuno/Tests)](https://github.com/pylons-tech/bdjuno/actions?query=workflow%3ATests)
[![Go Report Card](https://goreportcard.com/badge/github.com/pylons-tech/bdjuno)](https://goreportcard.com/report/github.com/pylons-tech/bdjuno)
![Codecov branch](https://img.shields.io/codecov/c/github/pylons-tech/bdjuno/cosmos/v0.43.x)

BDJuno (shorthand for BigDipper Juno) is the [Juno](https://github.com/pylons-tech/juno) implementation
for [BigDipper](https://github.com/pylons-tech/big-dipper).

It extends the custom Juno behavior by adding different handlers and custom operations to make it easier for BigDipper
showing the data inside the UI.

All the chains' data that are queried from the RPC and gRPC endpoints are stored inside
a [PostgreSQL](https://www.postgresql.org/) database on top of which [GraphQL](https://graphql.org/) APIs can then be
created using [Hasura](https://hasura.io/).

## Usage
To know how to setup and run BDJuno, please refer to
the [docs website](https://docs.bigdipper.live/cosmos-based/parser/overview/).

## Testing
If you want to test the code, you can do so by running

```shell
$ make test-unit
```

**Note**: Requires [Docker](https://docker.com).

This will:
1. Create a Docker container running a PostgreSQL database.
2. Run all the tests using that database as support.


