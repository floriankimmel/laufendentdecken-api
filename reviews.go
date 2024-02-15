package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Link struct {
	Link    string `json:"link"`
	AltText string `json:"altText"`
}

type Review struct {
	ID             string  `json:"id"`
	ProductName    string  `json:"productName"`
	LinksToProduct []*Link `json:"linksToProduct"`
	Brand          string  `json:"brand"`
	Weight         float64 `json:"weight"`
	Price          float64 `json:"price"`
	PicturesLinks  []*Link `json:"pictureLinks"`
	PodcastEpisode string  `json:"podcastEpisode"`
	Rating         float64 `json:"rating"`
	Statement      string  `json:"statement"`
	Shoe           *Shoe   `json:"shoe"`
}

type Shoe struct {
	ID   string `json:"id,omitempty"`
	Drop int64  `json:"drop"`
	Grip string `json:"grip"`
	Sole string `json:"sole"`
}

func (s *APIServer) handleReviewsByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetReviewsByID(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetReviewsByID(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]

	review, err := s.store.GetReviewByID(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, review)
}

func (s *SQLLiteStore) GetReviewByID(reviewID string) (*Review, error) {
	reviews, err := s.db.Query(`
        select
            id,
            product_name,
            brand,
            weight,
            price,
            podcast_episode,
            rating,
            statement
        from reviews where id = $1`, reviewID)
	if err != nil {
		return nil, err
	}

	if reviews.Next() {
		review, err := mapToReview(reviews)

		if err != nil {
			return nil, err
		}

		if review.Shoe, err = queryShoe(s, reviewID); err != nil {
			return nil, err
		}

		if review.LinksToProduct, err = queryLinks(s, reviewID, "product_links"); err != nil {
			return nil, err
		}

		if review.PicturesLinks, err = queryLinks(s, reviewID, "review_pictures"); err != nil {
			return nil, err
		}
		return review, nil
	}

	return nil, fmt.Errorf("review with uuid [%s] not found", reviewID)
}

func queryShoe(s *SQLLiteStore, id string) (*Shoe, error) {
	shoes, err := s.db.Query("select shoe_drop, grip, sole from review_shoes where review = $1", id)
	if err != nil {
		return nil, err
	}

	if shoes.Next() {
		shoe, err := mapToShoe(shoes)

		if err != nil {
			return nil, err
		}

		return shoe, nil
	}
	return nil, fmt.Errorf("shoe with uuid [%s] not found", id)
}

func queryLinks(s *SQLLiteStore, id string, table string) ([]*Link, error) {
	query := fmt.Sprintf("select link, alt_text from " + table + " where review = ?")
	rows, err := s.db.Query(query, id)

	if err != nil {
		return nil, err
	}

	links := []*Link{}
	for rows.Next() {
		link, err := mapToLink(rows)

		if err != nil {
			return nil, err
		}

		links = append(links, link)
	}

	return links, nil
}

func mapToReview(rows *sql.Rows) (*Review, error) {
	review := new(Review)
	err := rows.Scan(
		&review.ID,
		&review.ProductName,
		&review.Brand,
		&review.Weight,
		&review.Price,
		&review.PodcastEpisode,
		&review.Rating,
		&review.Statement,
	)

	return review, err
}

func mapToShoe(rows *sql.Rows) (*Shoe, error) {
	shoe := new(Shoe)
	err := rows.Scan(
		&shoe.Drop,
		&shoe.Grip,
		&shoe.Sole,
	)

	return shoe, err
}

func mapToLink(rows *sql.Rows) (*Link, error) {
	link := new(Link)
	err := rows.Scan(
		&link.Link,
		&link.AltText,
	)

	return link, err
}
