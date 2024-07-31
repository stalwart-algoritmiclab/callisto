/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package polls

import "github.com/stalwart-algoritmiclab/stwart-chain-go/x/polls/types"

// toMsgCreatePollsParamsDomain - mapping db model to model
func (m MsgCreatePollsParams) toMsgCreatePollsParamsDomain() types.MsgCreatePollsParams {
	return types.MsgCreatePollsParams{
		Creator:         m.Creator,
		MinDaysDuration: m.MinDaysDuration,
		MaxDaysDuration: m.MaxDaysDuration,
		MaxDaysPending:  m.MaxDaysPending,
		ProposerDeposit: m.ProposerDeposit,
		BurnVeto:        m.BurnVeto,
	}
}

// toDatabaseMsgCreatePollsParams - mapping func to a database model.
func toDatabaseMsgCreatePollsParams(hash string, m *types.MsgCreatePollsParams) MsgCreatePollsParams {
	return MsgCreatePollsParams{
		Creator:         m.Creator,
		MinDaysDuration: m.MinDaysDuration,
		MaxDaysDuration: m.MaxDaysDuration,
		MaxDaysPending:  m.MaxDaysPending,
		ProposerDeposit: m.ProposerDeposit,
		BurnVeto:        m.BurnVeto,
		TxHash:          hash,
	}
}

// toMsgUpdatePollsPayloadDomain - mapping db model to model
func (m MsgUpdatePollsParams) toMsgUpdatePollsPayloadDomain() types.MsgUpdatePollsParams {
	return types.MsgUpdatePollsParams{
		Creator:         m.Creator,
		MinDaysDuration: m.MinDaysDuration,
		MaxDaysDuration: m.MaxDaysDuration,
		MaxDaysPending:  m.MaxDaysPending,
		ProposerDeposit: m.ProposerDeposit,
		BurnVeto:        m.BurnVeto,
	}
}

// toDatabaseMsgUpdatePollsParams - mapping func to a database model.
func toDatabaseMsgUpdatePollsParams(hash string, m *types.MsgUpdatePollsParams) MsgUpdatePollsParams {
	return MsgUpdatePollsParams{
		Creator:         m.Creator,
		MinDaysDuration: m.MinDaysDuration,
		MaxDaysDuration: m.MaxDaysDuration,
		MaxDaysPending:  m.MaxDaysPending,
		ProposerDeposit: m.ProposerDeposit,
		BurnVeto:        m.BurnVeto,
		TxHash:          hash,
	}
}

// toMsgDeletePollsPayloadDomain - mapping db model to model
func (m MsgDeletePollsParams) toMsgDeletePollsPayloadDomain() types.MsgDeletePollsParams {
	return types.MsgDeletePollsParams{
		Creator: m.Creator,
	}
}

// toDatabaseMsgDeletePollsParams - mapping func to a database model.
func toDatabaseMsgDeletePollsParams(hash string, m *types.MsgDeletePollsParams) MsgDeletePollsParams {
	return MsgDeletePollsParams{
		Creator: m.Creator,
		TxHash:  hash,
	}
}

// toMsgCreatePollsPayloadDomain - mapping db model to model
func (m MsgCreatePolls) toMsgCreatePollsPayloadDomain() types.MsgCreatePoll {
	return types.MsgCreatePoll{
		Creator:            m.Creator,
		Title:              m.Title,
		Description:        m.Description,
		VotingStartTime:    m.VotingStartTime,
		VotingPeriod:       m.VotingPeriod,
		MinVoteAmount:      m.MinVoteAmount,
		MinAdressesCount:   m.MinAdressesCount,
		MinVoteCoinsAmount: m.MinVoteCoinsAmount,
	}
}

// toDatabaseMsgCreatePolls - mapping func to a database model.
func toDatabaseMsgCreatePolls(hash string, m *types.MsgCreatePoll) MsgCreatePolls {
	return MsgCreatePolls{
		Creator:            m.Creator,
		Title:              m.Title,
		Description:        m.Description,
		VotingStartTime:    m.VotingStartTime,
		VotingPeriod:       m.VotingPeriod,
		MinVoteAmount:      m.MinVoteAmount,
		MinAdressesCount:   m.MinAdressesCount,
		MinVoteCoinsAmount: m.MinVoteCoinsAmount,
		TxHash:             hash,
	}
}

// toOptionsDomain - mapping db model to model
func (m Options) toOptionsDomain() types.Options {
	return types.Options{
		Id:           m.ID,
		VotersCount:  m.VotersCount,
		TokensAmount: m.TokensAmount,
		IsVeto:       m.IsVeto,
		Text:         m.Text,
		IsVinner:     m.IsWinner,
	}
}

// toDatabaseOptions - mapping func to a database model.
func toDatabaseOptions(m types.Options) Options {
	return Options{
		ID:           m.Id,
		VotersCount:  m.VotersCount,
		TokensAmount: m.TokensAmount,
		IsVeto:       m.IsVeto,
		Text:         m.Text,
		IsWinner:     m.IsVinner,
	}
}

// toMsgVoteDomain - mapping db model to model
func (m MsgVote) toMsgVoteDomain() types.MsgVote {
	return types.MsgVote{
		Creator:  m.Creator,
		PollId:   m.PollID,
		OptionId: m.OptionID,
		Amount:   m.Amount,
	}
}

// toDatabaseMsgVote - mapping func to a database model.
func toDatabaseMsgVote(hash string, m *types.MsgVote) MsgVote {
	return MsgVote{
		Creator:  m.Creator,
		PollID:   m.PollId,
		OptionID: m.OptionId,
		Amount:   m.Amount,
		TxHash:   hash,
	}
}

// toMsgCreatePollsParamsDomainList - mapping db model to model
func toMsgCreatePollsParamsDomainList(m []MsgCreatePollsParams) []types.MsgCreatePollsParams {
	res := make([]types.MsgCreatePollsParams, 0, len(m))
	for _, msg := range m {
		res = append(res, msg.toMsgCreatePollsParamsDomain())
	}

	return res
}

// toMsgUpdatePollsParamsDomainList - mapping db model to model
func toMsgUpdatePollsParamsDomainList(m []MsgUpdatePollsParams) []types.MsgUpdatePollsParams {
	res := make([]types.MsgUpdatePollsParams, 0, len(m))
	for _, msg := range m {
		res = append(res, msg.toMsgUpdatePollsPayloadDomain())
	}

	return res
}

// toMsgDeletePollsParamsDomainList - mapping db model to model
func toMsgDeletePollsParamsDomainList(m []MsgDeletePollsParams) []types.MsgDeletePollsParams {
	res := make([]types.MsgDeletePollsParams, 0, len(m))
	for _, msg := range m {
		res = append(res, msg.toMsgDeletePollsPayloadDomain())
	}

	return res
}

// toMsgVoteDomainList - mapping db model to model
func toMsgVoteDomainList(m []MsgVote) []types.MsgVote {
	res := make([]types.MsgVote, 0, len(m))
	for _, msg := range m {
		res = append(res, msg.toMsgVoteDomain())
	}

	return res
}

// toOptionsDomainList - mapping db model to model
func toOptionsDomainList(m []Options) []types.Options {
	res := make([]types.Options, 0, len(m))
	for _, msg := range m {
		res = append(res, msg.toOptionsDomain())
	}

	return res
}

// toDatabaseOptionsList - mapping func to a database model.
func toDatabaseOptionsList(m []types.Options) []Options {
	res := make([]Options, 0, len(m))
	for _, msg := range m {
		res = append(res, toDatabaseOptions(msg))
	}

	return res
}
