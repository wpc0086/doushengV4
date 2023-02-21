package controller

var DemoVideos = []Video{
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

var DemoComments = []Comment{
	{
		Id:         1,
		User:       DemoUser,
		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

var DemoUser = User{
	Id:              1,
	Name:            "TestUser",
	FollowCount:     5,
	FollowerCount:   6,
	IsFollow:        false,
	Avatar:          "https://profile.csdnimg.cn/4/F/7/1_qq_41080854",
	BackgroundImage: "https://img.1ppt.com/uploads/allimg/2302/1_230214151655_1.JPG",
	Signature:       "轻松拿下对不队",
	TotalFavorited:  4,
	WorkCount:       3,
	FavoriteCount:   2,
}
