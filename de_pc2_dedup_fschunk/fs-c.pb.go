// Code generated by protoc-gen-gogo.
// source: fs-c.proto
// DO NOT EDIT!

/*
	Package de_pc2_dedup_fschunk is a generated protocol buffer package.

	It is generated from these files:
		fs-c.proto

	It has these top-level messages:
		File
		Chunk
*/
package de_pc2_dedup_fschunk

import proto "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "gogo.pb"

import io "io"
import fmt "fmt"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type File struct {
	Filename         *string `protobuf:"bytes,1,opt,name=filename" json:"filename,omitempty"`
	Fsize            *uint64 `protobuf:"varint,2,opt,name=fsize" json:"fsize,omitempty"`
	Type             *string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	ChunkCount       *uint32 `protobuf:"varint,4,opt,name=chunkCount" json:"chunkCount,omitempty"`
	Label            *string `protobuf:"bytes,5,opt,name=label" json:"label,omitempty"`
	Partial          *bool   `protobuf:"varint,6,opt,name=partial,def=0" json:"partial,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}

const Default_File_Partial bool = false

func (m *File) GetFilename() string {
	if m != nil && m.Filename != nil {
		return *m.Filename
	}
	return ""
}

func (m *File) GetFsize() uint64 {
	if m != nil && m.Fsize != nil {
		return *m.Fsize
	}
	return 0
}

func (m *File) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *File) GetChunkCount() uint32 {
	if m != nil && m.ChunkCount != nil {
		return *m.ChunkCount
	}
	return 0
}

func (m *File) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *File) GetPartial() bool {
	if m != nil && m.Partial != nil {
		return *m.Partial
	}
	return Default_File_Partial
}

type Chunk struct {
	// Usually a sha1/md5 fingerprint
	Fp    []byte  `protobuf:"bytes,2,opt,name=fp" json:"fp,omitempty"`
	Csize *uint32 `protobuf:"varint,3,opt,name=csize" json:"csize,omitempty"`
	// value of a rabin fingerprinter when the chunk was accepted. May not be set
	ChunkHash        *int64 `protobuf:"varint,4,opt,name=chunkHash" json:"chunkHash,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Chunk) Reset()         { *m = Chunk{} }
func (m *Chunk) String() string { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()    {}

func (m *Chunk) GetFp() []byte {
	if m != nil {
		return m.Fp
	}
	return nil
}

func (m *Chunk) GetCsize() uint32 {
	if m != nil && m.Csize != nil {
		return *m.Csize
	}
	return 0
}

func (m *Chunk) GetChunkHash() int64 {
	if m != nil && m.ChunkHash != nil {
		return *m.ChunkHash
	}
	return 0
}

func init() {
}
func (m *File) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
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
				return fmt.Errorf("proto: wrong wireType = %d for field Filename", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[index:postIndex])
			m.Filename = &s
			index = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fsize", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Fsize = &v
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[index:postIndex])
			m.Type = &s
			index = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChunkCount", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ChunkCount = &v
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Label", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[index:postIndex])
			m.Label = &s
			index = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Partial", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Partial = &b
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := github_com_gogo_protobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *Chunk) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fp", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fp = append(m.Fp, data[index:postIndex]...)
			index = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Csize", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Csize = &v
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChunkHash", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ChunkHash = &v
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := github_com_gogo_protobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *File) Size() (n int) {
	var l int
	_ = l
	if m.Filename != nil {
		l = len(*m.Filename)
		n += 1 + l + sovFsC(uint64(l))
	}
	if m.Fsize != nil {
		n += 1 + sovFsC(uint64(*m.Fsize))
	}
	if m.Type != nil {
		l = len(*m.Type)
		n += 1 + l + sovFsC(uint64(l))
	}
	if m.ChunkCount != nil {
		n += 1 + sovFsC(uint64(*m.ChunkCount))
	}
	if m.Label != nil {
		l = len(*m.Label)
		n += 1 + l + sovFsC(uint64(l))
	}
	if m.Partial != nil {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Chunk) Size() (n int) {
	var l int
	_ = l
	if m.Fp != nil {
		l = len(m.Fp)
		n += 1 + l + sovFsC(uint64(l))
	}
	if m.Csize != nil {
		n += 1 + sovFsC(uint64(*m.Csize))
	}
	if m.ChunkHash != nil {
		n += 1 + sovFsC(uint64(*m.ChunkHash))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovFsC(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFsC(x uint64) (n int) {
	return sovFsC(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *File) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *File) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Filename != nil {
		data[i] = 0xa
		i++
		i = encodeVarintFsC(data, i, uint64(len(*m.Filename)))
		i += copy(data[i:], *m.Filename)
	}
	if m.Fsize != nil {
		data[i] = 0x10
		i++
		i = encodeVarintFsC(data, i, uint64(*m.Fsize))
	}
	if m.Type != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintFsC(data, i, uint64(len(*m.Type)))
		i += copy(data[i:], *m.Type)
	}
	if m.ChunkCount != nil {
		data[i] = 0x20
		i++
		i = encodeVarintFsC(data, i, uint64(*m.ChunkCount))
	}
	if m.Label != nil {
		data[i] = 0x2a
		i++
		i = encodeVarintFsC(data, i, uint64(len(*m.Label)))
		i += copy(data[i:], *m.Label)
	}
	if m.Partial != nil {
		data[i] = 0x30
		i++
		if *m.Partial {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Chunk) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Chunk) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Fp != nil {
		data[i] = 0x12
		i++
		i = encodeVarintFsC(data, i, uint64(len(m.Fp)))
		i += copy(data[i:], m.Fp)
	}
	if m.Csize != nil {
		data[i] = 0x18
		i++
		i = encodeVarintFsC(data, i, uint64(*m.Csize))
	}
	if m.ChunkHash != nil {
		data[i] = 0x20
		i++
		i = encodeVarintFsC(data, i, uint64(*m.ChunkHash))
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64FsC(data []byte, offset int, v uint64) int {
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
func encodeFixed32FsC(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintFsC(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
