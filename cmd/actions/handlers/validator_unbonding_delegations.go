package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	actionstypes "github.com/Pylons-tech/bdjuno/cmd/actions/types"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func ValidatorUnbondingDelegations(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	var actionPayload actionstypes.Payload
	err = json.Unmarshal(reqBody, &actionPayload)
	if err != nil {
		http.Error(w, "invalid payload: failed to unmarshal json", http.StatusInternalServerError)
		return
	}

	result, err := getUnbondingDelegationsFromValidator(actionPayload.Input)
	if err != nil {
		errorHandler(w, err)
		return
	}

	data, _ := json.Marshal(result)
	w.Write(data)
}

func getUnbondingDelegationsFromValidator(input actionstypes.PayloadArgs) (actionstypes.UnbondingDelegationResponse, error) {
	parseCtx, sources, err := getCtxAndSources()
	if err != nil {
		return actionstypes.UnbondingDelegationResponse{}, err
	}

	// Get latest node height
	height, err := parseCtx.Node.LatestHeight()
	if err != nil {
		return actionstypes.UnbondingDelegationResponse{}, fmt.Errorf("error while getting chain latest block height: %s", err)
	}

	// Get all unbonding delegations from the given validator opr address
	unbondingDelegations, err := sources.StakingSource.GetUnbondingDelegationsFromValidator(
		height,
		input.Address,
		&query.PageRequest{
			Offset:     input.Offset,
			Limit:      input.Limit,
			CountTotal: input.CountTotal,
		},
	)
	if err != nil {
		return actionstypes.UnbondingDelegationResponse{},
			fmt.Errorf("error while getting all unbonding delegations from validator %s: %s", input.Address, err)
	}

	unbondingDelegationsList := make([]actionstypes.UnbondingDelegation, len(unbondingDelegations.UnbondingResponses))
	for index, del := range unbondingDelegations.UnbondingResponses {
		unbondingDelegationsList[index] = actionstypes.UnbondingDelegation{
			DelegatorAddress: del.DelegatorAddress,
			ValidatorAddress: del.ValidatorAddress,
			Entries:          del.Entries,
		}
	}

	return actionstypes.UnbondingDelegationResponse{
		UnbondingDelegations: unbondingDelegationsList,
		Pagination:           unbondingDelegations.Pagination,
	}, nil
}
