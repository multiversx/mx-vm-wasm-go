package hostCoretest

import (
	"bytes"
	"crypto/ed25519"
	"crypto/elliptic"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/multiversx/mx-chain-vm-go/crypto/hashing"
	"github.com/multiversx/mx-chain-vm-go/crypto/signing/secp256k1"
	mock "github.com/multiversx/mx-chain-vm-go/mock/context"
	"github.com/multiversx/mx-chain-vm-go/mock/contracts"
	"github.com/multiversx/mx-chain-vm-go/testcommon"
	test "github.com/multiversx/mx-chain-vm-go/testcommon"
	"github.com/multiversx/mx-chain-vm-go/vmhost"
	"github.com/multiversx/mx-chain-vm-go/vmhost/vmhooks"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/data/esdt"
	"github.com/multiversx/mx-chain-core-go/data/vm"
	"github.com/multiversx/mx-chain-scenario-go/worldmock"
	"github.com/multiversx/mx-chain-scenario-go/worldmock/esdtconvert"
	vmcommon "github.com/multiversx/mx-chain-vm-common-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var baseTestConfig = &testcommon.TestConfig{
	GasProvided:     1000,
	GasUsedByParent: 400,
	GasUsedByChild:  200,

	ParentBalance: 1000,
	ChildBalance:  1000,
}

func Test_ManagedIsESDTFrozen_NotFrozen(t *testing.T) {
	testManagedIsESDTFrozen(t, false)
}

func Test_ManagedIsESDTFrozen_Frozen(t *testing.T) {
	testManagedIsESDTFrozen(t, true)
}

func testManagedIsESDTFrozen(t *testing.T, isFrozen bool) {
	testConfig := baseTestConfig

	var addressHandle, tokenIDHandle int32
	var nonce int64

	expectedFrozen := int64(0)
	if isFrozen {
		expectedFrozen = 1
	}

	_, err := test.BuildMockInstanceCallTest(t).
		WithContracts(
			test.CreateMockContract(test.ParentAddress).
				WithBalance(testConfig.ParentBalance).
				WithConfig(testConfig).
				WithMethods(func(parentInstance *mock.InstanceMock, config interface{}) {
					parentInstance.AddMockMethod("testFunction", func() *mock.InstanceMock {
						host := parentInstance.Host

						managedTypes := host.ManagedTypes()
						addressHandle = managedTypes.NewManagedBufferFromBytes(test.ParentAddress)
						tokenIDHandle = managedTypes.NewManagedBufferFromBytes(test.ESDTTestTokenName)

						retValue := vmhooks.ManagedIsESDTFrozenWithHost(
							host,
							addressHandle,
							tokenIDHandle,
							nonce)

						host.Output().Finish(big.NewInt(int64(retValue)).Bytes())
						return parentInstance
					})
				}),
		).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithRecipientAddr(test.ParentAddress).
			WithGasProvided(testConfig.GasProvided).
			WithFunction("testFunction").
			Build()).
		WithSetup(func(host vmhost.VMHost, world *worldmock.MockWorld) {
			createMockBuiltinFunctions(t, host, world)
			err := world.BuiltinFuncs.SetTokenData(
				test.ParentAddress,
				test.ESDTTestTokenName,
				0,
				&esdt.ESDigitalToken{
					Value:      big.NewInt(100),
					Type:       uint32(core.Fungible),
					Properties: esdtconvert.MakeESDTUserMetadataBytes(isFrozen),
				})
			require.Nil(t, err)
		}).
		AndAssertResults(func(world *worldmock.MockWorld, verify *test.VMOutputVerifier) {
			verify.
				Ok().
				ReturnData(big.NewInt(expectedFrozen).Bytes())
		})
	assert.Nil(t, err)
}

func Test_ManagedIsESDTFrozen_IsPaused(t *testing.T) {
	testManagedIsESDTFrozenIsPaused(t, true)
}

func Test_ManagedIsESDTFrozen_IsNotPaused(t *testing.T) {
	testManagedIsESDTFrozenIsPaused(t, false)
}

func testManagedIsESDTFrozenIsPaused(t *testing.T, isPaused bool) {
	testConfig := baseTestConfig

	var tokenIDHandle int32

	expectedPaused := int64(0)
	if isPaused {
		expectedPaused = 1
	}

	_, err := test.BuildMockInstanceCallTest(t).
		WithContracts(
			test.CreateMockContract(test.ParentAddress).
				WithBalance(testConfig.ParentBalance).
				WithConfig(testConfig).
				WithMethods(func(parentInstance *mock.InstanceMock, config interface{}) {
					parentInstance.AddMockMethod("testFunction", func() *mock.InstanceMock {
						host := parentInstance.Host

						managedTypes := host.ManagedTypes()
						tokenIDHandle = managedTypes.NewManagedBufferFromBytes(test.ESDTTestTokenName)

						retValue := vmhooks.ManagedIsESDTPausedWithHost(
							host,
							tokenIDHandle)

						host.Output().Finish(big.NewInt(int64(retValue)).Bytes())
						return parentInstance
					})
				}),
		).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithRecipientAddr(test.ParentAddress).
			WithGasProvided(testConfig.GasProvided).
			WithFunction("testFunction").
			Build()).
		WithSetup(func(host vmhost.VMHost, world *worldmock.MockWorld) {
			world.IsPausedValue = isPaused
		}).
		AndAssertResults(func(world *worldmock.MockWorld, verify *test.VMOutputVerifier) {
			verify.
				Ok().
				ReturnData(big.NewInt(expectedPaused).Bytes())
		})
	assert.Nil(t, err)
}

func Test_ManagedIsESDTFrozen_IsLimitedTransfer(t *testing.T) {
	testManagedIsESDTFrozenIsLimitedTransfer(t, true)
}

func Test_ManagedIsESDTFrozen_IsNotLimitedTransfer(t *testing.T) {
	testManagedIsESDTFrozenIsLimitedTransfer(t, false)
}

func testManagedIsESDTFrozenIsLimitedTransfer(t *testing.T, isLimitedTransfer bool) {
	testConfig := baseTestConfig

	var tokenIDHandle int32

	expectedLimitedTransfer := int64(0)
	if isLimitedTransfer {
		expectedLimitedTransfer = 1
	}

	_, err := test.BuildMockInstanceCallTest(t).
		WithContracts(
			test.CreateMockContract(test.ParentAddress).
				WithBalance(testConfig.ParentBalance).
				WithConfig(testConfig).
				WithMethods(func(parentInstance *mock.InstanceMock, config interface{}) {
					parentInstance.AddMockMethod("testFunction", func() *mock.InstanceMock {
						host := parentInstance.Host

						managedTypes := host.ManagedTypes()
						tokenIDHandle = managedTypes.NewManagedBufferFromBytes(test.ESDTTestTokenName)

						retValue := vmhooks.ManagedIsESDTLimitedTransferWithHost(
							host,
							tokenIDHandle)

						host.Output().Finish(big.NewInt(int64(retValue)).Bytes())
						return parentInstance
					})
				}),
		).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithRecipientAddr(test.ParentAddress).
			WithGasProvided(testConfig.GasProvided).
			WithFunction("testFunction").
			Build()).
		WithSetup(func(host vmhost.VMHost, world *worldmock.MockWorld) {
			world.IsLimitedTransferValue = isLimitedTransfer
		}).
		AndAssertResults(func(world *worldmock.MockWorld, verify *test.VMOutputVerifier) {
			verify.
				Ok().
				ReturnData(big.NewInt(expectedLimitedTransfer).Bytes())
		})
	assert.Nil(t, err)
}

