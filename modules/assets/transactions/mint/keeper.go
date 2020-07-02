package mint

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/assets/constants"
	"github.com/persistenceOne/persistenceSDK/modules/assets/mapper"
	"github.com/persistenceOne/persistenceSDK/types"
)

type transactionKeeper struct {
	mapper mapper.Mapper
}

var _ types.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) error {
	message := messageFromInterface(msg)
	immutablePropertyList := message.Properties.GetList()
	hashID := mapper.GenerateHashID(immutablePropertyList)
	assetID := mapper.NewAssetID(message.ChainID, message.MaintainersID, message.ClassificationID, hashID)
	asset := mapper.NewAsset(assetID, message.Properties, message.Lock, message.Burn)
	assets := mapper.NewAssets(transactionKeeper.mapper, context).Fetch(assetID)
	if assets.Get(assetID) != nil {
		return constants.EntityAlreadyExistsCode
	}
	assets.Add(asset)
	return nil
}

func NewTransactionKeeper(mapper mapper.Mapper) types.TransactionKeeper {
	return transactionKeeper{mapper: mapper}
}
