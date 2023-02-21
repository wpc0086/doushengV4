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

package consts

const (
	UserTableName       = "user"
	SecretKey           = "byte dance 11111 return"
	UserServiceName     = "user"
	VideoMaxSize        = 20000000 //20M
	PublishServiceName  = "publish"
	InterActServiceName = "interact"
	ApiServiceName      = "api"
	MySQLDefaultDSN     = "root:root@tcp(172.30.154.234:3306)/dousheng?charset=utf8&parseTime=True&loc=Local"
	TCP                 = "tcp"
	UserServiceAddr     = "172.30.154.43:9010"
	PublishServiceAddr  = "172.30.154.43:9011"
	InterActServiceAddr = "172.30.154.43:9012"
	ETCDAddress         = "172.30.154.234:2379"
	EndPoint            = "172.30.154.234:9000"
	AccessKeyID         = "minioadmin"
	SecretAccessKey     = "minioadmin"
	UseSSL              = false
	BucketName          = "video"
	VideoContentType    = "video/mp4"
	ImageContentType    = "image/png"
	SaveFilePlace       = "./temp/"
	MinioVideoPrefex    = "http://172.30.154.234:9000/video/"
)
