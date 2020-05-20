// Copyright 2020 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cleanup

// Cleanup allows defers to be aborted when cleanup needs to happen
// conditionally. Usage:
// cu := Make(func() { f.Close() })
// defer cu.Clean() // any failure before release is called will close the file.
// ...
// cu.Release() // on success, aborts closing the file.
// return f
type Cleanup struct {
	clean func()
}

// Make creates a new Cleanup object.
func Make(f func()) Cleanup {
	return Cleanup{clean: f}
}

// Clean calls the cleanup function.
func (c *Cleanup) Clean() {
	if c.clean != nil {
		c.clean()
		c.clean = nil
	}
}

// Release releases the cleanup from its duties, i.e. cleanup function is not
// called after this point.
func (c *Cleanup) Release() {
	c.clean = nil
}
