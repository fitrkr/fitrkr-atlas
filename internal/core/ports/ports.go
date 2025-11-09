package ports

type Write struct {
	Exercise   ExerciseWriteGroup
	Equipment  EquipmentWriteGroup
	Attachment AttachmentWrite
	Muscle     MuscleWriteGroup
	Category   CategoryWriteGroup
}

type Read struct {
	Exercise   ExerciseReadGroup
	Equipment  EquipmentReadGroup
	Attachment AttachmentRead
	Muscle     MuscleReadGroup
	Category   CategoryReadGroup
}

type EquipmentWriteGroup struct {
	EquipmentWrite
	Attachment EquipmentAttachmentWrite
}

type EquipmentReadGroup struct {
	EquipmentRead
	Attachment EquipmentAttachmentRead
}

type MuscleReadGroup struct {
	MuscleRead
	Group MuscleGroupRead
}

type MuscleWriteGroup struct {
	MuscleWrite
	Group MuscleGroupWrite
}

type CategoryWriteGroup struct {
	CategoryWrite
	Subcategory SubcategoryWrite
}

type CategoryReadGroup struct {
	CategoryRead
	Subcategory SubcategoryRead
}

type ExerciseWriteGroup struct {
	ExerciseWrite
	Alias       ExerciseAliasWrite
	Equipment   ExerciseEquipmentWrite
	Muscle      ExerciseMuscleWrite
	Category    ExerciseCategoryWrite
	Media       ExerciseMediaWrite
	Instruction ExerciseInstructionWrite
}

type ExerciseReadGroup struct {
	ExerciseRead
	Alias       ExerciseAliasRead
	Equipment   ExerciseEquipmentRead
	Muscle      ExerciseMuscleRead
	Category    ExerciseCategoryRead
	Media       ExerciseMediaRead
	Instruction ExerciseInstructionRead
}
