package handlers

import (
	"encoding/json"
	"net/http"

	actionstypes "github.com/Pylons-tech/bdjuno/cmd/actions/types"
	"github.com/Pylons-tech/bdjuno/database"
	"github.com/Pylons-tech/bdjuno/modules"
	"github.com/Pylons-tech/bdjuno/types/config"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/forbole/juno/v2/cmd/parse"
	"github.com/forbole/juno/v2/modules/messages"

	nodeconfig "github.com/forbole/juno/v2/node/config"
	"github.com/forbole/juno/v2/node/remote"
)

func getCtxAndSources() (*parse.Context, *modules.Sources, error) {
	parseCfg := parse.NewConfig().
		WithDBBuilder(database.Builder).
		WithEncodingConfigBuilder(config.MakeEncodingConfig([]module.BasicManager{
			simapp.ModuleBasics,
		})).
		WithRegistrar(modules.NewRegistrar(messages.JoinMessageParsers(
			messages.CosmosMessageAddressesParser,
		)))

	parseCtx, err := parse.GetParsingContext(parseCfg)
	if err != nil {
		return nil, nil, err
	}

	node := nodeconfig.NewConfig(
		nodeconfig.TypeRemote,
		remote.NewDetails(
			remote.NewRPCConfig("hasura-actions", actionstypes.FlagRPC, 100),
			remote.NewGrpcConfig(actionstypes.FlagGRPC, !actionstypes.FlagSecure),
		),
	)

	sources, err := modules.BuildSources(node, parseCtx.EncodingConfig)
	if err != nil {
		return nil, nil, err
	}

	return parseCtx, sources, nil
}

func errorHandler(w http.ResponseWriter, err error) {
	errorObject := actionstypes.GraphQLError{
		Message: err.Error(),
	}
	errorBody, _ := json.Marshal(errorObject)
	w.WriteHeader(http.StatusBadRequest)
	w.Write(errorBody)
}
