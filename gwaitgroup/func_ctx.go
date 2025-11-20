// Package gwaitgroup
//
// ----------------develop info----------------
//
//	@Author Jimmy
//	@DateTime 2025-11-5 11:57
//
// --------------------------------------------
package gwaitgroup

import "context"

// GoCtx Extended usage of the "go" function, with parameter "Content"
func (wg *WaitGroup) GoCtx(f func(ctx context.Context), ctx context.Context) {
	wg.Go(func() {
		f(ctx)
	})
}

// GoCtxInt Extended usage of the "go" function, with parameter "Content" and "Int"
func (wg *WaitGroup) GoCtxInt(f func(ctx context.Context, index int), ctx context.Context, index int) {
	wg.Go(func() {
		f(ctx, index)
	})
}

// GoCtxInt64 Extended usage of the "go" function, with parameter "Content" and "Int64"
func (wg *WaitGroup) GoCtxInt64(f func(ctx context.Context, index int64), ctx context.Context, index int64) {
	wg.Go(func() {
		f(ctx, index)
	})
}

// GoCtxInt32 Extended usage of the "go" function, with parameter "Content" and "Int32"
func (wg *WaitGroup) GoCtxInt32(f func(ctx context.Context, index int32), ctx context.Context, index int32) {
	wg.Go(func() {
		f(ctx, index)
	})
}

// GoCtxInt16 Extended usage of the "go" function, with parameter "Content" and "Int16"
func (wg *WaitGroup) GoCtxInt16(f func(ctx context.Context, index int16), ctx context.Context, index int16) {
	wg.Go(func() {
		f(ctx, index)
	})
}

// GoCtxInt8 Extended usage of the "go" function, with parameter "Content" and "Int8"
func (wg *WaitGroup) GoCtxInt8(f func(ctx context.Context, index int8), ctx context.Context, index int8) {
	wg.Go(func() {
		f(ctx, index)
	})
}

// GoCtxUInt Extended usage of the "go" function, with parameter "Content" and "UInt"
func (wg *WaitGroup) GoCtxUInt(f func(ctx context.Context, index uint), ctx context.Context, index uint) {
	wg.Go(func() {
		f(ctx, index)
	})
}

// GoCtxUInt64 Extended usage of the "go" function, with parameter "Content" and "UInt64"
func (wg *WaitGroup) GoCtxUInt64(f func(ctx context.Context, index uint64), ctx context.Context, index uint64) {
	wg.Go(func() {
		f(ctx, index)
	})
}

// GoCtxUInt32 Extended usage of the "go" function, with parameter "Content" and "UInt32"
func (wg *WaitGroup) GoCtxUInt32(f func(ctx context.Context, index uint32), ctx context.Context, index uint32) {
	wg.Go(func() {
		f(ctx, index)
	})
}

// GoCtxUInt16 Extended usage of the "go" function, with parameter "Content" and "UInt16"
func (wg *WaitGroup) GoCtxUInt16(f func(ctx context.Context, index uint16), ctx context.Context, index uint16) {
	wg.Go(func() {
		f(ctx, index)
	})
}

// GoCtxUInt8 Extended usage of the "go" function, with parameter "Content" and "UInt8"
func (wg *WaitGroup) GoCtxUInt8(f func(ctx context.Context, index uint8), ctx context.Context, index uint8) {
	wg.Go(func() {
		f(ctx, index)
	})
}

// GoCtxString Extended usage of the "go" function, with parameter "Content" and "String"
func (wg *WaitGroup) GoCtxString(f func(ctx context.Context, str string), ctx context.Context, str string) {
	wg.Go(func() {
		f(ctx, str)
	})
}

// GoCtxBytes Extended usage of the "go" function, with parameter "Content" and "Bytes"
func (wg *WaitGroup) GoCtxBytes(f func(ctx context.Context, byteSlice []byte), ctx context.Context, byteSlice []byte) {
	wg.Go(func() {
		f(ctx, byteSlice)
	})
}
