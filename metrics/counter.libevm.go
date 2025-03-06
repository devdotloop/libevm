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

package metrics

// GetOrRegisterOrOverrideCounter searches for a metric already registered
// with the `name` given. If a metric is found and it is a [Counter], this one
// is returned. If a metric is found and it is not a [Counter], it is unregistered
// and replaced with a new registered [Counter]. If no metric is found, a new
// [Counter] is constructed and registered.
// This is especially useful for a metric defined in this project with a different
// type than a metric defined in a consumer of this project, for the same name.
func GetOrRegisterOrOverrideCounter(name string, r Registry) Counter {
	if r == nil {
		r = DefaultRegistry
	}
	counter, ok := r.GetOrRegister(name, NewCounter).(Counter)
	if ok {
		return counter
	}
	r.Unregister(name)
	counter = NewCounter()
	_ = r.Register(name, counter)
	return counter
}
