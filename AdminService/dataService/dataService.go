package dataService

import (
	"AdminService/model"
	"errors"
	"gorm.io/gorm"
)

type DataService struct {
	db *gorm.DB
}

func NewDataService(db *gorm.DB) *DataService {
	return &DataService{db}
}

func (repo *DataService) CreateReview(review *model.Review) error {
	var newReview model.Review

	newReview.ClientUsername = review.ClientUsername
	newReview.RepairmanUsername = review.RepairmanUsername
	newReview.Text = review.Text

	retValue := repo.db.Table("reviews").Save(&newReview)

	return retValue.Error
}

func (repo *DataService) ReportReview(id uint64) error {
	var newReview model.Review

	result := repo.db.Table("reviews").Where("id = ?", id).First(&newReview)

	if result.Error != nil {
		return errors.New("Review cannot be found!")
	}

	newReview.Report = true

	retValue := repo.db.Table("reviews").Save(&newReview)

	return retValue.Error
}

func (repo *DataService) CreateReviewResponse(review *model.Review, id string) error {
	var newReview model.Review

	newReview.ClientUsername = review.ClientUsername
	newReview.RepairmanUsername = review.RepairmanUsername
	newReview.Text = review.Text
	newReview.ResponseId = id

	retValue := repo.db.Table("reviews").Save(&newReview)

	return retValue.Error
}

func (repo *DataService) DeleteReview(id uint64) error {
	var newReview model.Review

	result := repo.db.Table("reviews").Where("id = ?", id).First(&newReview)

	if result.Error != nil {
		return errors.New("Review cannot be found!")
	}

	newReview.Deleted = true

	retValue := repo.db.Table("reviews").Save(&newReview)

	return retValue.Error
}

func (repo *DataService) FindAllReviews(username string) []model.Review {
	var allReviews []model.Review
	repo.db.Table("reviews").Where("deleted = false AND client_username = ? ", username).Find(&allReviews)

	return allReviews
}

func (repo *DataService) GetAllReviews() []model.Review {
	var allReviews []model.Review
	repo.db.Table("reviews").Where("deleted = false").Find(&allReviews)

	return allReviews
}
