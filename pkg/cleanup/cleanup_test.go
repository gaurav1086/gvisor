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

import "testing"

func testCleanupHelper(clean *bool) {
	cu := Make(func() {
		*clean = true
	})
	defer cu.Clean()
}

func TestCleanup(t *testing.T) {
	clean := false
	testCleanupHelper(&clean)
	if !clean {
		t.Fatalf("cleanup function was not called.")
	}
}

func TestRelease(t *testing.T) {
	cu := Make(func() {
		t.Fatalf("cleanup function should not be called")
	})
	defer cu.Clean()
	cu.Release()
}
