// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package rbac

type Access string

const (
	CREATE  Access = "CREATE"
	READ    Access = "READ"
	EXECUTE Access = "EXECUTE"
	UPDATE  Access = "UPDATE"
	DELETE  Access = "DELETE"
)

type Group struct {
	DefaultAccess map[string][]string `json:"defaultAccess,omitempty"` // Keep original
	Description   string              `json:"description,omitempty"`
	Id            string              `json:"id,omitempty"` // Keep omitempty for backward compatibility
	Roles         []Role              `json:"roles,omitempty"`
}

// Helper methods for type-safe access
func (g *Group) GetDefaultAccessTyped(key string) []Access {
	if g.DefaultAccess == nil {
		return nil
	}

	stringAccesses := g.DefaultAccess[key]
	accesses := make([]Access, len(stringAccesses))
	for i, s := range stringAccesses {
		accesses[i] = Access(s)
	}
	return accesses
}

func (g *Group) SetDefaultAccessTyped(key string, accesses []Access) {
	if g.DefaultAccess == nil {
		g.DefaultAccess = make(map[string][]string)
	}

	stringAccesses := make([]string, len(accesses))
	for i, a := range accesses {
		stringAccesses[i] = string(a)
	}
	g.DefaultAccess[key] = stringAccesses
}

func (g *Group) AddDefaultAccess(key string, access Access) {
	if g.DefaultAccess == nil {
		g.DefaultAccess = make(map[string][]string)
	}
	g.DefaultAccess[key] = append(g.DefaultAccess[key], string(access))
}

func (g *Group) HasDefaultAccess(key string, access Access) bool {
	if g.DefaultAccess == nil {
		return false
	}

	for _, a := range g.DefaultAccess[key] {
		if a == string(access) {
			return true
		}
	}
	return false
}
