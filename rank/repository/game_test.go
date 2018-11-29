package repository

import (
	"log"
	"testing"

	"github.com/juju/mgosession"
	"github.com/stretchr/testify/assert"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/entity"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/framework/config"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/util"
	mgo "gopkg.in/mgo.v2"
)

func TestFindAllGames(t *testing.T) {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	m := New(pool, config.MONGODB_DATABASE)

	t.Run("should have returned all games", func(t *testing.T) {

		pool.Session(nil).DB(config.MONGODB_DATABASE).C(config.GAME_COLLECTION).RemoveAll(nil)

		name := "Game 1"

		g1 := &entity.Game{
			Name: name,
		}

		m.StoreGame(g1)
		games, err := m.FindAllGames()
		assert.Nil(t, err)
		assert.Equal(t, 1, len(games))
		assert.Equal(t, name, games[0].Name)
	})

	t.Run("should have returned error", func(t *testing.T) {
		m = New(pool, "otherdatabase")
		games, err := m.FindAllGames()
		assert.NotNil(t, err)
		assert.Nil(t, games)
	})
}

func TestFindGameByID(t *testing.T) {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	m := New(pool, config.MONGODB_DATABASE)

	t.Run("should find certain Game by stored ID", func(t *testing.T) {

		name := "Game Name"

		g1 := &entity.Game{
			Name: name,
		}

		// TODO
		id, _ := m.StoreGame(g1)

		game, err := m.GetGameByID(id)
		assert.Equal(t, name, game.Name)
		assert.Nil(t, err)
		assert.NotNil(t, id)
		assert.Equal(t, id, game.ID)
		assert.Equal(t, true, util.IsValidID(id.String()))
		assert.Equal(t, true, util.IsValidID(game.ID.String()))
	})
}

func TestDeleteGameByID(t *testing.T) {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	m := New(pool, config.MONGODB_DATABASE)

	t.Run("should delete certain Game by stored ID", func(t *testing.T) {

		name := "Game Name"

		g1 := &entity.Game{
			Name: name,
		}

		// TODO
		id, _ := m.StoreGame(g1)

		game, errGetByID := m.GetGameByID(id)
		assert.Equal(t, id, game.ID)
		assert.Nil(t, errGetByID)

		err := m.DeleteGameByID(id)
		assert.Nil(t, err)

		game, errGetByID2 := m.GetGameByID(id)
		assert.Nil(t, game)
		assert.Nil(t, errGetByID2)
	})
}

func TestUpdateGame(t *testing.T) {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	m := New(pool, config.MONGODB_DATABASE)

	t.Run("should have updated a new game", func(t *testing.T) {
		pool.Session(nil).DB(config.MONGODB_DATABASE).C(config.GAME_COLLECTION).RemoveAll(nil)

		name := "Game Name"

		g1 := &entity.Game{
			Name: name,
		}

		// TODO
		id, err := m.StoreGame(g1)
		assert.Nil(t, err)

		game, errGetByID := m.GetGameByID(id)
		assert.Nil(t, errGetByID)
		assert.Equal(t, name, game.Name)

		differentName := "Different name"

		game.Name = differentName
		errUpdate := m.UpdateGame(game)
		assert.Nil(t, errUpdate)

		updatedGame, errGetByID2 := m.GetGameByID(id)

		assert.Nil(t, errGetByID2)
		assert.Equal(t, differentName, updatedGame.Name)
	})
}
