package executor

// Code generated by elrondapi generator. DO NOT EDIT.

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// !!!!!!!!!!!!!!!!!!!!!! AUTO-GENERATED FILE !!!!!!!!!!!!!!!!!!!!!!
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

// VMHooks contains all VM functions that can be called by the executor during SC execution.
type VMHooks interface {
	MainVMHooks
	ManagedVMHooks
	BigFloatVMHooks
	BigIntVMHooks
	ManagedBufferVMHooks
	SmallIntVMHooks
	CryptoVMHooks
}

type MainVMHooks interface {
	GetGasLeft() int64
	GetSCAddress(resultOffset MemPtr)
	GetOwnerAddress(resultOffset MemPtr)
	GetShardOfAddress(addressOffset MemPtr) int32
	IsSmartContract(addressOffset MemPtr) int32
	SignalError(messageOffset MemPtr, messageLength MemLength)
	GetExternalBalance(addressOffset MemPtr, resultOffset MemPtr)
	GetBlockHash(nonce int64, resultOffset MemPtr) int32
	GetESDTBalance(addressOffset MemPtr, tokenIDOffset MemPtr, tokenIDLen MemLength, nonce int64, resultOffset MemPtr) int32
	GetESDTNFTNameLength(addressOffset MemPtr, tokenIDOffset MemPtr, tokenIDLen MemLength, nonce int64) int32
	GetESDTNFTAttributeLength(addressOffset MemPtr, tokenIDOffset MemPtr, tokenIDLen MemLength, nonce int64) int32
	GetESDTNFTURILength(addressOffset MemPtr, tokenIDOffset MemPtr, tokenIDLen MemLength, nonce int64) int32
	GetESDTTokenData(addressOffset MemPtr, tokenIDOffset MemPtr, tokenIDLen MemLength, nonce int64, valueHandle int32, propertiesOffset MemPtr, hashOffset MemPtr, nameOffset MemPtr, attributesOffset MemPtr, creatorOffset MemPtr, royaltiesHandle int32, urisOffset MemPtr) int32
	GetESDTLocalRoles(tokenIdHandle int32) int64
	ValidateTokenIdentifier(tokenIdHandle int32) int32
	TransferValue(destOffset MemPtr, valueOffset MemPtr, dataOffset MemPtr, length MemLength) int32
	TransferValueExecute(destOffset MemPtr, valueOffset MemPtr, gasLimit int64, functionOffset MemPtr, functionLength MemLength, numArguments int32, argumentsLengthOffset MemPtr, dataOffset MemPtr) int32
	TransferESDTExecute(destOffset MemPtr, tokenIDOffset MemPtr, tokenIDLen MemLength, valueOffset MemPtr, gasLimit int64, functionOffset MemPtr, functionLength MemLength, numArguments int32, argumentsLengthOffset MemPtr, dataOffset MemPtr) int32
	TransferESDTNFTExecute(destOffset MemPtr, tokenIDOffset MemPtr, tokenIDLen MemLength, valueOffset MemPtr, nonce int64, gasLimit int64, functionOffset MemPtr, functionLength MemLength, numArguments int32, argumentsLengthOffset MemPtr, dataOffset MemPtr) int32
	MultiTransferESDTNFTExecute(destOffset MemPtr, numTokenTransfers int32, tokenTransfersArgsLengthOffset MemPtr, tokenTransferDataOffset MemPtr, gasLimit int64, functionOffset MemPtr, functionLength MemLength, numArguments int32, argumentsLengthOffset MemPtr, dataOffset MemPtr) int32
	CreateAsyncCall(destOffset MemPtr, valueOffset MemPtr, dataOffset MemPtr, dataLength MemLength, successOffset MemPtr, successLength MemLength, errorOffset MemPtr, errorLength MemLength, gas int64, extraGasForCallback int64) int32
	SetAsyncContextCallback(callback MemPtr, callbackLength MemLength, data MemPtr, dataLength MemLength, gas int64) int32
	UpgradeContract(destOffset MemPtr, gasLimit int64, valueOffset MemPtr, codeOffset MemPtr, codeMetadataOffset MemPtr, length MemLength, numArguments int32, argumentsLengthOffset MemPtr, dataOffset MemPtr)
	UpgradeFromSourceContract(destOffset MemPtr, gasLimit int64, valueOffset MemPtr, sourceContractAddressOffset MemPtr, codeMetadataOffset MemPtr, numArguments int32, argumentsLengthOffset MemPtr, dataOffset MemPtr)
	DeleteContract(destOffset MemPtr, gasLimit int64, numArguments int32, argumentsLengthOffset MemPtr, dataOffset MemPtr)
	AsyncCall(destOffset MemPtr, valueOffset MemPtr, dataOffset MemPtr, length MemLength)
	GetArgumentLength(id int32) int32
	GetArgument(id int32, argOffset MemPtr) int32
	GetFunction(functionOffset MemPtr) int32
	GetNumArguments() int32
	StorageStore(keyOffset MemPtr, keyLength MemLength, dataOffset MemPtr, dataLength MemLength) int32
	StorageLoadLength(keyOffset MemPtr, keyLength MemLength) int32
	StorageLoadFromAddress(addressOffset MemPtr, keyOffset MemPtr, keyLength MemLength, dataOffset MemPtr) int32
	StorageLoad(keyOffset MemPtr, keyLength MemLength, dataOffset MemPtr) int32
	SetStorageLock(keyOffset MemPtr, keyLength MemLength, lockTimestamp int64) int32
	GetStorageLock(keyOffset MemPtr, keyLength MemLength) int64
	IsStorageLocked(keyOffset MemPtr, keyLength MemLength) int32
	ClearStorageLock(keyOffset MemPtr, keyLength MemLength) int32
	GetCaller(resultOffset MemPtr)
	CheckNoPayment()
	GetCallValue(resultOffset MemPtr) int32
	GetESDTValue(resultOffset MemPtr) int32
	GetESDTValueByIndex(resultOffset MemPtr, index int32) int32
	GetESDTTokenName(resultOffset MemPtr) int32
	GetESDTTokenNameByIndex(resultOffset MemPtr, index int32) int32
	GetESDTTokenNonce() int64
	GetESDTTokenNonceByIndex(index int32) int64
	GetCurrentESDTNFTNonce(addressOffset MemPtr, tokenIDOffset MemPtr, tokenIDLen MemLength) int64
	GetESDTTokenType() int32
	GetESDTTokenTypeByIndex(index int32) int32
	GetNumESDTTransfers() int32
	GetCallValueTokenName(callValueOffset MemPtr, tokenNameOffset MemPtr) int32
	GetCallValueTokenNameByIndex(callValueOffset MemPtr, tokenNameOffset MemPtr, index int32) int32
	WriteLog(dataPointer MemPtr, dataLength MemLength, topicPtr MemPtr, numTopics int32)
	WriteEventLog(numTopics int32, topicLengthsOffset MemPtr, topicOffset MemPtr, dataOffset MemPtr, dataLength MemLength)
	GetBlockTimestamp() int64
	GetBlockNonce() int64
	GetBlockRound() int64
	GetBlockEpoch() int64
	GetBlockRandomSeed(pointer MemPtr)
	GetStateRootHash(pointer MemPtr)
	GetPrevBlockTimestamp() int64
	GetPrevBlockNonce() int64
	GetPrevBlockRound() int64
	GetPrevBlockEpoch() int64
	GetPrevBlockRandomSeed(pointer MemPtr)
	Finish(pointer MemPtr, length MemLength)
	ExecuteOnSameContext(gasLimit int64, addressOffset MemPtr, valueOffset MemPtr, functionOffset MemPtr, functionLength MemLength, numArguments int32, argumentsLengthOffset MemPtr, dataOffset MemPtr) int32
	ExecuteOnDestContext(gasLimit int64, addressOffset MemPtr, valueOffset MemPtr, functionOffset MemPtr, functionLength MemLength, numArguments int32, argumentsLengthOffset MemPtr, dataOffset MemPtr) int32
	ExecuteReadOnly(gasLimit int64, addressOffset MemPtr, functionOffset MemPtr, functionLength MemLength, numArguments int32, argumentsLengthOffset MemPtr, dataOffset MemPtr) int32
	CreateContract(gasLimit int64, valueOffset MemPtr, codeOffset MemPtr, codeMetadataOffset MemPtr, length MemLength, resultOffset MemPtr, numArguments int32, argumentsLengthOffset MemPtr, dataOffset MemPtr) int32
	DeployFromSourceContract(gasLimit int64, valueOffset MemPtr, sourceContractAddressOffset MemPtr, codeMetadataOffset MemPtr, resultAddressOffset MemPtr, numArguments int32, argumentsLengthOffset MemPtr, dataOffset MemPtr) int32
	GetNumReturnData() int32
	GetReturnDataSize(resultID int32) int32
	GetReturnData(resultID int32, dataOffset MemPtr) int32
	CleanReturnData()
	DeleteFromReturnData(resultID int32)
	GetOriginalTxHash(dataOffset MemPtr)
	GetCurrentTxHash(dataOffset MemPtr)
	GetPrevTxHash(dataOffset MemPtr)
}

