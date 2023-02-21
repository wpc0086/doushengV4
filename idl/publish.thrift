namespace go publish

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

struct ActionRequest {
    1: string token
    2: binary data
    3: string title
}

struct ActionResp {
    1: i32 status_code
    2: string status_msg
}

struct ListRequest {	
    1: i64 user_id
    2: string token
}

struct ListResp {
    1: i32 status_code
    2: string status_msg
    3: list<Video> video_list
}

struct FeedRequest {
    1: optional i64 latest_time
    2: optional string token
}

struct FeedResponse {
    1: i32 status_code
    2: string status_msg
    3: list<Video> video_list
    4: i64 next_time 
}

service PublishService {
    FeedResponse FeedPublish(1: FeedRequest req)
    ActionResp ActionPublish(1: ActionRequest req)
    ListResp ListPublish(1: ListRequest req)
}