package models

import (
	"backend/internal/entities"

	"gorm.io/gorm"
)

type PostModel interface {
	Create(post *entities.Post) error
	Update(post *entities.Post) error
	Delete(id uint) error
	GetPost(id uint) (*entities.Post, error)
	GetPosts(id uint) ([]entities.Post, error)
	GetAllPosts() ([]entities.Post, error)
}

type postModel struct {
	db *gorm.DB
}

func NewPostModel(db *gorm.DB) PostModel {
	return &postModel{db: db}
}

func (r *postModel) Create(post *entities.Post) error {
	return r.db.Preload("Author").Create(post).Error
}

func (r *postModel) Update(post *entities.Post) error {
	return r.db.Save(post).Error
}

func (r *postModel) Delete(id uint) error {
	return r.db.Delete(&entities.Post{}, id).Error
}

func (r *postModel) GetPost(id uint) (*entities.Post, error) {
	var post entities.Post
	err := r.db.Preload("Author").Find(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postModel) GetPosts(id uint) ([]entities.Post, error) {
	var posts []entities.Post
	err := r.db.Preload("Author").Where("author_id = ?", id).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *postModel) GetAllPosts() ([]entities.Post, error) {
	var posts []entities.Post
	err := r.db.Preload("Author").Find(&posts).Error
	return posts, err
}