type ManagedVMHooks interface {
	ManagedSCAddress(destinationHandle int32)
	ManagedOwnerAddress(destinationHandle int32)
	ManagedCaller(destinationHandle int32)
	ManagedSignalError(errHandle int32)
	ManagedWriteLog(topicsHandle int32, dataHandle int32)
	ManagedGetOriginalTxHash(resultHandle int32)
	ManagedGetStateRootHash(resultHandle int32)
	ManagedGetBlockRandomSeed(resultHandle int32)
	ManagedGetPrevBlockRandomSeed(resultHandle int32)
	ManagedGetReturnData(resultID int32, resultHandle int32)
	ManagedGetMultiESDTCallValue(multiCallValueHandle int32)
	ManagedGetESDTBalance(addressHandle int32, tokenIDHandle int32, nonce int64, valueHandle int32)
	ManagedGetESDTTokenData(addressHandle int32, tokenIDHandle int32, nonce int64, valueHandle int32, propertiesHandle int32, hashHandle int32, nameHandle int32, attributesHandle int32, creatorHandle int32, royaltiesHandle int32, urisHandle int32)
	ManagedAsyncCall(destHandle int32, valueHandle int32, functionHandle int32, argumentsHandle int32)
	ManagedCreateAsyncCall(destHandle int32, valueHandle int32, functionHandle int32, argumentsHandle int32, successOffset MemPtr, successLength MemLength, errorOffset MemPtr, errorLength MemLength, gas int64, extraGasForCallback int64, callbackClosureHandle int32) int32
	ManagedGetCallbackClosure(callbackClosureHandle int32)
	ManagedUpgradeFromSourceContract(destHandle int32, gas int64, valueHandle int32, addressHandle int32, codeMetadataHandle int32, argumentsHandle int32, resultHandle int32)
	ManagedUpgradeContract(destHandle int32, gas int64, valueHandle int32, codeHandle int32, codeMetadataHandle int32, argumentsHandle int32, resultHandle int32)
	ManagedDeleteContract(destHandle int32, gasLimit int64, argumentsHandle int32)
	ManagedDeployFromSourceContract(gas int64, valueHandle int32, addressHandle int32, codeMetadataHandle int32, argumentsHandle int32, resultAddressHandle int32, resultHandle int32) int32
	ManagedCreateContract(gas int64, valueHandle int32, codeHandle int32, codeMetadataHandle int32, argumentsHandle int32, resultAddressHandle int32, resultHandle int32) int32
	ManagedExecuteReadOnly(gas int64, addressHandle int32, functionHandle int32, argumentsHandle int32, resultHandle int32) int32
	ManagedExecuteOnSameContext(gas int64, addressHandle int32, valueHandle int32, functionHandle int32, argumentsHandle int32, resultHandle int32) int32
	ManagedExecuteOnDestContext(gas int64, addressHandle int32, valueHandle int32, functionHandle int32, argumentsHandle int32, resultHandle int32) int32
	ManagedMultiTransferESDTNFTExecute(dstHandle int32, tokenTransfersHandle int32, gasLimit int64, functionHandle int32, argumentsHandle int32) int32
	ManagedTransferValueExecute(dstHandle int32, valueHandle int32, gasLimit int64, functionHandle int32, argumentsHandle int32) int32
	ManagedIsESDTFrozen(addressHandle int32, tokenIDHandle int32, nonce int64) int32
	ManagedIsESDTLimitedTransfer(tokenIDHandle int32) int32
	ManagedIsESDTPaused(tokenIDHandle int32) int32
	ManagedBufferToHex(sourceHandle int32, destHandle int32)
}

