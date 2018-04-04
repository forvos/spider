# p2p

## 通信协议规范

采用二进制传输协议,大端字节顺序
规范：
|版本号   |magic  | type  | zip   |masgId     |len    | body
|--------|-------|-------|-------|----------|-------|----------|
 1        594C

版本号: 单字节 uint8 |**|***|***| => |大版本号|中版本号|小版本号|
magic: 2字节 uint16  594C 固定数字

data type: 单字节 uint8
压缩或者加密: 单字节 uint8
masgId:消息id
len: 8字节 uint64 body的数据长度
body:protobuf序列化的二进制数据


