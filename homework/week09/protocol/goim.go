package protocol

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
)

const (
	PackageLengthSize = 4
	HeaderLengthSize  = 2
	VersionSize       = 4
	OperationSize     = 4
	SequenceIdSize    = 4
	)

const 	MaxSize = int32(1 << 12)

type Message struct {
	HeaderLen  int16
	Version    int16
	Operation  int32
	SequenceId int32
	Body       string
}


func Pack(ctx context.Context, m *Message) []byte {
	packLen := HeaderLengthSize + len(m.Body)
	buf := make([]byte, 1000)
	binary.BigEndian.PutUint32(buf, uint32(packLen))
	binary.BigEndian.PutUint16(buf, uint16(m.HeaderLen))
	binary.BigEndian.PutUint16(buf, uint16(m.Version))
	binary.BigEndian.PutUint32(buf, uint32(m.Operation))
	binary.BigEndian.PutUint32(buf, uint32(m.SequenceId))
	if len(m.Body) > 0 {
		buf = append(buf, []byte(m.Body)...)
	}
	return buf
}

func ReadSocket(ctx context.Context, buffer *bytes.Buffer) (*Message, error) {
	lengthSize := PackageLengthSize + HeaderLengthSize + VersionSize + OperationSize + SequenceIdSize
	buf := buffer.Next(lengthSize)
	if len(buf) != lengthSize * 8 {
		return nil, errors.New("invalid message")
	}
	message := &Message{
		HeaderLen:  0,
		Version:    0,
		Operation:  0,
		SequenceId: 0,
		Body:       "",
	}
	packLenOffset := PackageLengthSize
	packLen := binary.BigEndian.Uint32(buf[0:PackageLengthSize])
	if packLen > uint32(MaxSize) {
		return nil, errors.New("max sie over limit")
	}
	
	headerLenOffset := packLenOffset + HeaderLengthSize
	message.HeaderLen = int16(binary.BigEndian.Uint16(buf[packLenOffset:headerLenOffset]))
	if message.HeaderLen != HeaderLengthSize {
		return nil, errors.New("invalid header length")
	}
	versionOffset := headerLenOffset + VersionSize
	message.Version = int16(binary.BigEndian.Uint16(buf[headerLenOffset:versionOffset]))
	OperationOffset := versionOffset + OperationSize
	message.Operation = int32(binary.BigEndian.Uint32(buf[versionOffset:OperationOffset]))
	sequenceIdOffset := OperationOffset + OperationSize
	message.Operation = int32(binary.BigEndian.Uint32(buf[OperationOffset:sequenceIdOffset]))

	bodyLen := int32(packLen) - int32(message.HeaderLen)
	if bodyLen > 0 {
		message.Body = string(buffer.Next(int(bodyLen)))
	}
	return message, nil
	
}