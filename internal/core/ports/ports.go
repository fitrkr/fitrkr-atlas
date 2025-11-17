// Package ports
package ports

type Write struct {
	Exercise   ExerciseWriteGroup
	Equipment  EquipmentWriteGroup
	Attachment AttachmentWrite
	Muscle     MuscleWrite
	Category   CategoryWrite
	View       ViewWrite
}

type Read struct {
	Exercise   ExerciseReadGroup
	Equipment  EquipmentReadGroup
	Attachment AttachmentRead
	Muscle     MuscleRead
	Category   CategoryRead
	View       ViewRead
}

type EquipmentWriteGroup struct {
	EquipmentWrite
	Attachment EquipmentAttachmentWrite
}

type EquipmentReadGroup struct {
	EquipmentRead
	Attachment EquipmentAttachmentRead
}

type ExerciseWriteGroup struct {
	ExerciseWrite
	Alias      ExerciseAliasWrite
	Attachment ExerciseAttachmentWrite
	Muscle     ExerciseMuscleWrite
	Category   ExerciseCategoryWrite
}

type ExerciseReadGroup struct {
	ExerciseRead
	Alias      ExerciseAliasRead
	Attachment ExerciseAttachmentRead
	Muscle     ExerciseMuscleRead
	Category   ExerciseCategoryRead
}
