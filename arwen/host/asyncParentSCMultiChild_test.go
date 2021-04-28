package host

import (
	"errors"
	"math/big"
	"testing"

	mock "github.com/ElrondNetwork/arwen-wasm-vm/mock/context"
	"github.com/ElrondNetwork/elrond-go/testscommon/txDataBuilder"
	"github.com/stretchr/testify/require"
)

func createTestAsyncMultiChildParentContract(t testing.TB, host *vmHost, imb *mock.InstanceBuilderMock, testConfig *asyncCallMultiChildTestConfig) {
	parentInstance := imb.CreateAndStoreInstanceMock(t, host, parentAddress, testConfig.parentBalance)
	addAsyncMultiChildParentMethodsToInstanceMock(parentInstance, testConfig)
}

func addAsyncMultiChildParentMethodsToInstanceMock(instanceMock *mock.InstanceMock, testConfig *asyncCallMultiChildTestConfig) {
	input := DefaultTestContractCallInput()
	input.GasProvided = testConfig.gasProvidedToChild

	instanceMock.AddMockMethodWithError("forwardAsyncCall", func() {
		host := instanceMock.Host
		instance := mock.GetMockInstance(host)
		t := instance.T
		arguments := host.Runtime().Arguments()
		destination := arguments[0]
		function := string(arguments[1])
		value := big.NewInt(testConfig.transferFromParentToChild).Bytes()

		host.Metering().UseGas(testConfig.gasUsedByParent)

		for childCall := 0; childCall < testConfig.childCalls; childCall++ {
			callData := txDataBuilder.NewBuilder()
			callData.Func(function)
			// recursiveChildCalls
			callData.BigInt(big.NewInt(1))

			err := host.Runtime().ExecuteAsyncCall(destination, callData.ToBytes(), value)
			require.Nil(t, err)
		}
	}, errors.New("breakpoint / failed to call function"))

	instanceMock.AddMockMethod("callBack", func() {
		host := instanceMock.Host
		host.Metering().UseGas(testConfig.gasUsedByCallback)
	})
}
