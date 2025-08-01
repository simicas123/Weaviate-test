//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2025 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package segmentindex

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/weaviate/weaviate/usecases/byteops"
)

const (
	// HeaderSize describes the general offset in a segment until the data
	// starts, it is composed of 2 bytes for level, 2 bytes for version,
	// 2 bytes for secondary index count, 2 bytes for strategy, 8 bytes
	// for the pointer to the index part
	HeaderSize = 16

	// ChecksumSize describes the length of the segment file checksum.
	// This is currently based on the CRC32 hashing algorithm.
	ChecksumSize = 4
)

type Header struct {
	Level            uint16
	Version          uint16
	SecondaryIndices uint16
	Strategy         Strategy
	IndexStart       uint64
}

func (h *Header) WriteTo(w io.Writer) (int64, error) {
	data := make([]byte, HeaderSize)
	rw := byteops.NewReadWriter(data)
	rw.WriteUint16(h.Level)
	rw.WriteUint16(h.Version)
	rw.WriteUint16(h.SecondaryIndices)
	rw.WriteUint16(uint16(h.Strategy))
	rw.WriteUint64(h.IndexStart)

	write, err := w.Write(data)
	if err != nil {
		return 0, err
	}
	if write != HeaderSize {
		return 0, fmt.Errorf("expected to write %d bytes, got %d", HeaderSize, write)
	}

	return int64(HeaderSize), nil
}

func (h *Header) PrimaryIndex(source []byte) ([]byte, error) {
	if h.SecondaryIndices == 0 {
		return source[h.IndexStart:], nil
	}

	offsets, err := h.parseSecondaryIndexOffsets(
		source[h.IndexStart:h.secondaryIndexOffsetsEnd()])
	if err != nil {
		return nil, err
	}

	// the beginning of the first secondary is also the end of the primary
	end := offsets[0]
	return source[h.secondaryIndexOffsetsEnd():end], nil
}

func (h *Header) secondaryIndexOffsetsEnd() uint64 {
	return h.IndexStart + (uint64(h.SecondaryIndices) * 8)
}

func (h *Header) parseSecondaryIndexOffsets(source []byte) ([]uint64, error) {
	r := bufio.NewReader(bytes.NewReader(source))

	offsets := make([]uint64, h.SecondaryIndices)
	if err := binary.Read(r, binary.LittleEndian, &offsets); err != nil {
		return nil, err
	}

	return offsets, nil
}

func (h *Header) SecondaryIndex(source []byte, indexID uint16) ([]byte, error) {
	if indexID >= h.SecondaryIndices {
		return nil, fmt.Errorf("retrieve index %d with len %d",
			indexID, h.SecondaryIndices)
	}

	offsets, err := h.parseSecondaryIndexOffsets(
		source[h.IndexStart:h.secondaryIndexOffsetsEnd()])
	if err != nil {
		return nil, err
	}

	start := offsets[indexID]
	if indexID == h.SecondaryIndices-1 {
		// this is the last index, return until EOF
		return source[start:], nil
	}

	end := offsets[indexID+1]
	return source[start:end], nil
}

func ParseHeader(data []byte) (*Header, error) {
	if len(data) != HeaderSize {
		return nil, fmt.Errorf("expected %d bytes, got %d", HeaderSize, len(data))
	}
	rw := byteops.NewReadWriter(data)
	out := &Header{}
	out.Level = rw.ReadUint16()
	out.Version = rw.ReadUint16()
	out.SecondaryIndices = rw.ReadUint16()
	out.Strategy = Strategy(rw.ReadUint16())
	out.IndexStart = rw.ReadUint64()

	if out.Version > CurrentSegmentVersion {
		return nil, fmt.Errorf("unsupported version %d", out.Version)
	}

	return out, nil
}
