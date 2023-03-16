// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ClusterQuery struct {
	_tab flatbuffers.Table
}

func GetRootAsClusterQuery(buf []byte, offset flatbuffers.UOffsetT) *ClusterQuery {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ClusterQuery{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsClusterQuery(buf []byte, offset flatbuffers.UOffsetT) *ClusterQuery {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ClusterQuery{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ClusterQuery) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ClusterQuery) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ClusterQuery) Ssql(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *ClusterQuery) SsqlLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ClusterQuery) SsqlBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ClusterQuery) MutateSsql(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *ClusterQuery) Hosts(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *ClusterQuery) HostsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ClusterQuery) Local() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ClusterQuery) MutateLocal(n int32) bool {
	return rcv._tab.MutateInt32Slot(8, n)
}

func ClusterQueryStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func ClusterQueryAddSsql(builder *flatbuffers.Builder, ssql flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(ssql), 0)
}
func ClusterQueryStartSsqlVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func ClusterQueryAddHosts(builder *flatbuffers.Builder, hosts flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(hosts), 0)
}
func ClusterQueryStartHostsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ClusterQueryAddLocal(builder *flatbuffers.Builder, local int32) {
	builder.PrependInt32Slot(2, local, 0)
}
func ClusterQueryEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}