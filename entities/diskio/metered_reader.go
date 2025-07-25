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

package diskio

import (
	"time"
)

type Reader interface {
	Read(p []byte) (n int, err error)
	ReadAt(p []byte, off int64) (n int, err error)
}

type MeteredReaderCallback func(read int64, nanoseconds int64)

type MeteredReader struct {
	file Reader
	cb   MeteredReaderCallback
}

// Read passes the read through to the underlying reader. On a successful read,
// it will trigger the attached callback and provide it with metrics. If no
// callback is set, it will ignore it.
func (m *MeteredReader) Read(p []byte) (n int, err error) {
	start := time.Now()
	n, err = m.file.Read(p)
	took := time.Since(start).Nanoseconds()
	if err != nil {
		return
	}

	if m.cb != nil {
		m.cb(int64(n), took)
	}

	return
}

// ReadAt passes the read through to the underlying reader. On a successful read,
// it will trigger the attached callback and provide it with metrics. If no
// callback is set, it will ignore it.
func (m *MeteredReader) ReadAt(p []byte, off int64) (n int, err error) {
	start := time.Now()
	n, err = m.file.ReadAt(p, off)
	took := time.Since(start).Nanoseconds()
	if err != nil {
		return
	}

	if m.cb != nil {
		m.cb(int64(n), took)
	}

	return
}

func NewMeteredReader(file Reader, cb MeteredReaderCallback) *MeteredReader {
	return &MeteredReader{file: file, cb: cb}
}
