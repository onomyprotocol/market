// Package client contains dao client implementation.
package client

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govrest "github.com/cosmos/cosmos-sdk/x/gov/client/rest"

	"github.com/pendulum-labs/market/x/market/client/cli"
)

var (
	// DenomMetadataProposalHandler is the cli handler used for the gov cli integration.
	DenomMetadataProposalHandler = govclient.NewProposalHandler(cli.CmdDenomMetadataProposal, emptyRestHandler) // nolint:gochecknoglobals // cosmos-sdk style
)

func emptyRestHandler(client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "unsupported-market-routes",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "Legacy REST Routes are not supported for Market proposals")
		},
	}
}
