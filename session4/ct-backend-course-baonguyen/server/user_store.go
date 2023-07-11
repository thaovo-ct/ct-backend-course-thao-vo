package main

import (
	"sync"
	"fmt"
	)

// TODO #1: implement in-memory user store

var userStore = NewUserStore()

func NewUserStore() *UserStore {
	return &UserStore{data: make(map[string]UserInfo)}
}

type UserStore struct {
	mu   sync.Mutex
	data map[string]UserInfo
}

func (u *UserStore) Save(info UserInfo) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	_, exist := u.data[info.Username]
	if exist {
		return fmt.Errorf("User exist")
	}
	u.data[info.Username] = info
	return nil
}

func (u *UserStore) Get(username string) (UserInfo, error) {
	info, exist := u.data[username]
	if exist {
		return info, nil
	} 
	return UserInfo{}, fmt.Errorf("User not exist")
}

type UserInfo struct {
	// TODO: implement me
	Address string `json:"address"`
	Username string `json:"username"`
	Password string `json:"password"`
	Fullname string `json:"full_name"`
}