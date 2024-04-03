package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type DopingBan struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Sport    string `json:"sport"`
	Ban      string `json:"ban"`
	BanStart string `json:"banStart"`
	BanEnd   string `json:"banEnd"`
	Reason   string `json:"reason"`
}

func (s *APIServer) handleDopingBans(w http.ResponseWriter, r *http.Request) error {
	if r.Method == methodGet {
		dopingBans, err := s.store.GetDopingBans()
		if err != nil {
			return err
		}

		return WriteJSON(w, http.StatusOK, dopingBans)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleDopingBanByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == methodGet {
		return s.handleGetDopingBanByID(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetDopingBanByID(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]

	dopingBan, err := s.store.GetDopingBanByID(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, dopingBan)
}

func (s *SQLLiteStore) GetDopingBans() (*[]DopingBan, error) {
	doingBans, err := s.db.Query(`
        select
            id,
            name,
            sport,
            ban,
            ban_start,
            ban_end,
            reason
        from doping_bans`)
	if err != nil {
		return nil, err
	}
	dopingBanArray := []DopingBan{}

	for doingBans.Next() {
		dopingBan, err := mapToDopingBan(doingBans)

		if err != nil {
			return nil, err
		}

		dopingBanArray = append(dopingBanArray, *dopingBan)
	}

	return &dopingBanArray, nil
}
func (s *SQLLiteStore) GetDopingBanByID(dopingBanID string) (*DopingBan, error) {
	doingBans, err := s.db.Query(`
        select
            id,
            name,
            sport,
            ban,
            ban_start,
            ban_end,
            reason
        from doping_bans where id = $1`, dopingBanID)
	if err != nil {
		return nil, err
	}

	if doingBans.Next() {
		dopingBan, err := mapToDopingBan(doingBans)

		if err != nil {
			return nil, err
		}

		return dopingBan, nil
	}

	return nil, fmt.Errorf("doping ban with uuid [%s] not found", dopingBanID)
}

func mapToDopingBan(rows *sql.Rows) (*DopingBan, error) {
	dopenBan := new(DopingBan)
	err := rows.Scan(
		&dopenBan.ID,
		&dopenBan.Name,
		&dopenBan.Sport,
		&dopenBan.Ban,
		&dopenBan.BanStart,
		&dopenBan.BanEnd,
		&dopenBan.Reason,
	)

	return dopenBan, err
}
