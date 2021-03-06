/**
 * @Author: zhu
 * @Date: 20-5-21 下午4:54
 */
package main

//传输层：tcp（面向连接的、可靠的、基于字节流的传输层通信协议）、udp（无连接的传输层协议，提供面向事务的简单不可靠信息传送服务）协议
//应用层：HTTP（超文本传输协议），FTP（文本传输协议）协议
//网络层 IP，ICMP，IGMP协议
//网络接口层：ARP（正向地址解析协议，通过已知的IP，寻找对应主机的MAC地址），RARP（反向地址转换协议，通过MAC地址确定IP地址）协议
//网络分层架构:每一层利用下一层提供的服务来为上一层提供服务，本层服务的实现细节对上层屏蔽.
//“IP地址+端口号”就对应一个socket,可以用Socket来描述网络连接的一对一关系.流式是一种面向连接的Socket，针对于面向连接的TCP服务应用；数据报式Socket是一种无连接的Socket，对应于无连接的UDP服务应用
//TCP连接是全双工的，因此每个方向都必须单独进行关闭。所以四次挥手。挥手等待时间： 2MSL即两倍的MSL，TCP的TIME_WAIT状态也称为2MSL等待状态，MSL是Maximum Segment Lifetime英文的缩写，中文可以译为“报文最大生存时间”
//msl时间子30s到2分钟之间。服务端的timeout + FIN的传输时间，为了保证可靠，采用更加保守的等待时间2MSL。
//HTTP协议通常承载于TCP协议之上，有时也承载于TLS或SSL协议层之上，这个时候，就成了我们常说的HTTPS。
//1XX 服务器已经接收了客户端请求，客户端可以继续发送请求
//2XX 表示客户端已经成功收到请求并进行处理 eg:200 客户端请求成功
//3XX 表示服务器要求客户端进行重定向
//4XX 表示客户端的请求有非法内容  eg:400 Bad Request请求报文有语法错误 401 Unauthorized未授权 403 Forbidden服务器拒绝服务 404 Not Found请求的资源不存在
//5XX 表示服务器未能正常处理客户端的请求而出现意外错误 eg:500 服务器内部错误 503服务器临时不能处理客户端请求(稍后可能可以)
func main()  {

}