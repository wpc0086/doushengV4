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
	"doushengV4/kitex_gen/interact"
	"doushengV4/kitex_gen/user"
)

// User pack user info
func User(u *user.User) *interact.User {
	if u == nil {
		return nil
	}

	return &interact.User{Id: u.Id, Name: u.Name, FavoriteCount: u.FollowCount,
		FollowerCount: u.FollowerCount, Avatar: "https://profile.csdnimg.cn/4/F/7/1_qq_41080854",
		Signature: "轻松拿下对不队", BackgroundImage: "https://img.1ppt.com/uploads/allimg/2302/1_230214151655_1.JPG",
		TotalFavorited: u.TotalFavorited, WorkCount: u.WorkCount, FollowCount: u.FollowCount, IsFollow: u.IsFollow}
}
