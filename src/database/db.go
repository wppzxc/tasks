package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
)

var DB *gorm.DB

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Users{})
	db.AutoMigrate(&Tasks{})
	db.AutoMigrate(&Commits{})
	db.AutoMigrate(&ServiceInfo{})
}

func GetCurrentServiceInfo() (*ServiceInfo, error) {
	svc := &ServiceInfo{}
	if err := DB.Last(svc).Error; err != nil {
		return nil, err
	}
	return svc, nil
}

func UpdateServiceInfo(svc *ServiceInfo) (*ServiceInfo, error) {
	if err := DB.Model(svc).Update(svc).Error; err != nil {
		return nil, err
	}
	newSvc, err := GetCurrentServiceInfo()
	if err != nil {
		return nil, err
	}
	return newSvc, nil
}

func GetUser(username string) (*Users, error) {
	user := &Users{}
	if err := DB.Where("name = ?", username).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByNameAndPassword(username, password string) (*Users, error) {
	user := &Users{}
	if err := DB.Where("name = ? and password = ?", username, password).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(user *Users) (*Users, error) {
	tx := DB.Begin()
	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	newUser, err := GetUser(user.Name)
	if err != nil {
		return nil, fmt.Errorf("Error in get user : %s, '%s' ", user.Name, err)
	}
	return newUser, nil
}

func UpdateUser(user *Users) (*Users, error) {
	if err := DB.Model(user).Where("name = ?", user.Name).Updates(user).Error; err != nil {
		return nil, err
	}
	newUser, err := GetUser(user.Name)
	if err != nil {
		return nil, fmt.Errorf("Error in get user : %s, '%s' ", user.Name, err)
	}
	return newUser, nil
}

func UpdateUsers(users []*Users) error {
	tx := DB.Begin()
	var err error
	for _, u := range users {
		if _, err = UpdateUser(u); err != nil {
			err = fmt.Errorf("%s, Error in update user : %s", err, u.Name)
		}
	}
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func GetCurrentTasks() (*Tasks, error) {
	task := &Tasks{}
	if err := DB.Last(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func GetTaskByID(taskID string) (*Tasks, error) {
	task := &Tasks{}
	if err := DB.Where("id = ?", taskID).First(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func CreateTask(task *Tasks) (*Tasks, error) {
	tx := DB.Begin()
	if err := tx.Create(task).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	newTask, err := GetCurrentTasks()
	if err != nil {
		return nil, fmt.Errorf("Error in get task : %d, '%s' ", task.ID, err)
	}
	return newTask, nil
}

func UpdateTask(task *Tasks) (*Tasks, error) {
	if err := DB.Model(task).Where("id = ?", task.ID).Updates(task).Error; err != nil {
		return nil, err
	}
	newTask, err := GetTaskByID(strconv.Itoa(int(task.ID)))
	if err != nil {
		return nil, fmt.Errorf("Error in get task : %d, '%s' ", task.ID, err)
	}
	return newTask, nil
}

func GetCommitByName(username string, state string) ([]*Commits, error) {
	var commits []*Commits
	var err error
	sql := DB.Order("created_at desc")
	if len(state) != 0 {
		sql = sql.Where("status = ?", state)
	}
	if len(username) == 0 {
		err = sql.Find(commits).Error
	} else {
		err = sql.Where("username = ?", username).Find(&commits).Error
	}
	if err != nil {
		return nil, err
	}
	return commits, nil
}

func GetCurrentCommits(username string, taskID string) (*Commits, error) {
	commit := &Commits{}
	if err := DB.Where("username = ? and task_id = ?", username, taskID).Last(commit).Error; err != nil {
		return nil, err
	}
	return commit, nil
}

func GetCommitByID(commitID string) (*Commits, error) {
	commit := &Commits{}
	if err := DB.Where("id = ?", commitID).First(commit).Error; err != nil {
		return nil, err
	}
	return commit, nil
}

func CreateCommit(commit *Commits) (*Commits, error) {
	tx := DB.Begin()
	if err := tx.Create(commit).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	newCommit, err := GetCurrentCommits(commit.Username, strconv.Itoa(int(commit.TaskID)))
	if err != nil {
		return nil, fmt.Errorf("Error in get commit : %d, '%s' ", commit.TaskID, err)
	}
	return newCommit, nil
}

func UpdateCommit(commit *Commits) (*Commits, error) {
	if err := DB.Model(commit).Where("id = ?", commit.ID).Updates(commit).Error; err != nil {
		return nil, err
	}
	newCommit, err := GetCommitByID(strconv.Itoa(int(commit.ID)))
	if err != nil {
		return nil, err
	}
	return newCommit, nil
}
