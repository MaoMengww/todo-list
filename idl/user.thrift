namespace go user 

include "model.thrift"

struct User {
    1: string name;
    2: string phone;
    3: string password;
}

// 注册
struct RegisterRequest {
    1: string name;
    2: string phone;
    3: string password;
}

struct RegisterResponse {
    1: model.RespBase base;
    2: i32 userId;
}

// 登录
struct LoginRequest {
    1: string name;
    2: string password;
}

struct LoginResponse {
    1: model.RespBase base;
    2: model.userInfo info;
}

// 获取用户信息
struct GetUserRequest {
    1: i32 userId;
}

struct GetUserResponse {
    1: model.userInfo info;
    2: model.RespBase base;
}

//改名
struct UpdateusernameRequest {
    1: i32 userId;
    2: string username;
}

struct UpdateusernameResponse {
    1: model.RespBase base;
}

service UserService {
    RegisterResponse Register(1: RegisterRequest req);
    LoginResponse Login(1: LoginRequest req);
    GetUserResponse GetUser(1: GetUserRequest req);
    UpdateusernameResponse Updateusername(1: UpdateusernameRequest req);
}