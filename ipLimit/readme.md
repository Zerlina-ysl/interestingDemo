场景：在一个高并发的web服务器中，要限制IP的频繁访问。现模拟100个IP同时并发访问服务器，每个IP要重复访问1000次。每个IP三分钟之内只能访问一次。修改以下代码完成该过程，要求能成功输出 success:100
- 并发访问---> go func mutex
- 三分钟内只能访问一次-->使用map记录三分钟内是否访问过，并在加入map的时候就异步删除
  ## 基本逻辑
  使用双重for循环，模拟一千次访问，且每次访问都需要一百个ip并发访问
  因为只有一百个ip，且每个ip三分钟内只能访问一次，所以最终只能访问100次，即success=100
```
  // 模拟访问次数
  for i := 0; i < 1000; i++ {
  // 模拟ip
  for j := 0; j < 100; j++ {
  对于每个ip三分钟内只能访问一次，可以使用以ip为key的map，每次轮到新的ip可以遍历，看此ip是否在map中，如果存在就不能访问，如果不存在就可以访问，加入map中，并异步开启线程，三分钟后从map删除ip
  // visit 判断此时访问的ip是否存在
  func (ban *Ban) visit(ip string) bool {
  mutex.Lock()
  defer mutex.Unlock()
  if _, ok := ban.visitIPs[ip]; ok {
  return true
  }
  ban.visitIPs[ip] = struct{}{}
  // 在加入ip后异步计算失效时间
  go ban.invalidAfter3Min(ip)
  return false
  }
  // invalidAfter3Min 开启的异步任务 定时删除
  func (ban *Ban) invalidAfter3Min(ip string) {

  time.Sleep(3 * time.Minute)
  mutex.Lock()
  visitedIPs := ban.visitIPs
  delete(visitedIPs, ip)
  ban.visitIPs = visitedIPs
  mutex.Unlock()
  }
```
  ## 并发实现
  对于每个具有多个请求同事访问的逻辑，都需要使用sync.mutex同步锁，包括 从map中查看ip是否存在、开启异步任务三分钟后从map删除ip等
  而对于一千个请求的并发访问，需要通过waitGroup
  包括对于sucess的数量更改，使用atomic包更改
  ## 实现方式
  通过map，如果未访问过的ip，加入map，而以后ip访问过，就会直接拦下。此程序执行速度很快，<3min,因此不会存在被删除的ip，因此最终只有100次
 ```
 type Ban struct {
  visitIPs map[string]struct{}
  }
  go func() {
  defer wg.Done()
  // j不会按序打印
  ip := fmt.Sprintf("192.168.1.%d", ipEnd)
  //fmt.Println(i)

  if !ban.visit(ip) {
  fmt.Println(ip)
  atomic.AddInt64(&success, 1)
  }
  }()
  func (ban *Ban) visit(ip string) bool {
  mutex.Lock()
  defer mutex.Unlock()
  if _, ok := ban.visitIPs[ip]; ok {
  return true    // 存在map中的ip直接拦下
  }
  ban.visitIPs[ip] = struct{}{}
  // 在加入ip后异步计算失效时间
  go ban.invalidAfter3Min(ip)
  return false

}
```
## 关于success
如果取消ipEnd的使用，就会发现success<=100;如果在goroutine中打印一些变量，success会更小。  
因为waitGroup的使用，如果此次goroutine运行时间较长，计数器不为1，其他的goroutine只能阻塞，j++  
而使用ipEnd赋值，每次重开一个ipEnd，就能确保每个ipEnd都可以进入goroutine
https://blog.csdn.net/qq_28163175/article/details/75287877