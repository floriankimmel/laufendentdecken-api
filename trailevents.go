package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type TrailEvent struct {
	ID string `json:"id"`
}

func (s *APIServer) handleTrailEventsByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
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
	rows, err := s.db.Query("select * from trail_event where id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return mapToTrailEvent(rows)
	}

	return nil, fmt.Errorf("trailEvent with uuid [%s] not found", id)
}

func mapToTrailEvent(rows *sql.Rows) (*TrailEvent, error) {
	trailEvent := new(TrailEvent)
	err := rows.Scan(
		&trailEvent.ID,
	)

	return trailEvent, err
}
