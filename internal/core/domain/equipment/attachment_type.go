package equipment

import "strings"

type AttachmentType int

const (
	Band AttachmentType = iota + 1
	Cable
	Plate
)

func NewAttachmentType(equipmentType string) (AttachmentType, error) {
	switch strings.TrimSpace(strings.ToLower(equipmentType)) {
	case "band":
		return Band, nil
	case "cable":
		return Cable, nil
	case "plate":
		return Plate, nil
	default:
		return 0, ErrInvalidAttachmentType
	}
}

func (a AttachmentType) ToString() string {
	switch a {
	case Band:
		return "band"
	case Cable:
		return "cable"
	case Plate:
		return "plate"
	default:
		return ""
	}
}
