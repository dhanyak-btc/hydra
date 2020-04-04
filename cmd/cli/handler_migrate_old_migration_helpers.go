package cli

import (
	"fmt"
	"github.com/gobuffalo/pop/v5"
	"github.com/ory/x/sqlcon"
	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"
	"strconv"
	"strings"
	"time"
)

type oldTableName string

const (
	MigrationTableName                     = "hydra_schema_migrations"
	clientMigrationTableName  oldTableName = "hydra_client_migration"
	jwkMigrationTableName     oldTableName = "hydra_jwk_migration"
	consentMigrationTableName oldTableName = "hydra_oauth2_authentication_consent_migration"
	oauth2MigrationTableName  oldTableName = "hydra_oauth2_migration"
)

func getMigrationRecords(c *pop.Connection, tableName oldTableName) ([]migrate.MigrationRecord, error) {
	var records []migrate.MigrationRecord

	/* #nosec G201 TableName is static */
	err := sqlcon.HandleError(
		c.RawQuery(fmt.Sprintf("SELECT * FROM %s", tableName)).
			Eager().
			All(&records))

	return records, err
}

func migrateOldMigrationTables(c *pop.Connection) error {
	if err := c.RawQuery(fmt.Sprintf("SELECT * FROM %s", clientMigrationTableName)); err != nil {
		// assume there are no old migration tables => done
		return nil
	}

	return sqlcon.HandleError(c.Transaction(func(tx *pop.Connection) error {
		// in this order the migrations only depend on already done ones
		for i, table := range []oldTableName{clientMigrationTableName, jwkMigrationTableName, consentMigrationTableName, oauth2MigrationTableName} {
			// get old migrations
			migrations, err := getMigrationRecords(tx, table)
			if err != nil {
				return err
			}

			// translate migrations
			for _, m := range migrations {
				if m.AppliedAt.Before(time.Now()) {
					// the migration was run already -> set it as run for fizz
					// fizz standard version pattern: YYYYMMDDhhmmss
					migrationNumber, err := strconv.ParseInt(strings.TrimSuffix(m.Id, ".sql"), 10, 0)
					if err != nil {
						return errors.WithStack(err)
					}

					if err := tx.RawQuery(
						fmt.Sprintf("INSERT INTO %s (version) VALUES (2019%02d%08d)", MigrationTableName, i+1, migrationNumber)).
						Exec(); err != nil {
						return sqlcon.HandleError(err)
					}
				}
			}

			// delete old migration table
			if err := tx.RawQuery(fmt.Sprintf("DROP TABLE %s", table)).Exec(); err != nil {
				return sqlcon.HandleError(err)
			}
		}

		return nil
	}))
}