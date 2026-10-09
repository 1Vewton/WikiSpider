package target

// TargetType defines types of target to store data in
type TargetType string

const (
	// JSONTarget defines the target that stores data in json file
	JSONTarget TargetType = "json_target"
	// MilvusTarget defines target that stores data in milvus
	MilvusTarget TargetType = "milvus_target"
	// NotSupported defines targets that are not supported
	NotSupported TargetType = "not_supported"
)

// ToTargetType converts string to TargetType
func ToTargetType(
	data string,
) TargetType {
	switch data {
	case "json_target":
		return JSONTarget
	case "milvus_target":
		return MilvusTarget
	case "not_supported":
		return NotSupported
	default:
		return NotSupported
	}
}

// ToString converts target type to string
func (target TargetType) ToString() string {
	switch target {
	case JSONTarget:
		return "json_target"
	default:
		return "target_type"
	}
}

// GetAllTypes gets all the target types
func GetAllTypes() []TargetType {
	return []TargetType{
		JSONTarget,
		MilvusTarget,
	}
}
