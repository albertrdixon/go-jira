package jiradata

/////////////////////////////////////////////////////////////////////////
// This Code is Generated by SlipScheme Project:
// https://github.com/coryb/slipscheme
//
// Generated with command:
// slipscheme -pkg jiradata -dir jiradata -overwrite schemas/IssueUpdate.json
/////////////////////////////////////////////////////////////////////////
//                            DO NOT EDIT                              //
/////////////////////////////////////////////////////////////////////////

// EntityProperty defined from schema:
// {
//   "title": "Entity Property",
//   "type": "object",
//   "properties": {
//     "key": {
//       "title": "key",
//       "type": "string"
//     },
//     "value": {
//       "title": "value"
//     }
//   }
// }
type EntityProperty struct {
	Key   string      `json:"key,omitempty" yaml:"key,omitempty"`
	Value interface{} `json:"value,omitempty" yaml:"value,omitempty"`
}