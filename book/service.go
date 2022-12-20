package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(book BooksRequest) (Book, error)
	Update(ID int, book BooksRequest) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	book, err := s.repository.FindAll()
	return book, err
}

func (s *service) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	return book, err
}

func (s *service) Create(bookRequest BooksRequest) (Book, error) {
	price, _ := bookRequest.Price.Int64()
	book := Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Description: bookRequest.Description,
		Rating:      bookRequest.Rating,
		Discount:    bookRequest.Discount,
	}

	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Update(ID int, bookRequest BooksRequest) (Book, error) {
	book, _ := s.repository.FindByID(ID)
	price, _ := bookRequest.Price.Int64()

	book.Title = bookRequest.Title
	book.Price = int(price)
	book.Description = bookRequest.Description
	book.Rating = bookRequest.Rating
	book.Discount = bookRequest.Discount

	newBook, err := s.repository.Update(book)
	return newBook, err
}
