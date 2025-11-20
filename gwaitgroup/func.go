// Package gwaitgroup
//
// ----------------develop info----------------
//
//	@Author Jimmy
//	@DateTime 2025-11-5 11:57
//
// --------------------------------------------
package gwaitgroup

// GoInt Extended usage of the "go" function, with parameter "Content" and "Int"
func (wg *WaitGroup) GoInt(f func(index int), index int) {
	wg.Go(func() {
		f(index)
	})
}

// GoInt64 Extended usage of the "go" function, with parameter "Content" and "Int64"
func (wg *WaitGroup) GoInt64(f func(index int64), index int64) {
	wg.Go(func() {
		f(index)
	})
}

// GoInt32 Extended usage of the "go" function, with parameter "Content" and "Int32"
func (wg *WaitGroup) GoInt32(f func(index int32), index int32) {
	wg.Go(func() {
		f(index)
	})
}

// GoInt16 Extended usage of the "go" function, with parameter "Content" and "Int16"
func (wg *WaitGroup) GoInt16(f func(index int16), index int16) {
	wg.Go(func() {
		f(index)
	})
}

// GoInt8 Extended usage of the "go" function, with parameter "Content" and "Int8"
func (wg *WaitGroup) GoInt8(f func(index int8), index int8) {
	wg.Go(func() {
		f(index)
	})
}

// GoUInt Extended usage of the "go" function, with parameter "Content" and "UInt"
func (wg *WaitGroup) GoUInt(f func(index uint), index uint) {
	wg.Go(func() {
		f(index)
	})
}

// GoUInt64 Extended usage of the "go" function, with parameter "Content" and "UInt64"
func (wg *WaitGroup) GoUInt64(f func(index uint64), index uint64) {
	wg.Go(func() {
		f(index)
	})
}

// GoUInt32 Extended usage of the "go" function, with parameter "Content" and "UInt32"
func (wg *WaitGroup) GoUInt32(f func(index uint32), index uint32) {
	wg.Go(func() {
		f(index)
	})
}

// GoUInt16 Extended usage of the "go" function, with parameter "Content" and "UInt16"
func (wg *WaitGroup) GoUInt16(f func(index uint16), index uint16) {
	wg.Go(func() {
		f(index)
	})
}

// GoUInt8 Extended usage of the "go" function, with parameter "Content" and "UInt8"
func (wg *WaitGroup) GoUInt8(f func(index uint8), index uint8) {
	wg.Go(func() {
		f(index)
	})
}

// GoString Extended usage of the "go" function, with parameter "Content" and "String"
func (wg *WaitGroup) GoString(f func(str string), str string) {
	wg.Go(func() {
		f(str)
	})
}

// GoBytes Extended usage of the "go" function, with parameter "Content" and "Bytes"
func (wg *WaitGroup) GoBytes(f func(byteSlice []byte), byteSlice []byte) {
	wg.Go(func() {
		f(byteSlice)
	})
}
