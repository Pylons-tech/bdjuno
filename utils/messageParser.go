package utils

import (
	"fmt"
	"strconv"

	"github.com/gogo/protobuf/proto"

	pylonsTypes "github.com/Pylons-tech/pylons/x/pylons/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MessageNotSupported returns an error telling that the given message is not supported
func MessageNotSupported(msg sdk.Msg) error {
	return fmt.Errorf("message type not supported: %s", proto.MessageName(msg))
}

// // MessageAddressesParser represents a function that extracts all the
// // involved addresses from a provided message (both accounts and validators)
// type MessageAddressesParser = func(cdc codec.Codec, msg sdk.Msg) ([]string, error)

// // JoinMessageParsers joins together all the given parsers, calling them in order
// func JoinMessageParsers(parsers ...MessageAddressesParser) MessageAddressesParser {
// 	return func(cdc codec.Codec, msg sdk.Msg) ([]string, error) {
// 		for _, parser := range parsers {
// 			// Try getting the addresses
// 			addresses, _ := parser(cdc, msg)

// 			// If some addresses are found, return them
// 			if len(addresses) > 0 {
// 				return addresses, nil
// 			}
// 		}
// 		return nil, MessageNotSupported(msg)
// 	}
// }

// // CosmosMessageAddressesParser represents a MessageAddressesParser that parses a
// // Cosmos message and returns all the involved addresses (both accounts and validators)
// var CosmosMessageAddressesParser = JoinMessageParsers(
// 	AuthzMessagesParser,
// 	BankMessagesParser,
// 	CrisisMessagesParser,
// 	DistributionMessagesParser,
// 	EvidenceMessagesParser,
// 	FeeGrantMessagesParser,
// 	GovMessagesParser,
// 	IBCTransferMessagesParser,
// 	SlashingMessagesParser,
// 	StakingMessagesParser,
// 	DefaultMessagesParser,
// 	PylonsMessageParser,
// )

// // DefaultMessagesParser represents the default messages parser that simply returns the list
// // of all the signers of a message
// func DefaultMessagesParser(_ codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {
// 	var signers = make([]string, len(cosmosMsg.GetSigners()))
// 	for index, signer := range cosmosMsg.GetSigners() {
// 		signers[index] = signer.String()
// 	}
// 	return signers, nil
// }

func PylonsMessageParser(cdc codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {
	switch msg := cosmosMsg.(type) {

	case *pylonsTypes.MsgBurnDebtToken:
		// return []string{msg., msg.Grantee}, nil
		return []string{msg.Creator, msg.RedeemInfo.Address, msg.RedeemInfo.ID, msg.RedeemInfo.ProcessorName}, nil
	case *pylonsTypes.MsgCreateAccount:
		return []string{msg.Creator, msg.Username}, nil
	case *pylonsTypes.MsgCompleteExecutionEarly:
		return []string{msg.Creator, msg.ID}, nil
	case *pylonsTypes.MsgCancelTrade:
		return []string{msg.Creator, string(msg.ID)}, nil
	case *pylonsTypes.MsgCreateCookbook:
		return []string{msg.Creator, msg.ID, msg.Name, msg.Description, msg.Developer, msg.Version, msg.SupportEmail, strconv.FormatBool(msg.Enabled)}, nil
	case *pylonsTypes.MsgCreateRecipe:
		msgReturnVal := []string{msg.Creator, msg.CookbookID, msg.ID, msg.Name, msg.Description, msg.Version}
		for i, coinInputsArray := range msg.CoinInputs {
			coinIndex := coinInputsArray.GetCoins()
			msgReturnVal = append(msgReturnVal, coinIndex.GetDenomByIndex(i))
		}
		return msgReturnVal, nil
	case *pylonsTypes.MsgCreateTrade:
		msgReturnVal := []string{msg.Creator}
		for i, coinInputsArray := range msg.CoinInputs {
			coinIndex := coinInputsArray.GetCoins()
			msgReturnVal = append(msgReturnVal, coinIndex.GetDenomByIndex(i))
		}
		for _, itemInputsArray := range msg.ItemInputs {
			msgReturnVal = append(msgReturnVal, itemInputsArray.String())
		}
		for _, coinOutputsArray := range msg.CoinOutputs {
			msgReturnVal = append(msgReturnVal, coinOutputsArray.GetDenom())
		}
		for _, itemOutputsArray := range msg.ItemOutputs {
			msgReturnVal = append(msgReturnVal, itemOutputsArray.String())
		}
		return append(msgReturnVal, msg.ExtraInfo), nil
	case *pylonsTypes.MsgExecuteRecipe:
		msgReturnVal := []string{msg.Creator, msg.CookbookID, msg.RecipeID, string(msg.CoinInputsIndex)}
		msgReturnVal = append(msgReturnVal, msg.ItemIDs...)
		for _, paymentInfoArray := range msg.PaymentInfos {
			msgReturnVal = append(msgReturnVal, paymentInfoArray.String())
		}
		return msgReturnVal, nil
	case *pylonsTypes.MsgFulfillTrade:
		msgReturnVal := []string{msg.Creator, string(msg.ID), string(msg.CoinInputsIndex)}
		for _, ItemArray := range msg.Items {
			msgReturnVal = append(msgReturnVal, ItemArray.String())
		}
		for _, paymentInfoArray := range msg.PaymentInfos {
			msgReturnVal = append(msgReturnVal, paymentInfoArray.String())
		}
		return msgReturnVal, nil
	case *pylonsTypes.MsgGoogleInAppPurchaseGetCoins:
		return []string{msg.Creator, msg.ProductID, msg.PurchaseToken, msg.ReceiptDataBase64, msg.Signature}, nil
	case *pylonsTypes.MsgSendItems:
		msgReturnVal := []string{msg.Creator, msg.Receiver}
		for _, ItemArray := range msg.Items {
			msgReturnVal = append(msgReturnVal, ItemArray.String())
		}
		return msgReturnVal, nil
		//Using basic info from here to below to see if it works
	case *pylonsTypes.MsgUpdateRecipe:
		msgReturnVal := []string{msg.Creator, msg.CookbookID, msg.ID, msg.Name, msg.Description, msg.Version}
		for i, coinInputsArray := range msg.CoinInputs {
			coinIndex := coinInputsArray.GetCoins()
			msgReturnVal = append(msgReturnVal, coinIndex.GetDenomByIndex(i))
		}
		for _, itemInputsArray := range msg.ItemInputs {
			msgReturnVal = append(msgReturnVal, itemInputsArray.String())
		}
		return msgReturnVal, nil
	case *pylonsTypes.MsgSetItemString:
		return []string{msg.Creator, msg.CookbookID}, nil
	case *pylonsTypes.MsgTransferCookbook:
		return []string{msg.Creator}, nil
	case *pylonsTypes.MsgUpdateAccount:
		return []string{msg.Creator}, nil
	case *pylonsTypes.MsgUpdateCookbook:
		return []string{msg.Creator}, nil

	}
	return nil, nil

}

// // AuthzMessagesParser returns the list of all the accounts involved in the given
// // message if it's related to the x/authz module
// func AuthzMessagesParser(cdc codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {
// 	switch msg := cosmosMsg.(type) {
// 	case *authztypes.MsgGrant:
// 		return []string{msg.Granter, msg.Grantee}, nil

// 	case *authztypes.MsgRevoke:
// 		return []string{msg.Granter, msg.Grantee}, nil

// 	case *authztypes.MsgExec:
// 		addresses := []string{msg.Grantee}

// 		for _, msg := range msg.Msgs {
// 			var sdkMsg sdk.Msg
// 			err := cdc.UnpackAny(msg, &sdkMsg)
// 			if err != nil {
// 				return nil, err
// 			}

// 			parsedAddresses, err := DefaultMessagesParser(cdc, sdkMsg)
// 			if err != nil {
// 				return nil, err
// 			}
// 			addresses = append(addresses, parsedAddresses...)
// 		}

// 		return addresses, nil
// 	}

// 	return nil, MessageNotSupported(cosmosMsg)
// }

// // BankMessagesParser returns the list of all the accounts involved in the given
// // message if it's related to the x/bank module
// func BankMessagesParser(_ codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {
// 	switch msg := cosmosMsg.(type) {
// 	case *banktypes.MsgSend:
// 		return []string{msg.ToAddress, msg.FromAddress}, nil

// 	case *banktypes.MsgMultiSend:
// 		var addresses []string
// 		for _, i := range msg.Inputs {
// 			addresses = append(addresses, i.Address)
// 		}
// 		for _, o := range msg.Outputs {
// 			addresses = append(addresses, o.Address)
// 		}
// 		return addresses, nil
// 	}

// 	return nil, MessageNotSupported(cosmosMsg)
// }

// // CrisisMessagesParser returns the list of all the accounts involved in the given
// // message if it's related to the x/crisis module
// func CrisisMessagesParser(_ codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {
// 	switch msg := cosmosMsg.(type) {

// 	case *crisistypes.MsgVerifyInvariant:
// 		return []string{msg.Sender}, nil

// 	}

// 	return nil, MessageNotSupported(cosmosMsg)
// }

// // DistributionMessagesParser returns the list of all the accounts involved in the given
// // message if it's related to the x/distribution module
// func DistributionMessagesParser(_ codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {
// 	switch msg := cosmosMsg.(type) {

// 	case *distrtypes.MsgSetWithdrawAddress:
// 		return []string{msg.DelegatorAddress, msg.WithdrawAddress}, nil

// 	case *distrtypes.MsgWithdrawDelegatorReward:
// 		return []string{msg.DelegatorAddress, msg.ValidatorAddress}, nil

// 	case *distrtypes.MsgWithdrawValidatorCommission:
// 		return []string{msg.ValidatorAddress}, nil

// 	case *distrtypes.MsgFundCommunityPool:
// 		return []string{msg.Depositor}, nil

// 	}

// 	return nil, MessageNotSupported(cosmosMsg)
// }

// // EvidenceMessagesParser returns the list of all the accounts involved in the given
// // message if it's related to the x/evidence module
// func EvidenceMessagesParser(_ codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {
// 	switch msg := cosmosMsg.(type) {

// 	case *evidencetypes.MsgSubmitEvidence:
// 		return []string{msg.Submitter}, nil

// 	}

// 	return nil, MessageNotSupported(cosmosMsg)
// }

// // FeeGrantMessagesParser returns the list of all the accounts involved in the given
// // message if it's related to the x/feegrant module
// func FeeGrantMessagesParser(cdc codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {
// 	switch msg := cosmosMsg.(type) {

// 	case *feegranttypes.MsgGrantAllowance:
// 		return []string{msg.Grantee, msg.Granter}, nil

// 	case *feegranttypes.MsgRevokeAllowance:
// 		return []string{msg.Granter, msg.Grantee}, nil

// 	}

// 	return nil, MessageNotSupported(cosmosMsg)
// }

// // GovMessagesParser returns the list of all the accounts involved in the given
// // message if it's related to the x/gov module
// func GovMessagesParser(cdc codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {
// 	switch msg := cosmosMsg.(type) {

// 	case *govtypes.MsgSubmitProposal:
// 		addresses := []string{msg.Proposer}

// 		var content govtypes.Content
// 		err := cdc.UnpackAny(msg.Content, &content)
// 		if err != nil {
// 			return nil, err
// 		}

// 		// Get addresses from contents
// 		switch content := content.(type) {
// 		case *distrtypes.CommunityPoolSpendProposal:
// 			addresses = append(addresses, content.Recipient)
// 		}

// 		return addresses, nil

// 	case *govtypes.MsgDeposit:
// 		return []string{msg.Depositor}, nil

// 	case *govtypes.MsgVote:
// 		return []string{msg.Voter}, nil

// 	}

// 	return nil, MessageNotSupported(cosmosMsg)
// }

// // IBCTransferMessagesParser returns the list of all the accounts involved in the given
// // message if it's related to the x/iBCTransfer module
// func IBCTransferMessagesParser(_ codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {
// 	switch msg := cosmosMsg.(type) {

// 	case *ibctransfertypes.MsgTransfer:
// 		return []string{msg.Sender, msg.Receiver}, nil

// 	}

// 	return nil, MessageNotSupported(cosmosMsg)
// }

// // SlashingMessagesParser returns the list of all the accounts involved in the given
// // message if it's related to the x/slashing module
// func SlashingMessagesParser(_ codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {
// 	switch msg := cosmosMsg.(type) {

// 	case *slashingtypes.MsgUnjail:
// 		return []string{msg.ValidatorAddr}, nil

// 	}

// 	return nil, MessageNotSupported(cosmosMsg)
// }

// // StakingMessagesParser returns the list of all the accounts involved in the given
// // message if it's related to the x/staking module
// func StakingMessagesParser(_ codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {
// 	switch msg := cosmosMsg.(type) {

// 	case *stakingtypes.MsgCreateValidator:
// 		return []string{msg.ValidatorAddress, msg.DelegatorAddress}, nil

// 	case *stakingtypes.MsgEditValidator:
// 		return []string{msg.ValidatorAddress}, nil

// 	case *stakingtypes.MsgDelegate:
// 		return []string{msg.DelegatorAddress, msg.ValidatorAddress}, nil

// 	case *stakingtypes.MsgBeginRedelegate:
// 		return []string{msg.DelegatorAddress, msg.ValidatorSrcAddress, msg.ValidatorDstAddress}, nil

// 	case *stakingtypes.MsgUndelegate:
// 		return []string{msg.DelegatorAddress, msg.ValidatorAddress}, nil

// 	}

// 	return nil, MessageNotSupported(cosmosMsg)
// }
