/*
 * Copyright (c) 2023, © Circle Internet Financial, LTD.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package keeper

import (
	"context"
	"strings"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/circlefin/noble-cctp/x/cctp/types"
)

func (k msgServer) UpdatePerMessageBurnLimit(goCtx context.Context, msg *types.MsgUpdatePerMessageBurnLimit) (*types.MsgUpdatePerMessageBurnLimitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	tokenController := k.GetTokenController(ctx)
	if tokenController != msg.From {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "this message sender cannot update the per message burn limit")
	}

	newPerMessageBurnLimit := types.PerMessageBurnLimit{
		Denom:  strings.ToLower(msg.Denom),
		Amount: msg.Amount,
	}
	k.SetPerMessageBurnLimit(ctx, newPerMessageBurnLimit)

	err := ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgUpdatePerMessageBurnLimitResponse{}, err
}
