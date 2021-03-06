package gov

import (
	"log"
	"net/http"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"

	_cli "github.com/ironman0x7b2/client/cli"
	"github.com/ironman0x7b2/client/handlers/common"
	"github.com/ironman0x7b2/client/messages"
	"github.com/ironman0x7b2/client/models"
	"github.com/ironman0x7b2/client/utils"
)

/**
 * @api {get} /proposals get all proposals
 * @apiDescription Used to get all proposals
 * @apiName GetAllProposals
 * @apiGroup gov
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */
const MODULE = "gov"

func getAllProposalsHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var limit uint64

		if l := r.URL.Query().Get("limit"); len(l) != 0 {
			i, err := strconv.ParseUint(l, 10, 64)
			if err != nil {
				utils.WriteErrorToResponse(w, 400, common.ErrorParseQueryParams(MODULE))

				log.Println(err.Error())
				return
			}
			limit = i
		}

		proposals, err := cli.GetAllProposals(limit)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, errorFailedToGetProposals())

			log.Println(err.Error())
			return
		}

		utils.WriteResultToResponse(w, 200, proposals)
	}
}

/**
 * @api {get} /proposals/{id} get proposal
 * @apiDescription Used to get Proposal
 * @apiName GetProposal
 * @apiGroup gov
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

func getProposalHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		var pid uint64
		id := vars["id"]
		if len(id) != 0 {
			_id, err := strconv.ParseUint(id, 10, 64)
			if err != nil {
				utils.WriteErrorToResponse(w, 400, common.ErrorParseQueryParams(MODULE))

				log.Println(err.Error())
				return
			}
			pid = _id
		}

		proposal, err := cli.GetProposal(pid)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, errorFailedToGetProposal())

			log.Println(err.Error())
			return
		}

		utils.WriteResultToResponse(w, 200, proposal)
	}
}

/**
 * @api {get} /proposals/{id}/votes get proposal votes
 * @apiDescription Used to get proposal votes
 * @apiName GetProposalVotes
 * @apiGroup gov
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

func getProposalVotesHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		var pid uint64
		id := vars["id"]
		if len(id) != 0 {
			_id, err := strconv.ParseUint(id, 10, 64)
			if err != nil {
				utils.WriteErrorToResponse(w, 400, common.ErrorParseQueryParams(MODULE))

				log.Println(err.Error())
				return
			}
			pid = _id
		}

		votes, err := cli.GetProposalVotes(pid)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, errorFailedToGetProposalVotes())

			log.Println(err.Error())
			return
		}

		_votes := models.NewVotesFromRaw(votes)
		utils.WriteResultToResponse(w, 200, _votes)
	}
}

/**
 * @api {get} //proposals/{id}/voters/{address} get proposal voter
 * @apiDescription Used to get proposal voter
 * @apiName GetProposalVoter
 * @apiGroup gov
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

func getProposalVoteHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		var pid uint64
		var vAddress sdk.AccAddress

		id := vars["id"]
		if len(id) != 0 {
			_id, err := strconv.ParseUint(id, 10, 64)
			if err != nil {
				utils.WriteErrorToResponse(w, 400, common.ErrorParseQueryParams(MODULE))

				log.Println(err.Error())
				return
			}
			pid = _id
		}

		address := vars["address"]
		if len(id) != 0 {
			_address, err := sdk.AccAddressFromHex(address)
			if err != nil {
				utils.WriteErrorToResponse(w, 400, common.ErrorDecodeAddress(MODULE))

				log.Println(err.Error())
				return
			}
			vAddress = _address
		}

		vote, err := cli.GetProposalVote(pid, vAddress)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, errorFailedToGetProposalVote())

			log.Println(err.Error())
			return
		}

		_vote := models.NewVoteFromRaw(vote)
		utils.WriteResultToResponse(w, 200, _vote)
	}
}

/**
 * @api {post} /proposals submit proposal
 * @apiDescription Used to submit proposal
 * @apiName SubmitProposal
 * @apiGroup gov
 * @apiParamExample {json} Request-Example:
 * {
 *	"from":"Name",
 *	"from_address":"4CC1DA947C678D6DD1E375D9AF1674C2B633D25B",
 *  "title":"Title",
 *  "description":"Description",
 *  "type":"Text",
 *	"amount":[{"denom":"tsent","value":10}],
 *	"gas":210000,
 *	"password":"password"
 * }
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

func submitProposalHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := newProposal(r)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, common.ErrorParseRequestBody(MODULE))

			log.Println(err.Error())
			return
		}

		if err = body.Validate(); err != nil {
			utils.WriteErrorToResponse(w, 400, common.ErrorValidateRequestBody(MODULE))

			log.Println(err.Error())
			return
		}

		msg, err := messages.NewProposal(body.FromAddress, body.Title, body.Description, body.Type, body.Amount).Raw()
		if err != nil {
			utils.WriteErrorToResponse(w, 400, common.ErrorFailedToPrepareMsg(MODULE))

			log.Println(err.Error())
			return
		}

		cli.CLIContext = cli.WithFromName(body.From)

		res, _err := cli.Tx([]sdk.Msg{msg}, body.Memo, body.Gas, body.GasAdjustment,
			body.GasPrices.Raw(), body.Fees.Raw(), body.Password)
		if _err != nil {
			utils.WriteErrorToResponse(w, 400, common.ErrorFailedToBroadcastTransaction(MODULE))

			log.Println(_err.Error())
			return
		}

		utils.WriteResultToResponse(w, 200, res)
	}
}

/**
 * @api {post} /proposals/{id}/deposits proposal deposits
 * @apiDescription Used to deposit amount for proposal
 * @apiName ProposalDeposits
 * @apiGroup gov
 * @apiParamExample {json} Request-Example:
 * {
 *	"from":"Name",
 *	"from_address":"4CC1DA947C678D6DD1E375D9AF1674C2B633D25B",
 *	"amount":[{"denom":"tsent","value":10}],
 *	"gas":210000,
 *	"password":"password"
 * }
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

func proposalDepositsHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		body, err := newProposalDeposits(r)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, common.ErrorParseRequestBody(MODULE))

			log.Println(err.Error())
			return
		}

		if err = body.Validate(); err != nil {
			utils.WriteErrorToResponse(w, 400, common.ErrorValidateRequestBody(MODULE))

			log.Println(err.Error())
			return
		}

		id, err := strconv.ParseUint(vars["id"], 10, 64)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, common.ErrorParseQueryParams(MODULE))

			log.Println(err.Error())
			return
		}

		msg, err := messages.NewProposalDeposits(body.FromAddress, id, body.Amount).Raw()
		if err != nil {
			utils.WriteErrorToResponse(w, 400, common.ErrorFailedToPrepareMsg(MODULE))

			log.Println(err.Error())
			return
		}

		cli.CLIContext = cli.WithFromName(body.From)

		res, _err := cli.Tx([]sdk.Msg{msg}, body.Memo, body.Gas, body.GasAdjustment,
			body.GasPrices.Raw(), body.Fees.Raw(), body.Password)
		if _err != nil {
			utils.WriteErrorToResponse(w, 400, common.ErrorFailedToBroadcastTransaction(MODULE))

			log.Println(_err.Error())
			return
		}

		utils.WriteResultToResponse(w, 200, res)
	}
}

/**
 * @api {post} /proposals/{id}/votes proposal votes
 * @apiDescription Used to submit the vote for proposal
 * @apiName ProposalVotes
 * @apiGroup gov
 * @apiParamExample {json} Request-Example:
 * {
 *	"from":"Name",
 *	"from_address":"4CC1DA947C678D6DD1E375D9AF1674C2B633D25B",
 *  "option":"yes",
 *	"gas":210000,
 *	"password":"password"
 * }
 * @apiSuccess {Boolean} success Success key.
 * @apiSuccess {object} result Success object.
 */