func Test_ManagedBufferToHex(t *testing.T) {
	testConfig := baseTestConfig

	asBytes := []byte{1, 2, 3}
	asString := "010203"

	_, err := test.BuildMockInstanceCallTest(t).
		WithContracts(
			test.CreateMockContract(test.ParentAddress).
				WithBalance(testConfig.ParentBalance).
				WithConfig(testConfig).
				WithMethods(func(parentInstance *mock.InstanceMock, config interface{}) {
					parentInstance.AddMockMethod("testFunction", func() *mock.InstanceMock {
						host := parentInstance.Host

						managedTypes := host.ManagedTypes()
						sourceHandle := managedTypes.NewManagedBufferFromBytes(asBytes)
						destHandle := managedTypes.NewManagedBuffer()

						vmhooks.ManagedBufferToHexWithHost(
							host,
							sourceHandle,
							destHandle)

						bytesResult, _ := managedTypes.GetBytes(destHandle)
						if string(bytesResult) != asString {
							host.Runtime().SignalUserError("assert failed")
							return parentInstance
						}

						return parentInstance
					})
				}),
		).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithRecipientAddr(test.ParentAddress).
			WithGasProvided(testConfig.GasProvided).
			WithFunction("testFunction").
			Build()).
		AndAssertResults(func(world *worldmock.MockWorld, verify *test.VMOutputVerifier) {
			verify.
				Ok()
		})
	assert.Nil(t, err)
}

func Test_BigIntToString(t *testing.T) {
	testConfig := baseTestConfig

	asBigInt := big.NewInt(1234567890)
	asString := "1234567890"

	_, err := test.BuildMockInstanceCallTest(t).
		WithContracts(
			test.CreateMockContract(test.ParentAddress).
				WithBalance(testConfig.ParentBalance).
				WithConfig(testConfig).
				WithMethods(func(parentInstance *mock.InstanceMock, config interface{}) {
					parentInstance.AddMockMethod("testFunction", func() *mock.InstanceMock {
						host := parentInstance.Host

						managedTypes := host.ManagedTypes()
						sourceHandle := managedTypes.NewBigInt(asBigInt)
						destHandle := managedTypes.NewManagedBuffer()

						vmhooks.BigIntToStringWithHost(
							host,
							sourceHandle,
							destHandle)

						bytesResult, _ := managedTypes.GetBytes(destHandle)
						if string(bytesResult) != asString {
							host.Runtime().SignalUserError("assert failed")
							return parentInstance
						}

						return parentInstance
					})
				}),
		).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithRecipientAddr(test.ParentAddress).
			WithGasProvided(testConfig.GasProvided).
			WithFunction("testFunction").
			Build()).
		AndAssertResults(func(world *worldmock.MockWorld, verify *test.VMOutputVerifier) {
			verify.
				Ok()
		})
	assert.Nil(t, err)
}

// Real contracts always check first that the big int fits.
// This special test case represents an intentionally badly written contract.
func bigIntToInt64MockContract(parentInstance *mock.InstanceMock, _ interface{}) {
	parentInstance.AddMockMethod("testFunction", func() *mock.InstanceMock {
		vmHooksImpl := vmhooks.NewVMHooksImpl(parentInstance.Host)

		inputHandle := int32(0)
		vmHooksImpl.BigIntGetSignedArgument(0, inputHandle)
		result := vmHooksImpl.BigIntGetInt64(inputHandle)
		if result < 0 {
			return parentInstance
		}

		vmHooksImpl.SmallIntFinishSigned(result)

		return parentInstance
	})
}

func Test_BigIntToInt64(t *testing.T) {
	testConfig := baseTestConfig

	bigIntArg := []byte{0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

	_, err := test.BuildMockInstanceCallTest(t).
		WithContracts(
			test.CreateMockContract(test.ParentAddress).
				WithBalance(testConfig.ParentBalance).
				WithConfig(testConfig).
				WithMethods(bigIntToInt64MockContract),
		).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithArguments(bigIntArg).
			WithRecipientAddr(test.ParentAddress).
			WithGasProvided(testConfig.GasProvided).
			WithFunction("testFunction").
			Build()).
		AndAssertResults(func(world *worldmock.MockWorld, verify *test.VMOutputVerifier) {
			verify.Ok()
			assert.Equal(t, verify.VmOutput.ReturnData, [][]byte{bigIntArg})
		})
	assert.Nil(t, err)
}

func Test_BigIntToInt64_NotRepresentable(t *testing.T) {
	testConfig := baseTestConfig

	bigIntArg := []byte{0x01, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

	_, err := test.BuildMockInstanceCallTest(t).
		WithContracts(
			test.CreateMockContract(test.ParentAddress).
				WithBalance(testConfig.ParentBalance).
				WithConfig(testConfig).
				WithMethods(bigIntToInt64MockContract),
		).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithArguments(bigIntArg).
			WithRecipientAddr(test.ParentAddress).
			WithGasProvided(testConfig.GasProvided + 20000).
			WithFunction("testFunction").
			Build()).
		AndAssertResults(func(world *worldmock.MockWorld, verify *test.VMOutputVerifier) {
			verify.ExecutionFailed()
			assert.Equal(t, verify.VmOutput.ReturnMessage, vmhost.ErrBigIntCannotBeRepresentedAsInt64.Error())
		})
	assert.Nil(t, err)
}

func Test_ManagedRipemd160(t *testing.T) {
	testConfig := baseTestConfig

	asBytes := []byte{1, 2, 3}
	asRipemd160, _ := hashing.NewHasher().Ripemd160(asBytes)

	_, err := test.BuildMockInstanceCallTest(t).
		WithContracts(
			test.CreateMockContract(test.ParentAddress).
				WithBalance(testConfig.ParentBalance).
				WithConfig(testConfig).
				WithMethods(func(parentInstance *mock.InstanceMock, config interface{}) {
					parentInstance.AddMockMethod("testFunction", func() *mock.InstanceMock {
						host := parentInstance.Host

						managedTypes := host.ManagedTypes()
						sourceHandle := managedTypes.NewManagedBufferFromBytes(asBytes)
						destHandle := managedTypes.NewManagedBuffer()

						vmhooks.ManagedRipemd160WithHost(
							host,
							sourceHandle,
							destHandle)

						bytesResult, _ := managedTypes.GetBytes(destHandle)
						if !bytes.Equal(bytesResult, asRipemd160) {
							host.Runtime().SignalUserError("assert failed")
							return parentInstance
						}

						return parentInstance
					})
				}),
		).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithRecipientAddr(test.ParentAddress).
			WithGasProvided(testConfig.GasProvided).
			WithFunction("testFunction").
			Build()).
		AndAssertResults(func(world *worldmock.MockWorld, verify *test.VMOutputVerifier) {
			verify.
				Ok()
		})
	assert.Nil(t, err)
}

const blsCheckOK = "3e886a4c6e109a151f4105aee65a5192d150ef1fa68d3cd76964a0b086006dbe4324c989deb0e4416c6d6706db1b1910eb2732f08842fb4886067b9ed191109ac2188d76002d2e11da80a3f0ea89fee6b59c834cc478a6bd49cb8a193b1abb16@e96bd0f36b70c5ccc0c4396343bd7d8255b8a526c55fa1e218511fafe6539b8e@04725db195e37aa237cdbbda76270d4a229b6e7a3651104dc58c4349c0388e8546976fe54a04240530b99064e434c90f"

