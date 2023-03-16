// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Uint8Column struct {
	_tab flatbuffers.Table
}

func GetRootAsUint8Column(buf []byte, offset flatbuffers.UOffsetT) *Uint8Column {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Uint8Column{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsUint8Column(buf []byte, offset flatbuffers.UOffsetT) *Uint8Column {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Uint8Column{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Uint8Column) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Uint8Column) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Uint8Column) Unicode() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Uint8Column) MutateUnicode(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *Uint8Column) Uint8(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *Uint8Column) Uint8Length() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Uint8Column) Uint8Bytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Uint8Column) MutateUint8(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func Uint8ColumnStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func Uint8ColumnAddUnicode(builder *flatbuffers.Builder, unicode uint32) {
	builder.PrependUint32Slot(0, unicode, 0)
}
func Uint8ColumnAddUint8(builder *flatbuffers.Builder, uint8 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(uint8), 0)
}
func Uint8ColumnStartUint8Vector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func Uint8ColumnEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}