package gasschedules

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// !!!!!!!!!!!!!!!!!!!!!! AUTO-GENERATED FILE !!!!!!!!!!!!!!!!!!!!!!
// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

// Please do not edit manually!
// Call `go generate` in `arwen-wasm-vm-v1_4/arwenmandos/gasSchedules` to update it.

const (
	gasScheduleV3 = `[BuiltInCost]
    ChangeOwnerAddress       = 5000000
    ClaimDeveloperRewards    = 5000000
    SaveUserName             = 1000000
    SaveKeyValue             = 250000
    ESDTTransfer             = 250000
    ESDTBurn                 = 250000
    ESDTLocalMint            = 250000
    ESDTLocalBurn            = 250000
    ESDTNFTCreate            = 500000
    ESDTNFTAddQuantity       = 500000
    ESDTNFTBurn              = 500000
    ESDTNFTTransfer          = 500000
    ESDTNFTChangeCreateOwner = 1000000
    ESDTNFTAddUri            = 500000
    ESDTNFTUpdateAttributes  = 500000
    ESDTNFTMultiTransfer     = 1000000

[MetaChainSystemSCsCost]
    Stake               = 5000000
    UnStake             = 5000000
    UnBond              = 5000000
    Claim               = 5000000
    Get                 = 5000000
    ChangeRewardAddress = 5000000
    ChangeValidatorKeys = 5000000
    UnJail              = 5000000
    DelegationOps       = 1000000
    DelegationMgrOps    = 50000000
    ESDTIssue           = 50000000
    ESDTOperations      = 50000000
    Proposal            = 5000000
    Vote                = 500000
    DelegateVote        = 1000000
    RevokeVote          = 500000
    CloseProposal       = 1000000
    GetAllNodeStates    = 20000000
    UnstakeTokens       = 5000000
    UnbondTokens        = 5000000

[BaseOperationCost]
    StorePerByte      = 50000
    ReleasePerByte    = 10000
    DataCopyPerByte   = 1000
    PersistPerByte    = 10000
    CompilePerByte    = 300
    AoTPreparePerByte = 300
    GetCode           = 1000000

[ElrondAPICost]
    GetSCAddress       = 100
    GetOwnerAddress    = 5000
    IsSmartContract    = 5000
    GetShardOfAddress  = 5000
    GetExternalBalance = 7000
    GetBlockHash       = 10000
    GetOriginalTxHash  = 10000
    TransferValue      = 150000
    GetArgument        = 100
    GetFunction        = 100
    GetNumArguments    = 100
    StorageStore       = 250000
    StorageLoad        = 100000
    CachedStorageLoad  = 100
    GetCaller          = 100
    GetCallValue       = 100
    Log                = 3750
    Finish             = 1
    SignalError        = 1
    GetBlockTimeStamp  = 10000
    GetGasLeft         = 100
    Int64GetArgument   = 100
    Int64StorageStore  = 250000
    Int64StorageLoad   = 100000
    Int64Finish        = 1000
    GetStateRootHash   = 10000
    GetBlockNonce      = 10000
    GetBlockEpoch      = 10000
    GetBlockRound      = 10000
    GetBlockRandomSeed = 10000
    ExecuteOnSameContext = 160000
    ExecuteOnDestContext = 160000
    DelegateExecution    = 160000
    AsyncCallStep        = 200000
    AsyncCallbackGasLock = 2000000
    ExecuteReadOnly      = 160000
    CreateContract       = 300000
    GetReturnData        = 100
    GetNumReturnData     = 100
    GetReturnDataSize    = 100
    CleanReturnData      = 100
    DeleteFromReturnData = 100    

[EthAPICost]
    UseGas              = 100
    GetAddress          = 100000
    GetExternalBalance  = 70000
    GetBlockHash        = 100000
    Call                = 160000
    CallDataCopy        = 200
    GetCallDataSize     = 100
    CallCode            = 160000
    CallDelegate        = 160000
    CallStatic          = 160000
    StorageStore        = 250000
    StorageLoad         = 100000
    GetCaller           = 100
    GetCallValue        = 100
    CodeCopy            = 1000
    GetCodeSize         = 100
    GetBlockCoinbase    = 100
    Create              = 320000
    GetBlockDifficulty  = 100
    ExternalCodeCopy    = 3000
    GetExternalCodeSize = 2500
    GetGasLeft          = 100
    GetBlockGasLimit    = 100000
    GetTxGasPrice       = 1000
    Log                 = 3750
    GetBlockNumber      = 100000
    GetTxOrigin         = 100000
    Finish              = 1
    Revert              = 1
    GetReturnDataSize   = 200
    ReturnDataCopy      = 500
    SelfDestruct        = 5000000
    GetBlockTimeStamp   = 100000

[BigIntAPICost]
    BigIntNew                = 2000
    BigIntByteLength         = 2000
    BigIntUnsignedByteLength = 2000
    BigIntSignedByteLength   = 2000
    BigIntGetBytes           = 2000
    BigIntGetUnsignedBytes   = 2000
    BigIntGetSignedBytes     = 2000
    BigIntSetBytes           = 2000
    BigIntSetUnsignedBytes   = 2000
    BigIntSetSignedBytes     = 2000
    BigIntIsInt64            = 2000
    BigIntGetInt64           = 2000
    BigIntSetInt64           = 2000
    BigIntAdd                = 2000
    BigIntSub                = 2000
    BigIntMul                = 6000
    BigIntSqrt               = 6000
    BigIntPow                = 6000
    BigIntLog                = 6000
    BigIntTDiv               = 6000
    BigIntTMod               = 6000
    BigIntEDiv               = 6000
    BigIntEMod               = 6000
    BigIntAbs                = 2000
    BigIntNeg                = 2000
    BigIntSign               = 2000
    BigIntCmp                = 2000
    BigIntNot                = 2000
    BigIntAnd                = 2000
    BigIntOr                 = 2000
    BigIntXor                = 2000
    BigIntShr                = 2000
    BigIntShl                = 2000
    BigIntFinishUnsigned     = 1000
    BigIntFinishSigned       = 1000
    BigIntStorageLoadUnsigned   = 100000
    BigIntStorageStoreUnsigned  = 250000
    BigIntGetArgument           = 1000
    BigIntGetUnsignedArgument   = 1000
    BigIntGetSignedArgument     = 1000
    BigIntGetCallValue          = 1000
    BigIntGetExternalBalance    = 10000
    CopyPerByteForTooBig        = 1000

    [BigFloatAPICost]
    BigFloatNewFromParts = 3000
    BigFloatAdd          = 7000
    BigFloatSub          = 7000
    BigFloatMul          = 7000
    BigFloatDiv          = 7000
    BigFloatTruncate     = 5000
    BigFloatNeg          = 5000
    BigFloatClone        = 5000
    BigFloatCmp          = 4000
    BigFloatAbs          = 5000
    BigFloatSqrt         = 7000
    BigFloatPow          = 10000
    BigFloatFloor        = 5000
    BigFloatCeil         = 5000
    BigFloatIsInt        = 3000
    BigFloatSetBigInt    = 3000
    BigFloatSetInt64     = 1000
    BigFloatGetConst     = 1000
    
[CryptoAPICost]
    SHA256                 = 1000000
    Keccak256              = 1000000
    Ripemd160              = 1000000
    VerifyBLS              = 5000000
    VerifyEd25519          = 2000000
    VerifySecp256k1        = 2000000
    EllipticCurveNew       = 10000
    AddECC                 = 75000
    DoubleECC              = 65000
    IsOnCurveECC           = 10000
    ScalarMultECC          = 400000
    MarshalECC             = 13000
    MarshalCompressedECC   = 15000
    UnmarshalECC           = 20000
    UnmarshalCompressedECC = 270000
    GenerateKeyECC         = 7000000
    EncodeDERSig           = 10000000

[ManagedBufferAPICost]
    MBufferNew                   = 2000
    MBufferNewFromBytes          = 4000
    MBufferGetLength             = 2000
    MBufferGetBytes              = 2000
    MBufferGetByteSlice          = 2000
    MBufferCopyByteSlice         = 2000
    MBufferSetBytes              = 2000
    MBufferAppend                = 2000
    MBufferAppendBytes           = 2000
    MBufferToBigIntUnsigned      = 4000
    MBufferToBigIntSigned        = 10000
    MBufferFromBigIntUnsigned    = 4000
    MBufferFromBigIntSigned      = 10000
    MBufferToBigFloat            = 2000
    MBufferFromBigFloat          = 2000
    MBufferStorageStore          = 75000
    MBufferStorageLoad           = 50000
    MBufferGetArgument           = 1000
    MBufferFinish                = 1000
    MBufferSetRandom             = 6000

[WASMOpcodeCost]
    Unreachable = 1
    Nop = 1
    Block = 1
    Loop = 1
    If = 1
    Else = 2
    End = 2
    Br = 2
    BrIf = 3
    BrTable = 2
    Return = 3
    Call = 3
    CallIndirect = 3
    Drop = 3
    Select = 3
    TypedSelect = 3
    LocalGet = 3
    LocalSet = 3
    LocalTee = 3
    GlobalGet = 3
    GlobalSet = 3
    I32Load = 3
    I64Load = 3
    F32Load = 6
    F64Load = 6
    I32Load8S = 3
    I32Load8U = 3
    I32Load16S = 3
    I32Load16U = 3
    I64Load8S = 3
    I64Load8U = 3
    I64Load16S = 3
    I64Load16U = 3
    I64Load32S = 3
    I64Load32U = 3
    I32Store = 3
    I64Store = 3
    F32Store = 12
    F64Store = 12
    I32Store8 = 3
    I32Store16 = 3
    I64Store8 = 3
    I64Store16 = 3
    I64Store32 = 3
    MemorySize = 5
    MemoryGrow = 5
    I32Const = 1
    I64Const = 1
    F32Const = 1
    F64Const = 1
    RefNull = 1
    RefIsNull = 1
    RefFunc = 1
    I32Eqz = 1
    I32Eq = 1
    I32Ne = 1
    I32LtS = 1
    I32LtU = 1
    I32GtS = 1
    I32GtU = 1
    I32LeS = 1
    I32LeU = 1
    I32GeS = 1
    I32GeU = 1
    I64Eqz = 1
    I64Eq = 1
    I64Ne = 1
    I64LtS = 1
    I64LtU = 1
    I64GtS = 1
    I64GtU = 1
    I64LeS = 1
    I64LeU = 1
    I64GeS = 1
    I64GeU = 1
    F32Eq = 6
    F32Ne = 6
    F32Lt = 6
    F32Gt = 6
    F32Le = 6
    F32Ge = 6
    F64Eq = 6
    F64Ne = 6
    F64Lt = 6
    F64Gt = 6
    F64Le = 6
    F64Ge = 6
    I32Clz = 100
    I32Ctz = 100
    I32Popcnt = 100
    I32Add = 1
    I32Sub = 1
    I32Mul = 3
    I32DivS = 18
    I32DivU = 18
    I32RemS = 18
    I32RemU = 18
    I32And = 1
    I32Or = 1
    I32Xor = 1
    I32Shl = 3
    I32ShrS = 3
    I32ShrU = 3
    I32Rotl = 5
    I32Rotr = 5
    I64Clz = 100
    I64Ctz = 100
    I64Popcnt = 100
    I64Add = 1
    I64Sub = 1
    I64Mul = 3
    I64DivS = 18
    I64DivU = 18
    I64RemS = 18
    I64RemU = 18
    I64And = 1
    I64Or = 1
    I64Xor = 1
    I64Shl = 3
    I64ShrS = 3
    I64ShrU = 3
    I64Rotl = 5
    I64Rotr = 5
    F32Abs = 5
    F32Neg = 5
    F32Ceil = 100
    F32Floor = 100
    F32Trunc = 100
    F32Nearest = 100
    F32Sqrt = 100
    F32Add = 5
    F32Sub = 5
    F32Mul = 15
    F32Div = 100
    F32Min = 15
    F32Max = 15
    F32Copysign = 5
    F64Abs = 5
    F64Neg = 5
    F64Ceil = 100
    F64Floor = 100
    F64Trunc = 100
    F64Nearest = 100
    F64Sqrt = 100
    F64Add = 5
    F64Sub = 5
    F64Mul = 15
    F64Div = 100
    F64Min = 15
    F64Max = 15
    F64Copysign = 5
    I32WrapI64 = 9
    I32TruncF32S = 100
    I32TruncF32U = 100
    I32TruncF64S = 100
    I32TruncF64U = 100
    I64ExtendI32S = 9
    I64ExtendI32U = 9
    I64TruncF32S = 100
    I64TruncF32U = 100
    I64TruncF64S = 100
    I64TruncF64U = 100
    F32ConvertI32S = 100
    F32ConvertI32U = 100
    F32ConvertI64S = 100
    F32ConvertI64U = 100
    F32DemoteF64 = 100
    F64ConvertI32S = 100
    F64ConvertI32U = 100
    F64ConvertI64S = 100
    F64ConvertI64U = 100
    F64PromoteF32 = 100
    I32ReinterpretF32 = 100
    I64ReinterpretF64 = 100
    F32ReinterpretI32 = 100
    F64ReinterpretI64 = 100
    I32Extend8S = 9
    I32Extend16S = 9
    I64Extend8S = 9
    I64Extend16S = 9
    I64Extend32S = 9
    I32TruncSatF32S = 100
    I32TruncSatF32U = 100
    I32TruncSatF64S = 100
    I32TruncSatF64U = 100
    I64TruncSatF32S = 100
    I64TruncSatF32U = 100
    I64TruncSatF64S = 100
    I64TruncSatF64U = 100
    MemoryInit = 5
    DataDrop = 5
    MemoryCopy = 5
    MemoryFill = 5
    TableInit = 10
    ElemDrop = 10
    TableCopy = 10
    TableFill = 10
    TableGet = 10
    TableSet = 10
    TableGrow = 10
    TableSize = 10
    AtomicNotify = 10
    I32AtomicWait = 10
    I64AtomicWait = 10
    AtomicFence = 10
    I32AtomicLoad = 15
    I64AtomicLoad = 15
    I32AtomicLoad8U = 15
    I32AtomicLoad16U = 15
    I64AtomicLoad8U = 15
    I64AtomicLoad16U = 15
    I64AtomicLoad32U = 15
    I32AtomicStore = 15
    I64AtomicStore = 15
    I32AtomicStore8 = 15
    I32AtomicStore16 = 15
    I64AtomicStore8 = 15
    I64AtomicStore16 = 15
    I64AtomicStore32 = 15
    I32AtomicRmwAdd = 20
    I64AtomicRmwAdd = 20
    I32AtomicRmw8AddU = 20
    I32AtomicRmw16AddU = 20
    I64AtomicRmw8AddU = 20
    I64AtomicRmw16AddU = 20
    I64AtomicRmw32AddU = 20
    I32AtomicRmwSub = 20
    I64AtomicRmwSub = 20
    I32AtomicRmw8SubU = 20
    I32AtomicRmw16SubU = 20
    I64AtomicRmw8SubU = 20
    I64AtomicRmw16SubU = 20
    I64AtomicRmw32SubU = 20
    I32AtomicRmwAnd = 15
    I64AtomicRmwAnd = 15
    I32AtomicRmw8AndU = 15
    I32AtomicRmw16AndU = 15
    I64AtomicRmw8AndU = 15
    I64AtomicRmw16AndU = 15
    I64AtomicRmw32AndU = 15
    I32AtomicRmwOr = 15
    I64AtomicRmwOr = 15
    I32AtomicRmw8OrU = 15
    I32AtomicRmw16OrU = 15
    I64AtomicRmw8OrU = 15
    I64AtomicRmw16OrU = 15
    I64AtomicRmw32OrU = 15
    I32AtomicRmwXor = 15
    I64AtomicRmwXor = 15
    I32AtomicRmw8XorU = 15
    I32AtomicRmw16XorU = 15
    I64AtomicRmw8XorU = 15
    I64AtomicRmw16XorU = 15
    I64AtomicRmw32XorU = 15
    I32AtomicRmwXchg = 30
    I64AtomicRmwXchg = 30
    I32AtomicRmw8XchgU = 30
    I32AtomicRmw16XchgU = 30
    I64AtomicRmw8XchgU = 30
    I64AtomicRmw16XchgU = 30
    I64AtomicRmw32XchgU = 30
    I32AtomicRmwCmpxchg = 30
    I64AtomicRmwCmpxchg = 30
    I32AtomicRmw8CmpxchgU = 30
    I32AtomicRmw16CmpxchgU = 30
    I64AtomicRmw8CmpxchgU = 30
    I64AtomicRmw16CmpxchgU = 30
    I64AtomicRmw32CmpxchgU = 30
    V128Load = 18
    V128Store = 18
    V128Const = 18
    I8x16Splat = 20
    I8x16ExtractLaneS = 20
    I8x16ExtractLaneU = 20
    I8x16ReplaceLane = 20
    I16x8Splat = 20
    I16x8ExtractLaneS = 20
    I16x8ExtractLaneU = 20
    I16x8ReplaceLane = 20
    I32x4Splat = 20
    I32x4ExtractLane = 20
    I32x4ReplaceLane = 20
    I64x2Splat = 20
    I64x2ExtractLane = 20
    I64x2ReplaceLane = 20
    F32x4Splat = 120
    F32x4ExtractLane = 120
    F32x4ReplaceLane = 120
    F64x2Splat = 120
    F64x2ExtractLane = 120
    F64x2ReplaceLane = 120
    I8x16Eq = 30
    I8x16Ne = 30
    I8x16LtS = 40
    I8x16LtU = 40
    I8x16GtS = 40
    I8x16GtU = 40
    I8x16LeS = 40
    I8x16LeU = 40
    I8x16GeS = 40
    I8x16GeU = 40
    I16x8Eq = 30
    I16x8Ne = 30
    I16x8LtS = 40
    I16x8LtU = 40
    I16x8GtS = 40
    I16x8GtU = 40
    I16x8LeS = 40
    I16x8LeU = 40
    I16x8GeS = 40
    I16x8GeU = 40
    I32x4Eq = 30
    I32x4Ne = 30
    I32x4LtS = 40
    I32x4LtU = 40
    I32x4GtS = 40
    I32x4GtU = 40
    I32x4LeS = 40
    I32x4LeU = 40
    I32x4GeS = 40
    I32x4GeU = 40
    F32x4Eq = 120
    F32x4Ne = 120
    F32x4Lt = 120
    F32x4Gt = 120
    F32x4Le = 120
    F32x4Ge = 120
    F64x2Eq = 120
    F64x2Ne = 120
    F64x2Lt = 120
    F64x2Gt = 120
    F64x2Le = 120
    F64x2Ge = 120
    V128Not = 40
    V128And = 40
    V128AndNot = 40
    V128Or = 40
    V128Xor = 40
    V128Bitselect = 40
    I8x16Neg = 20
    I8x16AnyTrue = 20
    I8x16AllTrue = 20
    I8x16Shl = 30
    I8x16ShrS = 30
    I8x16ShrU = 30
    I8x16Add = 20
    I8x16AddSaturateS = 20
    I8x16AddSaturateU = 20
    I8x16Sub = 20
    I8x16SubSaturateS = 20
    I8x16SubSaturateU = 20
    I8x16MinS = 40
    I8x16MinU = 40
    I8x16MaxS = 40
    I8x16MaxU = 40
    I8x16Mul = 80
    I16x8Neg = 40
    I16x8AnyTrue = 40
    I16x8AllTrue = 40
    I16x8Shl = 30
    I16x8ShrS = 30
    I16x8ShrU = 30
    I16x8Add = 20
    I16x8AddSaturateS = 20
    I16x8AddSaturateU = 20
    I16x8Sub = 20
    I16x8SubSaturateS = 20
    I16x8SubSaturateU = 20
    I16x8Mul = 40
    I16x8MinS = 40
    I16x8MinU = 40
    I16x8MaxS = 40
    I16x8MaxU = 40
    I32x4Neg = 20
    I32x4AnyTrue = 20
    I32x4AllTrue = 20
    I32x4Shl = 30
    I32x4ShrS = 30
    I32x4ShrU = 30
    I32x4Add = 20
    I32x4Sub = 20
    I32x4Mul = 80
    I32x4MinS = 40
    I32x4MinU = 40
    I32x4MaxS = 40
    I32x4MaxU = 40
    I64x2Neg = 40
    I64x2AnyTrue = 20
    I64x2AllTrue = 20
    I64x2Shl = 30
    I64x2ShrS = 30
    I64x2ShrU = 30
    I64x2Add = 20
    I64x2Sub = 20
    I64x2Mul = 80
    F32x4Abs = 200
    F32x4Neg = 200
    F32x4Sqrt = 1000
    F32x4Add = 200
    F32x4Sub = 200
    F32x4Mul = 800
    F32x4Div = 1000
    F32x4Min = 500
    F32x4Max = 500
    F64x2Abs = 500
    F64x2Neg = 400
    F64x2Sqrt = 1000
    F64x2Add = 200
    F64x2Sub = 200
    F64x2Mul = 800
    F64x2Div = 1000
    F64x2Min = 500
    F64x2Max = 500
    I32x4TruncSatF32x4S = 1000
    I32x4TruncSatF32x4U = 1000
    I64x2TruncSatF64x2S = 1000
    I64x2TruncSatF64x2U = 1000
    F32x4ConvertI32x4S = 1000
    F32x4ConvertI32x4U = 1000
    F64x2ConvertI64x2S = 1000
    F64x2ConvertI64x2U = 1000
    V8x16Swizzle = 1200
    V8x16Shuffle = 1200
    V8x16LoadSplat = 40
    V16x8LoadSplat = 40
    V32x4LoadSplat = 40
    V64x2LoadSplat = 40
    I8x16NarrowI16x8S = 800
    I8x16NarrowI16x8U = 800
    I16x8NarrowI32x4S = 800
    I16x8NarrowI32x4U = 800
    I16x8WidenLowI8x16S = 800
    I16x8WidenHighI8x16S = 800
    I16x8WidenLowI8x16U = 800
    I16x8WidenHighI8x16U = 800
    I32x4WidenLowI16x8S = 800
    I32x4WidenHighI16x8S = 800
    I32x4WidenLowI16x8U = 800
    I32x4WidenHighI16x8U = 800
    I16x8Load8x8S = 400
    I16x8Load8x8U = 400
    I32x4Load16x4S = 400
    I32x4Load16x4U = 400
    I64x2Load32x2S = 400
    I64x2Load32x2U = 400
    I8x16RoundingAverageU = 200
    I16x8RoundingAverageU = 200
    LocalAllocate = 2
    LocalsUnmetered = 100
    MaxMemoryGrow = 8
    MaxMemoryGrowDelta = 10

[DynamicStorageLoad]
    QuadraticCoefficient = 688
    SignOfQuadratic = 0
    LinearCoefficient = 31858
    SignOfLinear = 0
    ConstantCoefficient = 15287
    SignOfConstant = 0
    MinimumGasCost = 10000

`
	gasScheduleV4 = `[BuiltInCost]
    ChangeOwnerAddress       = 5000000
    ClaimDeveloperRewards    = 5000000
    SaveUserName             = 1000000
    SaveKeyValue             = 100000
    ESDTTransfer             = 200000
    ESDTBurn                 = 100000
    ESDTLocalMint            = 50000
    ESDTLocalBurn            = 50000
    ESDTNFTCreate            = 150000
    ESDTNFTAddQuantity       = 50000
    ESDTNFTBurn              = 50000
    ESDTNFTTransfer          = 200000
    ESDTNFTChangeCreateOwner = 1000000
    ESDTNFTAddUri            = 50000
    ESDTNFTUpdateAttributes  = 50000
    ESDTNFTMultiTransfer     = 200000

[MetaChainSystemSCsCost]
    Stake                 = 5000000
    UnStake               = 5000000
    UnBond                = 5000000
    Claim                 = 5000000
    Get                   = 5000000
    ChangeRewardAddress   = 5000000
    ChangeValidatorKeys   = 5000000
    UnJail                = 5000000
    DelegationOps         = 1000000
    DelegationMgrOps      = 50000000
    ValidatorToDelegation = 500000000
    ESDTIssue             = 50000000
    ESDTOperations        = 50000000
    Proposal              = 50000000
    Vote                  = 50000000
    DelegateVote          = 50000000
    RevokeVote            = 50000000
    CloseProposal         = 50000000
    GetAllNodeStates      = 20000000
    UnstakeTokens         = 5000000
    UnbondTokens          = 5000000
    FixWaitingListSize    = 500000000

[BaseOperationCost]
    StorePerByte      = 10000
    ReleasePerByte    = 1000
    DataCopyPerByte   = 100
    PersistPerByte    = 1000
    CompilePerByte    = 300
    AoTPreparePerByte = 100
    GetCode           = 1000000

[ElrondAPICost]
    GetSCAddress       = 100
    GetOwnerAddress    = 5000
    IsSmartContract    = 5000
    GetShardOfAddress  = 5000
    GetExternalBalance = 7000
    GetBlockHash       = 10000
    TransferValue      = 100000
    GetArgument        = 100
    GetFunction        = 100
    GetNumArguments    = 100
    StorageStore       = 75000
    StorageLoad        = 50000
    CachedStorageLoad  = 100    
    GetCaller          = 100
    GetCallValue       = 100
    Log                = 3750
    Finish             = 1
    SignalError        = 1
    GetBlockTimeStamp  = 10000
    GetGasLeft         = 100
    Int64GetArgument   = 100
    Int64StorageStore  = 75000
    Int64StorageLoad   = 50000
    Int64Finish        = 1000
    GetStateRootHash   = 10000
    GetBlockNonce      = 10000
    GetBlockEpoch      = 10000
    GetBlockRound      = 10000
    GetBlockRandomSeed = 10000
    ExecuteOnSameContext = 100000
    ExecuteOnDestContext = 100000
    DelegateExecution    = 100000
    AsyncCallStep        = 100000
    AsyncCallbackGasLock = 4000000
    ExecuteReadOnly      = 160000
    CreateContract       = 300000
    GetReturnData        = 100
    GetNumReturnData     = 100
    GetReturnDataSize    = 100
    CleanReturnData      = 100
    DeleteFromReturnData = 100
    GetOriginalTxHash    = 10000

[EthAPICost]
    UseGas              = 100
    GetAddress          = 100000
    GetExternalBalance  = 70000
    GetBlockHash        = 100000
    Call                = 160000
    CallDataCopy        = 200
    GetCallDataSize     = 100
    CallCode            = 160000
    CallDelegate        = 160000
    CallStatic          = 160000
    StorageStore        = 250000
    StorageLoad         = 100000
    GetCaller           = 100
    GetCallValue        = 100
    CodeCopy            = 1000
    GetCodeSize         = 100
    GetBlockCoinbase    = 100
    Create              = 320000
    GetBlockDifficulty  = 100
    ExternalCodeCopy    = 3000
    GetExternalCodeSize = 2500
    GetGasLeft          = 100
    GetBlockGasLimit    = 100000
    GetTxGasPrice       = 1000
    Log                 = 3750
    GetBlockNumber      = 100000
    GetTxOrigin         = 100000
    Finish              = 1
    Revert              = 1
    GetReturnDataSize   = 200
    ReturnDataCopy      = 500
    SelfDestruct        = 5000000
    GetBlockTimeStamp   = 100000

[BigIntAPICost]
    BigIntNew                = 2000
    BigIntByteLength         = 2000
    BigIntUnsignedByteLength = 2000
    BigIntSignedByteLength   = 2000
    BigIntGetBytes           = 2000
    BigIntGetUnsignedBytes   = 2000
    BigIntGetSignedBytes     = 2000
    BigIntSetBytes           = 2000
    BigIntSetUnsignedBytes   = 2000
    BigIntSetSignedBytes     = 2000
    BigIntIsInt64            = 2000
    BigIntGetInt64           = 2000
    BigIntSetInt64           = 2000
    BigIntAdd                = 2000
    BigIntSub                = 2000
    BigIntMul                = 6000
    BigIntSqrt               = 6000
    BigIntPow                = 6000
    BigIntLog                = 6000
    BigIntTDiv               = 6000
    BigIntTMod               = 6000
    BigIntEDiv               = 6000
    BigIntEMod               = 6000
    BigIntAbs                = 2000
    BigIntNeg                = 2000
    BigIntSign               = 2000
    BigIntCmp                = 2000
    BigIntNot                = 2000
    BigIntAnd                = 2000
    BigIntOr                 = 2000
    BigIntXor                = 2000
    BigIntShr                = 2000
    BigIntShl                = 2000
    BigIntFinishUnsigned     = 1000
    BigIntFinishSigned       = 1000
    BigIntStorageLoadUnsigned   = 50000
    BigIntStorageStoreUnsigned  = 75000
    BigIntGetArgument           = 1000
    BigIntGetUnsignedArgument   = 1000
    BigIntGetSignedArgument     = 1000
    BigIntGetCallValue          = 1000
    BigIntGetExternalBalance    = 10000
    CopyPerByteForTooBig        = 1000

    [BigFloatAPICost]
    BigFloatNewFromParts = 3000
    BigFloatAdd          = 7000
    BigFloatSub          = 7000
    BigFloatMul          = 7000
    BigFloatDiv          = 7000
    BigFloatTruncate     = 5000
    BigFloatNeg          = 5000
    BigFloatClone        = 5000
    BigFloatCmp          = 4000
    BigFloatAbs          = 5000
    BigFloatSqrt         = 7000
    BigFloatPow          = 10000
    BigFloatFloor        = 5000
    BigFloatCeil         = 5000
    BigFloatIsInt        = 3000
    BigFloatSetBigInt    = 3000
    BigFloatSetInt64     = 1000
    BigFloatGetConst     = 1000
    
[CryptoAPICost]
    SHA256                 = 1000000
    Keccak256              = 1000000
    Ripemd160              = 1000000
    VerifyBLS              = 5000000
    VerifyEd25519          = 2000000
    VerifySecp256k1        = 2000000
    EllipticCurveNew       = 10000
    AddECC                 = 75000
    DoubleECC              = 65000
    IsOnCurveECC           = 10000
    ScalarMultECC          = 400000
    MarshalECC             = 13000
    MarshalCompressedECC   = 15000
    UnmarshalECC           = 20000
    UnmarshalCompressedECC = 270000
    GenerateKeyECC         = 7000000
    EncodeDERSig           = 10000000

[ManagedBufferAPICost]
    MBufferNew                   = 2000
    MBufferNewFromBytes          = 2000
    MBufferGetLength             = 2000
    MBufferGetBytes              = 2000
    MBufferGetByteSlice          = 2000
    MBufferCopyByteSlice         = 2000
    MBufferSetBytes              = 2000
    MBufferAppend                = 2000
    MBufferAppendBytes           = 2000
    MBufferToBigIntUnsigned      = 2000
    MBufferToBigIntSigned        = 5000
    MBufferFromBigIntUnsigned    = 2000
    MBufferFromBigIntSigned      = 5000
    MBufferToBigFloat            = 2000
    MBufferFromBigFloat          = 2000
    MBufferStorageStore          = 75000
    MBufferStorageLoad           = 50000
    MBufferGetArgument           = 1000
    MBufferFinish                = 1000
    MBufferSetRandom             = 6000

[WASMOpcodeCost]
    Unreachable = 5
    Nop = 5
    Block = 5
    Loop = 5
    If = 5
    Else = 5
    End = 5
    Br = 5
    BrIf = 5
    BrTable = 5
    Return = 5
    Call = 5
    CallIndirect = 5
    Drop = 5
    Select = 5
    TypedSelect = 5
    LocalGet = 5
    LocalSet = 5
    LocalTee = 5
    GlobalGet = 5
    GlobalSet = 5
    I32Load = 5
    I64Load = 5
    F32Load = 6
    F64Load = 6
    I32Load8S = 5
    I32Load8U = 5
    I32Load16S = 5
    I32Load16U = 5
    I64Load8S = 5
    I64Load8U = 5
    I64Load16S = 5
    I64Load16U = 5
    I64Load32S = 5
    I64Load32U = 5
    I32Store = 5
    I64Store = 5
    F32Store = 12
    F64Store = 12
    I32Store8 = 5
    I32Store16 = 5
    I64Store8 = 5
    I64Store16 = 5
    I64Store32 = 5
    MemorySize = 5
    MemoryGrow = 5
    I32Const = 5
    I64Const = 5
    F32Const = 5
    F64Const = 5
    RefNull = 5
    RefIsNull = 5
    RefFunc = 5
    I32Eqz = 5
    I32Eq = 5
    I32Ne = 5
    I32LtS = 5
    I32LtU = 5
    I32GtS = 5
    I32GtU = 5
    I32LeS = 5
    I32LeU = 5
    I32GeS = 5
    I32GeU = 5
    I64Eqz = 5
    I64Eq = 5
    I64Ne = 5
    I64LtS = 5
    I64LtU = 5
    I64GtS = 5
    I64GtU = 5
    I64LeS = 5
    I64LeU = 5
    I64GeS = 5
    I64GeU = 5
    F32Eq = 6
    F32Ne = 6
    F32Lt = 6
    F32Gt = 6
    F32Le = 6
    F32Ge = 6
    F64Eq = 6
    F64Ne = 6
    F64Lt = 6
    F64Gt = 6
    F64Le = 6
    F64Ge = 6
    I32Clz = 100
    I32Ctz = 100
    I32Popcnt = 100
    I32Add = 5
    I32Sub = 5
    I32Mul = 5
    I32DivS = 18
    I32DivU = 18
    I32RemS = 18
    I32RemU = 18
    I32And = 5
    I32Or = 5
    I32Xor = 5
    I32Shl = 5
    I32ShrS = 5
    I32ShrU = 5
    I32Rotl = 5
    I32Rotr = 5
    I64Clz = 100
    I64Ctz = 100
    I64Popcnt = 100
    I64Add = 5
    I64Sub = 5
    I64Mul = 5
    I64DivS = 18
    I64DivU = 18
    I64RemS = 18
    I64RemU = 18
    I64And = 5
    I64Or = 5
    I64Xor = 5
    I64Shl = 5
    I64ShrS = 5
    I64ShrU = 5
    I64Rotl = 5
    I64Rotr = 5
    F32Abs = 5
    F32Neg = 5
    F32Ceil = 100
    F32Floor = 100
    F32Trunc = 100
    F32Nearest = 100
    F32Sqrt = 100
    F32Add = 5
    F32Sub = 5
    F32Mul = 15
    F32Div = 100
    F32Min = 15
    F32Max = 15
    F32Copysign = 5
    F64Abs = 5
    F64Neg = 5
    F64Ceil = 100
    F64Floor = 100
    F64Trunc = 100
    F64Nearest = 100
    F64Sqrt = 100
    F64Add = 5
    F64Sub = 5
    F64Mul = 15
    F64Div = 100
    F64Min = 15
    F64Max = 15
    F64Copysign = 5
    I32WrapI64 = 9
    I32TruncF32S = 100
    I32TruncF32U = 100
    I32TruncF64S = 100
    I32TruncF64U = 100
    I64ExtendI32S = 9
    I64ExtendI32U = 9
    I64TruncF32S = 100
    I64TruncF32U = 100
    I64TruncF64S = 100
    I64TruncF64U = 100
    F32ConvertI32S = 100
    F32ConvertI32U = 100
    F32ConvertI64S = 100
    F32ConvertI64U = 100
    F32DemoteF64 = 100
    F64ConvertI32S = 100
    F64ConvertI32U = 100
    F64ConvertI64S = 100
    F64ConvertI64U = 100
    F64PromoteF32 = 100
    I32ReinterpretF32 = 100
    I64ReinterpretF64 = 100
    F32ReinterpretI32 = 100
    F64ReinterpretI64 = 100
    I32Extend8S = 9
    I32Extend16S = 9
    I64Extend8S = 9
    I64Extend16S = 9
    I64Extend32S = 9
    I32TruncSatF32S = 100
    I32TruncSatF32U = 100
    I32TruncSatF64S = 100
    I32TruncSatF64U = 100
    I64TruncSatF32S = 100
    I64TruncSatF32U = 100
    I64TruncSatF64S = 100
    I64TruncSatF64U = 100
    MemoryInit = 5
    DataDrop = 5
    MemoryCopy = 5
    MemoryFill = 5
    TableInit = 10
    ElemDrop = 10
    TableCopy = 10
    TableFill = 10
    TableGet = 10
    TableSet = 10
    TableGrow = 10
    TableSize = 10
    AtomicNotify = 10
    I32AtomicWait = 10
    I64AtomicWait = 10
    AtomicFence = 10
    I32AtomicLoad = 15
    I64AtomicLoad = 15
    I32AtomicLoad8U = 15
    I32AtomicLoad16U = 15
    I64AtomicLoad8U = 15
    I64AtomicLoad16U = 15
    I64AtomicLoad32U = 15
    I32AtomicStore = 15
    I64AtomicStore = 15
    I32AtomicStore8 = 15
    I32AtomicStore16 = 15
    I64AtomicStore8 = 15
    I64AtomicStore16 = 15
    I64AtomicStore32 = 15
    I32AtomicRmwAdd = 20
    I64AtomicRmwAdd = 20
    I32AtomicRmw8AddU = 20
    I32AtomicRmw16AddU = 20
    I64AtomicRmw8AddU = 20
    I64AtomicRmw16AddU = 20
    I64AtomicRmw32AddU = 20
    I32AtomicRmwSub = 20
    I64AtomicRmwSub = 20
    I32AtomicRmw8SubU = 20
    I32AtomicRmw16SubU = 20
    I64AtomicRmw8SubU = 20
    I64AtomicRmw16SubU = 20
    I64AtomicRmw32SubU = 20
    I32AtomicRmwAnd = 15
    I64AtomicRmwAnd = 15
    I32AtomicRmw8AndU = 15
    I32AtomicRmw16AndU = 15
    I64AtomicRmw8AndU = 15
    I64AtomicRmw16AndU = 15
    I64AtomicRmw32AndU = 15
    I32AtomicRmwOr = 15
    I64AtomicRmwOr = 15
    I32AtomicRmw8OrU = 15
    I32AtomicRmw16OrU = 15
    I64AtomicRmw8OrU = 15
    I64AtomicRmw16OrU = 15
    I64AtomicRmw32OrU = 15
    I32AtomicRmwXor = 15
    I64AtomicRmwXor = 15
    I32AtomicRmw8XorU = 15
    I32AtomicRmw16XorU = 15
    I64AtomicRmw8XorU = 15
    I64AtomicRmw16XorU = 15
    I64AtomicRmw32XorU = 15
    I32AtomicRmwXchg = 30
    I64AtomicRmwXchg = 30
    I32AtomicRmw8XchgU = 30
    I32AtomicRmw16XchgU = 30
    I64AtomicRmw8XchgU = 30
    I64AtomicRmw16XchgU = 30
    I64AtomicRmw32XchgU = 30
    I32AtomicRmwCmpxchg = 30
    I64AtomicRmwCmpxchg = 30
    I32AtomicRmw8CmpxchgU = 30
    I32AtomicRmw16CmpxchgU = 30
    I64AtomicRmw8CmpxchgU = 30
    I64AtomicRmw16CmpxchgU = 30
    I64AtomicRmw32CmpxchgU = 30
    V128Load = 18
    V128Store = 18
    V128Const = 18
    I8x16Splat = 20
    I8x16ExtractLaneS = 20
    I8x16ExtractLaneU = 20
    I8x16ReplaceLane = 20
    I16x8Splat = 20
    I16x8ExtractLaneS = 20
    I16x8ExtractLaneU = 20
    I16x8ReplaceLane = 20
    I32x4Splat = 20
    I32x4ExtractLane = 20
    I32x4ReplaceLane = 20
    I64x2Splat = 20
    I64x2ExtractLane = 20
    I64x2ReplaceLane = 20
    F32x4Splat = 120
    F32x4ExtractLane = 120
    F32x4ReplaceLane = 120
    F64x2Splat = 120
    F64x2ExtractLane = 120
    F64x2ReplaceLane = 120
    I8x16Eq = 30
    I8x16Ne = 30
    I8x16LtS = 40
    I8x16LtU = 40
    I8x16GtS = 40
    I8x16GtU = 40
    I8x16LeS = 40
    I8x16LeU = 40
    I8x16GeS = 40
    I8x16GeU = 40
    I16x8Eq = 30
    I16x8Ne = 30
    I16x8LtS = 40
    I16x8LtU = 40
    I16x8GtS = 40
    I16x8GtU = 40
    I16x8LeS = 40
    I16x8LeU = 40
    I16x8GeS = 40
    I16x8GeU = 40
    I32x4Eq = 30
    I32x4Ne = 30
    I32x4LtS = 40
    I32x4LtU = 40
    I32x4GtS = 40
    I32x4GtU = 40
    I32x4LeS = 40
    I32x4LeU = 40
    I32x4GeS = 40
    I32x4GeU = 40
    F32x4Eq = 120
    F32x4Ne = 120
    F32x4Lt = 120
    F32x4Gt = 120
    F32x4Le = 120
    F32x4Ge = 120
    F64x2Eq = 120
    F64x2Ne = 120
    F64x2Lt = 120
    F64x2Gt = 120
    F64x2Le = 120
    F64x2Ge = 120
    V128Not = 40
    V128And = 40
    V128AndNot = 40
    V128Or = 40
    V128Xor = 40
    V128Bitselect = 40
    I8x16Neg = 20
    I8x16AnyTrue = 20
    I8x16AllTrue = 20
    I8x16Shl = 30
    I8x16ShrS = 30
    I8x16ShrU = 30
    I8x16Add = 20
    I8x16AddSaturateS = 20
    I8x16AddSaturateU = 20
    I8x16Sub = 20
    I8x16SubSaturateS = 20
    I8x16SubSaturateU = 20
    I8x16MinS = 40
    I8x16MinU = 40
    I8x16MaxS = 40
    I8x16MaxU = 40
    I8x16Mul = 80
    I16x8Neg = 40
    I16x8AnyTrue = 40
    I16x8AllTrue = 40
    I16x8Shl = 30
    I16x8ShrS = 30
    I16x8ShrU = 30
    I16x8Add = 20
    I16x8AddSaturateS = 20
    I16x8AddSaturateU = 20
    I16x8Sub = 20
    I16x8SubSaturateS = 20
    I16x8SubSaturateU = 20
    I16x8Mul = 40
    I16x8MinS = 40
    I16x8MinU = 40
    I16x8MaxS = 40
    I16x8MaxU = 40
    I32x4Neg = 20
    I32x4AnyTrue = 20
    I32x4AllTrue = 20
    I32x4Shl = 30
    I32x4ShrS = 30
    I32x4ShrU = 30
    I32x4Add = 20
    I32x4Sub = 20
    I32x4Mul = 80
    I32x4MinS = 40
    I32x4MinU = 40
    I32x4MaxS = 40
    I32x4MaxU = 40
    I64x2Neg = 40
    I64x2AnyTrue = 20
    I64x2AllTrue = 20
    I64x2Shl = 30
    I64x2ShrS = 30
    I64x2ShrU = 30
    I64x2Add = 20
    I64x2Sub = 20
    I64x2Mul = 80
    F32x4Abs = 200
    F32x4Neg = 200
    F32x4Sqrt = 1000
    F32x4Add = 200
    F32x4Sub = 200
    F32x4Mul = 800
    F32x4Div = 1000
    F32x4Min = 500
    F32x4Max = 500
    F64x2Abs = 500
    F64x2Neg = 400
    F64x2Sqrt = 1000
    F64x2Add = 200
    F64x2Sub = 200
    F64x2Mul = 800
    F64x2Div = 1000
    F64x2Min = 500
    F64x2Max = 500
    I32x4TruncSatF32x4S = 1000
    I32x4TruncSatF32x4U = 1000
    I64x2TruncSatF64x2S = 1000
    I64x2TruncSatF64x2U = 1000
    F32x4ConvertI32x4S = 1000
    F32x4ConvertI32x4U = 1000
    F64x2ConvertI64x2S = 1000
    F64x2ConvertI64x2U = 1000
    V8x16Swizzle = 1200
    V8x16Shuffle = 1200
    V8x16LoadSplat = 40
    V16x8LoadSplat = 40
    V32x4LoadSplat = 40
    V64x2LoadSplat = 40
    I8x16NarrowI16x8S = 800
    I8x16NarrowI16x8U = 800
    I16x8NarrowI32x4S = 800
    I16x8NarrowI32x4U = 800
    I16x8WidenLowI8x16S = 800
    I16x8WidenHighI8x16S = 800
    I16x8WidenLowI8x16U = 800
    I16x8WidenHighI8x16U = 800
    I32x4WidenLowI16x8S = 800
    I32x4WidenHighI16x8S = 800
    I32x4WidenLowI16x8U = 800
    I32x4WidenHighI16x8U = 800
    I16x8Load8x8S = 400
    I16x8Load8x8U = 400
    I32x4Load16x4S = 400
    I32x4Load16x4U = 400
    I64x2Load32x2S = 400
    I64x2Load32x2U = 400
    I8x16RoundingAverageU = 200
    I16x8RoundingAverageU = 200
    LocalAllocate = 5
    LocalsUnmetered = 100
    MaxMemoryGrowDelta = 1
    MaxMemoryGrow = 100

[DynamicStorageLoad]
    QuadraticCoefficient = 688
    SignOfQuadratic = 0
    LinearCoefficient = 31858
    SignOfLinear = 0
    ConstantCoefficient = 15287
    SignOfConstant = 0
    MinimumGasCost = 10000
`
)
