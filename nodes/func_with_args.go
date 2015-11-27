// Auto-generated from postgres/src/include/nodes/parsenodes.h - DO NOT EDIT

package pg_query

import "encoding/json"

type FuncWithArgs struct {
	Funcname []Node `json:"funcname"` /* qualified name of function */
	Funcargs []Node `json:"funcargs"` /* list of Typename nodes */
}

func (node FuncWithArgs) MarshalJSON() ([]byte, error) {
	type FuncWithArgsMarshalAlias FuncWithArgs
	return json.Marshal(map[string]interface{}{
		"FUNCWITHARGS": (*FuncWithArgsMarshalAlias)(&node),
	})
}

func (node *FuncWithArgs) UnmarshalJSON(input []byte) (err error) {
	var fields map[string]json.RawMessage

	err = json.Unmarshal(input, &fields)
	if err != nil {
		return
	}

	if fields["funcname"] != nil {
		node.Funcname, err = UnmarshalNodeArrayJSON(fields["funcname"])
		if err != nil {
			return
		}
	}

	if fields["funcargs"] != nil {
		node.Funcargs, err = UnmarshalNodeArrayJSON(fields["funcargs"])
		if err != nil {
			return
		}
	}

	return
}