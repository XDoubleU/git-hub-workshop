package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"check-in/api/internal/config"
	"check-in/api/internal/database"
	"check-in/api/internal/models"
	"check-in/api/internal/services"
	"check-in/api/internal/tests"
)

type FixtureData struct {
	Notes         []*models.Note
	AmountOfNotes int
}

var mainTestEnv *tests.MainTestEnv //nolint:gochecknoglobals //global var for tests
var cfg config.Config              //nolint:gochecknoglobals //global var for tests
var logger *log.Logger             //nolint:gochecknoglobals //global var for tests
var fixtureData FixtureData        //nolint:gochecknoglobals //global var for tests

func clearAll(services services.Services) error {
	notes, err := services.Notes.GetAll(context.Background())
	if err != nil {
		return err
	}

	for _, note := range notes {
		err = services.Notes.Delete(context.Background(), note.ID)
		if err != nil {
			return err
		}
	}

	fixtureData.AmountOfNotes = 0

	return nil
}

func noteFixtures(services services.Services) error {
	for i := 0; i < 20; i++ {
		note, err := services.Notes.Create(context.Background(),
			fmt.Sprintf("TestNote%d", i), "Some text")
		if err != nil {
			return err
		}
		fixtureData.Notes = append(fixtureData.Notes, note)
		fixtureData.AmountOfNotes++
	}

	return nil
}

func fixtures(tx database.DB) {
	services := services.New(tx)

	err := clearAll(services)
	if err != nil {
		panic(err)
	}

	err = noteFixtures(services)
	if err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	var err error

	cfg = config.New()
	cfg.Env = config.TestEnv

	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)

	mainTestEnv, err = tests.SetupGlobal(
		cfg.DB.Dsn,
		cfg.DB.MaxConns,
		cfg.DB.MaxIdleTime,
	)
	if err != nil {
		panic(err)
	}

	fixtures(mainTestEnv.TestTx)

	exitCode := m.Run()
	err = tests.TeardownGlobal(mainTestEnv)
	if err != nil {
		panic(err)
	}

	os.Exit(exitCode)
}

func setupTest(
	_ *testing.T,
	mainTestEnv *tests.MainTestEnv,
) (tests.TestEnv, *application) {
	// t.Parallel()
	testEnv := tests.SetupSingle(mainTestEnv)

	testApp := &application{
		config:   cfg,
		logger:   logger,
		services: services.New(testEnv.TestTx),
	}

	return testEnv, testApp
}
