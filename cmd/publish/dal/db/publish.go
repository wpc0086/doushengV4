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

type Video struct {
	Id            int64  `json:"id,omitempty" gorm:"AUTO_INCREMENT;primary_key"`
	AuthorId      int64  `json:"author_id,omitempty" gorm:"index"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount uint   `json:"favorite_count,omitempty"`
	CommentCount  uint   `json:"comment_count,omitempty"`
	Title         string `json:"title,omitempty"`
	CreatedAt     int64  `json:"created_at" gorm:"index"`
}

func SaveVideoAndImage(vedio *Video) (err error) {
	err = DB.Create(&vedio).Error
	// 响应
	if err != nil {
		return err
	}
	return
}

func GetVideoListByUserId(uid int64) (videos []*Video, err error) {
	err = DB.Find(&videos, "author_id = ?", uid).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func GetFeed(lastTime int64) (feedList []*Video, err error) {
	//获取视频
	err = DB.Order("created_at DESC").Where("created_at < ?", lastTime).Limit(30).Find(&feedList).Error
	if err != nil {
		return nil, err
	}
	return feedList, nil
}
