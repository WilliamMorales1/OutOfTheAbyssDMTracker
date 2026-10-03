package db

import (
	"cmp"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

const migrationTable = "schema_migrations"

type migration struct {
	version int
	name    string
	up      string
	down    string
}

func loadMigrations(dir string) ([]migration, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read migrations: %w", err)
	}

	migrations := make([]migration, 0, len(entries))
	seen := make(map[int]struct{}, len(entries))
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".up.sql") {
			continue
		}

		base := strings.TrimSuffix(entry.Name(), ".up.sql")
		prefix, name, ok := strings.Cut(base, "_")
		if !ok {
			return nil, fmt.Errorf("invalid migration filename %q", entry.Name())
		}
		version, err := strconv.Atoi(prefix)
		if err != nil || version < 1 {
			return nil, fmt.Errorf("invalid migration version in %q", entry.Name())
		}
		if _, exists := seen[version]; exists {
			return nil, fmt.Errorf("duplicate migration version %d", version)
		}

		up, err := os.ReadFile(filepath.Join(dir, entry.Name()))
		if err != nil {
			return nil, fmt.Errorf("read migration %q: %w", entry.Name(), err)
		}
		downName := base + ".down.sql"
		down, err := os.ReadFile(filepath.Join(dir, downName))
		if err != nil {
			return nil, fmt.Errorf("read migration %q: %w", downName, err)
		}

		seen[version] = struct{}{}
		migrations = append(migrations, migration{
			version: version,
			name:    name,
			up:      string(up),
			down:    string(down),
		})
	}

	if len(migrations) == 0 {
		return nil, fmt.Errorf("no migrations found in %q", dir)
	}
	slices.SortFunc(migrations, func(a, b migration) int {
		return cmp.Compare(a.version, b.version)
	})
	return migrations, nil
}

func openMigrationDB(path string) (*sql.DB, error) {
	db, err := Open(path)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}
	return db, nil
}

func ensureMigrationTable(db *sql.DB) error {
	_, err := db.Exec(`
CREATE TABLE IF NOT EXISTS schema_migrations (
	version INTEGER NOT NULL PRIMARY KEY,
	dirty BOOLEAN NOT NULL
)`)
	return err
}

func migrationState(db *sql.DB) (int, bool, error) {
	var version int
	var dirty bool
	err := db.QueryRow("SELECT version, dirty FROM "+migrationTable).Scan(&version, &dirty)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, false, nil
	}
	if err != nil {
		return 0, false, fmt.Errorf("read migration state: %w", err)
	}
	return version, dirty, nil
}

func setMigrationState(tx *sql.Tx, version int) error {
	if _, err := tx.Exec("DELETE FROM " + migrationTable); err != nil {
		return err
	}
	_, err := tx.Exec("INSERT INTO "+migrationTable+" (version, dirty) VALUES (?, FALSE)", version)
	return err
}

func applyMigration(db *sql.DB, m migration) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	if _, err := tx.Exec(m.up); err != nil {
		return fmt.Errorf("apply %03d_%s: %w", m.version, m.name, err)
	}
	if err := setMigrationState(tx, m.version); err != nil {
		return fmt.Errorf("record %03d_%s: %w", m.version, m.name, err)
	}
	return tx.Commit()
}

func revertMigration(db *sql.DB, m migration, previousVersion int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	if _, err := tx.Exec(m.down); err != nil {
		return fmt.Errorf("revert %03d_%s: %w", m.version, m.name, err)
	}
	if previousVersion == 0 {
		if _, err := tx.Exec("DELETE FROM " + migrationTable); err != nil {
			return fmt.Errorf("record revert %03d_%s: %w", m.version, m.name, err)
		}
	} else if err := setMigrationState(tx, previousVersion); err != nil {
		return fmt.Errorf("record revert %03d_%s: %w", m.version, m.name, err)
	}
	return tx.Commit()
}

// RunMigrations applies all pending SQL migrations in migrationsDir.
// Migration state uses the same schema_migrations table as golang-migrate,
// so existing databases upgrade without a manual conversion.
func RunMigrations(migrationsDir, dbPath string) error {
	migrations, err := loadMigrations(migrationsDir)
	if err != nil {
		return err
	}
	db, err := openMigrationDB(dbPath)
	if err != nil {
		return fmt.Errorf("open migrations database: %w", err)
	}
	defer db.Close()

	if err := ensureMigrationTable(db); err != nil {
		return fmt.Errorf("create migration table: %w", err)
	}
	current, dirty, err := migrationState(db)
	if err != nil {
		return err
	}
	if dirty {
		return fmt.Errorf("database has dirty migration state at version %d", current)
	}

	for _, m := range migrations {
		if m.version <= current {
			continue
		}
		if err := applyMigration(db, m); err != nil {
			return err
		}
	}
	return nil
}

// ResetMigrations reverts every applied migration, then reapplies all
// migrations. It preserves the existing migrate command's reset behavior.
func ResetMigrations(migrationsDir, dbPath string) error {
	migrations, err := loadMigrations(migrationsDir)
	if err != nil {
		return err
	}
	db, err := openMigrationDB(dbPath)
	if err != nil {
		return fmt.Errorf("open migrations database: %w", err)
	}
	defer db.Close()

	if err := ensureMigrationTable(db); err != nil {
		return fmt.Errorf("create migration table: %w", err)
	}
	current, dirty, err := migrationState(db)
	if err != nil {
		return err
	}
	if dirty {
		return fmt.Errorf("database has dirty migration state at version %d", current)
	}

	for current > 0 {
		index := slices.IndexFunc(migrations, func(m migration) bool { return m.version == current })
		if index < 0 {
			return fmt.Errorf("migration %d is not present in %q", current, migrationsDir)
		}
		previousVersion := 0
		if index > 0 {
			previousVersion = migrations[index-1].version
		}
		if err := revertMigration(db, migrations[index], previousVersion); err != nil {
			return err
		}
		current = previousVersion
	}

	return RunMigrations(migrationsDir, dbPath)
}
