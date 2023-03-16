// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SemiStructuredData struct {
	_tab flatbuffers.Table
}

func GetRootAsSemiStructuredData(buf []byte, offset flatbuffers.UOffsetT) *SemiStructuredData {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SemiStructuredData{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsSemiStructuredData(buf []byte, offset flatbuffers.UOffsetT) *SemiStructuredData {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &SemiStructuredData{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *SemiStructuredData) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SemiStructuredData) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SemiStructuredData) Sz() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SemiStructuredData) MutateSz(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *SemiStructuredData) Content(obj *Content) *Content {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Content)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *SemiStructuredData) ArrayBracket(obj *ArrayBracket) *ArrayBracket {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(ArrayBracket)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *SemiStructuredData) ArrayElement(obj *ArrayElement) *ArrayElement {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(ArrayElement)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *SemiStructuredData) Text(obj *TextColumnar) *TextColumnar {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(TextColumnar)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *SemiStructuredData) Bool(obj *BoolColumnar) *BoolColumnar {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(BoolColumnar)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *SemiStructuredData) Null(obj *NullColumnar) *NullColumnar {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(NullColumnar)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *SemiStructuredData) F64(obj *Float64Columnar) *Float64Columnar {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Float64Columnar)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *SemiStructuredData) F32(obj *Float32Columnar) *Float32Columnar {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Float32Columnar)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *SemiStructuredData) U64(obj *Uint64Columnar) *Uint64Columnar {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Uint64Columnar)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *SemiStructuredData) I64(obj *Int64Columnar) *Int64Columnar {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Int64Columnar)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *SemiStructuredData) D64(obj *Int64Columnar) *Int64Columnar {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Int64Columnar)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *SemiStructuredData) U32(obj *Uint32Columnar) *Uint32Columnar {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Uint32Columnar)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *SemiStructuredData) I32(obj *Int32Columnar) *Int32Columnar {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Int32Columnar)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *SemiStructuredData) D32(obj *Int32Columnar) *Int32Columnar {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Int32Columnar)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *SemiStructuredData) U16(obj *Uint16Columnar) *Uint16Columnar {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(34))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Uint16Columnar)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *SemiStructuredData) I16(obj *Int16Columnar) *Int16Columnar {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(36))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Int16Columnar)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *SemiStructuredData) U8(obj *Uint8Columnar) *Uint8Columnar {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(38))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Uint8Columnar)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *SemiStructuredData) I8(obj *Int8Columnar) *Int8Columnar {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(40))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Int8Columnar)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func SemiStructuredDataStart(builder *flatbuffers.Builder) {
	builder.StartObject(19)
}
func SemiStructuredDataAddSz(builder *flatbuffers.Builder, sz uint32) {
	builder.PrependUint32Slot(0, sz, 0)
}
func SemiStructuredDataAddContent(builder *flatbuffers.Builder, content flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(content), 0)
}
func SemiStructuredDataAddArrayBracket(builder *flatbuffers.Builder, arrayBracket flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(arrayBracket), 0)
}
func SemiStructuredDataAddArrayElement(builder *flatbuffers.Builder, arrayElement flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(arrayElement), 0)
}
func SemiStructuredDataAddText(builder *flatbuffers.Builder, text flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(text), 0)
}
func SemiStructuredDataAddBool(builder *flatbuffers.Builder, bool flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(bool), 0)
}
func SemiStructuredDataAddNull(builder *flatbuffers.Builder, null flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(null), 0)
}
func SemiStructuredDataAddF64(builder *flatbuffers.Builder, f64 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(f64), 0)
}
func SemiStructuredDataAddF32(builder *flatbuffers.Builder, f32 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(f32), 0)
}
func SemiStructuredDataAddU64(builder *flatbuffers.Builder, u64 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(u64), 0)
}
func SemiStructuredDataAddI64(builder *flatbuffers.Builder, i64 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(10, flatbuffers.UOffsetT(i64), 0)
}
func SemiStructuredDataAddD64(builder *flatbuffers.Builder, d64 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(11, flatbuffers.UOffsetT(d64), 0)
}
func SemiStructuredDataAddU32(builder *flatbuffers.Builder, u32 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(12, flatbuffers.UOffsetT(u32), 0)
}
func SemiStructuredDataAddI32(builder *flatbuffers.Builder, i32 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(13, flatbuffers.UOffsetT(i32), 0)
}
func SemiStructuredDataAddD32(builder *flatbuffers.Builder, d32 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(14, flatbuffers.UOffsetT(d32), 0)
}
func SemiStructuredDataAddU16(builder *flatbuffers.Builder, u16 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(15, flatbuffers.UOffsetT(u16), 0)
}
func SemiStructuredDataAddI16(builder *flatbuffers.Builder, i16 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(16, flatbuffers.UOffsetT(i16), 0)
}
func SemiStructuredDataAddU8(builder *flatbuffers.Builder, u8 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(17, flatbuffers.UOffsetT(u8), 0)
}
func SemiStructuredDataAddI8(builder *flatbuffers.Builder, i8 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(18, flatbuffers.UOffsetT(i8), 0)
}
func SemiStructuredDataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
