package entity

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