func blsSplitString(t testing.TB, str string) ([]byte, []byte, []byte) {
	split := strings.Split(str, "@")
	pkBuff, err := hex.DecodeString(split[0])
	require.Nil(t, err)

	msgBuff, err := hex.DecodeString(split[1])
	require.Nil(t, err)

	sigBuff, err := hex.DecodeString(split[2])
	require.Nil(t, err)

	return pkBuff, msgBuff, sigBuff
}

func Test_ManagedVerifyBLS(t *testing.T) {
	testConfig := baseTestConfig

	_, err := test.BuildMockInstanceCallTest(t).
		WithContracts(
			test.CreateMockContract(test.ParentAddress).
				WithBalance(testConfig.ParentBalance).
				WithConfig(testConfig).
				WithMethods(func(parentInstance *mock.InstanceMock, config interface{}) {
					parentInstance.AddMockMethod("testFunction", func() *mock.InstanceMock {
						host := parentInstance.Host

						key, message, sig := blsSplitString(t, blsCheckOK)
						managedTypes := host.ManagedTypes()
						keyHandle := managedTypes.NewManagedBufferFromBytes(key)
						messageHandle := managedTypes.NewManagedBufferFromBytes(message)
						sigHandle := managedTypes.NewManagedBufferFromBytes(sig)

						result := vmhooks.ManagedVerifyBLSWithHost(
							host,
							keyHandle,
							messageHandle,
							sigHandle)

						if result != 0 {
							host.Runtime().SignalUserError("assert failed")
							return parentInstance
						}

						return parentInstance
					})
				}),
		).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithRecipientAddr(test.ParentAddress).
			WithGasProvided(testConfig.GasProvided).
			WithFunction("testFunction").
			Build()).
		AndAssertResults(func(world *worldmock.MockWorld, verify *test.VMOutputVerifier) {
			verify.
				Ok()
		})
	assert.Nil(t, err)
}

func Test_ManagedVerifyEd25519(t *testing.T) {
	testConfig := baseTestConfig

	seed, _ := hex.DecodeString("1122334455667788990011223344556677889900112233445566778899001122")
	privateKey := ed25519.NewKeyFromSeed(seed)
	publicKey := privateKey.Public()
	message := []byte("test message!")
	sig := ed25519.Sign(privateKey, message)

	_, err := test.BuildMockInstanceCallTest(t).
		WithContracts(
			test.CreateMockContract(test.ParentAddress).
				WithBalance(testConfig.ParentBalance).
				WithConfig(testConfig).
				WithMethods(func(parentInstance *mock.InstanceMock, config interface{}) {
					parentInstance.AddMockMethod("testFunction", func() *mock.InstanceMock {
						host := parentInstance.Host

						managedTypes := host.ManagedTypes()
						keyHandle := managedTypes.NewManagedBufferFromBytes(publicKey.(ed25519.PublicKey))
						messageHandle := managedTypes.NewManagedBufferFromBytes(message)
						sigHandle := managedTypes.NewManagedBufferFromBytes(sig)

						result := vmhooks.ManagedVerifyEd25519WithHost(
							host,
							keyHandle,
							messageHandle,
							sigHandle)

						if result != 0 {
							host.Runtime().SignalUserError("assert failed")
							return parentInstance
						}

						return parentInstance
					})
				}),
		).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithRecipientAddr(test.ParentAddress).
			WithGasProvided(testConfig.GasProvided).
			WithFunction("testFunction").
			Build()).
		AndAssertResults(func(world *worldmock.MockWorld, verify *test.VMOutputVerifier) {
			verify.
				Ok()
		})
	assert.Nil(t, err)
}

func Test_VerifySecp256k1(t *testing.T) {
	testConfig := baseTestConfig

	key, _ := hex.DecodeString("04d2e670a19c6d753d1a6d8b20bd045df8a08fb162cf508956c31268c6d81ffdabab65528eefbb8057aa85d597258a3fbd481a24633bc9b47a9aa045c91371de52")
	msg, _ := hex.DecodeString("01020304")
	r, _ := hex.DecodeString("fef45d2892953aa5bbcdb057b5e98b208f1617a7498af7eb765574e29b5d9c2c")
	s, _ := hex.DecodeString("d47563f52aac6b04b55de236b7c515eb9311757db01e02cff079c3ca6efb063f")

	verifier := secp256k1.NewSecp256k1()
	sig := verifier.EncodeSecp256k1DERSignature(r, s)

	_, err := test.BuildMockInstanceCallTest(t).
		WithContracts(
			test.CreateMockContract(test.ParentAddress).
				WithBalance(testConfig.ParentBalance).
				WithConfig(testConfig).
				WithMethods(func(parentInstance *mock.InstanceMock, config interface{}) {
					parentInstance.AddMockMethod("testFunction", func() *mock.InstanceMock {
						host := parentInstance.Host

						managedTypes := host.ManagedTypes()
						keyHandle := managedTypes.NewManagedBufferFromBytes(key)
						messageHandle := managedTypes.NewManagedBufferFromBytes(msg)
						sigHandle := managedTypes.NewManagedBufferFromBytes(sig)

						result := vmhooks.ManagedVerifySecp256k1WithHost(
							host,
							keyHandle,
							messageHandle,
							sigHandle)

						if result != 0 {
							host.Runtime().SignalUserError("assert failed")
							return parentInstance
						}

						return parentInstance
					})
				}),
		).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithRecipientAddr(test.ParentAddress).
			WithGasProvided(testConfig.GasProvided).
			WithFunction("testFunction").
			Build()).
		AndAssertResults(func(world *worldmock.MockWorld, verify *test.VMOutputVerifier) {
			verify.
				Ok()
		})
	assert.Nil(t, err)
}

func Test_VerifyCustomSecp256k1(t *testing.T) {
	testConfig := baseTestConfig

	key, _ := hex.DecodeString("04d2e670a19c6d753d1a6d8b20bd045df8a08fb162cf508956c31268c6d81ffdabab65528eefbb8057aa85d597258a3fbd481a24633bc9b47a9aa045c91371de52")
	msg, _ := hex.DecodeString("01020304")
	r, _ := hex.DecodeString("fef45d2892953aa5bbcdb057b5e98b208f1617a7498af7eb765574e29b5d9c2c")
	s, _ := hex.DecodeString("d47563f52aac6b04b55de236b7c515eb9311757db01e02cff079c3ca6efb063f")

	verifier := secp256k1.NewSecp256k1()
	sig := verifier.EncodeSecp256k1DERSignature(r, s)

	_, err := test.BuildMockInstanceCallTest(t).
		WithContracts(
			test.CreateMockContract(test.ParentAddress).
				WithBalance(testConfig.ParentBalance).
				WithConfig(testConfig).
				WithMethods(func(parentInstance *mock.InstanceMock, config interface{}) {
					parentInstance.AddMockMethod("testFunction", func() *mock.InstanceMock {
						host := parentInstance.Host

						managedTypes := host.ManagedTypes()
						keyHandle := managedTypes.NewManagedBufferFromBytes(key)
						messageHandle := managedTypes.NewManagedBufferFromBytes(msg)
						sigHandle := managedTypes.NewManagedBufferFromBytes(sig)

						result := vmhooks.ManagedVerifyCustomSecp256k1WithHost(
							host,
							keyHandle,
							messageHandle,
							sigHandle,
							int32(secp256k1.ECDSADoubleSha256))

						if result != 0 {
							host.Runtime().SignalUserError("assert failed")
							return parentInstance
						}

						return parentInstance
					})
				}),
		).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithRecipientAddr(test.ParentAddress).
			WithGasProvided(testConfig.GasProvided).
			WithFunction("testFunction").
			Build()).
		AndAssertResults(func(world *worldmock.MockWorld, verify *test.VMOutputVerifier) {
			verify.
				Ok()
		})
	assert.Nil(t, err)
}

