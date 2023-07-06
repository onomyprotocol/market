package app_test

import (
	"testing"

	appProvider "github.com/cosmos/interchain-security/app/provider"
	e2e "github.com/cosmos/interchain-security/tests/e2e"
	icssimapp "github.com/cosmos/interchain-security/testutil/ibc_testing"
	appConsumer "github.com/pendulum-labs/market/app/consumer"
	"github.com/stretchr/testify/suite"
)

// Executes the standard group of ccv tests against a consumer and provider app.go implementation.
func TestCCVTestSuite(t *testing.T) {
	// Pass in concrete app types that implement the interfaces defined in /testutil/e2e/interfaces.go
	ccvSuite := e2e.NewCCVTestSuite[*appProvider.App, *appConsumer.App](
		// Pass in ibctesting.AppIniters for provider and consumer.
		icssimapp.ProviderAppIniter, DemocracyConsumerAppIniter,
		// put tests that we want to skip here
		[]string{"TestRewardsDistribution"})

	// Run tests
	suite.Run(t, ccvSuite)
}
