namespace go interact

enum ErrCode {
    SuccessCode                = 0
    ServiceErrCode             = 10001
    ParamErrCode               = 10002
    UserAlreadyExistErrCode    = 10003
    AuthorizationFailedErrCode = 10004
}

struct User {
    1: i64 id
    2: string name
    3: i64 follow_count
    4: i64 follower_count
    5: bool is_follow
    6: string avatar
    7: string background_image
    8: string signature
    9: i64 total_favorited
    10: i64 work_count
    11: i64 favorite_count
}

struct Video {
    1: i64 id
    2: User author
    3: string play_url 
    4: string cover_url 
    5: i64 favorite_count 
    6: i64 comment_count 
    7: bool is_favorite 
    8: string title
}

struct Comment{
 	1: i64 id
    2: User user
    3: string content 
    4: string create_date
}

struct FavoriteActionRequest {
    1: string token
    2: i64 video_id
    3: i32 action_type
}

struct FavoriteActionResponse {
    1: i32 status_code
    2: string status_msg
}

struct FavoriteListRequest {	
    1: i64 user_id
    2: string token
}

struct FavoriteListResponse {
    1: i32 status_code
    2: string status_msg
    3: list<Video> video_list
}

struct CommentActionRequest {
    1: string token
    2: i64 video_id
    3: i32 action_type
    4: string comment_text
    5: i64 comment_id
}

struct CommentActionResponse {
    1: i32 status_code
    2: string status_msg
    3: Comment comment
}

struct CommentListRequest {
    1: string token
    2: i64 video_id
}

struct CommentListResponse {
    1: i32 status_code
    2: string status_msg
    3: list<Comment> comment_list
}

service InteractService {
    FavoriteActionResponse FavoriteAction(1: FavoriteActionRequest req)
    FavoriteListResponse FavoriteList(1: FavoriteListRequest req)
    CommentActionResponse CommentAction(1: CommentActionRequest req)
    CommentListResponse CommentList(1: CommentListRequest req)
}
