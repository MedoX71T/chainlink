//go:build integration

package mercury_test

import (
	"testing"

	"github.com/smartcontractkit/chainlink/v2/core/config/env"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
)

func TestIntegration_MercuryV1_Plugin(t *testing.T) {
	t.Setenv(string(env.MercuryPlugin.Cmd), "chainlink-mercury")

	integration_MercuryV1(t)
}

func TestIntegration_MercuryV2_Plugin(t *testing.T) {
	testutils.SkipFlakey(t, "https://smartcontract-it.atlassian.net/browse/MERC-5697")
	t.Setenv(string(env.MercuryPlugin.Cmd), "chainlink-mercury")
	integration_MercuryV2(t)
}

func TestIntegration_MercuryV3_Plugin(t *testing.T) {
	t.Setenv(string(env.MercuryPlugin.Cmd), "chainlink-mercury")

	integration_MercuryV3(t)
}
