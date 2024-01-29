package vmjsonintegrationtest

import (
	"testing"
)

func TestDelegation_v0_2(t *testing.T) {
	if testing.Short() {
		t.Skip("not a short test")
	}

	ScenariosTest(t).
		Folder("delegation/v0_2").
		Run().
		CheckNoError()
}

func TestDelegation_v0_3(t *testing.T) {
	if testing.Short() {
		t.Skip("not a short test")
	}

	ScenariosTest(t).
		Folder("delegation/v0_3").
		Exclude("delegation/v0_3/test/integration/genesis/genesis.scen.json").
		Exclude("delegation/v0_3/test/fuzz_gen_purchase_stake_rewards_issue.scen.json"). // TODO fix this
		Exclude("delegation/v0_3/test/fuzz_gen_unstake.scen.json").                      // TODO fix this
		Run().
		CheckNoError()
}

func TestDelegation_v0_4_genesis(t *testing.T) {
	if testing.Short() {
		t.Skip("not a short test")
	}

	ScenariosTest(t).
		Folder("delegation/v0_4_genesis").
		Run().
		CheckNoError()
}

func TestDelegation_v0_5_latest(t *testing.T) {
	if testing.Short() {
		t.Skip("not a short test")
	}

	ScenariosTest(t).
		Folder("delegation/v0_5_latest").
		Run().
		CheckNoError()
}
