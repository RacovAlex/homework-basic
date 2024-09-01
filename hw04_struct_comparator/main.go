package main

import "fmt"

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float32
}

// NewBook является конструктором для Book.
func NewBook(id int, title string, author string, year int, size int, rate float32) *Book {
	return &Book{
		id:     id,
		title:  title,
		author: author,
		year:   year,
		size:   size,
		rate:   rate,
	}
}

type BookBuilder struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float32
}

func NewBookBuilder() *BookBuilder {
	return &BookBuilder{}
}

func (b *BookBuilder) SetID(id int) *BookBuilder {
	b.id = id
	return b
}

func (b *BookBuilder) SetTitle(title string) *BookBuilder {
	b.title = title
	return b
}

func (b *BookBuilder) SetAuthor(author string) *BookBuilder {
	b.author = author
	return b
}

func (b *BookBuilder) SetYear(year int) *BookBuilder {
	b.year = year
	return b
}

func (b *BookBuilder) SetSize(size int) *BookBuilder {
	b.size = size
	return b
}

func (b *BookBuilder) SetRate(rate float32) *BookBuilder {
	b.rate = rate
	return b
}

func (b *BookBuilder) Build() *Book {
	return NewBook(b.id, b.title, b.author, b.year, b.size, b.rate)
}

// режим сравнения для компаратора

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
func (bc BookComparator) Compare(b1, b2 Book) bool {
	switch bc.comparator {
	case CompareByYear:
		return b1.year > b2.year
	case CompareBySize:
		return b1.size > b2.size
	case CompareByRate:
		return b1.rate > b2.rate
	default:
		return false
	}
}

// геттеры для полей Book:

func (b *Book) GetID() int {
	return b.id
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) GetAuthor() string {
	return b.author
}

func (b *Book) GetYear() int {
	return b.year
}

func (b *Book) GetSize() int {
	return b.size
}

func (b *Book) GetRate() float32 {
	return b.rate
}

// сеттеры для полей Book:

func (b *Book) SetID(id int) {
	b.id = id
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b *Book) SetRate(rate float32) {
	b.rate = rate
}

func main() {
	AnnaKarenina := NewBook(
		1,
		"Анна Каренина",
		"Лев Толстой",
		1878,
		1124,
		4.75)

	MartinIden := Book{
		id:     2,
		title:  "Мартин Иден",
		author: "Джек Лондон",
		year:   1909,
		size:   448,
		rate:   4.60,
	}

	bookComparator := NewBookComparator(CompareByRate)
	fmt.Println(bookComparator.Compare(*AnnaKarenina, MartinIden))

	fmt.Println(NewBookComparator(CompareByYear).Compare(*AnnaKarenina, MartinIden))

	ebook := NewBookBuilder().
		SetID(3).
		SetAuthor("Илья Чепкин").
		SetRate(5.0).
		SetTitle("Как программировать на Go").
		Build()

	fmt.Printf("Книга: %+v \n", *ebook)
}
