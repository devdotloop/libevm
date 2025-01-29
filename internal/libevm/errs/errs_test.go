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

package errs

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIDOf(t *testing.T) {
	tests := []struct {
		err    error
		wantID ID
		wantOK bool
	}{
		{
			err: errors.New("x"),
		},
		{
			err: fmt.Errorf("x"),
		},
		{
			err:    ID(3).Error("x"),
			wantID: 3,
			wantOK: true,
		},
		{
			err:    ID(42).Errorf("x"),
			wantID: 42,
			wantOK: true,
		},
		{
			err:    ID(99).Errorf("%w", errors.New("x")),
			wantID: 99,
			wantOK: true,
		},
		{
			err:    ID(0).Errorf("%w %w", errors.New("x"), errors.New("y")),
			wantID: 0,
			wantOK: true,
		},
	}

	for _, tt := range tests {
		got, ok := IDOf(tt.err)
		assert.Equalf(t, tt.wantID, got, "IDOf(%T)", tt.err)
		assert.Equalf(t, tt.wantOK, ok, "IDOf(%T)", tt.err)
	}
}
