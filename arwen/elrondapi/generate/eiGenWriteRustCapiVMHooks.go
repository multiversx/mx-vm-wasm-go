package elrondapigenerate

import (
	"fmt"
)

func WriteRustCapiVMHooks(out *eiGenWriter, eiMetadata *EIMetadata) {
	autoGeneratedHeader(out)
	out.WriteString(`
use std::ffi::c_void;

use crate::capi_vm_hook_pointers::vm_exec_vm_hook_c_func_pointers;

#[derive(Debug)]
pub struct CapiVMHooks {
    pub vm_hooks_ptr: *mut c_void,
    pub c_func_pointers_ptr: vm_exec_vm_hook_c_func_pointers,
}

impl CapiVMHooks {
    pub fn new(c_func_pointers_ptr: vm_exec_vm_hook_c_func_pointers) -> Self {
        Self {
            vm_hooks_ptr: std::ptr::null_mut(),
            c_func_pointers_ptr,
        }
    }
}

#[rustfmt::skip]
impl elrond_exec_service::VMHooks for CapiVMHooks {
    fn set_vm_hooks_ptr(&mut self, vm_hooks_ptr: *mut c_void) {
        self.vm_hooks_ptr = vm_hooks_ptr;
    }
`)

	for _, funcMetadata := range eiMetadata.AllFunctions {
		out.WriteString(fmt.Sprintf(
			"\n    fn %s%s {\n",
			snakeCase(funcMetadata.Name),
			writeRustFnDeclarationArguments(
				"&self",
				funcMetadata,
			),
		))

		out.WriteString(fmt.Sprintf(
			"        (self.c_func_pointers_ptr.%s)%s\n",
			cgoFuncPointerFieldName(funcMetadata),
			writeRustFnCallArguments("self.vm_hooks_ptr", funcMetadata),
		))

		out.WriteString("    }\n")
	}

	out.WriteString(`}
`)
}