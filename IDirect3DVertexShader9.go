package d3d9

/*
#include "d3d9wrapper.h"

HRESULT IDirect3DVertexShader9GetDevice(
		IDirect3DVertexShader9* obj,
		IDirect3DDevice9** ppDevice) {
	return obj->lpVtbl->GetDevice(obj, ppDevice);
}

HRESULT IDirect3DVertexShader9GetFunction(
		IDirect3DVertexShader9* obj,
		void* pData,
		UINT* pSizeOfData) {
	return obj->lpVtbl->GetFunction(obj, pData, pSizeOfData);
}

void IDirect3DVertexShader9Release(IDirect3DVertexShader9* obj) {
	obj->lpVtbl->Release(obj);
}
*/
import "C"
import "unsafe"

type VertexShader struct {
	handle *C.IDirect3DVertexShader9
}

func (obj VertexShader) Release() {
	C.IDirect3DVertexShader9Release(obj.handle)
}

func (obj VertexShader) GetDevice() (ppDevice Device, err error) {
	var c_ppDevice *C.IDirect3DDevice9
	err = toErr(C.IDirect3DVertexShader9GetDevice(obj.handle, &c_ppDevice))
	ppDevice = Device{c_ppDevice}
	return
}

// GetFunction returns the shader data.
func (obj VertexShader) GetFunction() (pData []byte, err error) {
	// first get the needed buffer size, pass nil as the data pointer
	var c_pSizeOfDataInBytes C.UINT
	err = toErr(C.IDirect3DVertexShader9GetFunction(
		obj.handle,
		nil,
		&c_pSizeOfDataInBytes,
	))
	if err != nil {
		return
	}
	pData = make([]byte, c_pSizeOfDataInBytes)
	err = toErr(C.IDirect3DVertexShader9GetFunction(
		obj.handle,
		unsafe.Pointer(&pData[0]),
		&c_pSizeOfDataInBytes,
	))
	return
}