func Test_ManagedEncodeSecp256k1DerSignature(t *testing.T) {
	testConfig := baseTestConfig

	r, _ := hex.DecodeString("fef45d2892953aa5bbcdb057b5e98b208f1617a7498af7eb765574e29b5d9c2c")
	s, _ := hex.DecodeString("d47563f52aac6b04b55de236b7c515eb9311757db01e02cff079c3ca6efb063f")

	verifier := secp256k1.NewSecp256k1()
	sig := verifier.EncodeSecp256k1DERSignature(r, s)

	_, err := test.BuildMockInstanceCallTest(t).
		WithContracts(
			test.CreateMockContract(test.ParentAddress).
				WithBalance(testConfig.ParentBalance).
				WithConfig(testConfig).
				WithMethods(func(parentInstance *mock.InstanceMock, config interface{}) {
					parentInstance.AddMockMethod("testFunction", func() *mock.InstanceMock {
						host := parentInstance.Host

						managedTypes := host.ManagedTypes()
						rHandle := managedTypes.NewManagedBufferFromBytes(r)
						sHandle := managedTypes.NewManagedBufferFromBytes(s)
						sigHandle := managedTypes.NewManagedBuffer()

						retResult := vmhooks.ManagedEncodeSecp256k1DerSignatureWithHost(
							host,
							rHandle,
							sHandle,
							sigHandle)

						result, _ := managedTypes.GetBytes(sigHandle)
						if retResult != 0 || !bytes.Equal(result, sig) {
							host.Runtime().SignalUserError("assert failed")
							return parentInstance
						}

						return parentInstance
					})
				}),
		).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithRecipientAddr(test.ParentAddress).
			WithGasProvided(testConfig.GasProvided).
			WithFunction("testFunction").
			Build()).
		AndAssertResults(func(world *worldmock.MockWorld, verify *test.VMOutputVerifier) {
			verify.
				Ok()
		})
	assert.Nil(t, err)
}

func Test_ManagedScalarBaseMultEC(t *testing.T) {
	testConfig := baseTestConfig

	dataBytes, _ := hex.DecodeString("11839296a789a3bc0045c8a5fb42c7d1bd998f54449579b446817afbd17273e662c97ee72995ef42640c550b9013fad0761353c7086a272c24088be94769fd16650")

	_, err := test.BuildMockInstanceCallTest(t).
		WithContracts(
			test.CreateMockContract(test.ParentAddress).
				WithBalance(testConfig.ParentBalance).
				WithConfig(testConfig).
				WithMethods(func(parentInstance *mock.InstanceMock, config interface{}) {
					parentInstance.AddMockMethod("testFunction", func() *mock.InstanceMock {
						host := parentInstance.Host

						managedTypes := host.ManagedTypes()
						xResultHandle := managedTypes.NewBigIntFromInt64(0)
						yResultHandle := managedTypes.NewBigIntFromInt64(0)

						p521ec := elliptic.P521().Params()
						ecHandle := managedTypes.PutEllipticCurve(p521ec)
						dataHandle := managedTypes.NewManagedBufferFromBytes(dataBytes)

						retResult := vmhooks.ManagedScalarBaseMultECWithHost(
							host,
							xResultHandle,
							yResultHandle,
							ecHandle,
							dataHandle)

						if retResult != 0 {
							host.Runtime().SignalUserError("assert failed")
							return parentInstance
						}

						return parentInstance
					})
				}),
		).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithRecipientAddr(test.ParentAddress).
			WithGasProvided(testConfig.GasProvided).
			WithFunction("testFunction").
			Build()).
		AndAssertResults(func(world *worldmock.MockWorld, verify *test.VMOutputVerifier) {
			verify.
				Ok()
		})
	assert.Nil(t, err)
}

func Test_ManagedScalarMultEC(t *testing.T) {
	testConfig := baseTestConfig

	pointXBytes, _ := hex.DecodeString("010ba38127b62997b313aa2990a13fce55c46fc3ae751a7a7b91c41341719b57f13b9185edd96a0211acf922adb13aa9d7c64925664a9419ae6f5bc9cc4d25f91f50")
	pointYBytes, _ := hex.DecodeString("016967055bf964609b6fd853e0aa9b90d6e1e942066278a18e8604f9fcef5b64370412f20836767829ee7e0d3fc8e2e204e2a8ec4f9257a552d66647b2d1b9856223")
	dataBytes, _ := hex.DecodeString("f93e4ae433cc12cf2a43fc0ef26400c0e125508224cdb649380f25479148a4ad")

	_, err := test.BuildMockInstanceCallTest(t).
		WithContracts(
			test.CreateMockContract(test.ParentAddress).
				WithBalance(testConfig.ParentBalance).
				WithConfig(testConfig).
				WithMethods(func(parentInstance *mock.InstanceMock, config interface{}) {
					parentInstance.AddMockMethod("testFunction", func() *mock.InstanceMock {
						host := parentInstance.Host

						managedTypes := host.ManagedTypes()
						xResultHandle := managedTypes.NewBigIntFromInt64(0)
						yResultHandle := managedTypes.NewBigIntFromInt64(0)

						p521ec := elliptic.P521().Params()
						ecHandle := managedTypes.PutEllipticCurve(p521ec)
						pointXHandle := managedTypes.NewBigInt(big.NewInt(0).SetBytes(pointXBytes))
						pointYHandle := managedTypes.NewBigInt(big.NewInt(0).SetBytes(pointYBytes))
						dataHandle := managedTypes.NewManagedBufferFromBytes(dataBytes)

						retResult := vmhooks.ManagedScalarMultECWithHost(
							host,
							xResultHandle,
							yResultHandle,
							ecHandle,
							pointXHandle,
							pointYHandle,
							dataHandle)

						if retResult != 0 {
							host.Runtime().SignalUserError("assert failed")
							return parentInstance
						}

						return parentInstance
					})
				}),
		).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithRecipientAddr(test.ParentAddress).
			WithGasProvided(testConfig.GasProvided).
			WithFunction("testFunction").
			Build()).
		AndAssertResults(func(world *worldmock.MockWorld, verify *test.VMOutputVerifier) {
			verify.
				Ok()
		})
	assert.Nil(t, err)
}

