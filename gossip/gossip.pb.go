// Code generated by protoc-gen-gogo.
// source: cockroach/gossip/gossip.proto
// DO NOT EDIT!

/*
	Package gossip is a generated protocol buffer package.

	It is generated from these files:
		cockroach/gossip/gossip.proto

	It has these top-level messages:
		Request
		Response
		Info
*/
package gossip

import proto "github.com/gogo/protobuf/proto"
import cockroach_proto1 "github.com/cockroachdb/cockroach/proto"

// discarding unused import cockroach_proto "github.com/cockroachdb/cockroach/proto"
import cockroach_util "github.com/cockroachdb/cockroach/util"

// discarding unused import gogoproto "gogoproto"

import github_com_cockroachdb_cockroach_proto "github.com/cockroachdb/cockroach/proto"

import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
import errors "errors"

import io "io"
import fmt "fmt"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

// Request is the request struct passed with the Gossip RPC.
type Request struct {
	// Requesting node's ID.
	NodeID github_com_cockroachdb_cockroach_proto.NodeID `protobuf:"varint,1,opt,name=node_id,proto3,casttype=github.com/cockroachdb/cockroach/proto.NodeID" json:"node_id,omitempty"`
	// Address of the requesting client.
	Addr cockroach_util.UnresolvedAddr `protobuf:"bytes,2,opt,name=addr" json:"addr"`
	// Local address of client on requesting node (this is a kludge to
	// allow gossip to know when client connections are dropped).
	LAddr cockroach_util.UnresolvedAddr `protobuf:"bytes,3,opt,name=l_addr" json:"l_addr"`
	// Maximum sequence number of gossip from this peer.
	MaxSeq int64 `protobuf:"varint,4,opt,name=max_seq,proto3" json:"max_seq,omitempty"`
	// Delta of new Infos since last gossip.
	Delta map[string]*Info `protobuf:"bytes,5,rep,name=delta" json:"delta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

func (m *Request) GetAddr() cockroach_util.UnresolvedAddr {
	if m != nil {
		return m.Addr
	}
	return cockroach_util.UnresolvedAddr{}
}

func (m *Request) GetLAddr() cockroach_util.UnresolvedAddr {
	if m != nil {
		return m.LAddr
	}
	return cockroach_util.UnresolvedAddr{}
}

func (m *Request) GetDelta() map[string]*Info {
	if m != nil {
		return m.Delta
	}
	return nil
}

// Response is returned from the Gossip.Gossip RPC.
// Delta will be nil in the event that Alternate is set.
type Response struct {
	// Responding Node's ID.
	NodeID github_com_cockroachdb_cockroach_proto.NodeID `protobuf:"varint,1,opt,name=node_id,proto3,casttype=github.com/cockroachdb/cockroach/proto.NodeID" json:"node_id,omitempty"`
	// Address of the responding client.
	Addr cockroach_util.UnresolvedAddr `protobuf:"bytes,2,opt,name=addr" json:"addr"`
	// Non-nil means client should retry with this address.
	Alternate *cockroach_util.UnresolvedAddr `protobuf:"bytes,3,opt,name=alternate" json:"alternate,omitempty"`
	// Maximum sequence number of gossip from this peer.
	MaxSeq int64 `protobuf:"varint,4,opt,name=max_seq,proto3" json:"max_seq,omitempty"`
	// Requested delta of server's infostore.
	Delta map[string]*Info `protobuf:"bytes,5,rep,name=delta" json:"delta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

func (m *Response) GetAddr() cockroach_util.UnresolvedAddr {
	if m != nil {
		return m.Addr
	}
	return cockroach_util.UnresolvedAddr{}
}

func (m *Response) GetAlternate() *cockroach_util.UnresolvedAddr {
	if m != nil {
		return m.Alternate
	}
	return nil
}

func (m *Response) GetDelta() map[string]*Info {
	if m != nil {
		return m.Delta
	}
	return nil
}

// Info is the basic unit of information traded over the
// gossip network.
type Info struct {
	Value cockroach_proto1.Value `protobuf:"bytes,1,opt,name=value" json:"value"`
	// Wall time when info is to be discarded (Unix-nanos)
	TTLStamp int64 `protobuf:"varint,2,opt,name=ttl_stamp,proto3" json:"ttl_stamp,omitempty"`
	// Number of hops from originator
	Hops uint32 `protobuf:"varint,3,opt,name=hops,proto3" json:"hops,omitempty"`
	// Originating node's ID
	NodeID github_com_cockroachdb_cockroach_proto.NodeID `protobuf:"varint,4,opt,name=node_id,proto3,casttype=github.com/cockroachdb/cockroach/proto.NodeID" json:"node_id,omitempty"`
}

func (m *Info) Reset()         { *m = Info{} }
func (m *Info) String() string { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()    {}

func (m *Info) GetValue() cockroach_proto1.Value {
	if m != nil {
		return m.Value
	}
	return cockroach_proto1.Value{}
}

func (m *Request) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Request) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.NodeID != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintGossip(data, i, uint64(m.NodeID))
	}
	data[i] = 0x12
	i++
	i = encodeVarintGossip(data, i, uint64(m.Addr.Size()))
	n1, err := m.Addr.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	data[i] = 0x1a
	i++
	i = encodeVarintGossip(data, i, uint64(m.LAddr.Size()))
	n2, err := m.LAddr.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if m.MaxSeq != 0 {
		data[i] = 0x20
		i++
		i = encodeVarintGossip(data, i, uint64(m.MaxSeq))
	}
	if len(m.Delta) > 0 {
		keysForDelta := make([]string, 0, len(m.Delta))
		for k := range m.Delta {
			keysForDelta = append(keysForDelta, k)
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForDelta)
		for _, k := range keysForDelta {
			data[i] = 0x2a
			i++
			v := m.Delta[k]
			if v == nil {
				return 0, errors.New("proto: map has nil element")
			}
			msgSize := v.Size()
			mapSize := 1 + len(k) + sovGossip(uint64(len(k))) + 1 + msgSize + sovGossip(uint64(msgSize))
			i = encodeVarintGossip(data, i, uint64(mapSize))
			data[i] = 0xa
			i++
			i = encodeVarintGossip(data, i, uint64(len(k)))
			i += copy(data[i:], k)
			data[i] = 0x12
			i++
			i = encodeVarintGossip(data, i, uint64(v.Size()))
			n3, err := v.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n3
		}
	}
	return i, nil
}

