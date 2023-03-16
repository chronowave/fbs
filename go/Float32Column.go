// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Float32Column struct {
	_tab flatbuffers.Table
}

func GetRootAsFloat32Column(buf []byte, offset flatbuffers.UOffsetT) *Float32Column {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Float32Column{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsFloat32Column(buf []byte, offset flatbuffers.UOffsetT) *Float32Column {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Float32Column{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Float32Column) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Float32Column) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Float32Column) Unicode() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Float32Column) MutateUnicode(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *Float32Column) Float(j int) float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *Float32Column) FloatLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Float32Column) MutateFloat(j int, n float32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateFloat32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func Float32ColumnStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func Float32ColumnAddUnicode(builder *flatbuffers.Builder, unicode uint32) {
	builder.PrependUint32Slot(0, unicode, 0)
}
func Float32ColumnAddFloat(builder *flatbuffers.Builder, float flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(float), 0)
}
func Float32ColumnStartFloatVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func Float32ColumnEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
