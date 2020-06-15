## socket 编程

```
Socket是BSD UNIX的进程通信机制，通常也称作”套接字”，用于描述IP地址和端口，是一个通信链的句柄
Socket是应用层与TCP/IP协议族通信的中间软件抽象层
Socket又称“套接字”，应用程序通常通过“套接字”向网络发出请求或者应答网络请求
常用的Socket类型有两种：流式Socket和数据报式Socket，流式是一种面向连接的Socket，针对于面向连接的TCP服务应用，数据报式Socket是一种无连接的Socket，针对于无连接的UDP服务应用
TCP：比较靠谱，面向连接，比较慢
UDP：不是太靠谱，比较快
```