func proposalVotesHandler(cli *_cli.CLI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		body, err := newProposalVotes(r)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, common.ErrorParseRequestBody(MODULE))

			log.Println(err.Error())
			return
		}

		if err = body.Validate(); err != nil {
			utils.WriteErrorToResponse(w, 400, common.ErrorValidateRequestBody(MODULE))

			log.Println(err.Error())
			return
		}

		id, err := strconv.ParseUint(vars["id"], 10, 64)
		if err != nil {
			utils.WriteErrorToResponse(w, 400, common.ErrorParseQueryParams(MODULE))

			log.Println(err.Error())
			return
		}

		msg, err := messages.NewProposalVotes(body.FromAddress, id, body.Option).Raw()
		if err != nil {
			utils.WriteErrorToResponse(w, 400, common.ErrorFailedToPrepareMsg(MODULE))

			log.Println(err.Error())
			return
		}

		cli.CLIContext = cli.WithFromName(body.From)

		res, _err := cli.Tx([]sdk.Msg{msg}, body.Memo, body.Gas, body.GasAdjustment,
			body.GasPrices.Raw(), body.Fees.Raw(), body.Password)
		if _err != nil {
			utils.WriteErrorToResponse(w, 400, common.ErrorFailedToBroadcastTransaction(MODULE))

			log.Println(_err.Error())
			return
		}

		utils.WriteResultToResponse(w, 200, res)
	}
}
