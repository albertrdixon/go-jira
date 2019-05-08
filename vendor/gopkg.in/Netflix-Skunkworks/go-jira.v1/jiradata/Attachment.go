package jiradata

/////////////////////////////////////////////////////////////////////////
// This Code is Generated by SlipScheme Project:
// https://github.com/coryb/slipscheme
//
// Generated with command:
// slipscheme -dir jiradata -pkg jiradata -overwrite schemas/ListofAttachment.json
/////////////////////////////////////////////////////////////////////////
//                            DO NOT EDIT                              //
/////////////////////////////////////////////////////////////////////////

// Attachment defined from schema:
// {
//   "title": "Attachment",
//   "type": "object",
//   "properties": {
//     "author": {
//       "title": "User",
//       "type": "object",
//       "properties": {
//         "accountId": {
//           "title": "accountId",
//           "type": "string"
//         },
//         "active": {
//           "title": "active",
//           "type": "boolean"
//         },
//         "applicationRoles": {
//           "title": "Simple List Wrapper",
//           "type": "object",
//           "properties": {
//             "items": {
//               "type": "array",
//               "items": {
//                 "title": "Group",
//                 "type": "object",
//                 "properties": {
//                   "name": {
//                     "type": "string"
//                   },
//                   "self": {
//                     "type": "string"
//                   }
//                 }
//               }
//             },
//             "max-results": {
//               "type": "integer"
//             },
//             "size": {
//               "type": "integer"
//             }
//           }
//         },
//         "avatarUrls": {
//           "title": "avatarUrls",
//           "type": "object",
//           "patternProperties": {
//             ".+": {
//               "type": "string"
//             }
//           }
//         },
//         "displayName": {
//           "title": "displayName",
//           "type": "string"
//         },
//         "emailAddress": {
//           "title": "emailAddress",
//           "type": "string"
//         },
//         "expand": {
//           "title": "expand",
//           "type": "string"
//         },
//         "groups": {
//           "title": "Simple List Wrapper",
//           "type": "object",
//           "properties": {
//             "items": {
//               "type": "array",
//               "items": {
//                 "title": "Group",
//                 "type": "object",
//                 "properties": {
//                   "name": {
//                     "type": "string"
//                   },
//                   "self": {
//                     "type": "string"
//                   }
//                 }
//               }
//             },
//             "max-results": {
//               "type": "integer"
//             },
//             "size": {
//               "type": "integer"
//             }
//           }
//         },
//         "key": {
//           "title": "key",
//           "type": "string"
//         },
//         "locale": {
//           "title": "locale",
//           "type": "string"
//         },
//         "name": {
//           "title": "name",
//           "type": "string"
//         },
//         "self": {
//           "title": "self",
//           "type": "string"
//         },
//         "timeZone": {
//           "title": "timeZone",
//           "type": "string"
//         }
//       }
//     },
//     "content": {
//       "title": "content",
//       "type": "string"
//     },
//     "created": {
//       "title": "created",
//       "type": "string"
//     },
//     "filename": {
//       "title": "filename",
//       "type": "string"
//     },
//     "id": {
//       "title": "id",
//       "type": "integer"
//     },
//     "mimeType": {
//       "title": "mimeType",
//       "type": "string"
//     },
//     "properties": {
//       "title": "properties",
//       "type": "object",
//       "patternProperties": {
//         ".+": {}
//       }
//     },
//     "self": {
//       "title": "self",
//       "type": "string"
//     },
//     "size": {
//       "title": "size",
//       "type": "integer"
//     },
//     "thumbnail": {
//       "title": "thumbnail",
//       "type": "string"
//     }
//   }
// }
type Attachment struct {
	Author     *User                  `json:"author,omitempty" yaml:"author,omitempty"`
	Content    string                 `json:"content,omitempty" yaml:"content,omitempty"`
	Created    string                 `json:"created,omitempty" yaml:"created,omitempty"`
	Filename   string                 `json:"filename,omitempty" yaml:"filename,omitempty"`
	ID         IntOrString            `json:"id,omitempty" yaml:"id,omitempty"`
	MimeType   string                 `json:"mimeType,omitempty" yaml:"mimeType,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty" yaml:"properties,omitempty"`
	Self       string                 `json:"self,omitempty" yaml:"self,omitempty"`
	Size       int                    `json:"size,omitempty" yaml:"size,omitempty"`
	Thumbnail  string                 `json:"thumbnail,omitempty" yaml:"thumbnail,omitempty"`
}
