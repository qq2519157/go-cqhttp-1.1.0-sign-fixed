原项目地址为<a href="https://github.com/Mrs4s/go-cqhttp">https://github.com/Mrs4s/go-cqhttp</a>
修复了sign服务器不能在gocq中使用的问题

对config进行了一点添加,如下
<code lang="json">
# 略……
account: # 账号相关
  uin: 1233456 # QQ账号
  password: '' # 密码为空时使用扫码登录

  # 略……
  
  sign-server: '-'  # sign服务器的IP或域名
  sign-server-key: '114514'  # sign服务器的key
  cmdId: '810_9' # QQ发送登录包的CmdId和SubCmdId，例子中810是登陆CmdId，9是SubCmdId

# 略……
</code>
