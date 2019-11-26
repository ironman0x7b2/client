package gov

import (
	"net/http"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"

	_cli "github.com/ironman0x7b2/client/cli"
	"github.com/ironman0x7b2/client/messages"
	"github.com/ironman0x7b2/client/models"
	"github.com/ironman0x7b2/client/types"
	"github.com/ironman0x7b2/client/utils"
)

func getAllProposalsHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var limit uint64

		if l := r.URL.Query().Get("limit"); len(l) != 0 {
			i, err := strconv.ParseUint(l, 10, 64)
			if err != nil {
				utils.WriteErrorToResponse(w, 400, err)
				return
			}
			limit = i
		}

		proposals, err := cli.GetAllProposals(limit)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, err)
			return
		}

		utils.WriteResultToResponse(w, 200, proposals)
	}
}

func getProposalHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		var pid uint64
		id := vars["id"]
		if len(id) != 0 {
			_id, err := strconv.ParseUint(id, 10, 64)
			if err != nil {
				utils.WriteErrorToResponse(w, 400, err)
				return
			}
			pid = _id
		}

		proposal, err := cli.GetProposal(pid)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, err)
			return
		}

		utils.WriteResultToResponse(w, 200, proposal)
	}
}

func getProposalVotesHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		var pid uint64
		id := vars["id"]
		if len(id) != 0 {
			_id, err := strconv.ParseUint(id, 10, 64)
			if err != nil {
				utils.WriteErrorToResponse(w, 400, err)
				return
			}
			pid = _id
		}

		votes, err := cli.GetProposalVotes(pid)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, err)
			return
		}

		_votes := models.NewVotesFromRaw(votes)
		utils.WriteResultToResponse(w, 200, _votes)
	}
}

func getProposalVoteHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		var pid uint64
		var vAddress sdk.AccAddress

		id := vars["id"]
		if len(id) != 0 {
			_id, err := strconv.ParseUint(id, 10, 64)
			if err != nil {
				utils.WriteErrorToResponse(w, 400, err)
				return
			}
			pid = _id
		}

		address := vars["address"]
		if len(id) != 0 {
			_address, err := sdk.AccAddressFromHex(address)
			if err != nil {
				utils.WriteErrorToResponse(w, 400, err)
				return
			}
			vAddress = _address
		}

		vote, err := cli.GetProposalVote(pid, vAddress)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, err)
			return
		}

		_vote := models.NewVoteFromRaw(vote)
		utils.WriteResultToResponse(w, 200, _vote)
	}
}

func submitProposalHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := newProposal(r)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to parse the request body",
				Info:    err.Error(),
			})
			return
		}

		if err = body.Validate(); err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to validate request body",
				Info:    err.Error(),
			})
			return
		}

		msg, err := messages.NewProposal(body.FromAddress, body.Title, body.Description, body.Type, body.Amount).Raw()
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to prepare the transfer message",
				Info:    err.Error(),
			})
			return
		}

		cli.CLIContext = cli.WithFromName(body.From)

		res, err := cli.Tx([]sdk.Msg{msg}, body.Memo, body.Gas, body.GasAdjustment,
			body.GasPrices.Raw(), body.Fees.Raw(), body.Password)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to broadcast the transaction",
				Info:    err.Error(),
			})
			return
		}

		utils.WriteResultToResponse(w, 200, res)
	}
}

func proposalDepositsHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		body, err := newProposalDeposits(r)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to parse the request body",
				Info:    err.Error(),
			})
			return
		}

		if err = body.Validate(); err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to validate request body",
				Info:    err.Error(),
			})
			return
		}

		id, err := strconv.ParseUint(vars["id"], 10, 64)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to convert id type",
				Info:    err.Error(),
			})
			return
		}

		msg, err := messages.NewProposalDeposits(body.FromAddress, id, body.Amount).Raw()
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to prepare the transfer message",
				Info:    err.Error(),
			})
			return
		}

		cli.CLIContext = cli.WithFromName(body.From)

		res, err := cli.Tx([]sdk.Msg{msg}, body.Memo, body.Gas, body.GasAdjustment,
			body.GasPrices.Raw(), body.Fees.Raw(), body.Password)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to broadcast the transaction",
				Info:    err.Error(),
			})
			return
		}

		utils.WriteResultToResponse(w, 200, res)
	}
}

func proposalVotesHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		body, err := newProposalVotes(r)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to parse the request body",
				Info:    err.Error(),
			})
			return
		}

		if err = body.Validate(); err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to validate request body",
				Info:    err.Error(),
			})
			return
		}

		id, err := strconv.ParseUint(vars["id"], 10, 64)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to convert id type",
				Info:    err.Error(),
			})
			return
		}

		msg, err := messages.NewProposalVotes(body.FromAddress, id, body.Option).Raw()
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to prepare the transfer message",
				Info:    err.Error(),
			})
			return
		}

		cli.CLIContext = cli.WithFromName(body.From)

		res, err := cli.Tx([]sdk.Msg{msg}, body.Memo, body.Gas, body.GasAdjustment,
			body.GasPrices.Raw(), body.Fees.Raw(), body.Password)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, &types.Error{
				Message: "failed to broadcast the transaction",
				Info:    err.Error(),
			})
			return
		}

		utils.WriteResultToResponse(w, 200, res)
	}
}
