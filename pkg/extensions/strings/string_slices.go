/* Copyright 2019 DevFactory FZ LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License. */

package strings

// ContainsString checks if a string slice contains the string
func ContainsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

// RemoveString removes a string from a string slice and returns
// a new slice
func RemoveString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}

// CompareSlices compares two string slices and returns true if both are empty
// or have exactly the same list of elements (order doesn't matter)
func CompareSlices(first, second []string) bool {
	if len(first) != len(second) {
		return false
	}
	for _, fromFirst := range first {
		if !ContainsString(second, fromFirst) {
			return false
		}
	}
	return true
}

// GetCommon returns common part of first and second
func GetCommon(first, second []string) []string {
	var common []string
	for _, s := range second {
		if ContainsString(first, s) {
			common = append(common, s)
		}
	}
	return common
}
