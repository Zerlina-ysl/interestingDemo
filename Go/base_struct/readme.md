# slice
* 如果slice没有发生扩容，修改会在原来的内存中
* 如果slice发生扩容，修改会在新的内存中
* 预先分配内存，可以提升性能。
* 预先分配len，使用index赋值，而非append，可以提升性能。
## bounds checking elimination
indexOutOfBoundError
* 如果能确定访问到的slice长度，可以先执行一次下标检查
# string
* string不可变，[]byte可变.将string转换为[]byte时不能进行写操作。
# channel
* for+select closed channel会造成死循环
* channel仅用于传递通知，而不是传递值
# 内存模型与happen-before
happen-before定义了偏序关系：某种条件下，保证事件a一定在事件b之前发生