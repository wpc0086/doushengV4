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

package pack

import (
	"doushengV4/cmd/publish/dal/db"
	"doushengV4/kitex_gen/publish"
)

// Publish pack Publish info
func Publish(v *db.Video) *publish.Video {
	if v == nil {
		return nil
	}

	return &publish.Video{Id: v.Id, FavoriteCount: int64(v.FavoriteCount),
		CommentCount: int64(v.CommentCount), PlayUrl: v.PlayUrl,
		CoverUrl: v.CoverUrl, Title: v.Title}
}

// Publishs pack list of Publish info
func Publishs(us []*db.Video) []*publish.Video {
	Publishs := make([]*publish.Video, 0)
	for _, v := range us {
		if temp := Publish(v); temp != nil {
			Publishs = append(Publishs, temp)
		}
	}
	return Publishs
}