func (m *Response) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Response) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.NodeID != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintGossip(data, i, uint64(m.NodeID))
	}
	data[i] = 0x12
	i++
	i = encodeVarintGossip(data, i, uint64(m.Addr.Size()))
	n4, err := m.Addr.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	if m.Alternate != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintGossip(data, i, uint64(m.Alternate.Size()))
		n5, err := m.Alternate.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.MaxSeq != 0 {
		data[i] = 0x20
		i++
		i = encodeVarintGossip(data, i, uint64(m.MaxSeq))
	}
	if len(m.Delta) > 0 {
		keysForDelta := make([]string, 0, len(m.Delta))
		for k := range m.Delta {
			keysForDelta = append(keysForDelta, k)
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForDelta)
		for _, k := range keysForDelta {
			data[i] = 0x2a
			i++
			v := m.Delta[k]
			if v == nil {
				return 0, errors.New("proto: map has nil element")
			}
			msgSize := v.Size()
			mapSize := 1 + len(k) + sovGossip(uint64(len(k))) + 1 + msgSize + sovGossip(uint64(msgSize))
			i = encodeVarintGossip(data, i, uint64(mapSize))
			data[i] = 0xa
			i++
			i = encodeVarintGossip(data, i, uint64(len(k)))
			i += copy(data[i:], k)
			data[i] = 0x12
			i++
			i = encodeVarintGossip(data, i, uint64(v.Size()))
			n6, err := v.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n6
		}
	}
	return i, nil
}

func (m *Info) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Info) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGossip(data, i, uint64(m.Value.Size()))
	n7, err := m.Value.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n7
	if m.TTLStamp != 0 {
		data[i] = 0x10
		i++
		i = encodeVarintGossip(data, i, uint64(m.TTLStamp))
	}
	if m.Hops != 0 {
		data[i] = 0x18
		i++
		i = encodeVarintGossip(data, i, uint64(m.Hops))
	}
	if m.NodeID != 0 {
		data[i] = 0x20
		i++
		i = encodeVarintGossip(data, i, uint64(m.NodeID))
	}
	return i, nil
}

