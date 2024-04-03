package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type TrailEvent struct {
	ID             string                `json:"id"`
	Name           string                `json:"name"`
	Date           string                `json:"date"`
	Location       string                `json:"location"`
	PodcastEpisode string                `json:"podcastEpisode"`
	Distances      []*TrailEventDistance `json:"distances"`
}

type TrailEventDistance struct {
	ID           string  `json:"id,omitempty"`
	TrailEventID string  `json:"trailEventID,omitempty"`
	Distance     float64 `json:"distance"`
	GpxLink      string  `json:"gpxLink"`
}

func (s *APIServer) handleTrailEventsByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == methodGet {
		return s.handleGetTrailEventsByID(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetTrailEventsByID(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]

	review, err := s.store.GetTrailEventByID(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, review)
}

func (s *SQLLiteStore) GetTrailEventByID(id string) (*TrailEvent, error) {
	rows, err := s.db.Query(`select * from trail_events where id = $1`, id)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		trailEvent, err := mapToTrailEvent(rows)

		if err != nil {
			return nil, err
		}

		if trailEvent.Distances, err = queryDistances(s, trailEvent.ID); err != nil {
			return nil, err
		}

		return trailEvent, nil
	}

	return nil, fmt.Errorf("trailEvent with uuid [%s] not found", id)
}

func queryDistances(s *SQLLiteStore, id string) ([]*TrailEventDistance, error) {
	rows, err := s.db.Query("select id, distance, gpx_link from trail_event_distances where trail_event = ?", id)

	if err != nil {
		return nil, err
	}

	trailEventDistances := []*TrailEventDistance{}
	for rows.Next() {
		trailEventDistance, err := mapToTrailEventDistance(rows)

		if err != nil {
			return nil, err
		}

		trailEventDistances = append(trailEventDistances, trailEventDistance)
	}

	return trailEventDistances, nil
}

func mapToTrailEventDistance(rows *sql.Rows) (*TrailEventDistance, error) {
	trailEventDistance := new(TrailEventDistance)
	err := rows.Scan(
		&trailEventDistance.ID,
		&trailEventDistance.Distance,
		&trailEventDistance.GpxLink,
	)

	return trailEventDistance, err
}
func mapToTrailEvent(rows *sql.Rows) (*TrailEvent, error) {
	trailEvent := new(TrailEvent)
	err := rows.Scan(
		&trailEvent.ID,
		&trailEvent.Name,
		&trailEvent.Date,
		&trailEvent.Location,
		&trailEvent.PodcastEpisode,
	)

	return trailEvent, err
}
