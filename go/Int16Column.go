// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Int16Column struct {
	_tab flatbuffers.Table
}

func GetRootAsInt16Column(buf []byte, offset flatbuffers.UOffsetT) *Int16Column {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Int16Column{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsInt16Column(buf []byte, offset flatbuffers.UOffsetT) *Int16Column {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Int16Column{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Int16Column) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Int16Column) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Int16Column) Unicode() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Int16Column) MutateUnicode(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *Int16Column) Int16(j int) int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt16(a + flatbuffers.UOffsetT(j*2))
	}
	return 0
}

func (rcv *Int16Column) Int16Length() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Int16Column) MutateInt16(j int, n int16) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt16(a+flatbuffers.UOffsetT(j*2), n)
	}
	return false
}

func Int16ColumnStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func Int16ColumnAddUnicode(builder *flatbuffers.Builder, unicode uint32) {
	builder.PrependUint32Slot(0, unicode, 0)
}
func Int16ColumnAddInt16(builder *flatbuffers.Builder, int16 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(int16), 0)
}
func Int16ColumnStartInt16Vector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(2, numElems, 2)
}
func Int16ColumnEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
