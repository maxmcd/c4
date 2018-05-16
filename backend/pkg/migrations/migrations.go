package migrations

import (
	"regexp"
	"sort"
	"strconv"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/cockroachdb"
	bindata "github.com/golang-migrate/migrate/source/go-bindata"
	"github.com/golang/glog"
	"github.com/maxmcd/c4/backend/pkg/config"
)

func getMigrations() (*migrate.Migrate, []int) {
	// TODO: Debug why the migrations aren't running even though they show up in bindata
	s := bindata.Resource(AssetNames(),
		func(name string) ([]byte, error) {
			return Asset(name)
		})
	var migration_version = regexp.MustCompile(`^\d*`)
	var versions []int
	sort.Strings(s.Names)
	for _, name := range s.Names {
		if string_mv := migration_version.FindString(name); string_mv != "" {
			mv, err := strconv.Atoi(string_mv)
			if err == nil &&
				(len(versions) == 0) ||
				(len(versions) > 0 && versions[len(versions)-1] != mv) {
				versions = append(versions, mv)
			}
		}
	}
	sort.Ints(versions)
	d, err := bindata.WithInstance(s)
	if err != nil {
		glog.Fatal(err)
	}
	m, err := migrate.NewWithSourceInstance(
		"go-bindata", d, config.DatabaseUrl("cockroach"))
	if err != nil {
		glog.Fatal(err)
	}
	return m, versions
}

func RunMigrations() {
	m, _ := getMigrations()
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		glog.Fatal(err)
	}
}

func DropAllMigrations() {
	m, _ := getMigrations()
	if err := m.Drop(); err != nil {
		glog.Fatal(err)
	}
}

func DownAllMigrations() {
	m, versions := getMigrations()
	// m.Drop()
	// force the newest version as we require down migrations to
	// check for existence
	if err := m.Force(versions[len(versions)-1]); err != nil {
		glog.Fatal(err)
	}
	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		glog.Fatal(err)
	}
}
