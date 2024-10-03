package entity

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Year   int
	Size   int
	Rate   float64
}

func (b *Book) MarshalJSON() ([]byte, error) {
	err := validStruct(b.Year, b.Size, b.Rate)
	if err != nil {
		return nil, err
	}

	temp := struct {
		ID     int     `json:"id"`
		Title  string  `json:"title"`
		Author string  `json:"author"`
		Year   int     `json:"year"`
		Size   int     `json:"size"`
		Rate   float64 `json:"rate"`
	}{
		ID:     b.ID,
		Title:  b.Title,
		Author: b.Author,
		Year:   b.Year,
		Size:   b.Size,
		Rate:   b.Rate,
	}
	j, err := json.Marshal(temp)
	if err != nil {
		return nil, fmt.Errorf("book marshaling: %w", err)
	}

	return j, nil
}

func (b *Book) UnmarshalJSON(j []byte) error {
	if len(j) == 0 {
		return fmt.Errorf("cannot unmarshal empty JSON")
	}
	temp := struct {
		ID     int     `json:"id"`
		Title  string  `json:"title"`
		Author string  `json:"author"`
		Year   int     `json:"year"`
		Size   int     `json:"size"`
		Rate   float64 `json:"rate"`
	}{}

	err := json.Unmarshal(j, &temp)
	if err != nil {
		return fmt.Errorf("book unmarshalling: %w", err)
	}

	err = validStruct(temp.Year, temp.Size, temp.Rate)
	if err != nil {
		return err
	}

	b.ID = temp.ID
	b.Title = temp.Title
	b.Author = temp.Author
	b.Year = temp.Year
	b.Size = temp.Size
	b.Rate = temp.Rate

	return nil
}

func validStruct(year, size int, rate float64) error {
	if year <= 0 {
		return fmt.Errorf("invalid year: %d", year)
	}
	if size <= 0 {
		return fmt.Errorf("invalid size: %d", size)
	}
	if rate < 0 || rate > 5 {
		return fmt.Errorf("rate should be from 0 to 5: %f", rate)
	}
	return nil
}

// Фиг поймешь надо сериализовать весь слайс объектов или иметь слайс
// сериализованных объектов. На всякий случай пишу оба варианта.

// Сериализация и десериализация массива объектов в массив JSON объектов

func SliceToBytesMarshalJSON(s []*Book) ([][]byte, error) {
	sliceJSON := make([][]byte, 0, len(s))
	for _, v := range s {
		j, err := v.MarshalJSON()
		if err != nil {
			return nil, fmt.Errorf("marshal json slice: %w", err)
		}
		sliceJSON = append(sliceJSON, j)
	}
	return sliceJSON, nil
}

func SliceFromBytesUnmarshalJSON(s [][]byte) ([]*Book, error) {
	sliceBook := make([]*Book, 0, len(s))
	for _, v := range s {
		var book Book
		err := book.UnmarshalJSON(v)
		if err != nil {
			return nil, fmt.Errorf("unmarshal json slice: %w", err)
		}
		sliceBook = append(sliceBook, &book)
	}
	return sliceBook, nil
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
