# go-chatroom
golang聊天室


C/S架构 tcp编程
  

clinet包  客户端
  main包  显示页面，提示输入，根据用户输入内容调度不同的方法
  process包  业务逻辑层 登录，注册，发消息等，注意其中有个server文件是与服务器另起一条协程保持通信
  utils包 客户端工具包
  
   
server包  服务端
  main包 主函数入口 初始化redis线程池，有连接时分配一个协程进行客户端连接通信，总调度层
  process包 子调度层，根据客户端消息类型调度不同方法去处理
  models包  模型层，对数据库的增删改查
  utils包  server端工具包 封装一些工具方法，如读写操作

common包  前后端通用的工具包
