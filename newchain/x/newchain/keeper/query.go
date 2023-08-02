package keeper

import (
	"newchain/x/newchain/types"
)

var _ types.QueryServer = Keeper{}
