package entity

import (
	"reflect"
	"testing"
)

func TestMarshalUnmarshalJSON(t *testing.T) {
	book := &Book{
		ID:     1,
		Title:  "Go Programming",
		Author: "John Doe",
		Year:   2022,
		Size:   320,
		Rate:   4.5,
	}

	// Marshal the book to JSON
	data, err := book.MarshalJSON()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Unmarshal the JSON back to a Book struct
	var result Book
	err = result.UnmarshalJSON(data)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Compare the original and unmarshalled books
	if !reflect.DeepEqual(book, &result) {
		t.Errorf("expected %v, got %v", book, &result)
	}
}

func TestSliceMarshalUnmarshalJSON(t *testing.T) {
	books := []*Book{
		{ID: 1, Title: "Go Programming", Author: "John Doe", Year: 2022, Size: 320, Rate: 4.5},
		{ID: 2, Title: "Python Basics", Author: "Jane Smith", Year: 2021, Size: 280, Rate: 4.8},
	}

	// Marshal slice of books to a single JSON array
	data, err := SliceMarshalJSON(books)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Unmarshal the JSON back to a slice of Book structs
	result, err := SliceUnmarshalJSON(data)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Compare the original and unmarshalled slices
	if !reflect.DeepEqual(books, result) {
		t.Errorf("expected %v, got %v", books, result)
	}
}

func TestSliceMarshalUnmarshalBytes(t *testing.T) {
	books := []*Book{
		{ID: 1, Title: "Go Programming", Author: "John Doe", Year: 2022, Size: 320, Rate: 4.5},
		{ID: 2, Title: "Python Basics", Author: "Jane Smith", Year: 2021, Size: 280, Rate: 4.8},
	}

	// Marshal each book in the slice to JSON bytes
	data, err := SliceToBytesMarshalJSON(books)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Unmarshal the JSON bytes back to a slice of Book structs
	result, err := SliceFromBytesUnmarshalJSON(data)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Compare the original and unmarshalled slices
	if !reflect.DeepEqual(books, result) {
		t.Errorf("expected %v, got %v", books, result)
	}
}
