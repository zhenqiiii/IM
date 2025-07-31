# IM-GO
### 项目概述
使用GO语言编写的一个即时通信（Instant Messaging）系统 <br/>

### 技术栈及其使用场景
1. HTML+TailwindCSS+JS : 前端页面（使用Gin框架静态托管）
2. Gin : 后端框架
3. Redis : 存放验证码
4. MySQL : 数据存储
5. JWT : 无状态用户认证
6. WebSocket : 接发消息
7. air : 后端热重载
8. `TODO` H5 ajax  : 获取音频
9. go官方smtp库: 注册验证邮件发送服务
10. bcrypt : 密码加密
11. `TODO`RabbitMQ : 
<br/>

`TODO`技术特点：借助Go语言`channel/goroutine`提高并发性

### 功能实现

- 聊天方式（核心）：
  - [ ] 访客模式
  - [ ] 私聊
  - [ ] 群聊
  - [ ] 广播
- 辅助功能：
  - [ ] 心跳检测下线
  - [ ] 快捷回复
  - [ ] 撤回记录
  - [ ] 拉黑



