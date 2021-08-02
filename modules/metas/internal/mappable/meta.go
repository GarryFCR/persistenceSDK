/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

var _ mappables.Meta = (*Meta)(nil)

func (meta Meta) GetData() types.Data { return meta.Data }
func (meta Meta) GetID() types.ID     { return meta.ID }
func (meta Meta) GetKey() helpers.Key {
	return key.FromID(meta.GetID())
}
func (Meta) RegisterCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, Meta{})
}

func NewMeta(data types.Data) mappables.Meta {
	return Meta{
		ID:   key.GenerateMetaID(data),
		Data: data,
	}
}
