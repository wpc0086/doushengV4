package db

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	Id         int64  `json:"id,omitempty" gorm:"primary_key;AUTO_INCREMENT"` //评论ID
	VideoId    int64  `json:"video_id,omitempty" gorm:"index:vd,priority:3"`
	UserId     int64  `json:"user_id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
	CreateTime int64  `json:"create_time,omitempty" gorm:"index"`
	DelateAt   bool   `json:"delate_at" gorm:"index" gorm:"index:vd,priority:4"`
}

func SaveCommon(videoId int64, userId int64, content string) (comment *Comment, err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		//1、保存到评论表中
		comment = &Comment{UserId: userId, VideoId: videoId, Content: content,
			CreateDate: time.Now().Format("01") + "-" + time.Now().Format("02"),
			CreateTime: time.Now().Unix(),
			DelateAt:   false}
		err = tx.Save(&comment).Error
		if err != nil {
			return err
		}
		//2、操作video表
		video := Video{}
		err = tx.Find(&video, "id = ?", videoId).Error
		if err != nil {
			return err
		}
		commentNum := int64(video.CommentCount)
		commentCount := commentNum + 1
		video = Video{}
		err = tx.Model(&video).Where("id = ?", videoId).Update("comment_count", commentCount).Error
		if err != nil {
			return err
		}
		return nil
	})
	return comment, err
}

func DelCommonById(id int64, videoId int64) (err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		//1、到评论表中删除
		common := Comment{}
		//通过主键软性删除
		err = DB.Model(&common).Where("id = ?", id).Update("delate_at", true).Error
		if err != nil {
			return err
		}
		//2、操作video表
		video := Video{}
		err = tx.Find(&video, "id = ?", videoId).Error
		if err != nil {
			return err
		}
		commentNum := int64(video.CommentCount)
		commentCount := commentNum - 1
		video = Video{}
		err = tx.Model(&video).Where("id = ?", videoId).Update("comment_count", commentCount).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func GetCommonsByVideoID(videoId int64) (comments []*Comment, err error) {
	comments = make([]*Comment, 0)
	err = DB.Order("create_time desc").Find(&comments, "video_id = ? And delate_at = ?", videoId, false).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}
