package wasmer2

// #cgo LDFLAGS: -Wl,-rpath,${SRCDIR} -L${SRCDIR}
// #cgo linux,amd64 LDFLAGS:-lvmexeccapi
// #include "./libvmexeccapi.h"
//
import "C"
import (
	"fmt"
	"unsafe"
)

type cBool C.bool
type cChar C.char
type cInt C.int
type cUchar C.uchar
type cUint C.uint
type cUint32T C.uint32_t
type cUint8T C.uint8_t

type cWasmerByteArray C.vm_exec_byte_array

// type cWasmerExportFuncT C.vm_exec_export_func_t
// type cWasmerExportT C.vm_exec_export_t
// type cWasmerExportsT C.vm_exec_exports_t
// type cWasmerImportExportKind C.vm_exec_import_export_kind
// type cWasmerImportExportValue C.vm_exec_import_export_value
type cWasmerImportFuncT C.vm_exec_import_func_t
type cWasmerImportT C.vm_exec_import_t

// type cWasmerInstanceContextT C.vm_exec_instance_context_t
type cWasmerInstanceT C.vm_exec_instance_t

// type cWasmerMemoryT C.vm_exec_memory_t
type cWasmerResultT C.vm_exec_result_t

// type cWasmerValueT C.vm_exec_value_t
type cWasmerValueTag C.vm_exec_value_tag

type cWasmerCompilationOptions C.vm_exec_compilation_options_t

// const cWasmFunction = C.WASM_FUNCTION
// const cWasmGlobal = C.WASM_GLOBAL
const cWasmI32 = C.VM_EXEC_VALUE_I32
const cWasmI64 = C.VM_EXEC_VALUE_I64

// const cWasmMemory = C.WASM_MEMORY
// const cWasmTable = C.WASM_TABLE
const cWasmerOk = C.VM_EXEC_OK

func cNewWasmerImportT(moduleName string, importName string, function *cWasmerImportFuncT) cWasmerImportT {
	var importedFunction C.vm_exec_import_t
	importedFunction.module_name = (C.vm_exec_byte_array)(cGoStringToWasmerByteArray(moduleName))
	importedFunction.import_name = (C.vm_exec_byte_array)(cGoStringToWasmerByteArray(importName))
	// importedFunction.tag = cWasmFunction

	var pointer = (**C.vm_exec_import_func_t)(unsafe.Pointer(&importedFunction.import_func))
	*pointer = (*C.vm_exec_import_func_t)(function)

	return (cWasmerImportT)(importedFunction)
}

// func cWasmerInstanceGetPointsUsed(instance *cWasmerInstanceT) uint64 {
// 	return uint64(C.vm_exec_instance_get_points_used(
// 		(*C.vm_exec_instance_t)(instance),
// 	))
// }

// func cWasmerInstanceSetPointsUsed(instance *cWasmerInstanceT, points uint64) {
// 	C.vm_exec_instance_set_points_used(
// 		(*C.vm_exec_instance_t)(instance),
// 		(C.uint64_t)(points),
// 	)
// }

// func cWasmerInstanceSetGasLimit(instance *cWasmerInstanceT, gasLimit uint64) {
// 	C.vm_exec_instance_set_points_limit(
// 		(*C.vm_exec_instance_t)(instance),
// 		(C.uint64_t)(gasLimit),
// 	)
// }

// func cWasmerInstanceSetBreakpointValue(instance *cWasmerInstanceT, value uint64) {
// 	C.vm_exec_instance_set_runtime_breakpoint_value(
// 		(*C.vm_exec_instance_t)(instance),
// 		(C.uint64_t)(value),
// 	)
// }

// func cWasmerInstanceGetBreakpointValue(instance *cWasmerInstanceT) uint64 {
// 	return uint64(C.vm_exec_instance_get_runtime_breakpoint_value(
// 		(*C.vm_exec_instance_t)(instance),
// 	))
// }

// func cWasmerInstanceIsFunctionImported(instance *cWasmerInstanceT, name string) bool {
// 	var functionName = cCString(name)
// 	return bool(C.vm_exec_instance_is_function_imported(
// 		(*C.vm_exec_instance_t)(instance),
// 		(*C.char)(unsafe.Pointer(functionName)),
// 	))
// }

// func cWasmerInstanceEnableRkyv() {
// 	C.vm_exec_instance_enable_rkyv()
// }

// func cWasmerInstanceDisableRkyv() {
// 	C.vm_exec_instance_disable_rkyv()
// }

func cWasmerSetSIGSEGVPassthrough() {
	fmt.Println("Go: calling wasmer_set_sigsegv_passthrough")
	C.vm_exec_set_sigsegv_passthrough()

	// _ = C.do_something((C.uint32_t)(5))
	// fmt.Println(result)
}

// func cWasmerInstanceCache(
// 	instance *cWasmerInstanceT,
// 	cacheBytes **cUchar,
// 	cacheLen *cUint32T,
// ) cWasmerResultT {
// 	return (cWasmerResultT)(C.vm_exec_instance_cache(
// 		(*C.vm_exec_instance_t)(instance),
// 		(**C.uchar)(unsafe.Pointer(cacheBytes)),
// 		(*C.uint32_t)(cacheLen),
// 	))
// }

