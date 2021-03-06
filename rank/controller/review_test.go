package controller

import (
	"log"
	"testing"

	"github.com/juju/mgosession"
	"github.com/stretchr/testify/assert"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/entity"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/middlewares/config"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/repository"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/util"
	mgo "gopkg.in/mgo.v2"
)

func TestFindAllReviews(t *testing.T) {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	repo := repository.New(pool, config.MONGODB_DATABASE)

	controller := newReviewController(repo)

	pool.Session(nil).DB(config.MONGODB_DATABASE).C(config.REVIEW_COLLECTION).RemoveAll(nil)

	r1 := &entity.Review{
		Title: "Title 1",
	}

	controller.StoreReview(r1)

	t.Run("should return inserted review with 'Title 1' as title", func(t *testing.T) {
		reviews, err := controller.FindAllReviews()
		assert.Nil(t, err)
		assert.Equal(t, 1, len(reviews))
		assert.Equal(t, "Title 1", reviews[0].Title)
	})
}

func TestFindAllUnpublishedReviews(t *testing.T) {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	repo := repository.New(pool, config.MONGODB_DATABASE)

	controller := newReviewController(repo)

	pool.Session(nil).DB(config.MONGODB_DATABASE).C(config.REVIEW_COLLECTION).RemoveAll(nil)

	r1 := &entity.Review{
		Title: "Title 1",
	}

	controller.StoreReview(r1)

	t.Run("should return inserted review with 'Title 1' as title and IsPublished as false", func(t *testing.T) {
		reviews, err := controller.FindAllUnpublishedReviews()
		assert.Nil(t, err)
		assert.Equal(t, 1, len(reviews))
		assert.Equal(t, "Title 1", reviews[0].Title)
		assert.False(t, reviews[0].IsPublished)
	})
}

func TestStoreReview(t *testing.T) {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	repo := repository.New(pool, config.MONGODB_DATABASE)

	controller := newReviewController(repo)

	pool.Session(nil).DB(config.MONGODB_DATABASE).C(config.REVIEW_COLLECTION).RemoveAll(nil)

	r1 := &entity.Review{
		Title: "Title 1",
	}

	t.Run("should return inserted ID", func(t *testing.T) {
		id, _ := controller.StoreReview(r1) // TODO
		assert.Equal(t, true, util.IsValidID(id.String()))
	})

	t.Run("should have inserted new review", func(t *testing.T) {
		reviews, errFindAll := controller.FindAllReviews()
		assert.Nil(t, errFindAll)
		assert.Equal(t, 1, len(reviews))
	})
}

func TestGetByIDReview(t *testing.T) {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	repo := repository.New(pool, config.MONGODB_DATABASE)

	controller := newReviewController(repo)

	t.Run("should return Review from inserted ID", func(t *testing.T) {
		r1 := &entity.Review{
			Title: "Title Test",
		}

		id, _ := controller.StoreReview(r1) // TODO
		assert.Equal(t, true, util.IsValidID(id.String()))

		review, err := controller.GetReviewByID(id)
		assert.Equal(t, true, util.IsValidID(review.ID.String()))
		assert.Equal(t, review.Title, "Title Test")
		assert.Nil(t, err)
	})
}

func TestDeleteByIDReview(t *testing.T) {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	repo := repository.New(pool, config.MONGODB_DATABASE)

	controller := newReviewController(repo)

	t.Run("should delete Review from inserted ID", func(t *testing.T) {
		r1 := &entity.Review{
			Title: "Title Test",
		}

		id, _ := controller.StoreReview(r1) // TODO

		err := controller.DeleteReviewByID(id)
		assert.Nil(t, err)

		review, errGetByID := controller.GetReviewByID(id)
		assert.Nil(t, review)
		assert.Nil(t, errGetByID)
	})
}

func TestUpdateReview(t *testing.T) {
	session, err := mgo.Dial(config.MONGODB_HOST)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()

	pool := mgosession.NewPool(nil, session, config.MONGODB_CONNECTION_POOL)
	defer pool.Close()

	repo := repository.New(pool, config.MONGODB_DATABASE)

	controller := newReviewController(repo)

	pool.Session(nil).DB(config.MONGODB_DATABASE).C(config.REVIEW_COLLECTION).RemoveAll(nil)

	t.Run("should update Review title", func(t *testing.T) {

		r1 := &entity.Review{
			Title: "Title 1",
		}

		id, _ := controller.StoreReview(r1) // TODO

		review, errGetByID := controller.GetReviewByID(id)
		assert.Nil(t, errGetByID)
		assert.Equal(t, "Title 1", review.Title)

		review.Title = "Different title"
		err := controller.UpdateReview(review)
		assert.Nil(t, err)

		updatedReview, errGetByID2 := controller.GetReviewByID(id)
		assert.Nil(t, errGetByID2)
		assert.Equal(t, "Different title", updatedReview.Title)
	})
}
