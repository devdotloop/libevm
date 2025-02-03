// Copyright 2025 the libevm authors.
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

// Package errs provides a mechanism for [testing error semantics] through
// unique identifiers, instead of depending on error messages that may result in
// change-detector tests.
//
// [testing error semantics]: https://google.github.io/styleguide/go/decisions#test-error-semantics
package errs

import (
	"errors"
	"fmt"
)

// An ID is a distinct numeric identifier for an error. It has no effect on the
// error message and can only be accessed with this package.
type ID int

// Error returns a new error with the ID.
func (id ID) Error(msg string) error {
	return noWrap{errors.New(msg), id}
}

type noWrap struct {
	error
	id ID
}

// Errorf is the formatted equivalent of [ID.Error], supporting the same
// wrapping semantics as [fmt.Errorf].
func (id ID) Errorf(format string, a ...any) error {
	switch err := fmt.Errorf(format, a...).(type) {
	case singleWrapper:
		return single{err, id}
	case multiWrapper:
		return multi{err, id}
	default:
		return noWrap{err, id}
	}
}

type singleWrapper interface {
	error
	Unwrap() error
}

type single struct {
	singleWrapper
	id ID
}

type multiWrapper interface {
	error
	Unwrap() []error
}

type multi struct {
	multiWrapper
	id ID
}

// IDOf returns the ID of the error, if one exists, and a flag to indicate as
// such. IDOf does not unwrap errors.
func IDOf(err error) (ID, bool) {
	switch err := err.(type) {
	case noWrap:
		return err.id, true
	case single:
		return err.id, true
	case multi:
		return err.id, true
	default:
		return 0, false
	}
}