type BigFloatVMHooks interface {
	BigFloatNewFromParts(integralPart int32, fractionalPart int32, exponent int32) int32
	BigFloatNewFromFrac(numerator int64, denominator int64) int32
	BigFloatNewFromSci(significand int64, exponent int64) int32
	BigFloatAdd(destinationHandle int32, op1Handle int32, op2Handle int32)
	BigFloatSub(destinationHandle int32, op1Handle int32, op2Handle int32)
	BigFloatMul(destinationHandle int32, op1Handle int32, op2Handle int32)
	BigFloatDiv(destinationHandle int32, op1Handle int32, op2Handle int32)
	BigFloatNeg(destinationHandle int32, opHandle int32)
	BigFloatClone(destinationHandle int32, opHandle int32)
	BigFloatCmp(op1Handle int32, op2Handle int32) int32
	BigFloatAbs(destinationHandle int32, opHandle int32)
	BigFloatSign(opHandle int32) int32
	BigFloatSqrt(destinationHandle int32, opHandle int32)
	BigFloatPow(destinationHandle int32, opHandle int32, exponent int32)
	BigFloatFloor(destBigIntHandle int32, opHandle int32)
	BigFloatCeil(destBigIntHandle int32, opHandle int32)
	BigFloatTruncate(destBigIntHandle int32, opHandle int32)
	BigFloatSetInt64(destinationHandle int32, value int64)
	BigFloatIsInt(opHandle int32) int32
	BigFloatSetBigInt(destinationHandle int32, bigIntHandle int32)
	BigFloatGetConstPi(destinationHandle int32)
	BigFloatGetConstE(destinationHandle int32)
}

