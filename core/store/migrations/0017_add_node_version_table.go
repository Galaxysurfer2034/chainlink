package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

const up17 = `
CREATE TABLE IF NOT EXISTS "node_versions" (
    "version" TEXT PRIMARY KEY,
    "created_at" timestamp without time zone NOT NULL
);
`

const down17 = `
DROP TABLE IF EXISTS "node_versions";
`

func init() {
	Migrations = append(Migrations, &gormigrate.Migration{
		ID: "0017_add_node_version_table",
		Migrate: func(db *gorm.DB) error {
			return db.Exec(up17).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.Exec(down17).Error
		},
	})
}