func encodeFixed64Gossip(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Gossip(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGossip(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *Request) Size() (n int) {
	var l int
	_ = l
	if m.NodeID != 0 {
		n += 1 + sovGossip(uint64(m.NodeID))
	}
	l = m.Addr.Size()
	n += 1 + l + sovGossip(uint64(l))
	l = m.LAddr.Size()
	n += 1 + l + sovGossip(uint64(l))
	if m.MaxSeq != 0 {
		n += 1 + sovGossip(uint64(m.MaxSeq))
	}
	if len(m.Delta) > 0 {
		for k, v := range m.Delta {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
			}
			mapEntrySize := 1 + len(k) + sovGossip(uint64(len(k))) + 1 + l + sovGossip(uint64(l))
			n += mapEntrySize + 1 + sovGossip(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *Response) Size() (n int) {
	var l int
	_ = l
	if m.NodeID != 0 {
		n += 1 + sovGossip(uint64(m.NodeID))
	}
	l = m.Addr.Size()
	n += 1 + l + sovGossip(uint64(l))
	if m.Alternate != nil {
		l = m.Alternate.Size()
		n += 1 + l + sovGossip(uint64(l))
	}
	if m.MaxSeq != 0 {
		n += 1 + sovGossip(uint64(m.MaxSeq))
	}
	if len(m.Delta) > 0 {
		for k, v := range m.Delta {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
			}
			mapEntrySize := 1 + len(k) + sovGossip(uint64(len(k))) + 1 + l + sovGossip(uint64(l))
			n += mapEntrySize + 1 + sovGossip(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *Info) Size() (n int) {
	var l int
	_ = l
	l = m.Value.Size()
	n += 1 + l + sovGossip(uint64(l))
	if m.TTLStamp != 0 {
		n += 1 + sovGossip(uint64(m.TTLStamp))
	}
	if m.Hops != 0 {
		n += 1 + sovGossip(uint64(m.Hops))
	}
	if m.NodeID != 0 {
		n += 1 + sovGossip(uint64(m.NodeID))
	}
	return n
}

func sovGossip(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGossip(x uint64) (n int) {
	return sovGossip(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Request) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeID", wireType)
			}
			m.NodeID = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.NodeID |= (github_com_cockroachdb_cockroach_proto.NodeID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthGossip
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Addr.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LAddr", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthGossip
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LAddr.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSeq", wireType)
			}
			m.MaxSeq = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.MaxSeq |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthGossip
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if stringLenmapkey < 0 {
				return ErrInvalidLengthGossip
			}
			postStringIndexmapkey := iNdEx + int(stringLenmapkey)
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(data[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			var valuekey uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				valuekey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var mapmsglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				mapmsglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postmsgIndex := iNdEx + mapmsglen
			if mapmsglen < 0 {
				return ErrInvalidLengthGossip
			}
			if postmsgIndex > l {
				return io.ErrUnexpectedEOF
			}
			mapvalue := &Info{}
			if err := mapvalue.Unmarshal(data[iNdEx:postmsgIndex]); err != nil {
				return err
			}
			iNdEx = postmsgIndex
			if m.Delta == nil {
				m.Delta = make(map[string]*Info)
			}
			m.Delta[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipGossip(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGossip
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *Response) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeID", wireType)
			}
			m.NodeID = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.NodeID |= (github_com_cockroachdb_cockroach_proto.NodeID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthGossip
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Addr.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alternate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthGossip
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Alternate == nil {
				m.Alternate = &cockroach_util.UnresolvedAddr{}
			}
			if err := m.Alternate.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSeq", wireType)
			}
			m.MaxSeq = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.MaxSeq |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthGossip
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if stringLenmapkey < 0 {
				return ErrInvalidLengthGossip
			}
			postStringIndexmapkey := iNdEx + int(stringLenmapkey)
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(data[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			var valuekey uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				valuekey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var mapmsglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				mapmsglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postmsgIndex := iNdEx + mapmsglen
			if mapmsglen < 0 {
				return ErrInvalidLengthGossip
			}
			if postmsgIndex > l {
				return io.ErrUnexpectedEOF
			}
			mapvalue := &Info{}
			if err := mapvalue.Unmarshal(data[iNdEx:postmsgIndex]); err != nil {
				return err
			}
			iNdEx = postmsgIndex
			if m.Delta == nil {
				m.Delta = make(map[string]*Info)
			}
			m.Delta[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipGossip(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGossip
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *Info) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthGossip
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Value.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TTLStamp", wireType)
			}
			m.TTLStamp = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.TTLStamp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hops", wireType)
			}
			m.Hops = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Hops |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeID", wireType)
			}
			m.NodeID = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.NodeID |= (github_com_cockroachdb_cockroach_proto.NodeID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipGossip(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGossip
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func skipGossip(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for {
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGossip
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGossip(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGossip = fmt.Errorf("proto: negative length found during unmarshaling")
)