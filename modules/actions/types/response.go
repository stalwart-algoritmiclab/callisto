/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package types

import (
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	stakingtype "github.com/cosmos/cosmos-sdk/x/staking/types"
	polls "github.com/stalwart-algoritmiclab/stwart-chain-go/x/polls/types"
)

type Coin struct {
	Amount string `json:"amount"`
	Denom  string `json:"denom"`
}

func ConvertCoins(coins sdk.Coins) []Coin {
	amount := make([]Coin, 0)
	for _, coin := range coins {
		amount = append(amount, Coin{Amount: coin.Amount.String(), Denom: coin.Denom})
	}
	return amount
}

func ConvertDecCoins(coins sdk.DecCoins) []Coin {
	amount := make([]Coin, 0)
	for _, coin := range coins {
		amount = append(amount, Coin{Amount: coin.Amount.String(), Denom: coin.Denom})
	}
	return amount
}

// ========================= Withdraw Address Response =========================

type Address struct {
	Address string `json:"address"`
}

// ========================= Account Balance Response =========================

type Balance struct {
	Coins []Coin `json:"coins"`
}

// ========================= Delegation Response =========================

type DelegationResponse struct {
	Delegations []Delegation        `json:"delegations"`
	Pagination  *query.PageResponse `json:"pagination"`
}

type Delegation struct {
	DelegatorAddress string `json:"delegator_address"`
	ValidatorAddress string `json:"validator_address"`
	Coins            []Coin `json:"coins"`
}

// ========================= Delegation Reward Response =========================

type DelegationReward struct {
	Coins            []Coin `json:"coins"`
	ValidatorAddress string `json:"validator_address"`
}

// ========================= Validator Commission Response =========================

type ValidatorCommissionAmount struct {
	Coins []Coin `json:"coins"`
}

// ========================= Unbonding Delegation Response =========================

type UnbondingDelegationResponse struct {
	UnbondingDelegations []UnbondingDelegation `json:"unbonding_delegations"`
	Pagination           *query.PageResponse   `json:"pagination"`
}

type UnbondingDelegation struct {
	DelegatorAddress string                                 `json:"delegator_address"`
	ValidatorAddress string                                 `json:"validator_address"`
	Entries          []stakingtype.UnbondingDelegationEntry `json:"entries"`
}

// ========================= Redelegation Response =========================

type RedelegationResponse struct {
	Redelegations []Redelegation      `json:"redelegations"`
	Pagination    *query.PageResponse `json:"pagination"`
}

type Redelegation struct {
	DelegatorAddress    string              `json:"delegator_address"`
	ValidatorSrcAddress string              `json:"validator_src_address"`
	ValidatorDstAddress string              `json:"validator_dst_address"`
	RedelegationEntries []RedelegationEntry `json:"entries"`
}

type RedelegationEntry struct {
	CompletionTime time.Time `json:"completion_time"`
	Balance        math.Int  `json:"balance"`
}

// ========================= Polls Response =========================

// PollsResponse represents the response of the PollsAll query
type PollsResponse struct {
	Polls      []Poll              `json:"polls"`
	Pagination *query.PageResponse `json:"pagination"`
}

// Poll represents the poll
type Poll struct {
	ID                  uint64   `json:"id"`
	Title               string   `json:"title"`
	Description         string   `json:"description"`
	ProposerAddress     string   `json:"proposer_address"`
	VotingStartTime     string   `json:"voting_start_time"`
	VotingEndTime       string   `json:"voting_end_time"`
	VotingPeriod        string   `json:"voting_period"`
	MinVoteAmount       []Coin   `json:"min_vote_amount"`
	Status              string   `json:"status"`
	FailureReason       string   `json:"failure_reason"`
	MinAddressesCount   uint64   `json:"min_addresses_count"`
	MinVotedCoinsAmount []Coin   `json:"min_voted_coins_amount"`
	Options             []Option `json:"options"`
}

// Option represents the option of a poll
type Option struct {
	ID           uint64 `json:"id"`
	VotersCount  uint64 `json:"voters_count"`
	TokensAmount []Coin `json:"tokens_amount"`
	IsVeto       bool   `json:"is_veto"`
	Text         string `json:"text"`
	IsWinner     bool   `json:"is_winner"`
}

// ToPollResponse converts a Polls object to a Poll object
func ToPollResponse(poll polls.Polls) Poll {
	return Poll{
		ID:                  poll.Id,
		Title:               poll.Title,
		Description:         poll.Description,
		ProposerAddress:     poll.ProposerAddress,
		VotingStartTime:     poll.VotingStartTime,
		VotingEndTime:       poll.VotingEndTime,
		VotingPeriod:        poll.VotingPeriod,
		MinVoteAmount:       ConvertCoins(poll.MinVoteAmount),
		Status:              poll.Status,
		FailureReason:       poll.FailureReason,
		MinAddressesCount:   poll.MinAddressesCount,
		MinVotedCoinsAmount: ConvertCoins(poll.MinVotedCoinsAmount),
		Options:             ToOptionsResponse(poll.Options),
	}
}

// ToPollsResponse converts a slice of Polls objects to a slice of Poll objects
func ToPollsResponse(polls []polls.Polls) []Poll {
	result := make([]Poll, 0, len(polls))
	for _, poll := range polls {
		result = append(result, ToPollResponse(poll))
	}
	return result
}

// ToOptionResponse converts a Options object to a Option object
func ToOptionResponse(option *polls.Options) Option {
	return Option{
		ID:           option.Id,
		VotersCount:  option.VotersCount,
		TokensAmount: ConvertCoins(option.TokensAmount),
		IsVeto:       option.IsVeto,
		Text:         option.Text,
		IsWinner:     option.IsWinner,
	}
}

// ToOptionsResponse converts a slice of Options objects to a slice of Option objects
func ToOptionsResponse(options []*polls.Options) []Option {
	result := make([]Option, 0, len(options))
	for _, option := range options {
		result = append(result, ToOptionResponse(option))
	}
	return result
}
