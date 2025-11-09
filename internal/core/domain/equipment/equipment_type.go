package equipment

import "strings"

type EquipmentType int

const (
	BodyWeight EquipmentType = iota + 1
	FreeWeight
	Machine
)

func NewEquipmentType(equipmentType string) (EquipmentType, error) {
	switch strings.TrimSpace(strings.ToLower(equipmentType)) {
	case "body_weight":
		return BodyWeight, nil
	case "free_weight":
		return FreeWeight, nil
	case "machine":
		return Machine, nil
	default:
		return 0, ErrInvalidEquipmentType
	}
}

func (e EquipmentType) ToString() string {
	switch e {
	case BodyWeight:
		return "body_weight"
	case FreeWeight:
		return "free_weight"
	case Machine:
		return "machine"
	default:
		return "unknown"
	}
}
