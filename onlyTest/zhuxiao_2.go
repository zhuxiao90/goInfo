/**
 * @Author: zhu
 * @Date: 20-5-14 下午3:47
 */
package main

//select and poll 都需要在返回后，通过遍历文件描述符来获取已经就绪的socket。同时连接大量客户端在同一时刻也只有很少的处于就绪状态，因此随着监视的描述符数量的增长，其效率也会线性下降。
//epoll de  epoll_create(int size) 参数size并不是限制了epoll所能监听的描述符最大个数，只是对内核初始分配内部数据结构的一个建议.
//epoll 工作模式默认LT模式，当epoll_wait检测到描述符事件发生并将此事通知应用程序，应用程序可以不立刻处理该事件，下次调用epoll_wait时，会再次向应用程序并同事此事件，ET模式则相反。
//在select、poll中，进程只有在调用一定方法后，内核才对所有监视的文件描述符进行扫描，二epoll事先通过epoll_ctl()来注册一个文件描述符，一旦基于某个文件描述符就绪时
//内核会采用类似callback的毁掉机制，迅速激活这个文件描述符，当进程调用epoll_wait时便得到通知（即去掉了遍历文件描述符，而是通过监听回调的机制来获取已就绪的socket）
//epoll优点：1 监听的描述符数量不受限制，支持的FD上限是最大可以打开的文件数目，和内存有关，1G设备内存，大概10W个。
//2 IO的效率不会随着监视FD的数量增长而下降，epoll不同于select和poll轮询的方式，而是通过每个fd定义的回调函数来实现的，只有就绪的fd才会执行回调函数。
//3 如果没有大量的idle -connection或者dead-connection，epoll的效率并不会比select/poll高很多，但是当遇到大量的idle- connection，就会发现epoll的效率大大高于select/poll
//epoll通过红黑树（RB-tree）搜索被监视的文件描述符，在epoll实例上注册时间时，epoll会将该事件添加到epoll实例的红黑树并注册一个回调函数，当事件发生时会将事件添加到就绪链表中
func main()  {

}
