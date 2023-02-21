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
	"gorm.io/gorm"
)

type Favorite struct {
	Id         int64 `json:"id,omitempty" gorm:"primary_key;AUTO_INCREMENT"`
	VideoId    int64 `json:"video_id" gorm:"index:vu,priority:1" ` //复合索引
	UserId     int64 `json:"user_id" gorm:"index:vu,priority:2" gorm:"index:ua,priority:3"`
	ActionType int32 `json:"action_type" gorm:"index:ua,priority:4"` //1-点赞，2-取消点赞
	DelateAt   bool  `json:"delate_at" gorm:"index"`
}

func AddFavorite(videoId int64, userId int64, actionType int32, addNum int64) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		//1、操作favorite表
		var favorite Favorite
		//1.1、是否存在数据，判断要更新还是要创建
		tx.Debug().First(&favorite, "video_id = ? AND user_id = ?", videoId, userId)
		if favorite.Id == 0 {
			favorite.UserId = userId
			favorite.VideoId = videoId
			favorite.ActionType = actionType
		} else {
			favorite.ActionType = actionType
		}
		//1.2、保存
		err := tx.Debug().Save(&favorite).Error
		if err != nil {
			return err
		}
		//2、操作video表
		//GetFavoriteNumFromVideo
		video := Video{}
		err = tx.Debug().Find(&video, "id = ?", videoId).Error
		if err != nil {
			return err
		}
		favriteNum := int64(video.FavoriteCount)
		//UpdateFavoriteToVideo
		favoriteCount := favriteNum + addNum
		video = Video{}
		err = tx.Model(&video).Where("id = ?", videoId).Update("favorite_count", favoriteCount).Error
		if err != nil {
			return err
		}
		return nil
	})
}

func SaveOrUpdateFavorite(videoId int64, userId int64, actionType int32) (err error) {
	return DB.Transaction(func(tx *gorm.DB) error {
		//查询用户是否在数据库中，判断要更新还是要创建
		favorite := new(Favorite)
		err = tx.First(&favorite, "video_id = ? AND user_id = ?", videoId, userId).Error
		if err != nil {
			return err
		}
		if favorite != nil {
			favorite.ActionType = actionType
		} else {
			favorite = new(Favorite)
			favorite.UserId = userId
			favorite.VideoId = videoId
			favorite.ActionType = actionType
		}

		err = tx.Save(&favorite).Error
		if err != nil {
			return err
		}
		return nil
	})
}

func GetFavorite(videoId int64, userId int64) (favorite *Favorite, err error) {
	favorite = new(Favorite)
	err = DB.First(&favorite, "video_id = ? AND user_id = ?", videoId, userId).Error
	if err != nil {
		return nil, err
	}
	return favorite, nil
}

func GetVideosByUserId(userId int64) (videos []*Video, err error) {
	favorites := make([]*Favorite, 0)
	err = DB.Find(&favorites, "user_id = ? AND action_type = ?", userId, 1).Error
	if err != nil {
		return nil, err
	}
	videoIDs := []int64{}
	for _, f := range favorites {
		videoIDs = append(videoIDs, f.VideoId)
	}
	videos = make([]*Video, 0)
	for _, videoID := range videoIDs {
		var video *Video
		err := DB.First(&video, "id = ?", videoID).Error
		if err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}
	return videos, nil
}

func GetVideosIdByUserId(userId int64) (videoIDs []int64, err error) {
	favorites := make([]*Favorite, 0)
	//根据用户id并且点赞的
	err = DB.Find(&favorites, "user_id = ? AND action_type = ?", userId, 1).Error
	if err != nil {
		return nil, err
	}
	videoIDs = []int64{}
	for _, f := range favorites {
		videoIDs = append(videoIDs, f.VideoId)
	}
	return videoIDs, nil
}
func GetVideosByVideoIDs(videoIDs []int64) (videos []*Video, err error) {
	videos = make([]*Video, 0)
	for _, videoID := range videoIDs {
		var video *Video
		err := DB.Debug().First(&video, "id = ?", videoID).Error
		if err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}
	return videos, nil
}
