package subcontexts

import (
	"unsafe"

	arwen "github.com/ElrondNetwork/arwen-wasm-vm/arwen"
	context "github.com/ElrondNetwork/arwen-wasm-vm/arwen/context"
	"github.com/ElrondNetwork/arwen-wasm-vm/wasmer"
	vmcommon "github.com/ElrondNetwork/elrond-vm-common"
)

type Runtime struct {
	host            arwen.VMContext
	blockChainHook  vmcommon.BlockchainHook
	instance        *wasmer.Instance
	instanceContext *wasmer.InstanceContext
	vmInput         *vmcommon.VMInput
	scAddress       []byte
	callFunction    string
	vmType          []byte
	readOnly        bool
	stateStack      []*Runtime
}

func NewRuntimeSubcontext(host arwen.VMContext, blockChainHook vmcommon.BlockchainHook) (*Runtime, error) {
	runtime := &Runtime{
		host:            host,
		blockChainHook:  blockChainHook,
		instance:        nil,
		instanceContext: nil,
		vmInput:         nil,
		stateStack:      make([]*Runtime, 0),
	}
	return runtime, nil
}

func (runtime *Runtime) InitializeFromInput(input *vmcommon.ContractCallInput) error {
	runtime.vmInput = &input.VMInput
	runtime.scAddress = input.RecipientAddr
	runtime.callFunction = input.Function

	contract, err := runtime.blockChainHook.GetCode(runtime.scAddress)
	gasProvided := runtime.vmInput.GasProvided
	instance, err := wasmer.NewMeteredInstance(contract, gasProvided)
	if err != nil {
		return err
	}
	runtime.instance = instance
	return nil
}

func (runtime *Runtime) PushState() {
	newState := &Runtime{
		vmInput:      runtime.vmInput,
		scAddress:    runtime.scAddress,
		callFunction: runtime.callFunction,
		readOnly:     runtime.readOnly,
	}

	runtime.stateStack = append(runtime.stateStack, newState)
}

func (runtime *Runtime) PopState() error {
	stateStackLen := len(runtime.stateStack)
	if stateStackLen < 1 {
		return context.StateStackUnderflow
	}

	prevState := runtime.stateStack[stateStackLen-1]
	runtime.stateStack = runtime.stateStack[:stateStackLen-1]

	runtime.vmInput = prevState.vmInput
	runtime.scAddress = prevState.scAddress
	runtime.callFunction = prevState.callFunction
	runtime.readOnly = prevState.readOnly

	return nil
}

func (runtime *Runtime) LoadFromStateCopy(otherRuntime *Runtime) {
	runtime.vmInput = otherRuntime.vmInput
	runtime.scAddress = otherRuntime.scAddress
	runtime.callFunction = otherRuntime.callFunction
	runtime.readOnly = otherRuntime.readOnly
}

func (runtime *Runtime) GetVMInput() *vmcommon.VMInput {
	return runtime.vmInput
}

func (runtime *Runtime) GetSCAddress() []byte {
	return runtime.scAddress
}

func (runtime *Runtime) Function() string {
	return runtime.callFunction
}

func (runtime *Runtime) Arguments() [][]byte {
	return runtime.vmInput.Arguments
}

func (runtime *Runtime) SignalUserError() {
	// SignalUserError() remains in Runtime, and won't be moved into Output,
	// because there will be extra handling added here later, which requires
	// information from Runtime (e.g. runtime breakpoints)
	runtime.host.Output().SetReturnCode(vmcommon.UserError)
}

func (runtime *Runtime) SetRuntimeBreakpointValue(value arwen.BreakpointValue) {
	runtime.instance.SetBreakpointValue(uint64(value))
}

func (runtime *Runtime) GetRuntimeBreakpointValue() arwen.BreakpointValue {
	return arwen.BreakpointValue(runtime.instance.GetBreakpointValue())
}

func (runtime *Runtime) GetPointsUsed() uint64 {
	return runtime.instance.GetPointsUsed()
}

func (runtime *Runtime) SetPointsUsed(gasPoints uint64) {
	runtime.instance.SetPointsUsed(gasPoints)
}

func (runtime *Runtime) CallData() []byte {
	panic("not implemented")
}

func (runtime *Runtime) ReadOnly() bool {
	return runtime.readOnly
}

func (runtime *Runtime) SetReadOnly(readOnly bool) {
	runtime.readOnly = readOnly
}

func (runtime *Runtime) ExecuteOnSameContext(input *vmcommon.ContractCallInput) error {
	panic("not implemented")
}

func (runtime *Runtime) ExecuteOnDestContext(input *vmcommon.ContractCallInput) error {
	panic("not implemented")
}

func (runtime *Runtime) SetInstanceContextId(id int) {
	runtime.instance.SetContextData(unsafe.Pointer(&id))
}

func (runtime *Runtime) SetInstanceContext(instCtx *wasmer.InstanceContext) {
	runtime.instanceContext = instCtx
}

func (runtime *Runtime) GetInstanceContext() *wasmer.InstanceContext {
	return runtime.instanceContext
}

func (runtime *Runtime) MemStore(offset int32, data []byte) error {
	memory := runtime.instanceContext.Memory()
	return arwen.StoreBytes(memory, offset, data)
}

func (runtime *Runtime) MemLoad(offset int32, length int32) ([]byte, error) {
	memory := runtime.instanceContext.Memory()
	return arwen.LoadBytes(memory, offset, length)
}

func (runtime *Runtime) Clean() {
	runtime.instance.Clean()
}
