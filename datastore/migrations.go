package datastore

import (
	"fmt"

	migrate "src.techknowlogick.com/xormigrate"
	"xorm.io/xorm"

	"github.com/octo-5/karrot-api/model"
)

func Migrate(db *xorm.Engine) error {
	migrator := migrate.New(db, migration)
	return migrator.Migrate()
}

func RollbackLast(db *xorm.Engine) error {
	migrator := migrate.New(db, migration)
	return migrator.RollbackLast()
}

var (
	tables = []interface{}{
		&model.User{},
		&model.Todo{},
	}
	migration = []*migrate.Migration{
		{
			ID:          "20210204",
			Description: "initalizing database",
			Migrate: func(tx *xorm.Engine) error {
				if err := tx.Sync2(tables...); err != nil {
					return err
				}

				queries := []string{
					"DROP TABLE IF EXISTS todos, users",
					`CREATE TABLE users (
							id					BIGSERIAL PRIMARY KEY,
							name				VARCHAR(25) NOT NULL UNIQUE,
							mobile			VARCHAR(11) NOT NULL UNIQUE,
							password		TEXT NOT NULL,
							role				VARCHAR(20) NOT NULL DEFAULT 'USER' CHECK (role = 'ADMIN' OR role = 'USER'),
							created_at	TIMESTAMP NOT NULL DEFAULT TIMEZONE('utc', NOW())
					)`,
					`CREATE TABLE todos (
							id					BIGSERIAL PRIMARY KEY,
							text				TEXT NOT NULL,
							author_id		INT8 NOT NULL REFERENCES users (id),
							created_at	TIMESTAMP NOT NULL DEFAULT TIMEZONE('utc', NOW())
					)`,
					"DROP INDEX IF EXISTS idx_todos_author_id, idx_todos_created_at",
					"CREATE INDEX idx_todos_author_id ON todos(author_id)",
					"CREATE INDEX idx_todos_created_at ON todos(created_at)",
				}

				for _, q := range queries {
					if _, err := tx.Exec(q); err != nil {
						return fmt.Errorf("exec: %s: %w", q, err)
					}
				}
				return nil
			},
			Rollback: func(tx *xorm.Engine) error {
				queries := []string{
					"DROP TABLE IF EXISTS todos, users",
				}

				for _, q := range queries {
					if _, err := tx.Exec(q); err != nil {
						return fmt.Errorf("exec: %s: %w", q, err)
					}
				}
				return nil
			},
		},
	}
)
