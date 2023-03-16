// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Int32Column struct {
	_tab flatbuffers.Table
}

func GetRootAsInt32Column(buf []byte, offset flatbuffers.UOffsetT) *Int32Column {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Int32Column{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsInt32Column(buf []byte, offset flatbuffers.UOffsetT) *Int32Column {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Int32Column{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Int32Column) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Int32Column) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Int32Column) Unicode() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Int32Column) MutateUnicode(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *Int32Column) Int32(j int) int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *Int32Column) Int32Length() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Int32Column) MutateInt32(j int, n int32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func Int32ColumnStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func Int32ColumnAddUnicode(builder *flatbuffers.Builder, unicode uint32) {
	builder.PrependUint32Slot(0, unicode, 0)
}
func Int32ColumnAddInt32(builder *flatbuffers.Builder, int32 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(int32), 0)
}
func Int32ColumnStartInt32Vector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func Int32ColumnEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}