// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package streams

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go/windows/foundation"
)

const GUIDIOutputStream string = "905a0fe6-bc53-11df-8c49-001e4fc686da"
const SignatureIOutputStream string = "{905a0fe6-bc53-11df-8c49-001e4fc686da}"

type IOutputStream struct {
	ole.IInspectable
}

type IOutputStreamVtbl struct {
	ole.IInspectableVtbl

	WriteAsync uintptr
	FlushAsync uintptr
}

func (v *IOutputStream) VTable() *IOutputStreamVtbl {
	return (*IOutputStreamVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IOutputStream) WriteAsync(buffer *IBuffer) (*foundation.IAsyncOperationWithProgress, error) {
	var out *foundation.IAsyncOperationWithProgress
	hr, _, _ := syscall.SyscallN(
		v.VTable().WriteAsync,
		uintptr(unsafe.Pointer(v)),      // this
		uintptr(unsafe.Pointer(buffer)), // in IBuffer
		uintptr(unsafe.Pointer(&out)),   // out foundation.IAsyncOperationWithProgress
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *IOutputStream) FlushAsync() (*foundation.IAsyncOperation, error) {
	var out *foundation.IAsyncOperation
	hr, _, _ := syscall.SyscallN(
		v.VTable().FlushAsync,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out foundation.IAsyncOperation
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}
