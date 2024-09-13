package entity

import "errors"

var ErrEntityNotFound = errors.New("entity not found")

type Compare int

const (
	CompareByYear Compare = iota
	CompareBySize
	CompareByRate
)

// BookComparator структура, наверное, не лучшее решение,
// лучше реализовать при помощи интерфесов, но мы пока их не проходили.
type BookComparator struct {
	comparator Compare
}

// NewBookComparator - конструктор для BookComparator.
func NewBookComparator(comparator Compare) *BookComparator {
	return &BookComparator{comparator}
}

// Compare позволяет сравнивать две структуры Book по заданному
// полю в зависимости от выбранного компаратора (сравнителя) в
// конструкторе структуры BookComparator.
func (bc BookComparator) Compare(b1, b2 Book) (bool, error) {
	if b1 == (Book{}) && b2 == (Book{}) {
		return false, ErrEntityNotFound
	}
	switch bc.comparator {
	case CompareByYear:
		return b1.year > b2.year, nil
	case CompareBySize:
		return b1.size > b2.size, nil
	case CompareByRate:
		return b1.rate > b2.rate, nil
	default:
		return false, nil
	}
}
