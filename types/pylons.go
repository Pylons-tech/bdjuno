package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// type definitions for recipe and cookbook live here
type Cookbook struct {
	Creator      string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	ID           string `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	NodeVersion  uint64 `protobuf:"varint,3,opt,name=nodeVersion,proto3" json:"nodeVersion,omitempty"`
	Name         string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description  string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Developer    string `protobuf:"bytes,6,opt,name=developer,proto3" json:"developer,omitempty"`
	Version      string `protobuf:"bytes,7,opt,name=version,proto3" json:"version,omitempty"`
	SupportEmail string `protobuf:"bytes,8,opt,name=supportEmail,proto3" json:"supportEmail,omitempty"`
	Enabled      bool   `protobuf:"varint,9,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

type Execution struct {
	Creator             string       `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	ID                  string       `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	RecipeID            string       `protobuf:"bytes,3,opt,name=recipeID,proto3" json:"recipeID,omitempty"`
	CookbookID          string       `protobuf:"bytes,4,opt,name=cookbookID,proto3" json:"cookbookID,omitempty"`
	RecipeVersion       string       `protobuf:"bytes,5,opt,name=recipeVersion,proto3" json:"recipeVersion,omitempty"`
	NodeVersion         uint64       `protobuf:"varint,6,opt,name=nodeVersion,proto3" json:"nodeVersion,omitempty"`
	BlockHeight         int64        `protobuf:"varint,7,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	ItemInputs          []ItemRecord `protobuf:"bytes,8,rep,name=itemInputs,proto3" json:"itemInputs"`
	CoinInputs          sdk.Coins    `protobuf:"bytes,9,rep,name=coinInputs,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coinInputs"`
	CoinOutputs         sdk.Coins    `protobuf:"bytes,10,rep,name=coinOutputs,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coinOutputs"`
	ItemOutputIDs       []string     `protobuf:"bytes,11,rep,name=itemOutputIDs,proto3" json:"itemOutputIDs,omitempty"`
	ItemModifyOutputIDs []string     `protobuf:"bytes,12,rep,name=itemModifyOutputIDs,proto3" json:"itemModifyOutputIDs,omitempty"`
}

type EventCompleteExecution struct {
	Creator       string    `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	ID            string    `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	BurnCoins     sdk.Coins `protobuf:"bytes,3,rep,name=burnCoins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"burnCoins"`
	PayCoins      sdk.Coins `protobuf:"bytes,4,rep,name=payCoins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"payCoins"`
	TransferCoins sdk.Coins `protobuf:"bytes,5,rep,name=transferCoins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"transferCoins"`
	FeeCoins      sdk.Coins `protobuf:"bytes,6,rep,name=feeCoins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"feeCoins"`
	CoinOutputs   sdk.Coins `protobuf:"bytes,7,rep,name=coinOutputs,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coinOutputs"`
	MintItems     []Item    `protobuf:"bytes,8,rep,name=mintItems,proto3" json:"mintItems"`
	ModifyItems   []Item    `protobuf:"bytes,9,rep,name=modifyItems,proto3" json:"modifyItems"`
}

type ItemRecord struct {
	ID      string           `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Doubles []DoubleKeyValue `protobuf:"bytes,2,rep,name=doubles,proto3" json:"doubles"`
	Longs   []LongKeyValue   `protobuf:"bytes,3,rep,name=longs,proto3" json:"longs"`
	Strings []StringKeyValue `protobuf:"bytes,4,rep,name=strings,proto3" json:"strings"`
}
