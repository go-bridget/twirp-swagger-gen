package swagger

var typeAliases = map[string]struct {
	Type, Format string
}{
	"google.protobuf.Timestamp": {
		Type:   "string",
		Format: "date-time",
	},
	"google.protobuf.Duration": {
		Type: "string",
	},
	"google.protobuf.StringValue": {
		Type: "string",
	},
	"google.protobuf.BytesValue": {
		Type:   "string",
		Format: "byte",
	},
	"google.protobuf.Int32Value": {
		Type:   "integer",
		Format: "int32",
	},
	"google.protobuf.UInt32Value": {
		Type:   "integer",
		Format: "int64",
	},
	"google.protobuf.Int64Value": {
		Type:   "string",
		Format: "int64",
	},
	"google.protobuf.UInt64Value": {
		Type:   "string",
		Format: "uint64",
	},
	"google.protobuf.FloatValue": {
		Type:   "number",
		Format: "float",
	},
	"google.protobuf.DoubleValue": {
		Type:   "number",
		Format: "double",
	},
	"google.protobuf.BoolValue": {
		Type:   "boolean",
		Format: "boolean",
	},
	"google.protobuf.Empty": {},
}
