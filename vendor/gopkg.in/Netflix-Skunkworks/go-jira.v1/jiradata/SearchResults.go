package jiradata

/////////////////////////////////////////////////////////////////////////
// This Code is Generated by SlipScheme Project:
// https://github.com/coryb/slipscheme
//
// Generated with command:
// slipscheme -dir jiradata -pkg jiradata -overwrite schemas/SearchResults.json
/////////////////////////////////////////////////////////////////////////
//                            DO NOT EDIT                              //
/////////////////////////////////////////////////////////////////////////

// SearchResults defined from schema:
// {
//   "title": "Search Results",
//   "id": "https://docs.atlassian.com/jira/REST/schema/search-results#",
//   "type": "object",
//   "definitions": {
//     "field-meta": {
//       "title": "Field Meta",
//       "type": "object",
//       "properties": {
//         "allowedValues": {
//           "type": "array",
//           "items": {}
//         },
//         "autoCompleteUrl": {
//           "type": "string"
//         },
//         "defaultValue": {},
//         "hasDefaultValue": {
//           "type": "boolean"
//         },
//         "key": {
//           "type": "string"
//         },
//         "name": {
//           "type": "string"
//         },
//         "operations": {
//           "type": "array",
//           "items": {
//             "type": "string"
//           }
//         },
//         "required": {
//           "type": "boolean"
//         },
//         "schema": {
//           "$ref": "#/definitions/json-type"
//         }
//       }
//     },
//     "history-metadata-participant": {
//       "title": "History Metadata Participant",
//       "type": "object",
//       "properties": {
//         "avatarUrl": {
//           "type": "string"
//         },
//         "displayName": {
//           "type": "string"
//         },
//         "displayNameKey": {
//           "type": "string"
//         },
//         "id": {
//           "type": "string"
//         },
//         "type": {
//           "type": "string"
//         },
//         "url": {
//           "type": "string"
//         }
//       }
//     },
//     "json-type": {
//       "title": "Json Type",
//       "type": "object",
//       "properties": {
//         "custom": {
//           "type": "string"
//         },
//         "customId": {
//           "type": "integer"
//         },
//         "items": {
//           "type": "string"
//         },
//         "system": {
//           "type": "string"
//         },
//         "type": {
//           "type": "string"
//         }
//       }
//     },
//     "link-group": {
//       "title": "Link Group",
//       "type": "object",
//       "properties": {
//         "groups": {
//           "type": "array",
//           "items": {
//             "$ref": "#/definitions/link-group"
//           }
//         },
//         "header": {
//           "$ref": "#/definitions/simple-link"
//         },
//         "id": {
//           "type": "string"
//         },
//         "links": {
//           "type": "array",
//           "items": {
//             "$ref": "#/definitions/simple-link"
//           }
//         },
//         "styleClass": {
//           "type": "string"
//         },
//         "weight": {
//           "type": "integer"
//         }
//       }
//     },
//     "simple-link": {
//       "title": "Simple Link",
//       "type": "object",
//       "properties": {
//         "href": {
//           "type": "string"
//         },
//         "iconClass": {
//           "type": "string"
//         },
//         "id": {
//           "type": "string"
//         },
//         "label": {
//           "type": "string"
//         },
//         "styleClass": {
//           "type": "string"
//         },
//         "title": {
//           "type": "string"
//         },
//         "weight": {
//           "type": "integer"
//         }
//       }
//     }
//   },
//   "properties": {
//     "expand": {
//       "title": "expand",
//       "type": "string"
//     },
//     "issues": {
//       "title": "issues",
//       "type": "array",
//       "items": {
//         "title": "Issue",
//         "type": "object",
//         "properties": {
//           "changelog": {
//             "title": "Changelog",
//             "type": "object",
//             "properties": {
//               "histories": {
//                 "title": "histories",
//                 "type": "array",
//                 "items": {
//                   "title": "Change History",
//                   "type": "object",
//                   "properties": {
//                     "author": {
//                       "title": "User",
//                       "type": "object",
//                       "properties": {
//                         "accountId": {
//                           "title": "accountId",
//                           "type": "string"
//                         },
//                         "active": {
//                           "title": "active",
//                           "type": "boolean"
//                         },
//                         "avatarUrls": {
//                           "title": "avatarUrls",
//                           "type": "object",
//                           "patternProperties": {
//                             ".+": {
//                               "type": "string"
//                             }
//                           }
//                         },
//                         "displayName": {
//                           "title": "displayName",
//                           "type": "string"
//                         },
//                         "emailAddress": {
//                           "title": "emailAddress",
//                           "type": "string"
//                         },
//                         "key": {
//                           "title": "key",
//                           "type": "string"
//                         },
//                         "name": {
//                           "title": "name",
//                           "type": "string"
//                         },
//                         "self": {
//                           "title": "self",
//                           "type": "string"
//                         },
//                         "timeZone": {
//                           "title": "timeZone",
//                           "type": "string"
//                         }
//                       }
//                     },
//                     "created": {
//                       "title": "created",
//                       "type": "string"
//                     },
//                     "historyMetadata": {
//                       "title": "History Metadata",
//                       "type": "object",
//                       "properties": {
//                         "activityDescription": {
//                           "title": "activityDescription",
//                           "type": "string"
//                         },
//                         "activityDescriptionKey": {
//                           "title": "activityDescriptionKey",
//                           "type": "string"
//                         },
//                         "actor": {
//                           "title": "History Metadata Participant",
//                           "type": "object",
//                           "properties": {
//                             "avatarUrl": {
//                               "type": "string"
//                             },
//                             "displayName": {
//                               "type": "string"
//                             },
//                             "displayNameKey": {
//                               "type": "string"
//                             },
//                             "id": {
//                               "type": "string"
//                             },
//                             "type": {
//                               "type": "string"
//                             },
//                             "url": {
//                               "type": "string"
//                             }
//                           }
//                         },
//                         "cause": {
//                           "title": "History Metadata Participant",
//                           "type": "object",
//                           "properties": {
//                             "avatarUrl": {
//                               "type": "string"
//                             },
//                             "displayName": {
//                               "type": "string"
//                             },
//                             "displayNameKey": {
//                               "type": "string"
//                             },
//                             "id": {
//                               "type": "string"
//                             },
//                             "type": {
//                               "type": "string"
//                             },
//                             "url": {
//                               "type": "string"
//                             }
//                           }
//                         },
//                         "description": {
//                           "title": "description",
//                           "type": "string"
//                         },
//                         "descriptionKey": {
//                           "title": "descriptionKey",
//                           "type": "string"
//                         },
//                         "emailDescription": {
//                           "title": "emailDescription",
//                           "type": "string"
//                         },
//                         "emailDescriptionKey": {
//                           "title": "emailDescriptionKey",
//                           "type": "string"
//                         },
//                         "extraData": {
//                           "title": "extraData",
//                           "type": "object",
//                           "patternProperties": {
//                             ".+": {
//                               "type": "string"
//                             }
//                           }
//                         },
//                         "generator": {
//                           "title": "History Metadata Participant",
//                           "type": "object",
//                           "properties": {
//                             "avatarUrl": {
//                               "type": "string"
//                             },
//                             "displayName": {
//                               "type": "string"
//                             },
//                             "displayNameKey": {
//                               "type": "string"
//                             },
//                             "id": {
//                               "type": "string"
//                             },
//                             "type": {
//                               "type": "string"
//                             },
//                             "url": {
//                               "type": "string"
//                             }
//                           }
//                         },
//                         "type": {
//                           "title": "type",
//                           "type": "string"
//                         }
//                       }
//                     },
//                     "id": {
//                       "title": "id",
//                       "type": "string"
//                     },
//                     "items": {
//                       "title": "items",
//                       "type": "array",
//                       "items": {
//                         "title": "Change Item",
//                         "type": "object",
//                         "properties": {
//                           "field": {
//                             "title": "field",
//                             "type": "string"
//                           },
//                           "fieldId": {
//                             "title": "fieldId",
//                             "type": "string"
//                           },
//                           "fieldtype": {
//                             "title": "fieldtype",
//                             "type": "string"
//                           },
//                           "from": {
//                             "title": "from",
//                             "type": "string"
//                           },
//                           "fromString": {
//                             "title": "fromString",
//                             "type": "string"
//                           },
//                           "to": {
//                             "title": "to",
//                             "type": "string"
//                           },
//                           "toString": {
//                             "title": "toString",
//                             "type": "string"
//                           }
//                         }
//                       }
//                     }
//                   }
//                 }
//               },
//               "maxResults": {
//                 "title": "maxResults",
//                 "type": "integer"
//               },
//               "startAt": {
//                 "title": "startAt",
//                 "type": "integer"
//               },
//               "total": {
//                 "title": "total",
//                 "type": "integer"
//               }
//             }
//           },
//           "editmeta": {
//             "title": "Edit Meta",
//             "type": "object",
//             "properties": {
//               "fields": {
//                 "title": "fields",
//                 "type": "object",
//                 "patternProperties": {
//                   ".+": {
//                     "title": "Field Meta",
//                     "type": "object",
//                     "properties": {
//                       "allowedValues": {
//                         "type": "array",
//                         "items": {}
//                       },
//                       "autoCompleteUrl": {
//                         "type": "string"
//                       },
//                       "defaultValue": {},
//                       "hasDefaultValue": {
//                         "type": "boolean"
//                       },
//                       "key": {
//                         "type": "string"
//                       },
//                       "name": {
//                         "type": "string"
//                       },
//                       "operations": {
//                         "type": "array",
//                         "items": {
//                           "type": "string"
//                         }
//                       },
//                       "required": {
//                         "type": "boolean"
//                       },
//                       "schema": {
//                         "$ref": "#/definitions/json-type"
//                       }
//                     }
//                   }
//                 }
//               }
//             }
//           },
//           "expand": {
//             "title": "expand",
//             "type": "string"
//           },
//           "fields": {
//             "title": "fields",
//             "type": "object",
//             "patternProperties": {
//               ".+": {}
//             }
//           },
//           "fieldsToInclude": {
//             "title": "Included Fields",
//             "type": "object"
//           },
//           "id": {
//             "title": "id",
//             "type": "string"
//           },
//           "key": {
//             "title": "key",
//             "type": "string"
//           },
//           "names": {
//             "title": "names",
//             "type": "object",
//             "patternProperties": {
//               ".+": {
//                 "type": "string"
//               }
//             }
//           },
//           "operations": {
//             "title": "Opsbar",
//             "type": "object",
//             "properties": {
//               "linkGroups": {
//                 "title": "linkGroups",
//                 "type": "array",
//                 "items": {
//                   "title": "Link Group",
//                   "type": "object",
//                   "properties": {
//                     "groups": {
//                       "type": "array",
//                       "items": {
//                         "$ref": "#/definitions/link-group"
//                       }
//                     },
//                     "header": {
//                       "$ref": "#/definitions/simple-link"
//                     },
//                     "id": {
//                       "type": "string"
//                     },
//                     "links": {
//                       "type": "array",
//                       "items": {
//                         "$ref": "#/definitions/simple-link"
//                       }
//                     },
//                     "styleClass": {
//                       "type": "string"
//                     },
//                     "weight": {
//                       "type": "integer"
//                     }
//                   }
//                 }
//               }
//             }
//           },
//           "properties": {
//             "title": "Properties",
//             "type": "object",
//             "properties": {
//               "properties": {
//                 "title": "properties",
//                 "type": "object",
//                 "patternProperties": {
//                   ".+": {
//                     "type": "string"
//                   }
//                 }
//               }
//             }
//           },
//           "renderedFields": {
//             "title": "renderedFields",
//             "type": "object",
//             "patternProperties": {
//               ".+": {}
//             }
//           },
//           "schema": {
//             "title": "schema",
//             "type": "object",
//             "patternProperties": {
//               ".+": {
//                 "title": "Json Type",
//                 "type": "object",
//                 "properties": {
//                   "custom": {
//                     "type": "string"
//                   },
//                   "customId": {
//                     "type": "integer"
//                   },
//                   "items": {
//                     "type": "string"
//                   },
//                   "system": {
//                     "type": "string"
//                   },
//                   "type": {
//                     "type": "string"
//                   }
//                 }
//               }
//             }
//           },
//           "self": {
//             "title": "self",
//             "type": "string"
//           },
//           "transitions": {
//             "title": "transitions",
//             "type": "array",
//             "items": {
//               "title": "Transition",
//               "type": "object",
//               "properties": {
//                 "expand": {
//                   "title": "expand",
//                   "type": "string"
//                 },
//                 "fields": {
//                   "title": "fields",
//                   "type": "object",
//                   "patternProperties": {
//                     ".+": {
//                       "title": "Field Meta",
//                       "type": "object",
//                       "properties": {
//                         "allowedValues": {
//                           "type": "array",
//                           "items": {}
//                         },
//                         "autoCompleteUrl": {
//                           "type": "string"
//                         },
//                         "defaultValue": {},
//                         "hasDefaultValue": {
//                           "type": "boolean"
//                         },
//                         "key": {
//                           "type": "string"
//                         },
//                         "name": {
//                           "type": "string"
//                         },
//                         "operations": {
//                           "type": "array",
//                           "items": {
//                             "type": "string"
//                           }
//                         },
//                         "required": {
//                           "type": "boolean"
//                         },
//                         "schema": {
//                           "$ref": "#/definitions/json-type"
//                         }
//                       }
//                     }
//                   }
//                 },
//                 "hasScreen": {
//                   "title": "hasScreen",
//                   "type": "boolean"
//                 },
//                 "id": {
//                   "title": "id",
//                   "type": "string"
//                 },
//                 "name": {
//                   "title": "name",
//                   "type": "string"
//                 },
//                 "to": {
//                   "title": "Status",
//                   "type": "object",
//                   "properties": {
//                     "description": {
//                       "title": "description",
//                       "type": "string"
//                     },
//                     "iconUrl": {
//                       "title": "iconUrl",
//                       "type": "string"
//                     },
//                     "id": {
//                       "title": "id",
//                       "type": "string"
//                     },
//                     "name": {
//                       "title": "name",
//                       "type": "string"
//                     },
//                     "self": {
//                       "title": "self",
//                       "type": "string"
//                     },
//                     "statusCategory": {
//                       "title": "Status Category",
//                       "type": "object",
//                       "properties": {
//                         "colorName": {
//                           "title": "colorName",
//                           "type": "string"
//                         },
//                         "id": {
//                           "title": "id",
//                           "type": "integer"
//                         },
//                         "key": {
//                           "title": "key",
//                           "type": "string"
//                         },
//                         "name": {
//                           "title": "name",
//                           "type": "string"
//                         },
//                         "self": {
//                           "title": "self",
//                           "type": "string"
//                         }
//                       }
//                     },
//                     "statusColor": {
//                       "title": "statusColor",
//                       "type": "string"
//                     }
//                   }
//                 }
//               }
//             }
//           },
//           "versionedRepresentations": {
//             "title": "versionedRepresentations",
//             "type": "object",
//             "patternProperties": {
//               ".+": {
//                 "type": "object",
//                 "patternProperties": {
//                   ".+": {}
//                 }
//               }
//             }
//           }
//         }
//       }
//     },
//     "maxResults": {
//       "title": "maxResults",
//       "type": "integer"
//     },
//     "names": {
//       "title": "names",
//       "type": "object",
//       "patternProperties": {
//         ".+": {
//           "type": "string"
//         }
//       }
//     },
//     "schema": {
//       "title": "schema",
//       "type": "object",
//       "patternProperties": {
//         ".+": {
//           "title": "Json Type",
//           "type": "object",
//           "properties": {
//             "custom": {
//               "type": "string"
//             },
//             "customId": {
//               "type": "integer"
//             },
//             "items": {
//               "type": "string"
//             },
//             "system": {
//               "type": "string"
//             },
//             "type": {
//               "type": "string"
//             }
//           }
//         }
//       }
//     },
//     "startAt": {
//       "title": "startAt",
//       "type": "integer"
//     },
//     "total": {
//       "title": "total",
//       "type": "integer"
//     },
//     "warningMessages": {
//       "title": "warningMessages",
//       "type": "array",
//       "items": {
//         "type": "string"
//       }
//     }
//   }
// }
type SearchResults struct {
	Expand          string            `json:"expand,omitempty" yaml:"expand,omitempty"`
	Issues          Issues            `json:"issues,omitempty" yaml:"issues,omitempty"`
	MaxResults      int               `json:"maxResults,omitempty" yaml:"maxResults,omitempty"`
	Names           map[string]string `json:"names,omitempty" yaml:"names,omitempty"`
	Schema          JSONTypeMap       `json:"schema,omitempty" yaml:"schema,omitempty"`
	StartAt         int               `json:"startAt,omitempty" yaml:"startAt,omitempty"`
	Total           int               `json:"total,omitempty" yaml:"total,omitempty"`
	WarningMessages WarningMessages   `json:"warningMessages,omitempty" yaml:"warningMessages,omitempty"`
}
