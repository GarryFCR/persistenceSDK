/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package super

import "github.com/persistenceOne/persistenceSDK/schema/helpers"

var _ helpers.AuxiliaryResponse = (*auxiliaryResponse)(nil)

func (auxiliaryResponse auxiliaryResponse) IsSuccessful() bool {
	return auxiliaryResponse.Success
}
func (auxiliaryResponse auxiliaryResponse) GetError() error {
	return auxiliaryResponse.Error
}
func newAuxiliaryResponse(error error) helpers.AuxiliaryResponse {
	success := true
	if error != nil {
		success = false
	}

	return auxiliaryResponse{
		Success: success,
		Error:   error,
	}
}
