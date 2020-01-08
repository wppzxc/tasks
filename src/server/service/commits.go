package service

import (
	"fmt"
	"github.com/wppzxc/tasks/src/database"
	"strconv"
)

func GetCommitByID(id uint) (*database.Commits, error) {
	cmtID := strconv.Itoa(int(id))
	commit, err := database.GetCommitByID(cmtID)
	if err != nil {
		return nil, err
	}
	return commit, nil
}

func GetCommitsByName(username string, state string) ([]*database.Commits, error) {
	commits, err := database.GetCommitByName(username, state)
	if err != nil {
		return nil, err
	}
	return commits, nil
}

func GetCurrentCommit(username string, taskID string) (*database.Commits, error) {
	commit, err := database.GetCurrentCommits(username, taskID)
	if err != nil {
		return nil, err
	}
	return commit, nil
}

func CreateCommit(commit *database.Commits) (*database.Commits, error) {
	newCommit, err := database.CreateCommit(commit)
	if err != nil {
		return nil, err
	}
	return newCommit, nil
}

func UpdateCommits(commit *database.Commits) (*database.Commits, error) {
	newCommit, err := database.UpdateCommit(commit)
	if err != nil {
		return nil, err
	}
	return newCommit, nil
}

func AddBonus(taskID string, username string) error {
	task, err := GetTaskByID(taskID)
	if err != nil {
		return err
	}
	var users []*database.Users
	user, err := GetUserByName(username)
	if err != nil {
		return err
	}
	user.Balance = user.Balance + task.Bonus
	users = append(users, user)
	if len(user.Parent) > 0 {
		parent, err := GetUserByName(user.Parent)
		if err != nil {
			return err
		}
		parent.Balance = parent.Balance + task.ParentBonus
		users = append(users, parent)
		if len(parent.Parent) > 0 {
			grand, err := GetUserByName(parent.Parent)
			if err != nil {
				return err
			}
			grand.Balance = grand.Balance + task.GrandParentBonus
			users = append(users, grand)
		}
	}
	for _, u := range users {
		if _, err := UpdateUser(u); err != nil {
			fmt.Println(err)
		}
	}
	return nil
}
