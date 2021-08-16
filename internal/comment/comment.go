package comment

import "github.com/jinzhu/gorm"

// Service – the struct for our comment service
type Service struct {
	DB *gorm.DB
}

// Comment – defines our comment structure
type Comment struct {
	gorm.Model
	Slug   string
	Body   string
	Author string
}

// CommentService – the interface for our comment service
type CommentService interface {
	GetComment(ID int) (Comment, error)
	GetCommentBySlug(slug string) ([]Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(ID int, newComment Comment) (Comment, error)
	DeleteComment(ID int) error
	GetAllComments() ([]Comment, error)
}

// NewService – returns a new comment Service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

// GetComment – retrieves comments by their ID from the database
func (s *Service) GetComment(ID uint) (Comment, error) {
	var comment Comment
	result := s.DB.First(&comment, ID)
	if result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

// GetCommentBySlug – retrieves all comments by slug
func (s *Service) GetCommentBySlug(slug string) ([]Comment, error) {
	var comments []Comment
	result := s.DB.Find(&comments).Where("slug = ?", slug)
	if result.Error != nil {
		return []Comment{}, result.Error
	}
	return comments, nil
}

// PostComment – adds a new comment to the database
func (s *Service) PostComment(comment Comment) (Comment, error) {
	result := s.DB.Save(&comment)
	if result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

// UpdateComment – updates comment in the database using ID
func (s *Service) UpdateComment(ID uint, newComment Comment) (Comment, error) {
	comment, err := s.GetComment(ID)
	if err != nil {
		return Comment{}, err
	}
	result := s.DB.Model(&comment).Updates(newComment)
	if result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

// DeleteComment – deletes a comment from the database by ID
func (s *Service) DeleteComment(ID uint) error {
	result := s.DB.Delete(&Comment{}, ID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllComments – gets all the comments from the database
func (s *Service) GetAllComments() ([]Comment, error) {
	var comments []Comment
	result := s.DB.Find(&comments)
	if result.Error != nil {
		return []Comment{}, result.Error
	}
	return comments, nil
}
