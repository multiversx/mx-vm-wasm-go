package elrondapigenerate

import (
	"fmt"
	"os"
)

func WriteEIInterface(eiMetadata *EIMetadata, out *os.File) {
	out.WriteString("package executor \n\n")
	out.WriteString("// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n")
	out.WriteString("// !!!!!!!!!!!!!!!!!!!!!! AUTO-GENERATED FILE !!!!!!!!!!!!!!!!!!!!!!\n")
	out.WriteString("// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n")
	out.WriteString("\n")
	out.WriteString("type ImportsInterface interface {\n")

	for _, funcMetadata := range eiMetadata.AllFunctions {
		out.WriteString(fmt.Sprintf("\t%s(", funcMetadata.CapitalizedName))
		for argIndex, arg := range funcMetadata.Arguments {
			if argIndex > 0 {
				out.WriteString(", ")
			}
			out.WriteString(fmt.Sprintf("%s %s", arg.Name, arg.Type))
		}
		out.WriteString(")")
		if funcMetadata.Result != nil {
			out.WriteString(fmt.Sprintf(" %s", funcMetadata.Result.Type))
		}

		out.WriteString("\n")
	}

	out.WriteString("}\n")
}

func externResult(functResult *EIFunctionResult) string {
	if functResult == nil {
		return "void"
	}
	return cgoType(functResult.Type)
}

func cgoType(goType string) string {
	if goType == "int32" {
		return "int32_t"
	}
	if goType == "int64" {
		return "long long"
	}
	return goType
}

func cgoFuncName(funcMetadata *EIFunction) string {
	return fmt.Sprintf("wasmer1_%s", funcMetadata.LowerCaseName)
}

func cgoImportName(funcMetadata *EIFunction) string {
	return fmt.Sprintf("C.%s", cgoFuncName(funcMetadata))
}

func WriteCAPIFunctions(eiMetadata *EIMetadata, out *os.File) {
	writeHeader(out, eiMetadata)
	writeImports(out, eiMetadata)
}

func writeHeader(out *os.File, eiMetadata *EIMetadata) {
	out.WriteString("package wasmer\n\n")
	out.WriteString("// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n")
	out.WriteString("// !!!!!!!!!!!!!!!!!!!!!! AUTO-GENERATED FILE !!!!!!!!!!!!!!!!!!!!!!\n")
	out.WriteString("// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n")
	out.WriteString("\n")

	out.WriteString("// // Declare the function signatures (see [cgo](https://golang.org/cmd/cgo/)).\n")
	out.WriteString("//\n")
	out.WriteString("// #include <stdlib.h>\n")
	out.WriteString("// typedef int int32_t;\n")
	out.WriteString("//\n")

	for _, funcMetadata := range eiMetadata.AllFunctions {
		out.WriteString(fmt.Sprintf("// extern %-9s %s(void* context",
			externResult(funcMetadata.Result),
			cgoFuncName(funcMetadata),
		))
		for _, arg := range funcMetadata.Arguments {
			out.WriteString(fmt.Sprintf(", %s %s", cgoType(arg.Type), arg.Name))
		}
		out.WriteString(");\n")
	}

	out.WriteString("import \"C\"\n\n")
	out.WriteString("import (\n")
	out.WriteString("\t\"unsafe\"\n\n")
	out.WriteString("\t\"github.com/ElrondNetwork/wasm-vm/executor\"\n")
	out.WriteString(")\n\n")
}

func writeImports(out *os.File, eiMetadata *EIMetadata) {
	writeImportsHeader(out, eiMetadata)
	writeImportsBody(eiMetadata, out)
}

func writeImportsHeader(out *os.File, eiMetadata *EIMetadata) {
	out.WriteString("// ElrondEIImports populates imports with the ElrondEI API methods\n")
	out.WriteString("func ElrondEIImports(imports executor.ImportFunctionReceiver) error {\n")
	out.WriteString("\timports.Namespace(\"env\")\n\n")
	out.WriteString("\tvar err error\n")
	for _, funcMetadata := range eiMetadata.AllFunctions {
		out.WriteString(fmt.Sprintf("\terr = imports.Append(\"%s\", %s, %s)\n",
			funcMetadata.LowerCaseName,
			cgoFuncName(funcMetadata),
			cgoImportName(funcMetadata),
		))
		out.WriteString("\tif err != nil {\n")
		out.WriteString("\t\treturn err\n")
		out.WriteString("\t}\n\n")
	}
	out.WriteString("\treturn nil\n")
	out.WriteString("}\n")
}

func writeImportsBody(eiMetadata *EIMetadata, out *os.File) {
	for _, funcMetadata := range eiMetadata.AllFunctions {
		out.WriteString(fmt.Sprintf("\n// export %s\n",
			cgoFuncName(funcMetadata),
		))
		out.WriteString(fmt.Sprintf("func %s(context unsafe.Pointer",
			cgoFuncName(funcMetadata),
		))
		for _, arg := range funcMetadata.Arguments {
			out.WriteString(fmt.Sprintf(", %s %s", arg.Name, arg.Type))
		}
		out.WriteString(")")
		if funcMetadata.Result != nil {
			out.WriteString(fmt.Sprintf(" %s", funcMetadata.Result.Type))
		}
		out.WriteString(" {\n")
		out.WriteString("\tcallbacks := importsInterfaceFromRaw(context)\n")
		out.WriteString("\t")
		if funcMetadata.Result != nil {
			out.WriteString("return ")
		}
		out.WriteString(fmt.Sprintf("callbacks.%s(",
			funcMetadata.CapitalizedName,
		))
		for argIndex, arg := range funcMetadata.Arguments {
			if argIndex > 0 {
				out.WriteString(", ")
			}
			out.WriteString(arg.Name)
		}
		out.WriteString(")\n")

		out.WriteString("}\n")
	}
}