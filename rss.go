package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/eduncan911/podcast"
)

func (s *APIServer) GenerateRssFeed(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/xml")

	pubDate := time.Date(2016, time.December, 31, 0, 0, 0, 0, time.UTC)
	logo := "https://laufendentdecken-podcast.at/wp-content/uploads/2019/01/itunes-logo.jpg"
	updatedDate := time.Now()

	lep := podcast.New(
		"Laufend Entdecken",
		"https://laufendentdecken-podcast.at",
		"Der österreichische Laufpodcast",
		&pubDate,
		&updatedDate,
	)
	lep.AddAtomLink("https://laufendentdecken-podcast.at/feed/test")
	lep.AddImage(logo)
	lep.IExplicit = "no"

	episodes, _ := s.store.GetEpisodes()

	for _, episode := range *episodes {
		episodePubDate, _ := time.Parse("2006-01-02 15:04:05", episode.Date)
		name := episode.Name
		slug := episode.Slug

		item := podcast.Item{
			Title:       episode.Name,
			Link:        "https://rssfeed.laufendentdecken-podcast.at/data/" + slug + ".m4a",
			Description: "Test for Episode " + name,
			PubDate:     &episodePubDate,
		}
		item.PubDateFormatted = episode.Date
		item.AddImage(logo)
		item.AddSummary("Listen and give Feedback")
		item.AddEnclosure("https://rssfeed.laufendentdecken-podcast.at/data/"+slug+".m4a", podcast.M4A, episode.LengthInBytes)

		if _, err := lep.AddItem(item); err != nil {
			fmt.Println(item.Title, ": error", err.Error())
			return nil
		}
	}

	if err := lep.Encode(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return nil
}

type Episode struct {
	Name          string
	Slug          string
	Date          string
	LengthInBytes int64
}

func (s *SQLLiteStore) GetEpisodes() (*[]Episode, error) {
	response, err := s.db.Query("select name, slug, date, lengthInBytes from test_episodes")

	if err != nil {
		return nil, err
	}

	episodes := []Episode{}
	for response.Next() {
		episode, _ := mapToEpisode(response)
		episodes = append(episodes, *episode)

	}

	return &episodes, nil
}

func mapToEpisode(rows *sql.Rows) (*Episode, error) {
	episode := new(Episode)
	err := rows.Scan(
		&episode.Name,
		&episode.Slug,
		&episode.Date,
		&episode.LengthInBytes,
	)

	return episode, err
}
