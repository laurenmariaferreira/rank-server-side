package repository

import (
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/entity"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/framework/config"
	"github.coventry.ac.uk/340CT-1819SEPJAN/ferrei28-server-side/rank/util"
)

// Review defines the methods must be implemented by injected layer.
type Review interface {
	DeleteByID(util.Identifier) error
	FindAll() ([]*entity.Review, error)
	GetByID(util.Identifier) (*entity.Review, error)
	Store(*entity.Review) (util.Identifier, error)
	Update(*entity.Review) error
}

// DeleteByID deletes a Review by its ID.
func (m *MongoDB) DeleteByID(id util.Identifier) error {
	return m.pool.Session(nil).DB(m.db).C(config.REVIEW_COLLECTION).RemoveId(id)
}

// FindAll returns all Reviews from the database sorted by ID.
func (m *MongoDB) FindAll() ([]*entity.Review, error) {
	var reviews []*entity.Review

	session := m.pool.Session(nil)
	collection := session.DB(m.db).C(config.REVIEW_COLLECTION)
	if err := collection.Find(nil).Sort("id").All(&reviews); err != nil {
		return nil, err
	}

	return reviews, nil
}

// GetByID finds a Review by its ID.
func (m *MongoDB) GetByID(id util.Identifier) (*entity.Review, error) {
	session := m.pool.Session(nil)
	coll := session.DB(m.db).C(config.REVIEW_COLLECTION)

	var review *entity.Review

	coll.FindId(id).One(&review)

	return review, nil
}

// Store inserts a new Review in the database.
func (m *MongoDB) Store(review *entity.Review) (util.Identifier, error) {
	session := m.pool.Session(nil)
	coll := session.DB(m.db).C(config.REVIEW_COLLECTION)

	review.ID = util.NewID()

	coll.Insert(review)

	return review.ID, nil
}

// Update updates a new Review in the database.
func (m *MongoDB) Update(review *entity.Review) error {
	session := m.pool.Session(nil)
	coll := session.DB(m.db).C(config.REVIEW_COLLECTION)

	_, err := coll.UpsertId(review.ID, review) // TODO - avoid null Reviews
	return err
}