type BigIntVMHooks interface {
	BigIntGetUnsignedArgument(id int32, destinationHandle int32)
	BigIntGetSignedArgument(id int32, destinationHandle int32)
	BigIntStorageStoreUnsigned(keyOffset MemPtr, keyLength MemLength, sourceHandle int32) int32
	BigIntStorageLoadUnsigned(keyOffset MemPtr, keyLength MemLength, destinationHandle int32) int32
	BigIntGetCallValue(destinationHandle int32)
	BigIntGetESDTCallValue(destination int32)
	BigIntGetESDTCallValueByIndex(destinationHandle int32, index int32)
	BigIntGetExternalBalance(addressOffset MemPtr, result int32)
	BigIntGetESDTExternalBalance(addressOffset MemPtr, tokenIDOffset MemPtr, tokenIDLen MemLength, nonce int64, resultHandle int32)
	BigIntNew(smallValue int64) int32
	BigIntUnsignedByteLength(referenceHandle int32) int32
	BigIntSignedByteLength(referenceHandle int32) int32
	BigIntGetUnsignedBytes(referenceHandle int32, byteOffset MemPtr) int32
	BigIntGetSignedBytes(referenceHandle int32, byteOffset MemPtr) int32
	BigIntSetUnsignedBytes(destinationHandle int32, byteOffset MemPtr, byteLength MemLength)
	BigIntSetSignedBytes(destinationHandle int32, byteOffset MemPtr, byteLength MemLength)
	BigIntIsInt64(destinationHandle int32) int32
	BigIntGetInt64(destinationHandle int32) int64
	BigIntSetInt64(destinationHandle int32, value int64)
	BigIntAdd(destinationHandle int32, op1Handle int32, op2Handle int32)
	BigIntSub(destinationHandle int32, op1Handle int32, op2Handle int32)
	BigIntMul(destinationHandle int32, op1Handle int32, op2Handle int32)
	BigIntTDiv(destinationHandle int32, op1Handle int32, op2Handle int32)
	BigIntTMod(destinationHandle int32, op1Handle int32, op2Handle int32)
	BigIntEDiv(destinationHandle int32, op1Handle int32, op2Handle int32)
	BigIntEMod(destinationHandle int32, op1Handle int32, op2Handle int32)
	BigIntSqrt(destinationHandle int32, opHandle int32)
	BigIntPow(destinationHandle int32, op1Handle int32, op2Handle int32)
	BigIntLog2(op1Handle int32) int32
	BigIntAbs(destinationHandle int32, opHandle int32)
	BigIntNeg(destinationHandle int32, opHandle int32)
	BigIntSign(opHandle int32) int32
	BigIntCmp(op1Handle int32, op2Handle int32) int32
	BigIntNot(destinationHandle int32, opHandle int32)
	BigIntAnd(destinationHandle int32, op1Handle int32, op2Handle int32)
	BigIntOr(destinationHandle int32, op1Handle int32, op2Handle int32)
	BigIntXor(destinationHandle int32, op1Handle int32, op2Handle int32)
	BigIntShr(destinationHandle int32, opHandle int32, bits int32)
	BigIntShl(destinationHandle int32, opHandle int32, bits int32)
	BigIntFinishUnsigned(referenceHandle int32)
	BigIntFinishSigned(referenceHandle int32)
	BigIntToString(bigIntHandle int32, destinationHandle int32)
}

