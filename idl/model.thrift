namespace go model

struct RespBase {
    1: i32 code;
    2: string message;
}

struct userInfo {
    1: i32 userId;
    2: string name;
}