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

type Favorite struct {
	Id         int64 `json:"id,omitempty" gorm:"primary_key;AUTO_INCREMENT"`
	VideoId    int64 `json:"video_id" gorm:"index:vu,priority:1" ` //复合索引
	UserId     int64 `json:"user_id" gorm:"index:vu,priority:2" gorm:"index:ua,priority:3"`
	ActionType int32 `json:"action_type" gorm:"index:ua,priority:4"` //1-点赞，2-取消点赞
	DelateAt   bool  `json:"delate_at" gorm:"index"`
}

func GetFavoritesByUserId(userId int64) (affected int64, err error) {
	favorite := make([]*Favorite, 0)
	err = DB.Where("user_id = ? and action_type = ?", userId, 1).Find(&favorite).Error
	// 响应
	if err != nil {
		return 0, err
	}
	return int64(len(favorite)), nil
}
