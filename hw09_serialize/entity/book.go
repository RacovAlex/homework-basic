package entity

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   int     `json:"year"`
	Size   int     `json:"size"`
	Rate   float64 `json:"rate"`
}

func (b *Book) MarshalJSON() ([]byte, error) {
	err := validateBook(b)
	if err != nil {
		return nil, err
	}

	return json.Marshal(*b)
}

func (b *Book) UnmarshalJSON(data []byte) error {
	type Alias Book
	temp := &Alias{}

	err := json.Unmarshal(data, temp)
	if err != nil {
		return fmt.Errorf("book unmarshalling: %w", err)
	}

	*b = Book(*temp)

	err = validateBook(b)
	if err != nil {
		return err
	}

	return nil
}

func validateBook(b *Book) error {
	if b.Year <= 0 {
		return fmt.Errorf("invalid year: %d", b.Year)
	}
	if b.Size <= 0 {
		return fmt.Errorf("invalid size: %d", b.Size)
	}
	if b.Rate < 0 || b.Rate > 5 {
		return fmt.Errorf("rate should be from 0 to 5: %f", b.Rate)
	}
	return nil
}

// Сериализация и десериализация в единый JSON массив

func SliceMarshalJSON(s []*Book) ([]byte, error) {
	return json.Marshal(s)
}

func SliceUnmarshalJSON(data []byte) ([]*Book, error) {
	var books []*Book
	err := json.Unmarshal(data, &books)
	if err != nil {
		return nil, fmt.Errorf("unmarshal json slice: %w", err)
	}
	return books, nil
}
