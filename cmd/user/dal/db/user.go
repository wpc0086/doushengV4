// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package db

import (
	"doushengV4/pkg/consts"
)

type User struct {
	Id            int64  `json:"id,omitempty" gorm:"AUTO_INCREMENT" gorm:"primary_key"`
	Name          string `json:"name,omitempty"`
	Password      string `json:"password,omitempty"`
	FollowCount   int64  `json:"follow_count"` //添加omitempty，0值就不显示了
	FollowerCount int64  `json:"follower_count"`
}

func (u *User) TableName() string {
	return consts.UserTableName
}

/**
 * @Description 增加user
 * @Param
 * @return
 **/
func CreateUser(user *User) (err error) {
	err = DB.Create(&user).Error
	// 响应
	if err != nil {
		return err
	}
	return
}

/**
 * @Description 查询user 通过 名字
 * @Param
 * @return
 **/
func GetUserByName(uname string) (user []*User, err error) {
	user = make([]*User, 0)
	err = DB.Where("name = ?", uname).Find(&user).Error
	// 响应
	if err != nil {
		return nil, err
	}
	return user, nil
}

/**
 * @Description 查询user 通过 id
 * @Param
 * @return
 **/
func GetUserById(uid int64) (user *User, err error) {
	user = new(User)
	err = DB.First(&user, uid).Error
	// 响应
	if err != nil { // 未找到
		return user, err
	}
	return user, nil // 找到

}

/**
 * @Description 查询账号密码
 * @Param
 * @return
 **/
func CheckUser(uname string, pwd string) (user *User, err error) {
	user = new(User)
	err = DB.Where("name = ? AND password = ?", uname, pwd).Find(&user).Error
	// 响应
	if err != nil {
		return nil, err
	}
	return user, nil
}
