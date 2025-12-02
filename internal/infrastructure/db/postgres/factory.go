package postgres

import (
	"database/sql"

	"github.com/fitrkr/atlas/internal/core/ports"
)

type PostgresPortsProvider struct {
	db *sql.DB
}

func NewProvider(db *sql.DB) *PostgresPortsProvider {
	return &PostgresPortsProvider{db: db}
}

func (p *PostgresPortsProvider) CreatePorts() (ports.Read, ports.Write) {
	read := ports.Read{
		Exercise: ports.ExerciseReadGroup{
			ExerciseRead: newExerciseReader(p.db),
			Alias:        newExerciseAliasReader(p.db),
			Attachment:   newExerciseAttachmentReader(p.db),
			Muscle:       newExerciseMuscleReader(p.db),
			Category:     newExerciseCategoryReader(p.db),
		},
		Equipment: ports.EquipmentReadGroup{
			EquipmentRead: newEquipmentReader(p.db),
			Attachment:    newEquipmentAttachmentReader(p.db),
		},
		Attachment: newAttachmentReader(p.db),
		Muscle:     newMuscleReader(p.db),
		Category:   newCategoryReader(p.db),
		View:       newViewReader(p.db),
	}
	write := ports.Write{
		Exercise: ports.ExerciseWriteGroup{
			ExerciseWrite: newExerciseWriter(p.db),
			Alias:         newExerciseAliasWriter(p.db),
			Attachment:    newExerciseAttachmentWriter(p.db),
			Muscle:        newExerciseMuscleWriter(p.db),
			Category:      newExerciseCategoryWriter(p.db),
		},
		Equipment: ports.EquipmentWriteGroup{
			EquipmentWrite: newEquipmentWriter(p.db),
			Attachment:     newEquipmentAttachmentWriter(p.db),
		},
		Attachment: newAttachmentWriter(p.db),
		Muscle:     newMuscleWriter(p.db),
		Category:   newCategoryWriter(p.db),
		View:       newViewWriter(p.db),
	}

	return read, write
}
