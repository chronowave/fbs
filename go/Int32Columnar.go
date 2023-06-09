// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Int32Columnar struct {
	_tab flatbuffers.Table
}

func GetRootAsInt32Columnar(buf []byte, offset flatbuffers.UOffsetT) *Int32Columnar {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Int32Columnar{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsInt32Columnar(buf []byte, offset flatbuffers.UOffsetT) *Int32Columnar {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Int32Columnar{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Int32Columnar) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Int32Columnar) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Int32Columnar) Metadata(obj *Metadata) *Metadata {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Metadata)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Int32Columnar) Value(j int) int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *Int32Columnar) ValueLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Int32Columnar) MutateValue(j int, n int32) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt32(a+flatbuffers.UOffsetT(j*4), n)
	}
	return false
}

func Int32ColumnarStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func Int32ColumnarAddMetadata(builder *flatbuffers.Builder, metadata flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(metadata), 0)
}
func Int32ColumnarAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(value), 0)
}
func Int32ColumnarStartValueVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func Int32ColumnarEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
