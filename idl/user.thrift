namespace go user

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

struct RegisterUserRequest {
    1: string username (vt.min_size = "1")
    2: string password (vt.min_size = "6")
}

struct RegisterResp {
    1: i32 status_code
    2: string status_msg
    3: i64 user_id
    4: string token
}

struct LoginUserRequest {
    1: string username (vt.min_size = "1")
    2: string password (vt.min_size = "6")
}

struct LoginResp {
    1: i32 status_code
    2: string status_msg
    3: i64 user_id
    4: string token
}

struct InfoUserRequest {
    1: i64 user_id (vt.min_size = "1")
    2: string token (vt.min_size = "1")
}

struct InfoUserResponse {
    1: i32 status_code
    2: string status_msg
    3: User user
}

service UserService {
    RegisterResp RegisterUser(1: RegisterUserRequest req)
    LoginResp LoginUser(1: LoginUserRequest req)
    InfoUserResponse InforUser(1: InfoUserRequest req)
}