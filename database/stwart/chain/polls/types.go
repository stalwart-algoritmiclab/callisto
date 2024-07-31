/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package polls

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	tableMsgCreatePollsParams = "stwart_polls_msg_create_polls_params"
	tableMsgUpdatePollsParams = "stwart_polls_msg_update_polls_params"
	tableMsgDeletePollsParams = "stwart_polls_msg_delete_polls_params"
	tableMsgCreatePolls       = "stwart_polls_msg_create_polls"
	tablePollsOptions         = "stwart_polls_options"
	tableMsgVote              = "stwart_polls_msg_vote"
)

type (
	// MsgCreatePollsParams - db model for 'stwart_polls_msg_create_polls_params'
	MsgCreatePollsParams struct {
		ID              uint64    `db:"id"`
		Creator         string    `db:"creator"`
		MinDaysDuration string    `db:"min_days_duration"`
		MaxDaysDuration string    `db:"max_days_duration"`
		MaxDaysPending  string    `db:"max_days_pending"`
		ProposerDeposit sdk.Coins `db:"proposer_deposit"`
		BurnVeto        bool      `db:"burn_veto"`
		TxHash          string    `db:"tx_hash"`
	}

	// MsgUpdatePollsParams - db model for 'stwart_polls_msg_update_polls_params'
	MsgUpdatePollsParams struct {
		ID              uint64    `db:"id"`
		Creator         string    `db:"creator"`
		MinDaysDuration string    `db:"min_days_duration"`
		MaxDaysDuration string    `db:"max_days_duration"`
		MaxDaysPending  string    `db:"max_days_pending"`
		ProposerDeposit sdk.Coins `db:"proposer_deposit"`
		BurnVeto        bool      `db:"burn_veto"`
		TxHash          string    `db:"tx_hash"`
	}

	// MsgDeletePollsParams - db model for 'stwart_polls_msg_delete_polls_params'
	MsgDeletePollsParams struct {
		ID      uint64 `db:"id"`
		Creator string `db:"creator"`
		TxHash  string `db:"tx_hash"`
	}

	// MsgCreatePolls - db model for 'stwart_polls_msg_create_polls'
	MsgCreatePolls struct {
		ID                 uint64 `db:"id"`
		Creator            string `db:"creator"`
		Title              string `db:"title"`
		Description        string `db:"description"`
		VotingStartTime    string `db:"voting_start_time"`
		VotingPeriod       string `db:"voting_period"`
		MinVoteAmount      uint64 `db:"min_vote_amount"`
		MinAdressesCount   uint64 `db:"min_adresses_count"`
		MinVoteCoinsAmount uint64 `db:"min_vote_coins_amount"`
		TxHash             string `db:"tx_hash"`
	}

	// Options - db model for 'stwart_polls_options'
	Options struct {
		ID           uint64    `db:"id"`
		PollID       uint64    `db:"poll_id"`
		VotersCount  uint64    `db:"voters_count"`
		TokensAmount sdk.Coins `db:"tokens_amount"`
		IsVeto       bool      `db:"is_veto"`
		Text         string    `db:"text"`
		IsWinner     bool      `db:"is_winner"`
	}

	// MsgVote - db model for 'stwart_polls_msg_vote'
	MsgVote struct {
		ID       uint64    `db:"id"`
		Creator  string    `db:"creator"`
		PollID   uint64    `db:"poll_id"`
		OptionID uint64    `db:"option_id"`
		Amount   sdk.Coins `db:"amount"`
		TxHash   string    `db:"tx_hash"`
	}
)
