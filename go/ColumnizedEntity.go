// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ColumnizedEntity struct {
	_tab flatbuffers.Table
}

func GetRootAsColumnizedEntity(buf []byte, offset flatbuffers.UOffsetT) *ColumnizedEntity {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ColumnizedEntity{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsColumnizedEntity(buf []byte, offset flatbuffers.UOffsetT) *ColumnizedEntity {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ColumnizedEntity{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ColumnizedEntity) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ColumnizedEntity) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ColumnizedEntity) I8(obj *Int8Column, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ColumnizedEntity) I8Length() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ColumnizedEntity) U8(obj *Uint8Column, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ColumnizedEntity) U8Length() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ColumnizedEntity) I16(obj *Int16Column, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ColumnizedEntity) I16Length() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ColumnizedEntity) U16(obj *Uint16Column, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ColumnizedEntity) U16Length() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ColumnizedEntity) I32(obj *Int32Column, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ColumnizedEntity) I32Length() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ColumnizedEntity) U32(obj *Uint32Column, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ColumnizedEntity) U32Length() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ColumnizedEntity) F32(obj *Float32Column, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ColumnizedEntity) F32Length() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ColumnizedEntity) D32(obj *Int32Column, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ColumnizedEntity) D32Length() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ColumnizedEntity) I64(obj *Int64Column, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ColumnizedEntity) I64Length() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ColumnizedEntity) U64(obj *Uint64Column, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ColumnizedEntity) U64Length() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ColumnizedEntity) F64(obj *Float64Column, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ColumnizedEntity) F64Length() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ColumnizedEntity) D64(obj *Int64Column, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ColumnizedEntity) D64Length() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ColumnizedEntity) Text(obj *Uint32Column, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ColumnizedEntity) TextLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ColumnizedEntity) Bool(obj *BoolColumn, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ColumnizedEntity) BoolLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ColumnizedEntity) Null(obj *NullColumn, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ColumnizedEntity) NullLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ColumnizedEntity) ArrayBracket(obj *Uint16Column, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(34))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ColumnizedEntity) ArrayBracketLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(34))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ColumnizedEntity) ArrayElement(obj *Uint16Column, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(36))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ColumnizedEntity) ArrayElementLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(36))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func ColumnizedEntityStart(builder *flatbuffers.Builder) {
	builder.StartObject(17)
}
func ColumnizedEntityAddI8(builder *flatbuffers.Builder, i8 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(i8), 0)
}
func ColumnizedEntityStartI8Vector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ColumnizedEntityAddU8(builder *flatbuffers.Builder, u8 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(u8), 0)
}
func ColumnizedEntityStartU8Vector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ColumnizedEntityAddI16(builder *flatbuffers.Builder, i16 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(i16), 0)
}
func ColumnizedEntityStartI16Vector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ColumnizedEntityAddU16(builder *flatbuffers.Builder, u16 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(u16), 0)
}
func ColumnizedEntityStartU16Vector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ColumnizedEntityAddI32(builder *flatbuffers.Builder, i32 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(i32), 0)
}
func ColumnizedEntityStartI32Vector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ColumnizedEntityAddU32(builder *flatbuffers.Builder, u32 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(u32), 0)
}
func ColumnizedEntityStartU32Vector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ColumnizedEntityAddF32(builder *flatbuffers.Builder, f32 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(f32), 0)
}
func ColumnizedEntityStartF32Vector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ColumnizedEntityAddD32(builder *flatbuffers.Builder, d32 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(d32), 0)
}
func ColumnizedEntityStartD32Vector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ColumnizedEntityAddI64(builder *flatbuffers.Builder, i64 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(i64), 0)
}
func ColumnizedEntityStartI64Vector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ColumnizedEntityAddU64(builder *flatbuffers.Builder, u64 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(u64), 0)
}
func ColumnizedEntityStartU64Vector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ColumnizedEntityAddF64(builder *flatbuffers.Builder, f64 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(10, flatbuffers.UOffsetT(f64), 0)
}
func ColumnizedEntityStartF64Vector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ColumnizedEntityAddD64(builder *flatbuffers.Builder, d64 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(11, flatbuffers.UOffsetT(d64), 0)
}
func ColumnizedEntityStartD64Vector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ColumnizedEntityAddText(builder *flatbuffers.Builder, text flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(12, flatbuffers.UOffsetT(text), 0)
}
func ColumnizedEntityStartTextVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ColumnizedEntityAddBool(builder *flatbuffers.Builder, bool flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(13, flatbuffers.UOffsetT(bool), 0)
}
func ColumnizedEntityStartBoolVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ColumnizedEntityAddNull(builder *flatbuffers.Builder, null flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(14, flatbuffers.UOffsetT(null), 0)
}
func ColumnizedEntityStartNullVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ColumnizedEntityAddArrayBracket(builder *flatbuffers.Builder, arrayBracket flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(15, flatbuffers.UOffsetT(arrayBracket), 0)
}
func ColumnizedEntityStartArrayBracketVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ColumnizedEntityAddArrayElement(builder *flatbuffers.Builder, arrayElement flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(16, flatbuffers.UOffsetT(arrayElement), 0)
}
func ColumnizedEntityStartArrayElementVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ColumnizedEntityEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
