package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Storage interface {
	GetReviewByID(id string) (*Review, error)
	GetTrailEventByID(id string) (*TrailEvent, error)
	GetEpisodes() (*[]Episode, error)
}

type SQLLiteStore struct {
	db *sql.DB
}

func NewSQLLiteStore() (*SQLLiteStore, error) {
	db, err := sql.Open("sqlite3", "laufendentdeckendb.db")
	if err != nil {
		log.Panic(err)
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &SQLLiteStore{
		db: db,
	}, nil
}

func (s *SQLLiteStore) Init() error {
	log.Println("Initializing database")

	query := `create table if not exists product_links (
		id varchar(500) primary key,
        review varchar(500),
        link varchar(500),
        alt_text varchar(500),
        foreign key (review) references reviews(id)
    );`

	if _, err := s.db.Exec(query); err != nil {
		return err
	}

	query = `create table if not exists review_pictures (
		id varchar(500) primary key,
        review varchar(500),
        link varchar(500),
        alt_text varchar(500),
        foreign key (review) references reviews(id)
    );`

	if _, err := s.db.Exec(query); err != nil {
		return err
	}

	query = `create table if not exists review_shoes (
		id varchar(500) primary key,
        review varchar(500),
        shoe_drop number,
        grip varchar(500),
        sole varchar(500),
        foreign key (review) references reviews(id)
    );`

	if _, err := s.db.Exec(query); err != nil {
		return err
	}

	query = `create table if not exists reviews (
		id varchar(500) primary key,
        product_name varchar(500),
        brand varchar(500),
        weight number,
        price number,
        podcast_episode varchar(500),
        rating number,
        statement varchar(500)
    );`

	if _, err := s.db.Exec(query); err != nil {
		return err
	}

	query = `create table if not exists trail_event_distances (
		id varchar(500) primary key,
        trail_event varchar(500),
        distance number,
        gpx_link varchar(500),
        foreign key (trail_event) references trail_events(id)
    );`

	if _, err := s.db.Exec(query); err != nil {
		return err
	}

	query = `create table if not exists trail_events (
		id varchar(500) primary key,
        name varchar(500),
        date varchar(500),
        location varchar(500),
        podcast_episode varchar(500)
    );`

	if _, err := s.db.Exec(query); err != nil {
		return err
	}

	query = `create table if not exists test_episodes (
		id varchar(500) primary key,
        name varchar(500),
        slug varchar(500),
        date varchar(500),
        lengthInBytes number
    );`

	if _, err := s.db.Exec(query); err != nil {
		return err
	}

	return nil
}
