package repository

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"
)

type inmenRepository struct {
	trips     map[string]*domain.TripModel
	rideFares map[string]*domain.RideFareModel
}

func NewInmenRepository() *inmenRepository {
	return &inmenRepository{
		trips:     make(map[string]*domain.TripModel),
		rideFares: make(map[string]*domain.RideFareModel),
	}
}

func (r *inmenRepository) CreateTrip(ctx context.Context, trip *domain.TripModel) (*domain.TripModel, error) {
	r.trips[trip.ID.Hex()] = trip
	return trip, nil
}
