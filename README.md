原项目地址为<a href="https://github.com/Mrs4s/go-cqhttp">https://github.com/Mrs4s/go-cqhttp</a>
修复了sign服务器不能在gocq中使用的问题

对config进行了一点添加,如下:
<pre><code># 略……
account:
  uin: 1233456
  password: '' 

  # 略……
  
  sign-server: 'http://127.0.0.1:8080' # sign服务器的IP或域名
  sign-server-key: '114514'  # sign服务器的key
  cmdId: '810_9' # QQ发送登录包的CmdId和SubCmdId，例子中810是登陆CmdId，9是SubCmdId，这里810_9不建议改

# 略……
</code></pre>
