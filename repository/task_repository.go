package repository

import (
	"fmt"
	"go-rest-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITaskRepository interface {
	GetAllTasks(tasks *[]model.Task, userId uint) error
	GetTaskByID(tasks *model.Task, userId uint, taskId uint) error
	CreateTask(tasks *model.Task) error
	UpdateTask(tasks *model.Task, userId uint, taskId uint) error
	DeleteTask(userId uint, taskId uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

func (tr *taskRepository) GetAllTasks(task *[]model.Task, userId uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(task).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) GetTaskByID(task *model.Task, userId uint, taskId uint) error {
	if err := tr.db.Joins("User").Where("user_id=? AND id=?", userId, taskId).First(task).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) CreateTask(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) UpdateTask(task *model.Task, userId uint, taskId uint) error {
	// https://www.notion.so/Udemy-Study-golang-api-b4a9ef387623488eb9d07704ec7193dc?pvs=4#63422d25b4624707b81a6d5ce364bde8
	result := tr.db.Model(task).Clauses(clause.Returning{}).Where("user_id=? AND id=?", userId, taskId).Update("title", task.Title)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (tr *taskRepository) DeleteTask(userId uint, taskId uint) error {
	result := tr.db.Where("id=? AND user_id=?", taskId, userId).Delete(&model.Task{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}