// func cWasmerInstanceFromCache(
// 	instance **cWasmerInstanceT,
// 	cacheBytes *cUchar,
// 	cacheLen cUint32T,
// 	options *cWasmerCompilationOptions,
// ) cWasmerResultT {
// 	return (cWasmerResultT)(C.vm_exec_instance_from_cache(
// 		(**C.vm_exec_instance_t)(unsafe.Pointer(instance)),
// 		(*C.uchar)(cacheBytes),
// 		(C.uint32_t)(cacheLen),
// 		(*C.vm_exec_compilation_options_t)(options),
// 	))
// }

func cWasmerSetImports(
	imports *cWasmerImportT,
	importsLength cInt,
) cWasmerResultT {
	return (cWasmerResultT)(C.vm_exec_set_imports(
		(*C.vm_exec_import_t)(imports),
		(C.uint32_t)(importsLength),
	))
}

// func cWasmerSetOpcodeCosts(opcode_costs *[OPCODE_COUNT]uint32) {
// 	C.vm_exec_set_opcode_costs(
// 		(*C.uint32_t)(unsafe.Pointer(opcode_costs)),
// 	)
// }

// func cWasmerExportFuncParams(function *cWasmerExportFuncT, parameters *cWasmerValueTag, parametersLength cUint32T) cWasmerResultT {
// 	return (cWasmerResultT)(C.vm_exec_export_func_params(
// 		(*C.vm_exec_export_func_t)(function),
// 		(*C.vm_exec_value_tag)(parameters),
// 		(C.uint32_t)(parametersLength),
// 	))
// }

// func cWasmerExportFuncParamsArity(function *cWasmerExportFuncT, result *cUint32T) cWasmerResultT {
// 	return (cWasmerResultT)(C.vm_exec_export_func_params_arity(
// 		(*C.vm_exec_export_func_t)(function),
// 		(*C.uint32_t)(result),
// 	))
// }

// func cWasmerExportFuncResultsArity(function *cWasmerExportFuncT, result *cUint32T) cWasmerResultT {
// 	return (cWasmerResultT)(C.vm_exec_export_func_returns_arity(
// 		(*C.vm_exec_export_func_t)(function),
// 		(*C.uint32_t)(result),
// 	))
// }

// func cWasmerExportKind(export *cWasmerExportT) cWasmerImportExportKind {
// 	return (cWasmerImportExportKind)(C.vm_exec_export_kind(
// 		(*C.vm_exec_export_t)(export),
// 	))
// }

// func cWasmerExportName(export *cWasmerExportT) cWasmerByteArray {
// 	return (cWasmerByteArray)(C.vm_exec_export_name(
// 		(*C.vm_exec_export_t)(export),
// 	))
// }

// func cWasmerExportToFunc(export *cWasmerExportT) *cWasmerExportFuncT {
// 	return (*cWasmerExportFuncT)(C.vm_exec_export_to_func(
// 		(*C.vm_exec_export_t)(export),
// 	))
// }

// func cWasmerExportToMemory(export *cWasmerExportT, memory **cWasmerMemoryT) cWasmerResultT {
// 	return (cWasmerResultT)(C.vm_exec_export_to_memory(
// 		(*C.vm_exec_export_t)(export),
// 		(**C.vm_exec_memory_t)(unsafe.Pointer(memory)),
// 	))
// }

// func cWasmerExportsDestroy(exports *cWasmerExportsT) {
// 	C.vm_exec_exports_destroy(
// 		(*C.vm_exec_exports_t)(exports),
// 	)
// }

// func cWasmerExportsGet(exports *cWasmerExportsT, index cInt) *cWasmerExportT {
// 	return (*cWasmerExportT)(C.vm_exec_exports_get(
// 		(*C.vm_exec_exports_t)(exports),
// 		(C.int)(index),
// 	))
// }

// func cWasmerExportsLen(exports *cWasmerExportsT) cInt {
// 	return (cInt)(C.vm_exec_exports_len(
// 		(*C.vm_exec_exports_t)(exports),
// 	))
// }

func cWasmerImportFuncDestroy(function *cWasmerImportFuncT) {
	C.vm_exec_import_func_destroy(
		(*C.vm_exec_import_func_t)(function),
	)
}

func cWasmerImportFuncNew(
	function unsafe.Pointer,
	parametersSignature *cWasmerValueTag,
	parametersLength cUint,
	resultsSignature *cWasmerValueTag,
	resultsLength cUint,
) *cWasmerImportFuncT {
	return (*cWasmerImportFuncT)(C.vm_exec_import_func_new(
		(*[0]byte)(function),
		(*C.vm_exec_value_tag)(parametersSignature),
		(C.uint)(parametersLength),
		(*C.vm_exec_value_tag)(resultsSignature),
		(C.uint)(resultsLength),
	))
}

