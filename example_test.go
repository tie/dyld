package dyld_test

import (
	"fmt"

	"github.com/tie/dyld"
)

func ExampleIsLoadable_dyld() {
	// dyld is the Mach-O dynamic loader executable.
	// We can’t load it because it has wrong DYLINKER file type.
	ok, err := dyld.IsLoadable("dyld")
	fmt.Println(!ok && err != nil)
	// Output: true
}

func ExampleOpen_scopeLocal() {
	lib, err := dyld.Open("libffi.dylib", dyld.ScopeLocal)
	if err != nil {
		fmt.Printf("open: %v", err)
		return
	}

	sym, err := lib.Lookup("ffi_call")
	if err != nil {
		fmt.Printf("lookup: %v", err)
		return
	}

	// The symbols from dylib are not globally visible.
	if sym, err := dyld.Lookup(sym.Name); err == nil {
		fmt.Printf("Oops, symbol %s is global!", sym.Name)
		return
	}

	fmt.Println("Found", sym.Name)
	// Output:
	// Found ffi_call
}

func ExampleLookup_scopeGlobal() {
	sym, err := dyld.Lookup("dlopen")
	if err != nil {
		fmt.Printf("lookup: %v", err)
		return
	}

	fmt.Println("Found", sym.Name)
	// Output:
	// Found dlopen
}

func ExampleAddr_lookup() {
	sym, err := dyld.Lookup("dlopen")
	if err != nil {
		fmt.Printf("lookup: %v", err)
		return
	}

	info, err := dyld.Addr(sym.Addr)
	if err != nil {
		fmt.Printf("addr: %v", err)
		return
	}

	fmt.Println("Found", info.Sname, "in", info.Fname)
	// Output:
	// Found dlopen in /usr/lib/system/libdyld.dylib
}