func Test_ManagedMarshalEC(t *testing.T) {
	testConfig := baseTestConfig

	pointXBytes, _ := hex.DecodeString("010ba38127b62997b313aa2990a13fce55c46fc3ae751a7a7b91c41341719b57f13b9185edd96a0211acf922adb13aa9d7c64925664a9419ae6f5bc9cc4d25f91f50")
	pointYBytes, _ := hex.DecodeString("016967055bf964609b6fd853e0aa9b90d6e1e942066278a18e8604f9fcef5b64370412f20836767829ee7e0d3fc8e2e204e2a8ec4f9257a552d66647b2d1b9856223")
	marshalled, _ := hex.DecodeString("04010ba38127b62997b313aa2990a13fce55c46fc3ae751a7a7b91c41341719b57f13b9185edd96a0211acf922adb13aa9d7c64925664a9419ae6f5bc9cc4d25f91f50016967055bf964609b6fd853e0aa9b90d6e1e942066278a18e8604f9fcef5b64370412f20836767829ee7e0d3fc8e2e204e2a8ec4f9257a552d66647b2d1b9856223")

	_, err := test.BuildMockInstanceCallTest(t).
		WithContracts(
			test.CreateMockContract(test.ParentAddress).
				WithBalance(testConfig.ParentBalance).
				WithConfig(testConfig).
				WithMethods(func(parentInstance *mock.InstanceMock, config interface{}) {
					parentInstance.AddMockMethod("testFunction", func() *mock.InstanceMock {
						host := parentInstance.Host

						managedTypes := host.ManagedTypes()
						resultHandle := managedTypes.NewBigIntFromInt64(0)

						p521ec := elliptic.P521().Params()
						ecHandle := managedTypes.PutEllipticCurve(p521ec)
						xPairHandle := managedTypes.NewBigInt(big.NewInt(0).SetBytes(pointXBytes))
						yPairHandle := managedTypes.NewBigInt(big.NewInt(0).SetBytes(pointYBytes))

						retResult := vmhooks.ManagedMarshalECWithHost(
							host,
							xPairHandle,
							yPairHandle,
							ecHandle,
							resultHandle)

						resultBytes, _ := managedTypes.GetBytes(resultHandle)

						if retResult == -1 || !bytes.Equal(resultBytes, marshalled) {
							host.Runtime().SignalUserError("assert failed")
							return parentInstance
						}

						return parentInstance
					})
				}),
		).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithRecipientAddr(test.ParentAddress).
			WithGasProvided(testConfig.GasProvided).
			WithFunction("testFunction").
			Build()).
		AndAssertResults(func(world *worldmock.MockWorld, verify *test.VMOutputVerifier) {
			verify.
				Ok()
		})
	assert.Nil(t, err)
}

func Test_ManagedUnmarshalEC(t *testing.T) {
	testConfig := baseTestConfig

	pointXBytes, _ := hex.DecodeString("010ba38127b62997b313aa2990a13fce55c46fc3ae751a7a7b91c41341719b57f13b9185edd96a0211acf922adb13aa9d7c64925664a9419ae6f5bc9cc4d25f91f50")
	pointYBytes, _ := hex.DecodeString("016967055bf964609b6fd853e0aa9b90d6e1e942066278a18e8604f9fcef5b64370412f20836767829ee7e0d3fc8e2e204e2a8ec4f9257a552d66647b2d1b9856223")
	dataBytes, _ := hex.DecodeString("04010ba38127b62997b313aa2990a13fce55c46fc3ae751a7a7b91c41341719b57f13b9185edd96a0211acf922adb13aa9d7c64925664a9419ae6f5bc9cc4d25f91f50016967055bf964609b6fd853e0aa9b90d6e1e942066278a18e8604f9fcef5b64370412f20836767829ee7e0d3fc8e2e204e2a8ec4f9257a552d66647b2d1b9856223")

	_, err := test.BuildMockInstanceCallTest(t).
		WithContracts(
			test.CreateMockContract(test.ParentAddress).
				WithBalance(testConfig.ParentBalance).
				WithConfig(testConfig).
				WithMethods(func(parentInstance *mock.InstanceMock, config interface{}) {
					parentInstance.AddMockMethod("testFunction", func() *mock.InstanceMock {
						host := parentInstance.Host

						managedTypes := host.ManagedTypes()

						p521ec := elliptic.P521().Params()
						ecHandle := managedTypes.PutEllipticCurve(p521ec)
						xResultHandle := managedTypes.NewBigIntFromInt64(0)
						yResultHandle := managedTypes.NewBigIntFromInt64(0)
						dataHandle := managedTypes.NewManagedBufferFromBytes(dataBytes)

						retResult := vmhooks.ManagedUnmarshalECWithHost(
							host,
							xResultHandle,
							yResultHandle,
							ecHandle,
							dataHandle)

						xResult, _ := managedTypes.GetBigInt(xResultHandle)
						yResult, _ := managedTypes.GetBigInt(yResultHandle)

						if retResult == -1 ||
							!bytes.Equal(xResult.Bytes(), pointXBytes) ||
							!bytes.Equal(yResult.Bytes(), pointYBytes) {
							host.Runtime().SignalUserError("assert failed")
							return parentInstance
						}

						return parentInstance
					})
				}),
		).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithRecipientAddr(test.ParentAddress).
			WithGasProvided(testConfig.GasProvided).
			WithFunction("testFunction").
			Build()).
		AndAssertResults(func(world *worldmock.MockWorld, verify *test.VMOutputVerifier) {
			verify.
				Ok()
		})
	assert.Nil(t, err)
}

func Test_ManagedMarshalCompressedEC(t *testing.T) {
	testConfig := baseTestConfig

	pointXBytes, _ := hex.DecodeString("010ba38127b62997b313aa2990a13fce55c46fc3ae751a7a7b91c41341719b57f13b9185edd96a0211acf922adb13aa9d7c64925664a9419ae6f5bc9cc4d25f91f50")
	pointYBytes, _ := hex.DecodeString("016967055bf964609b6fd853e0aa9b90d6e1e942066278a18e8604f9fcef5b64370412f20836767829ee7e0d3fc8e2e204e2a8ec4f9257a552d66647b2d1b9856223")
	marshalled, _ := hex.DecodeString("03010ba38127b62997b313aa2990a13fce55c46fc3ae751a7a7b91c41341719b57f13b9185edd96a0211acf922adb13aa9d7c64925664a9419ae6f5bc9cc4d25f91f50")

	_, err := test.BuildMockInstanceCallTest(t).
		WithContracts(
			test.CreateMockContract(test.ParentAddress).
				WithBalance(testConfig.ParentBalance).
				WithConfig(testConfig).
				WithMethods(func(parentInstance *mock.InstanceMock, config interface{}) {
					parentInstance.AddMockMethod("testFunction", func() *mock.InstanceMock {
						host := parentInstance.Host

						managedTypes := host.ManagedTypes()
						resultHandle := managedTypes.NewBigIntFromInt64(0)

						p521ec := elliptic.P521().Params()
						ecHandle := managedTypes.PutEllipticCurve(p521ec)
						xPairHandle := managedTypes.NewBigInt(big.NewInt(0).SetBytes(pointXBytes))
						yPairHandle := managedTypes.NewBigInt(big.NewInt(0).SetBytes(pointYBytes))

						retResult := vmhooks.ManagedMarshalCompressedECWithHost(
							host,
							xPairHandle,
							yPairHandle,
							ecHandle,
							resultHandle)

						resultBytes, _ := managedTypes.GetBytes(resultHandle)

						if retResult == -1 || !bytes.Equal(resultBytes, marshalled) {
							host.Runtime().SignalUserError("assert failed")
							return parentInstance
						}

						return parentInstance
					})
				}),
		).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithRecipientAddr(test.ParentAddress).
			WithGasProvided(testConfig.GasProvided).
			WithFunction("testFunction").
			Build()).
		AndAssertResults(func(world *worldmock.MockWorld, verify *test.VMOutputVerifier) {
			verify.
				Ok()
		})
	assert.Nil(t, err)
}

