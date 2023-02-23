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

package service

import (
	"context"
	"doushengV4/cmd/publish/dal/db"
	"doushengV4/cmd/publish/dal/redis"
	"doushengV4/cmd/publish/util"
	"doushengV4/kitex_gen/publish"
	"doushengV4/pkg/consts"
	"doushengV4/pkg/errno"
	"doushengV4/pkg/mw"
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

type ActionPublishService struct {
	ctx context.Context
}

func NewActionPublishService(ctx context.Context) *ActionPublishService {
	return &ActionPublishService{ctx: ctx}
}

// CreateNote create note info
func (s *ActionPublishService) ActionPulish(req *publish.ActionRequest) error {
	claims, ok := mw.CheckToken(req.Token)
	if ok != true {
		return errno.AuthorizationFailedErr
	}
	fileName := strconv.Itoa(int(uuid.New().ID())) + ".mp4"
	saveFile := filepath.Join(consts.SaveFilePlace, fileName)
	//发布视频时，把发布列表缓存清除 -- 延迟双删
	redis.RemovePublishList(s.ctx, claims.UserId)
	//保存到本地的临时文件
	err := SaveUploadedFile(req.Data, saveFile)
	if err != nil {
		return err
	}
	//多线程
	var wg sync.WaitGroup
	wg.Add(2)
	//1、上传到minio
	videoName := time.Now().Format("20060102") + "/" + strconv.FormatInt(claims.UserId, 10) + "/" + fileName //strconv:字符串和基本数据类型的转换
	go func(c context.Context, videoName string, saveFile string) {
		err = util.VideoToMinio(c, videoName, saveFile)
		if err != nil {
			return
		}
		defer wg.Done()
	}(s.ctx, videoName, saveFile)
	//2、获取第一帧图片
	split := strings.Split(fileName, ".")
	saveImg := filepath.Join(consts.SaveFilePlace, split[0]+".png")
	go func(videoPath, snapshotPath string) {
		_, err = util.GetSnapshot(saveFile, saveImg, 1)
		if err != nil {
			return
		}
		defer wg.Done()
	}(saveFile, saveImg)
	wg.Wait()
	//2.1、将图片上传到minio
	wg.Add(1)
	imageName := time.Now().Format("20060102") + "/" + strconv.FormatInt(claims.UserId, 10) + "/" + split[0] + ".png"
	go func(c context.Context, imageName string, saveImg string) {
		err = util.ImageToMinio(c, imageName, saveImg)
		if err != nil {
			return
		}
		defer wg.Done()
	}(s.ctx, imageName, saveImg)
	////保存到数据库
	video := db.Video{
		AuthorId:  claims.UserId,
		Title:     req.Title,
		PlayUrl:   consts.MinioVideoPrefex + videoName,
		CoverUrl:  consts.MinioVideoPrefex + imageName,
		CreatedAt: time.Now().Unix(),
	}
	err = db.SaveVideoAndImage(&video)
	if err != nil {
		return err
	}
	wg.Wait()
	//删除临时文件
	defer func(saveFile string, saveImg string) { //删除临时文件
		err := os.Remove(saveFile)
		if err != nil {
			return
		}
		err = os.Remove(saveImg)
		if err != nil {
			return
		}
	}(saveFile, saveImg)
	//判断用户info是否存在redis，有就更新
	userExistRedis, _ := redis.GetInfo(s.ctx, claims.UserId)
	if userExistRedis == true {
		redis.UpdateInfo(s.ctx, claims.UserId)
	}
	//将视频插入视频feed的zset中
	redis.SaveVideoAndImage(s.ctx, &video)
	//发布视频时，把发布列表缓存清除 -- 延迟双删
	redis.RemovePublishList(s.ctx, claims.UserId)
	return nil
}

func SaveUploadedFile(file []byte, dst string) error {
	err := os.WriteFile(dst, file, 0666)
	if err != nil {
		return err
	}
	return nil
}