type ManagedBufferVMHooks interface {
	MBufferNew() int32
	MBufferNewFromBytes(dataOffset MemPtr, dataLength MemLength) int32
	MBufferGetLength(mBufferHandle int32) int32
	MBufferGetBytes(mBufferHandle int32, resultOffset MemPtr) int32
	MBufferGetByteSlice(sourceHandle int32, startingPosition int32, sliceLength int32, resultOffset MemPtr) int32
	MBufferCopyByteSlice(sourceHandle int32, startingPosition int32, sliceLength int32, destinationHandle int32) int32
	MBufferEq(mBufferHandle1 int32, mBufferHandle2 int32) int32
	MBufferSetBytes(mBufferHandle int32, dataOffset MemPtr, dataLength MemLength) int32
	MBufferSetByteSlice(mBufferHandle int32, startingPosition int32, dataLength MemLength, dataOffset MemPtr) int32
	MBufferAppend(accumulatorHandle int32, dataHandle int32) int32
	MBufferAppendBytes(accumulatorHandle int32, dataOffset MemPtr, dataLength MemLength) int32
	MBufferToBigIntUnsigned(mBufferHandle int32, bigIntHandle int32) int32
	MBufferToBigIntSigned(mBufferHandle int32, bigIntHandle int32) int32
	MBufferFromBigIntUnsigned(mBufferHandle int32, bigIntHandle int32) int32
	MBufferFromBigIntSigned(mBufferHandle int32, bigIntHandle int32) int32
	MBufferToBigFloat(mBufferHandle int32, bigFloatHandle int32) int32
	MBufferFromBigFloat(mBufferHandle int32, bigFloatHandle int32) int32
	MBufferStorageStore(keyHandle int32, sourceHandle int32) int32
	MBufferStorageLoad(keyHandle int32, destinationHandle int32) int32
	MBufferStorageLoadFromAddress(addressHandle int32, keyHandle int32, destinationHandle int32)
	MBufferGetArgument(id int32, destinationHandle int32) int32
	MBufferFinish(sourceHandle int32) int32
	MBufferSetRandom(destinationHandle int32, length int32) int32
}

type SmallIntVMHooks interface {
	SmallIntGetUnsignedArgument(id int32) int64
	SmallIntGetSignedArgument(id int32) int64
	SmallIntFinishUnsigned(value int64)
	SmallIntFinishSigned(value int64)
	SmallIntStorageStoreUnsigned(keyOffset MemPtr, keyLength MemLength, value int64) int32
	SmallIntStorageStoreSigned(keyOffset MemPtr, keyLength MemLength, value int64) int32
	SmallIntStorageLoadUnsigned(keyOffset MemPtr, keyLength MemLength) int64
	SmallIntStorageLoadSigned(keyOffset MemPtr, keyLength MemLength) int64
	Int64getArgument(id int32) int64
	Int64finish(value int64)
	Int64storageStore(keyOffset MemPtr, keyLength MemLength, value int64) int32
	Int64storageLoad(keyOffset MemPtr, keyLength MemLength) int64
}