func cWasmerInstantiateWithOptions(
	instance **cWasmerInstanceT,
	wasmBytes *cUchar,
	wasmBytesLength cUint,
	options *cWasmerCompilationOptions,
) cWasmerResultT {
	return (cWasmerResultT)(C.vm_exec_new_instance(
		(**C.vm_exec_instance_t)(unsafe.Pointer(instance)),
		(*C.uchar)(wasmBytes),
		(C.uint)(wasmBytesLength),
		(*C.vm_exec_compilation_options_t)(options),
	))
}

func cWasmerInstanceCall(
	instance *cWasmerInstanceT,
	name *cChar,
) cWasmerResultT {
	return (cWasmerResultT)(C.vm_exec_instance_call(
		(*C.vm_exec_instance_t)(instance),
		(*C.char)(name),
	))
}

func cWasmerInstanceHasFunction(
	instance *cWasmerInstanceT,
	name *cChar,
) cInt {
	return (cInt)(C.vm_exec_instance_has_function(
		(*C.vm_exec_instance_t)(instance),
		(*C.char)(name),
	))
}

func cWasmerInstanceExportedFunctionNamesLength(instance *cWasmerInstanceT) cInt {
	return (cInt)(C.vm_exported_function_names_length(
		(*C.vm_exec_instance_t)(instance),
	))
}

func cWasmerInstanceExportedFunctionNames(instance *cWasmerInstanceT, buffer *cChar, length cInt) cInt {
	return (cInt)(C.vm_exported_function_names(
		(*C.vm_exec_instance_t)(instance),
		(*C.char)(buffer),
		(C.int)(length),
	))
}

// func cWasmerInstanceContextGet(instance *cWasmerInstanceT) *cWasmerInstanceContextT {
// 	return (*cWasmerInstanceContextT)(C.vm_exec_instance_context_get(
// 		(*C.vm_exec_instance_t)(instance),
// 	))
// }

// func cWasmerInstanceContextDataGet(instanceContext *cWasmerInstanceContextT) unsafe.Pointer {
// 	return unsafe.Pointer(C.vm_exec_instance_context_data_get(
// 		(*C.vm_exec_instance_context_t)(instanceContext),
// 	))
// }

// func cWasmerInstanceContextDataSet(instance *cWasmerInstanceT, dataPointer unsafe.Pointer) {
// 	C.vm_exec_instance_context_data_set(
// 		(*C.vm_exec_instance_t)(instance),
// 		dataPointer,
// 	)
// }

// func cWasmerInstanceContextMemory(instanceContext *cWasmerInstanceContextT) *cWasmerMemoryT {
// 	return (*cWasmerMemoryT)(C.vm_exec_instance_context_memory(
// 		(*C.vm_exec_instance_context_t)(instanceContext),
// 		0,
// 	))
// }

func cWasmerInstanceDestroy(instance *cWasmerInstanceT) {
	C.vm_exec_instance_destroy(
		(*C.vm_exec_instance_t)(instance),
	)
}

func cWasmerLastErrorLength() cInt {
	return (cInt)(C.vm_exec_last_error_length())
}

func cWasmerLastErrorMessage(buffer *cChar, length cInt) cInt {
	return (cInt)(C.vm_exec_last_error_message(
		(*C.char)(buffer),
		(C.int)(length),
	))
}

func cWasmerExcutionInfoLength() cInt {
	return (cInt)(C.vm_exec_execution_info_length())
}

func cWasmerExcutionInfoFlush(buffer *cChar, length cInt) cInt {
	return (cInt)(C.vm_exec_execution_info_flush(
		(*C.char)(buffer),
		(C.int)(length),
	))
}

// func cWasmerMemoryData(memory *cWasmerMemoryT) *cUint8T {
// 	return (*cUint8T)(C.vm_exec_memory_data(
// 		(*C.vm_exec_memory_t)(memory),
// 	))
// }

// func cWasmerMemoryDataLength(memory *cWasmerMemoryT) cUint32T {
// 	return (cUint32T)(C.vm_exec_memory_data_length(
// 		(*C.vm_exec_memory_t)(memory),
// 	))
// }

// func cWasmerMemoryGrow(memory *cWasmerMemoryT, numberOfPages cUint32T) cWasmerResultT {
// 	return (cWasmerResultT)(C.vm_exec_memory_grow(
// 		(*C.vm_exec_memory_t)(memory),
// 		(C.uint32_t)(numberOfPages),
// 	))
// }

// func cWasmerMemoryDestroy(memory *cWasmerMemoryT) {
// 	C.vm_exec_memory_destroy(
// 		(*C.vm_exec_memory_t)(memory),
// 	)
// }

func cCString(string string) *cChar {
	return (*cChar)(C.CString(string))
}

func cFree(pointer unsafe.Pointer) {
	C.free(pointer)
}

func cGoString(string *cChar) string {
	return C.GoString((*C.char)(string))
}

func cGoStringN(string *cChar, length cInt) string {
	return C.GoStringN((*C.char)(string), (C.int)(length))
}

func cGoStringToWasmerByteArray(string string) cWasmerByteArray {
	var cString = cCString(string)

	var byteArray cWasmerByteArray
	byteArray.bytes = (*C.uchar)(unsafe.Pointer(cString))
	byteArray.bytes_len = (C.uint)(len(string))

	return byteArray
}
