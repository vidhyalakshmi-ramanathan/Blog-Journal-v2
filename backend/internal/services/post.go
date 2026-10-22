package services

import (
	"backend/internal/entities"
	"backend/internal/models"
	"errors"
)

type PostService interface {
	Create(post *entities.Post) error
	Update(postID, userID uint, post *entities.Post) error
	Delete(postID, id uint) error
	GetPost(id uint) (*entities.PostResponse, error)
	GetPosts(id uint) (map[uint]entities.PostResponse, error)
	GetAllPosts() (map[uint]entities.PostResponse, error)
}

type postService struct {
	post models.PostModel
}

func NewPostService(post models.PostModel) PostService {
	return &postService{post: post}
}

func (s *postService) Create(post *entities.Post) error {
	if post.Title == "" || post.Content == "" {
		return errors.New("both title and content are required")
	}
	err := s.post.Create(post)
	if err != nil {
		return err
	}
	return nil
}

func (s *postService) Update(postID, userID uint, post *entities.Post) error {
	if post.Title == "" || post.Content == "" {
		return errors.New("both title and content are required")
	}
	oldpost, err := s.post.GetPost(postID)
	if err != nil {
		return err
	}
	if oldpost.AuthorID != userID {
		return errors.New("you are not allowed to edit this post")
	}
	oldpost.Title = post.Title
	oldpost.Content = post.Content

	err = s.post.Update(oldpost)
	if err != nil {
		return err
	}
	return nil
}

func (s *postService) Delete(postID, userID uint) error {
	post, err := s.post.GetPost(postID)
	if err != nil {
		return err
	}

	if post.AuthorID != userID {
		return errors.New("you are not allowed to delete this post")
	}
	return s.post.Delete(postID)
}
func (s *postService) GetPost(id uint) (*entities.PostResponse, error) {
	post, err := s.post.GetPost(id)
	if err != nil {
		return nil, err
	}
	return &entities.PostResponse{
		ID:       post.ID,
		Title:    post.Title,
		Content:  post.Content,
		CreateAt: post.CreatedAt.String(),
		AuthorID: post.AuthorID,
	}, nil
}

func (s *postService) GetPosts(id uint) (map[uint]entities.PostResponse, error) {
	posts, err := s.post.GetPosts(id)
	if err != nil {
		return nil, err
	}

	allPosts := make(map[uint]entities.PostResponse)

	for _, post := range posts {
		allPosts[post.ID] = entities.PostResponse{
			ID:       post.ID,
			Title:    post.Title,
			Content:  post.Content,
			CreateAt: post.CreatedAt.String(),
			AuthorID: post.AuthorID,
		}
	}
	/* postID1: PostRes1, postID2: PostRes2 */

	return allPosts, nil
}

func (s *postService) GetAllPosts() (map[uint]entities.PostResponse, error) {
	posts, err := s.post.GetAllPosts()
	if err != nil {
		return nil, err
	}
	allPosts := make(map[uint]entities.PostResponse)

	for _, post := range posts {
		allPosts[post.ID] = entities.PostResponse{
			ID:       post.ID,
			Title:    post.Title,
			Content:  post.Content,
			CreateAt: post.CreatedAt.String(),
			AuthorID: post.AuthorID,
		}
	}
	return allPosts, nil
}
