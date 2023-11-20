package ient_test

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"testing"
	"time"

	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/lib/pq"
	dockertest "github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/storage/ient"
	entgen "github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/storage/ient/gen"
	"github.com/ruslanSorokin/human-data-enricher/internal/human-data-enricher/provider/test_person"
	"github.com/ruslanSorokin/human-data-enricher/internal/pkg/util"
)

const (
	imageName    = "postgres"
	imageVersion = "16.0"
)

const (
	dbname   = "human-data-enricher_test"
	user     = "u"
	password = "p"
)

type IntegrationSuite struct {
	*suite.Suite
	*test_person.PSuite

	resource *dockertest.Resource
	pool     *dockertest.Pool
	driver   *entgen.Client
}

func NewIntegrationSuite(s *suite.Suite) *IntegrationSuite {
	return &IntegrationSuite{
		Suite:    s,
		PSuite:   nil,
		resource: nil,
		pool:     nil,
		driver:   nil,
	}
}

func TestIntegration_PostgresPersonStorage(t *testing.T) {
	suite.Run(t, NewIntegrationSuite(&suite.Suite{Assertions: assert.New(t)}))
}

func (s *IntegrationSuite) SetupSuite() {
	t := s.T()

	flag.Parse()
	if testing.Short() {
		t.Skip()
	}

	p := util.MustXX(dockertest.NewPool(""))

	s.pool = p

	r := util.MustXX(LaunchPostgresContainer(p, imageVersion))
	s.resource = r

	x := strings.Split(r.GetHostPort("5432/tcp"), ":")

	host, port := x[0], util.MustXX(strconv.Atoi(x[1]))

	driver := util.MustXX(ient.NewDriver(&ient.Config{
		DBName:   dbname,
		Username: user,
		Password: password,
		Hostname: host,
		Port:     port,
	}))
	s.driver = driver

	log := slog.Default()

	repo := ient.NewPersonStorage(log, driver)
	s.PSuite = test_person.NewSuite(s, repo)
}

func (s *IntegrationSuite) TearDownSuite() {
	if testing.Verbose() && s.Suite.T().Failed() {
		// If tests failed with "-v" flag, keep container running.
		return
	}
	util.MustX(s.driver.Close())
	util.MustX(RemovePostgresContainer(s.resource))
}

func (s *IntegrationSuite) SetupTest() {
	util.MustX(s.driver.Schema.Create(
		context.Background(),
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
	))
}

func (s *IntegrationSuite) TearDownTest() {
	if testing.Verbose() && s.Suite.T().Failed() {
		// If tests failed with "-v" flag, keep database intact.
		return
	}
	util.MustXX(FlushPostgres(s.driver))
}

func LaunchPostgresContainer(
	p *dockertest.Pool,
	v string,
) (*dockertest.Resource, error) {
	r, err := p.Run(imageName, v, []string{
		fmt.Sprintf("POSTGRES_PASSWORD=%s", password),
		fmt.Sprintf("POSTGRES_USER=%s", user),
		fmt.Sprintf("POSTGRES_DB=%s", dbname),
		"listen_addresses = '*'",
	})
	if err != nil {
		return nil, err
	}

	databaseURL := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		user,
		password,
		r.GetHostPort("5432/tcp"),
		dbname,
	)

	p.MaxWait = 120 * time.Second

	if err = p.Retry(func() error {
		db, err := sql.Open("postgres", databaseURL)
		if err != nil {
			return err
		}
		defer db.Close()
		return db.Ping()
	}); err != nil {
		return r, err
	}

	return r, err
}

func RemovePostgresContainer(r *dockertest.Resource) error {
	return r.Close()
}

func FlushPostgres(c *entgen.Client) (int, error) {
	return c.Person.
		Delete().
		Exec(context.TODO())
}