func Test_ManagedUnmarshalCompressedEC(t *testing.T) {
	testConfig := baseTestConfig

	pointXBytes, _ := hex.DecodeString("010ba38127b62997b313aa2990a13fce55c46fc3ae751a7a7b91c41341719b57f13b9185edd96a0211acf922adb13aa9d7c64925664a9419ae6f5bc9cc4d25f91f50")
	pointYBytes, _ := hex.DecodeString("016967055bf964609b6fd853e0aa9b90d6e1e942066278a18e8604f9fcef5b64370412f20836767829ee7e0d3fc8e2e204e2a8ec4f9257a552d66647b2d1b9856223")
	dataBytes, _ := hex.DecodeString("03010ba38127b62997b313aa2990a13fce55c46fc3ae751a7a7b91c41341719b57f13b9185edd96a0211acf922adb13aa9d7c64925664a9419ae6f5bc9cc4d25f91f50")

	_, err := test.BuildMockInstanceCallTest(t).
		WithContracts(
			test.CreateMockContract(test.ParentAddress).
				WithBalance(testConfig.ParentBalance).
				WithConfig(testConfig).
				WithMethods(func(parentInstance *mock.InstanceMock, config interface{}) {
					parentInstance.AddMockMethod("testFunction", func() *mock.InstanceMock {
						host := parentInstance.Host

						managedTypes := host.ManagedTypes()

						p521ec := elliptic.P521().Params()
						ecHandle := managedTypes.PutEllipticCurve(p521ec)
						xResultHandle := managedTypes.NewBigIntFromInt64(0)
						yResultHandle := managedTypes.NewBigIntFromInt64(0)
						dataHandle := managedTypes.NewManagedBufferFromBytes(dataBytes)

						retResult := vmhooks.ManagedUnmarshalCompressedECWithHost(
							host,
							xResultHandle,
							yResultHandle,
							ecHandle,
							dataHandle)

						xResult, _ := managedTypes.GetBigInt(xResultHandle)
						yResult, _ := managedTypes.GetBigInt(yResultHandle)

						if retResult == -1 ||
							!bytes.Equal(xResult.Bytes(), pointXBytes) ||
							!bytes.Equal(yResult.Bytes(), pointYBytes) {
							host.Runtime().SignalUserError("assert failed")
							return parentInstance
						}

						return parentInstance
					})
				}),
		).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithRecipientAddr(test.ParentAddress).
			WithGasProvided(testConfig.GasProvided).
			WithFunction("testFunction").
			Build()).
		AndAssertResults(func(world *worldmock.MockWorld, verify *test.VMOutputVerifier) {
			verify.
				Ok()
		})
	assert.Nil(t, err)
}

func Test_ManagedGenerateKeyEC(t *testing.T) {
	testConfig := baseTestConfig

	pointXBytes, _ := hex.DecodeString("010ba38127b62997b313aa2990a13fce55c46fc3ae751a7a7b91c41341719b57f13b9185edd96a0211acf922adb13aa9d7c64925664a9419ae6f5bc9cc4d25f91f50")
	pointYBytes, _ := hex.DecodeString("016967055bf964609b6fd853e0aa9b90d6e1e942066278a18e8604f9fcef5b64370412f20836767829ee7e0d3fc8e2e204e2a8ec4f9257a552d66647b2d1b9856223")
	expectedResultBytes, _ := hex.DecodeString("00ddb81d205713945e203848e2f5c312067649f9a40727ca26b672b164cd1f9108f564958b20312146bb9750b74757d97cfbbba2aedebaba3a68fe3f2d669a992fab")

	_, err := test.BuildMockInstanceCallTest(t).
		WithContracts(
			test.CreateMockContract(test.ParentAddress).
				WithBalance(testConfig.ParentBalance).
				WithConfig(testConfig).
				WithMethods(func(parentInstance *mock.InstanceMock, config interface{}) {
					parentInstance.AddMockMethod("testFunction", func() *mock.InstanceMock {
						host := parentInstance.Host

						managedTypes := host.ManagedTypes()

						p521ec := elliptic.P521().Params()
						ecHandle := managedTypes.PutEllipticCurve(p521ec)
						xResultHandle := managedTypes.NewBigInt(big.NewInt(0).SetBytes(pointXBytes))
						yResultHandle := managedTypes.NewBigInt(big.NewInt(0).SetBytes(pointYBytes))
						resultHandle := managedTypes.NewManagedBuffer()

						retResult := vmhooks.ManagedGenerateKeyECWithHost(
							host,
							xResultHandle,
							yResultHandle,
							ecHandle,
							resultHandle)

						resultBytes, _ := managedTypes.GetBytes(resultHandle)
						if retResult != 0 ||
							!bytes.Equal(resultBytes, expectedResultBytes) {
							host.Runtime().SignalUserError("assert failed")
							return parentInstance
						}

						return parentInstance
					})
				}),
		).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithRecipientAddr(test.ParentAddress).
			WithGasProvided(testConfig.GasProvided).
			WithFunction("testFunction").
			Build()).
		AndAssertResults(func(world *worldmock.MockWorld, verify *test.VMOutputVerifier) {
			verify.
				Ok()
		})
	assert.Nil(t, err)
}

func Test_ManagedCreateEC(t *testing.T) {
	testConfig := baseTestConfig

	_, err := test.BuildMockInstanceCallTest(t).
		WithContracts(
			test.CreateMockContract(test.ParentAddress).
				WithBalance(testConfig.ParentBalance).
				WithConfig(testConfig).
				WithMethods(func(parentInstance *mock.InstanceMock, config interface{}) {
					parentInstance.AddMockMethod("testFunction", func() *mock.InstanceMock {
						host := parentInstance.Host

						p224ec := elliptic.P224().Params()
						if !checkCreateECSuccess(host, "p224", p224ec) {
							return parentInstance
						}

						p256ec := elliptic.P256().Params()
						if !checkCreateECSuccess(host, "p256", p256ec) {
							return parentInstance
						}

						p384ec := elliptic.P384().Params()
						if !checkCreateECSuccess(host, "p384", p384ec) {
							return parentInstance
						}

						p521ec := elliptic.P521().Params()
						if !checkCreateECSuccess(host, "p521", p521ec) {
							return parentInstance
						}

						return parentInstance
					})
				}),
		).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithRecipientAddr(test.ParentAddress).
			WithGasProvided(testConfig.GasProvided).
			WithFunction("testFunction").
			Build()).
		AndAssertResults(func(world *worldmock.MockWorld, verify *test.VMOutputVerifier) {
			verify.
				Ok()
		})
	assert.Nil(t, err)
}

func checkCreateECSuccess(host vmhost.VMHost, name string, ecParams *elliptic.CurveParams) bool {
	managedTypes := host.ManagedTypes()
	dataHandle := managedTypes.NewManagedBufferFromBytes([]byte(name))

	retResult := vmhooks.ManagedCreateECWithHost(
		host,
		dataHandle)

	resultEC, _ := managedTypes.GetEllipticCurve(retResult)
	if resultEC.Params().Name != ecParams.Name {
		host.Runtime().SignalUserError("assert failed")
		return false
	}

	return true
}

