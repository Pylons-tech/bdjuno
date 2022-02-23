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

// DoubleInputParam describes the bounds on an item input/output parameter of type float64
type DoubleInputParam struct {
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The minimum legal value of this parameter.
	MinValue sdk.Dec `protobuf:"bytes,2,opt,name=minValue,proto3,customtype=github.com/cosmos/cosmos-sdk/t.Dec" json:"minValue"`
	// The maximum legal value of this parameter.
	MaxValue sdk.Dec `protobuf:"bytes,3,opt,name=maxValue,proto3,customtype=github.com/cosmos/cosmos-sdk/t.Dec" json:"maxValue"`
}

// LongInputParam describes the bounds on an item input/output parameter of type int64
type LongInputParam struct {
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The minimum legal value of this parameter.
	MinValue int64 `protobuf:"varint,2,opt,name=minValue,proto3" json:"minValue,omitempty"`
	// The maximum legal value of this parameter.
	MaxValue int64 `protobuf:"varint,3,opt,name=maxValue,proto3" json:"maxValue,omitempty"`
}

// StringInputParam describes the bounds on an item input/output parameter of type string
type StringInputParam struct {
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The value of the parameter
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

// ItemInput is a struct for describing an input item
type ItemInput struct {
	ID      string             `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Doubles []DoubleInputParam `protobuf:"bytes,2,rep,name=doubles,proto3" json:"doubles"`
	Longs   []LongInputParam   `protobuf:"bytes,3,rep,name=longs,proto3" json:"longs"`
	Strings []StringInputParam `protobuf:"bytes,4,rep,name=strings,proto3" json:"strings"`
}

// DoubleWeightRange describes weight range that produce double value
type DoubleWeightRange struct {
	Lower  sdk.Dec `protobuf:"bytes,1,opt,name=lower,proto3,customtype=github.com/cosmos/cosmos-sdk/t.Dec" json:"lower"`
	Upper  sdk.Dec `protobuf:"bytes,2,opt,name=upper,proto3,customtype=github.com/cosmos/cosmos-sdk/t.Dec" json:"upper"`
	Weight uint64  `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
}

// DoubleParam describes the bounds on an item input/output parameter of type float64
type DoubleParam struct {
	Key          string              `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	WeightRanges []DoubleWeightRange `protobuf:"bytes,2,rep,name=weightRanges,proto3" json:"weightRanges"`
	// When program is not empty, weightRanges is ignored
	Program string `protobuf:"bytes,3,opt,name=program,proto3" json:"program,omitempty"`
}

// IntWeightRange describes weight range that produce int value
type IntWeightRange struct {
	Lower  int64  `protobuf:"varint,1,opt,name=lower,proto3" json:"lower,omitempty"`
	Upper  int64  `protobuf:"varint,2,opt,name=upper,proto3" json:"upper,omitempty"`
	Weight uint64 `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
}

// LongParam describes the bounds on an item input/output parameter of type int64
type LongParam struct {
	Key          string           `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	WeightRanges []IntWeightRange `protobuf:"bytes,2,rep,name=weightRanges,proto3" json:"weightRanges"`
	// When program is not empty, weightRanges is ignored
	Program string `protobuf:"bytes,3,opt,name=program,proto3" json:"program,omitempty"`
}

// StringParam describes an item input/output parameter of type string
type StringParam struct {
	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// When program is not empty, value is ignored
	Program string `protobuf:"bytes,3,opt,name=program,proto3" json:"program,omitempty"`
}

// CoinOutput models the continuum of valid outcomes for coin generation in recipes
type CoinOutput struct {
	ID      string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Coin    sdk.Coin `protobuf:"bytes,2,opt,name=coin,proto3" json:"coin"`
	Program string   `protobuf:"bytes,3,opt,name=program,proto3" json:"program,omitempty"`
}

// StringKeyValue describes string key/value set
type StringKeyValue struct {
	Key   string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

// ItemOutput models the continuum of valid outcomes for item generation in recipes
type ItemOutput struct {
	ID      string        `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Doubles []DoubleParam `protobuf:"bytes,2,rep,name=doubles,proto3" json:"doubles"`
	Longs   []LongParam   `protobuf:"bytes,3,rep,name=longs,proto3" json:"longs"`
	Strings []StringParam `protobuf:"bytes,4,rep,name=strings,proto3" json:"strings"`
	// defines a list of mutable strings whose value can be customized by the user
	MutableStrings []StringKeyValue `protobuf:"bytes,5,rep,name=mutableStrings,proto3" json:"mutableStrings"`
	TransferFee    []sdk.Coin       `protobuf:"bytes,6,rep,name=transferFee,proto3" json:"transferFee"`
	// The percentage of a trade sale retained by the cookbook owner. In the range (0.0, 1.0).
	TradePercentage sdk.Dec `protobuf:"bytes,7,opt,name=tradePercentage,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"tradePercentage"`
	// quantity defines the maximum amount of these items that can be created. A 0 value indicates an infinite supply
	Quantity     uint64 `protobuf:"varint,8,opt,name=quantity,proto3" json:"quantity,omitempty"`
	AmountMinted uint64 `protobuf:"varint,9,opt,name=amountMinted,proto3" json:"amountMinted,omitempty"`
	Tradeable    bool   `protobuf:"varint,10,opt,name=tradeable,proto3" json:"tradeable,omitempty"`
}

// ItemModifyOutput describes what is modified from item input
type ItemModifyOutput struct {
	ID           string        `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ItemInputRef string        `protobuf:"bytes,2,opt,name=itemInputRef,proto3" json:"itemInputRef,omitempty"`
	Doubles      []DoubleParam `protobuf:"bytes,3,rep,name=doubles,proto3" json:"doubles"`
	Longs        []LongParam   `protobuf:"bytes,4,rep,name=longs,proto3" json:"longs"`
	Strings      []StringParam `protobuf:"bytes,5,rep,name=strings,proto3" json:"strings"`
	// defines a list of mutable strings whose value can be customized by the user
	MutableStrings []StringKeyValue `protobuf:"bytes,6,rep,name=mutableStrings,proto3" json:"mutableStrings"`
	TransferFee    []sdk.Coin       `protobuf:"bytes,7,rep,name=transferFee,proto3" json:"transferFee"`
	// The percentage of a trade sale retained by the cookbook owner. In the range (0.0, 1.0).
	TradePercentage sdk.Dec `protobuf:"bytes,8,opt,name=tradePercentage,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"tradePercentage"`
	// quantity defines the maximum amount of these items that can be created. A 0 value indicates an infinite supply
	Quantity     uint64 `protobuf:"varint,9,opt,name=quantity,proto3" json:"quantity,omitempty"`
	AmountMinted uint64 `protobuf:"varint,10,opt,name=amountMinted,proto3" json:"amountMinted,omitempty"`
	Tradeable    bool   `protobuf:"varint,11,opt,name=tradeable,proto3" json:"tradeable,omitempty"`
}

// EntriesList is a struct to keep list of items and coins
type EntriesList struct {
	CoinOutputs       []CoinOutput       `protobuf:"bytes,1,rep,name=coinOutputs,proto3" json:"coinOutputs"`
	ItemOutputs       []ItemOutput       `protobuf:"bytes,2,rep,name=itemOutputs,proto3" json:"itemOutputs"`
	ItemModifyOutputs []ItemModifyOutput `protobuf:"bytes,3,rep,name=itemModifyOutputs,proto3" json:"itemModifyOutputs"`
}

// WeightedOutputs is to make structs which is using weight to be based on
type WeightedOutputs struct {
	EntryIDs []string `protobuf:"bytes,1,rep,name=entryIDs,proto3" json:"entryIDs,omitempty"`
	Weight   uint64   `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
}

type CoinInput struct {
	Coins sdk.Coins `protobuf:"bytes,1,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
}

type Recipe struct {
	CookbookID    string            `protobuf:"bytes,1,opt,name=cookbookID,proto3" json:"cookbookID,omitempty"`
	ID            string            `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	NodeVersion   uint64            `protobuf:"varint,3,opt,name=nodeVersion,proto3" json:"nodeVersion,omitempty"`
	Name          string            `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description   string            `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Version       string            `protobuf:"bytes,6,opt,name=version,proto3" json:"version,omitempty"`
	CoinInputs    []CoinInput       `protobuf:"bytes,7,rep,name=coinInputs,proto3" json:"coinInputs"`
	ItemInputs    []ItemInput       `protobuf:"bytes,8,rep,name=itemInputs,proto3" json:"itemInputs"`
	Entries       EntriesList       `protobuf:"bytes,9,opt,name=entries,proto3" json:"entries"`
	Outputs       []WeightedOutputs `protobuf:"bytes,10,rep,name=outputs,proto3" json:"outputs"`
	BlockInterval int64             `protobuf:"varint,11,opt,name=blockInterval,proto3" json:"blockInterval,omitempty"`
	CostPerBlock  sdk.Coin          `protobuf:"bytes,12,opt,name=costPerBlock,proto3" json:"costPerBlock"`
	Enabled       bool              `protobuf:"varint,13,opt,name=enabled,proto3" json:"enabled,omitempty"`
	ExtraInfo     string            `protobuf:"bytes,14,opt,name=extraInfo,proto3" json:"extraInfo,omitempty"`
}

// Execution proto definition - used to create the response you'd see from the chain
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

type Item struct {
	Owner          string           `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	CookbookID     string           `protobuf:"bytes,2,opt,name=cookbookID,proto3" json:"cookbookID,omitempty"`
	ID             string           `protobuf:"bytes,3,opt,name=ID,proto3" json:"ID,omitempty"`
	NodeVersion    uint64           `protobuf:"varint,4,opt,name=nodeVersion,proto3" json:"nodeVersion,omitempty"`
	Doubles        []DoubleKeyValue `protobuf:"bytes,5,rep,name=doubles,proto3" json:"doubles"`
	Longs          []LongKeyValue   `protobuf:"bytes,6,rep,name=longs,proto3" json:"longs"`
	Strings        []StringKeyValue `protobuf:"bytes,7,rep,name=strings,proto3" json:"strings"`
	MutableStrings []StringKeyValue `protobuf:"bytes,8,rep,name=mutableStrings,proto3" json:"mutableStrings"`
	Tradeable      bool             `protobuf:"varint,9,opt,name=tradeable,proto3" json:"tradeable,omitempty"`
	LastUpdate     int64            `protobuf:"varint,10,opt,name=lastUpdate,proto3" json:"lastUpdate,omitempty"`
	TransferFee    []sdk.Coin       `protobuf:"bytes,11,rep,name=transferFee,proto3" json:"transferFee"`
	// The percentage of a trade sale retained by the cookbook owner. In the range (0.0, 1.0).
	TradePercentage sdk.Dec `protobuf:"bytes,12,opt,name=tradePercentage,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"tradePercentage"`
}

type ItemRecord struct {
	ID      string           `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Doubles []DoubleKeyValue `protobuf:"bytes,2,rep,name=doubles,proto3" json:"doubles"`
	Longs   []LongKeyValue   `protobuf:"bytes,3,rep,name=longs,proto3" json:"longs"`
	Strings []StringKeyValue `protobuf:"bytes,4,rep,name=strings,proto3" json:"strings"`
}

type DoubleKeyValue struct {
	Key   string  `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Value sdk.Dec `protobuf:"bytes,2,opt,name=Value,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"Value"`
}

type LongKeyValue struct {
	Key   string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Value int64  `protobuf:"varint,2,opt,name=Value,proto3" json:"Value,omitempty"`
}
