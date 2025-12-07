namespace go model

struct RespBase {
    1: i32 code;
    2: string message;
}

struct userInfo {
    1: i64 userId;
    2: string name;
}

struct todoInfo {
    1: i64 id;
    2: string title;
    3: string content;
    4: bool completed;
    5: i64 userId;
    6: string createdAt;
    7: string diedAt;
    8: i64 priority;
}