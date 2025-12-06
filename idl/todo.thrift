namespace go todo

include "model.thrift"

struct Todo {
    1: i32 id;
    2: string title;
    3: string content;
    4: bool completed;
    5: i32 userId;
    6: string createdAt;
    7: string diedAt;
    8: i32 priority;
}

//增加Todo请求
struct AddTodoRequest {
    1: string title;
    2: string content;
    3: i32 userId;
    4: string diedAt;
    5: i32 priority;
}

//增加Todo响应
struct AddTodoResponse {
    1: i32 id;
    2: model.RespBase base;
}

//删除todo请求
struct DeleteTodoRequest {
    1: i32 id;
}

//删除todo响应
struct DeleteTodoResponse {
    1: bool success;
    2: model.RespBase base;
}

//更新todo请求
struct UpdateTodoRequest {
    1: i32 id;
    2: optional string title;
    3: optional string content;
    4: optional bool completed;
    5: optional string diedAt;
    6: optional i32 priority;
}

//更新todo响应
struct UpdateTodoResponse {
    1: model.RespBase base;
}

//获取todo请求
struct GetTodoRequest {
    1: i32 id;
}     

//获取todo响应
struct GetTodoResponse {         
    1: Todo todo;
    2: model.RespBase base;
}

//列出todo请求
struct ListTodoRequest {
    1: i32 userId;
}

//列出todo响应  
struct ListTodoResponse {
    1: list<Todo> todos;
    2: model.RespBase base;
}

service TodoService {
    AddTodoResponse AddTodo(1: AddTodoRequest request);
    DeleteTodoResponse DeleteTodo(1: DeleteTodoRequest request);
    UpdateTodoResponse UpdateTodo(1: UpdateTodoRequest request);
    GetTodoResponse GetTodo(1: GetTodoRequest request);
    ListTodoResponse ListTodo(1: ListTodoRequest request);
}