func Test_ManagedDeleteContract(t *testing.T) {
	testConfig := baseTestConfig

	_, err := test.BuildMockInstanceCallTest(t).
		WithContracts(
			test.CreateMockContract(test.ParentAddress).
				WithBalance(testConfig.ParentBalance).
				WithConfig(testConfig).
				WithCodeMetadata([]byte{vmcommon.MetadataUpgradeable, 0}).
				WithOwnerAddress(test.ParentAddress).
				WithMethods(func(parentInstance *mock.InstanceMock, config interface{}) {
					parentInstance.AddMockMethod("testFunction", func() *mock.InstanceMock {
						host := parentInstance.Host
						managedTypes := host.ManagedTypes()

						argumentsHandle := managedTypes.NewManagedBuffer()
						_ = managedTypes.WriteManagedVecOfManagedBuffers([][]byte{{1, 2}, {3, 4}}, argumentsHandle)

						destHandle := managedTypes.NewManagedBufferFromBytes(test.ParentAddress)

						vmhooks.ManagedDeleteContractWithHost(
							host,
							destHandle,
							100000,
							argumentsHandle)

						return parentInstance
					})
				}),
		).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithRecipientAddr(test.ParentAddress).
			WithGasProvided(testConfig.GasProvided).
			WithFunction("testFunction").
			Build()).
		WithSetup(func(host vmhost.VMHost, world *worldmock.MockWorld) {
			setZeroCodeCosts(host)
			setAsyncCosts(host, testConfig.GasLockCost)
		}).
		AndAssertResults(func(world *worldmock.MockWorld, verify *test.VMOutputVerifier) {
			verify.
				Ok().
				DeletedAccounts(test.ParentAddress)
		})
	assert.Nil(t, err)
}

func Test_ManagedDeleteContract_CrossShard(t *testing.T) {
	testConfig := makeTestConfig()

	_, err := test.BuildMockInstanceCallTest(t).
		WithContracts(
			test.CreateMockContractOnShard(test.ChildAddress, 1).
				WithBalance(testConfig.ChildBalance).
				WithConfig(testConfig).
				WithCodeMetadata([]byte{vmcommon.MetadataUpgradeable, 0}).
				WithOwnerAddress(test.ParentAddress).
				WithMethods(contracts.WasteGasChildMock),
		).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithCallerAddr(test.ParentAddress).
			WithRecipientAddr(test.ChildAddress).
			WithCallValue(testConfig.TransferFromParentToChild).
			WithGasProvided(testConfig.GasProvided).
			WithFunction(vmhost.DeleteFunctionName).
			WithArguments(
				[]byte{0}, // placeholder for data used by async framework
				[]byte{0}, // placeholder for data used by async framework
				big.NewInt(testConfig.TransferToThirdParty).Bytes(),
				[]byte(contracts.AsyncChildData),
				[]byte{0}).
			WithCallType(vm.AsynchronousCall).
			Build()).
		WithSetup(func(host vmhost.VMHost, world *worldmock.MockWorld) {
			world.SelfShardID = 1
			if world.CurrentBlockInfo == nil {
				world.CurrentBlockInfo = &worldmock.BlockInfo{}
			}
			world.CurrentBlockInfo.BlockRound = 1
			setZeroCodeCosts(host)
			setAsyncCosts(host, testConfig.GasLockCost)
		}).
		AndAssertResults(func(world *worldmock.MockWorld, verify *test.VMOutputVerifier) {
			verify.Ok().
				DeletedAccounts(test.ChildAddress)
		})
	assert.Nil(t, err)
}

func TestBaseOpsAPI_NFTNonceOverflow(t *testing.T) {
	testConfig := makeTestConfig()

	MaxUint := ^uint64(0)
	MaxInt := int64(MaxUint >> 1)

	OverflowedMaxInt := uint64(MaxInt) + 1

	tokenValue := int64(100)
	_, err := test.BuildMockInstanceCallTest(t).
		WithContracts(
			test.CreateMockContract(test.ParentAddress).
				WithBalance(testConfig.ParentBalance).
				WithConfig(testConfig).
				WithMethods(func(parentInstance *mock.InstanceMock, config interface{}) {
					parentInstance.AddMockMethod("testFunction", func() *mock.InstanceMock {
						host := parentInstance.Host
						managed := host.ManagedTypes()

						addressHandle := managed.NewManagedBufferFromBytes(test.ParentAddress)
						tokenIDHandle := managed.NewManagedBufferFromBytes(test.ESDTTestTokenName)

						nonce := int64(OverflowedMaxInt)

						valueHandle := managed.NewBigIntFromInt64(0)
						propertiesHandle := managed.NewManagedBuffer()
						hashHandle := managed.NewManagedBuffer()
						nameHandle := managed.NewManagedBuffer()
						attributesHandle := managed.NewManagedBuffer()
						creatorHandle := managed.NewManagedBuffer()
						royaltiesHandle := managed.NewManagedBuffer()
						urisHandle := managed.NewManagedBuffer()

						vmhooks.ManagedGetESDTTokenDataWithHost(host,
							addressHandle,
							tokenIDHandle,
							nonce,
							valueHandle, propertiesHandle, hashHandle, nameHandle, attributesHandle, creatorHandle, royaltiesHandle, urisHandle)

						value, err := managed.GetBigInt(valueHandle)
						if err != nil {
							host.Runtime().SignalUserError(err.Error())
							return parentInstance
						}
						host.Output().Finish(value.Bytes())

						return parentInstance
					})
				}),
		).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithRecipientAddr(test.ParentAddress).
			WithGasProvided(testConfig.GasProvided).
			WithFunction("testFunction").
			Build()).
		WithSetup(func(host vmhost.VMHost, world *worldmock.MockWorld) {
			createMockBuiltinFunctions(t, host, world)
			setZeroCodeCosts(host)
			err := world.BuiltinFuncs.SetTokenData(
				test.ParentAddress,
				test.ESDTTestTokenName,
				OverflowedMaxInt,
				&esdt.ESDigitalToken{
					Value:      big.NewInt(tokenValue),
					Type:       uint32(core.Fungible),
					Properties: esdtconvert.MakeESDTUserMetadataBytes(false),
				})
			assert.Nil(t, err)
		}).
		AndAssertResults(func(world *worldmock.MockWorld, verify *test.VMOutputVerifier) {
			verify.
				Ok().
				ReturnData(big.NewInt(tokenValue).Bytes())
		})
	assert.Nil(t, err)
}

func Test_ManagedGetCodeMetadata(t *testing.T) {
	testConfig := baseTestConfig

	metadata := []byte{0, vmcommon.MetadataPayable}

	_, err := test.BuildMockInstanceCallTest(t).
		WithContracts(
			test.CreateMockContract(test.ParentAddress).
				WithBalance(testConfig.ParentBalance).
				WithConfig(testConfig).
				WithCodeMetadata(metadata).
				WithMethods(func(parentInstance *mock.InstanceMock, config interface{}) {
					parentInstance.AddMockMethod("testFunction", func() *mock.InstanceMock {
						host := parentInstance.Host

						managedTypes := host.ManagedTypes()
						addressHandle := managedTypes.NewManagedBufferFromBytes(test.ParentAddress)
						destHandle := managedTypes.NewManagedBuffer()

						vmhooks.ManagedGetCodeMetadataWithHost(
							host,
							addressHandle,
							destHandle)

						bytesResult, _ := managedTypes.GetBytes(destHandle)
						if !bytes.Equal(metadata, bytesResult) {
							host.Runtime().SignalUserError("assert failed")
							return parentInstance
						}

						return parentInstance
					})
				}),
		).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithRecipientAddr(test.ParentAddress).
			WithGasProvided(testConfig.GasProvided).
			WithFunction("testFunction").
			Build()).
		AndAssertResults(func(world *worldmock.MockWorld, verify *test.VMOutputVerifier) {
			verify.
				Ok()
		})
	assert.Nil(t, err)
}