type CryptoVMHooks interface {
	Sha256(dataOffset MemPtr, length MemLength, resultOffset MemPtr) int32
	ManagedSha256(inputHandle int32, outputHandle int32) int32
	Keccak256(dataOffset MemPtr, length MemLength, resultOffset MemPtr) int32
	ManagedKeccak256(inputHandle int32, outputHandle int32) int32
	Ripemd160(dataOffset MemPtr, length MemLength, resultOffset MemPtr) int32
	ManagedRipemd160(inputHandle int32, outputHandle int32) int32
	VerifyBLS(keyOffset MemPtr, messageOffset MemPtr, messageLength MemLength, sigOffset MemPtr) int32
	ManagedVerifyBLS(keyHandle int32, messageHandle int32, sigHandle int32) int32
	VerifyEd25519(keyOffset MemPtr, messageOffset MemPtr, messageLength MemLength, sigOffset MemPtr) int32
	ManagedVerifyEd25519(keyHandle int32, messageHandle int32, sigHandle int32) int32
	VerifyCustomSecp256k1(keyOffset MemPtr, keyLength MemLength, messageOffset MemPtr, messageLength MemLength, sigOffset MemPtr, hashType int32) int32
	ManagedVerifyCustomSecp256k1(keyHandle int32, messageHandle int32, sigHandle int32, hashType int32) int32
	VerifySecp256k1(keyOffset MemPtr, keyLength MemLength, messageOffset MemPtr, messageLength MemLength, sigOffset MemPtr) int32
	ManagedVerifySecp256k1(keyHandle int32, messageHandle int32, sigHandle int32) int32
	EncodeSecp256k1DerSignature(rOffset MemPtr, rLength MemLength, sOffset MemPtr, sLength MemLength, sigOffset MemPtr) int32
	ManagedEncodeSecp256k1DerSignature(rHandle int32, sHandle int32, sigHandle int32) int32
	AddEC(xResultHandle int32, yResultHandle int32, ecHandle int32, fstPointXHandle int32, fstPointYHandle int32, sndPointXHandle int32, sndPointYHandle int32)
	DoubleEC(xResultHandle int32, yResultHandle int32, ecHandle int32, pointXHandle int32, pointYHandle int32)
	IsOnCurveEC(ecHandle int32, pointXHandle int32, pointYHandle int32) int32
	ScalarBaseMultEC(xResultHandle int32, yResultHandle int32, ecHandle int32, dataOffset MemPtr, length MemLength) int32
	ManagedScalarBaseMultEC(xResultHandle int32, yResultHandle int32, ecHandle int32, dataHandle int32) int32
	ScalarMultEC(xResultHandle int32, yResultHandle int32, ecHandle int32, pointXHandle int32, pointYHandle int32, dataOffset MemPtr, length MemLength) int32
	ManagedScalarMultEC(xResultHandle int32, yResultHandle int32, ecHandle int32, pointXHandle int32, pointYHandle int32, dataHandle int32) int32
	MarshalEC(xPairHandle int32, yPairHandle int32, ecHandle int32, resultOffset MemPtr) int32
	ManagedMarshalEC(xPairHandle int32, yPairHandle int32, ecHandle int32, resultHandle int32) int32
	MarshalCompressedEC(xPairHandle int32, yPairHandle int32, ecHandle int32, resultOffset MemPtr) int32
	ManagedMarshalCompressedEC(xPairHandle int32, yPairHandle int32, ecHandle int32, resultHandle int32) int32
	UnmarshalEC(xResultHandle int32, yResultHandle int32, ecHandle int32, dataOffset MemPtr, length MemLength) int32
	ManagedUnmarshalEC(xResultHandle int32, yResultHandle int32, ecHandle int32, dataHandle int32) int32
	UnmarshalCompressedEC(xResultHandle int32, yResultHandle int32, ecHandle int32, dataOffset MemPtr, length MemLength) int32
	ManagedUnmarshalCompressedEC(xResultHandle int32, yResultHandle int32, ecHandle int32, dataHandle int32) int32
	GenerateKeyEC(xPubKeyHandle int32, yPubKeyHandle int32, ecHandle int32, resultOffset MemPtr) int32
	ManagedGenerateKeyEC(xPubKeyHandle int32, yPubKeyHandle int32, ecHandle int32, resultHandle int32) int32
	CreateEC(dataOffset MemPtr, dataLength MemLength) int32
	ManagedCreateEC(dataHandle int32) int32
	GetCurveLengthEC(ecHandle int32) int32
	GetPrivKeyByteLengthEC(ecHandle int32) int32
	EllipticCurveGetValues(ecHandle int32, fieldOrderHandle int32, basePointOrderHandle int32, eqConstantHandle int32, xBasePointHandle int32, yBasePointHandle int32) int32
}
