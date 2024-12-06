// Copyright 2024 the libevm authors.
//
// The libevm additions to go-ethereum are free software: you can redistribute
// them and/or modify them under the terms of the GNU Lesser General Public License
// as published by the Free Software Foundation, either version 3 of the License,
// or (at your option) any later version.
//
// The libevm additions are distributed in the hope that they will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser
// General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see
// <http://www.gnu.org/licenses/>.

package rlp

import (
	"io"
	"math/bits"
)

func writeTo(w io.Writer, b ...byte) error {
	_, err := w.Write(b)
	return err
}

// EncodeRLP implements the [Encoder] interface and is the inverse of
// [ParseTree].
func (b ByteNode) EncodeRLP(w io.Writer) error {
	if b < 128 {
		return writeTo(w, byte(b))
	}
	return writeTo(w, 0x81, byte(b))
}

func (b ByteNode) headerAndDataLengths() (uint64, uint64) {
	if b < 128 {
		return 1, 0
	}
	return 1, 1
}

func byteLength(n uint64) uint64 {
	return uint64(bits.Len64(n)+7) / 8
}

func headerLength(data uint64) uint64 {
	if data <= 55 {
		return 1
	}
	return 1 + byteLength(data)
}

func writeNodeHeader(w io.Writer, dataLen uint64, short, long byte) error {
	var (
		tag   byte
		beLen []byte
	)
	if dataLen <= 55 {
		tag = short + byte(dataLen)
	} else {
		tag = long + byte(byteLength(dataLen))
		beLen = bigEndian(dataLen)
	}
	return writeTo(w, append(
		[]byte{tag},
		beLen...,
	)...)
}

// EncodeRLP implements the [Encoder] interface and is the inverse of
// [ParseTree].
func (l ListNode) EncodeRLP(w io.Writer) error {
	if err := writeNodeHeader(w, l.dataLength(), 0xc0, 0xf7); err != nil {
		return err
	}
	for _, it := range l {
		if err := it.EncodeRLP(w); err != nil {
			return err
		}
	}
	return nil
}

func (l ListNode) dataLength() uint64 {
	var n uint64
	for _, it := range l {
		a, b := it.headerAndDataLengths()
		n += a + b
	}
	return n
}

func (l ListNode) headerAndDataLengths() (uint64, uint64) {
	n := l.dataLength()
	return headerLength(n), n
}

// EncodeRLP implements the [Encoder] interface and is the inverse of
// [ParseTree].
func (s StringNode) EncodeRLP(w io.Writer) error {
	if err := writeNodeHeader(w, uint64(len(s)), 0x80, 0xb7); err != nil {
		return err
	}
	return writeTo(w, []byte(s)...)
}

func (s StringNode) headerAndDataLengths() (uint64, uint64) {
	data := uint64(len(s))
	return headerLength(data), data
}

// bigEndian returns the big-endian representation of n, without leading zeroes.
func bigEndian(n uint64) []byte {
	buf := make([]byte, byteLength(n))
	for i := len(buf) - 1; i >= 0; i-- {
		buf[i] = byte(n & 0xff)
		n >>= 8
	}
	return buf
}