func Test_ManagedIsBuiltinFunction(t *testing.T) {
	testConfig := baseTestConfig

	_, err := test.BuildMockInstanceCallTest(t).
		WithContracts(
			test.CreateMockContract(test.ParentAddress).
				WithBalance(testConfig.ParentBalance).
				WithConfig(testConfig).
				WithMethods(func(parentInstance *mock.InstanceMock, config interface{}) {
					parentInstance.AddMockMethod("testFunction", func() *mock.InstanceMock {
						host := parentInstance.Host

						managedTypes := host.ManagedTypes()
						functionNameHandle := managedTypes.NewManagedBufferFromBytes([]byte("ESDTTransfer"))

						returnValue := vmhooks.ManagedIsBuiltinFunctionWithHost(
							host,
							functionNameHandle)
						if returnValue != 1 {
							host.Runtime().SignalUserError("assert failed")
							return parentInstance
						}

						functionNameHandle = managedTypes.NewManagedBufferFromBytes([]byte("NotABuiltInFunction"))

						returnValue = vmhooks.ManagedIsBuiltinFunctionWithHost(
							host,
							functionNameHandle)
						if returnValue != 0 {
							host.Runtime().SignalUserError("assert failed")
							return parentInstance
						}

						return parentInstance
					})
				}),
		).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithRecipientAddr(test.ParentAddress).
			WithGasProvided(testConfig.GasProvided).
			WithFunction("testFunction").
			Build()).
		WithSetup(func(host vmhost.VMHost, world *worldmock.MockWorld) {
			createMockBuiltinFunctions(t, host, world)
		}).
		AndAssertResults(func(world *worldmock.MockWorld, verify *test.VMOutputVerifier) {
			verify.
				Ok()
		})
	assert.Nil(t, err)
}

func Test_Direct_ManagedGetBackTransfers(t *testing.T) {
	testConfig := makeTestConfig()
	egldTransfer := big.NewInt(2)
	initialESDTTokenBalance := uint64(100)
	testConfig.ESDTTokensToTransfer = 5

	_, err := test.BuildMockInstanceCallTest(t).
		WithContracts(
			test.CreateMockContract(test.ParentAddress).
				WithBalance(testConfig.ParentBalance).
				WithConfig(testConfig).
				WithMethods(func(parentInstance *mock.InstanceMock, config interface{}) {
					parentInstance.AddMockMethod("callChild", func() *mock.InstanceMock {
						host := parentInstance.Host
						input := test.DefaultTestContractCallInput()
						input.GasProvided = testConfig.GasProvidedToChild
						input.CallerAddr = test.ParentAddress
						input.RecipientAddr = test.ChildAddress
						input.Function = "childFunction"
						returnValue := contracts.ExecuteOnDestContextInMockContracts(host, input)
						if returnValue != 0 {
							host.Runtime().FailExecution(fmt.Errorf("return value %d", returnValue))
						}
						managedTypes := host.ManagedTypes()
						esdtTransfers, egld := managedTypes.GetBackTransfers()
						assert.Equal(t, 1, len(esdtTransfers))
						assert.Equal(t, test.ESDTTestTokenName, esdtTransfers[0].ESDTTokenName)
						assert.Equal(t, big.NewInt(0).SetUint64(testConfig.ESDTTokensToTransfer), esdtTransfers[0].ESDTValue)
						assert.Equal(t, egld, egldTransfer)
						return parentInstance
					})
				}),
			test.CreateMockContract(test.ChildAddress).
				WithBalance(testConfig.ChildBalance).
				WithConfig(testConfig).
				WithMethods(func(parentInstance *mock.InstanceMock, config interface{}) {
					parentInstance.AddMockMethod("childFunction", func() *mock.InstanceMock {
						host := parentInstance.Host

						valueBytes := egldTransfer.Bytes()
						err := host.Output().Transfer(
							test.ParentAddress,
							test.ChildAddress, 0, 0, big.NewInt(0).SetBytes(valueBytes), nil, []byte{}, vm.DirectCall)
						if err != nil {
							host.Runtime().FailExecution(err)
						}

						transfer := &vmcommon.ESDTTransfer{
							ESDTValue:      big.NewInt(int64(testConfig.ESDTTokensToTransfer)),
							ESDTTokenName:  test.ESDTTestTokenName,
							ESDTTokenType:  0,
							ESDTTokenNonce: 0,
						}

						ret := vmhooks.TransferESDTNFTExecuteWithTypedArgs(
							host,
							test.ParentAddress,
							[]*vmcommon.ESDTTransfer{transfer},
							int64(testConfig.GasProvidedToChild),
							nil,
							nil)
						if ret != 0 {
							host.Runtime().FailExecution(fmt.Errorf("Transfer ESDT failed"))
						}

						return parentInstance
					})
				}),
		).
		WithSetup(func(host vmhost.VMHost, world *worldmock.MockWorld) {
			childAccount := world.AcctMap.GetAccount(test.ChildAddress)
			_ = childAccount.SetTokenBalanceUint64(test.ESDTTestTokenName, 0, initialESDTTokenBalance)
			createMockBuiltinFunctions(t, host, world)
			setZeroCodeCosts(host)
		}).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithRecipientAddr(test.ParentAddress).
			WithGasProvided(testConfig.GasProvided).
			WithFunction("callChild").
			Build()).
		AndAssertResults(func(world *worldmock.MockWorld, verify *test.VMOutputVerifier) {
			verify.
				Ok()
		})
	assert.Nil(t, err)
}

func Test_Async_ManagedGetBackTransfers(t *testing.T) {
	testConfig := makeTestConfig()
	initialESDTTokenBalance := uint64(100)
	testConfig.GasProvided = 10_000
	testConfig.GasProvidedToChild = 1000
	testConfig.ESDTTokensToTransfer = 5
	testConfig.SuccessCallback = "myCallback"
	testConfig.ErrorCallback = "myCallback"
	testConfig.TransferFromChildToParent = 2
	testConfig.ParentAddress = test.ParentAddress
	testConfig.ChildAddress = test.ChildAddress
	testConfig.NephewAddress = test.NephewAddress

	_, err := test.BuildMockInstanceCallTest(t).
		WithContracts(
			test.CreateMockContract(test.ParentAddress).
				WithBalance(testConfig.ParentBalance).
				WithConfig(testConfig).
				WithCodeMetadata([]byte{0, 0}).
				WithMethods(contracts.BackTransfer_ParentCallsChild),
			test.CreateMockContract(test.ChildAddress).
				WithBalance(testConfig.ChildBalance).
				WithConfig(testConfig).
				WithMethods(
					contracts.BackTransfer_ChildMakesAsync,
					contracts.BackTransfer_ChildCallback,
				),
			test.CreateMockContract(test.NephewAddress).
				WithBalance(testConfig.ChildBalance).
				WithConfig(testConfig).
				WithMethods(contracts.WasteGasChildMock),
		).
		WithSetup(func(host vmhost.VMHost, world *worldmock.MockWorld) {
			childAccount := world.AcctMap.GetAccount(test.ChildAddress)
			_ = childAccount.SetTokenBalanceUint64(test.ESDTTestTokenName, 0, initialESDTTokenBalance)
			createMockBuiltinFunctions(t, host, world)
			setZeroCodeCosts(host)
			host.Metering().GasSchedule().BaseOpsAPICost.AsyncCallbackGasLock = 0
		}).
		WithInput(test.CreateTestContractCallInputBuilder().
			WithRecipientAddr(test.ParentAddress).
			WithGasProvided(testConfig.GasProvided).
			WithFunction("callChild").
			WithArguments([]byte{1}).
			Build()).
		AndAssertResults(func(world *worldmock.MockWorld, verify *test.VMOutputVerifier) {
			verify.
				Ok()
		})
	assert.Nil(t, err)
}
