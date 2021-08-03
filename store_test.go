package main

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite" // The "testify/suite" package is used to make the test suite
)

type StoreSuite struct {
	suite.Suite /*
		The suite is defined as a struct, with the store and db as its
		attributes. Any variables that are to be shared between tests in a
		suite should be stored as attributes of the suite instance
	*/
	store *dbStore
	db    *sql.DB
}

func (s *StoreSuite) SetupSuite() {
	// The database connection is opened in the setup, and stored as an instance variable, as is the higher level `store`, that wraps the `db`
	connString := "dbname=bird_encyclopedia sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		s.T().Fatal(err)
	}
	s.db = db
	s.store = &dbStore{db: db}
}

func (s *StoreSuite) SetupTest() {
	_, err := s.db.Query("DELETE FROM birds") /*
		We delete all entries from the table before each test runs, to ensure a
		consistent state before our tests run. In more complex applications, this
		is sometimes achieved in the form of migrations
	*/
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *StoreSuite) TearDownSuite() {
	s.db.Close() // Close the connection after all tests in the suite finish
}

func TestStoreSuite(t *testing.T) { // This is the actual "test" as seen by Go, which runs the tests defined below
	s := new(StoreSuite)
	suite.Run(t, s)
}

func (s *StoreSuite) TestCreateBird() {
	s.store.CreateBird(&Bird{
		Description: "test description",
		Species:     "test species",
	})
	// Query the database for the entry we just created
	res, err := s.db.Query(`SELECT COUNT(*) FROM birds WHERE description='test description' AND SPECIES='test species'`)
	if err != nil {
		s.T().Fatal(err)
	}

	// Get the count result
	var count int
	for res.Next() {
		err := res.Scan(&count)
		if err != nil {
			s.T().Error(err)
		}
	}

	if count != 1 {
		s.T().Errorf("incorrect count,wanted 1,got %d", count)
	}
}

func (s *StoreSuite) TestGetBird() {
	_, err := s.db.Query(`INSERT INTO birds(species,description) VALUES('BIRD','DESCRIPTION')`)
	if err != nil {
		s.T().Fatal(err)
	}

	//Get the list of birds through the stores 'GetBirds' method
	birds, err := s.store.GetBirds()
	if err != nil {
		s.T().Fatal(err)
	}

	nBirds := len(birds)
	if nBirds != 1 {
		s.T().Errorf("incorrect count, wanted 1, got %d", nBirds)
	}

	expectedBird := Bird{"BIRD", "DESCRIPTION"}
	if *birds[0] != expectedBird {
		s.T().Errorf("incorrect details,expected %v got %v", expectedBird, *birds[0])
	}
}
