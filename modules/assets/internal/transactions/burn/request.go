/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package burn

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

type transactionRequest struct {
	BaseReq rest.BaseReq `json:"baseReq"`
	FromID  string       `json:"fromID" valid:"required~required field fromID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field fromID "`
	AssetID string       `json:"assetID" valid:"required~required field assetID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field assetID "`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

func (transactionRequest transactionRequest) Validate() error {
	_, Error := govalidator.ValidateStruct(transactionRequest)
	return Error
}
func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, cliContext client.Context) (helpers.TransactionRequest, error) {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(cliContext),
		cliCommand.ReadString(flags.FromID),
		cliCommand.ReadString(flags.AssetID),
	), nil
}
func (transactionRequest transactionRequest) FromJSON(rawMessage json.RawMessage) (helpers.TransactionRequest, error) {
	if Error := json.Unmarshal(rawMessage, &transactionRequest); Error != nil {
		return nil, Error
	}

	return transactionRequest, nil
}
func (transactionRequest transactionRequest) GetBaseReq() rest.BaseReq {
	return transactionRequest.BaseReq
}
func (transactionRequest transactionRequest) MakeMsg() (sdkTypes.Msg, error) {
	from, Error := sdkTypes.AccAddressFromBech32(transactionRequest.GetBaseReq().From)
	if Error != nil {
		return nil, Error
	}

	return newMessage(
		from,
		base.NewID(transactionRequest.FromID),
		base.NewID(transactionRequest.AssetID),
	), nil
}
func (transactionRequest) RegisterCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, transactionRequest{})
}
func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}
func newTransactionRequest(baseReq rest.BaseReq, fromID string, assetID string) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq: baseReq,
		FromID:  fromID,
		AssetID: assetID,
	}
}
