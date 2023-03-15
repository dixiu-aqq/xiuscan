package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var results []int
	ip := os.Args[1:2]
	var wg sync.WaitGroup
	for i := 1; i < 65536; i++ {
		wg.Add(1)
		go func(s []string, j int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", ip, j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			results = append(results, j)
		}(ip, i)
	}
	wg.Wait()
	for _, v := range results {
		switch v {
		case 21:
			fmt.Printf("%d 端口打开,ftp,可尝试\u001B[36m匿名上传、下载、爆破/嗅探/后门\u001B[m\n", v)
		case 22:
			fmt.Printf("%d 端口打开,ssh,可尝试\u001B[36m爆破、内网代理、openssh漏洞\u001B[m\n", v)
		case 23:
			fmt.Printf("%d 端口打开,telnet,可尝试\u001B[36m爆破、弱口令\u001B[m\n", v)
		case 25:
			fmt.Printf("%d 端口打开,smtp,\u001B[36m邮件伪造\u001B[m\n", v)
		case 53:
			fmt.Printf("%d 端口打开,dns,可尝试\u001B[36m域传送、DNS劫持、缓存投毒,欺骗\u001B[m\n", v)
		case 67, 68:
			fmt.Printf("%d 端口打开,dhcp,\u001B[36m劫持、欺骗\u001B[m\n", v)
		case 80, 443:
			fmt.Printf("%d 端口打开,web服务,\u001B[36m访问地址尝试相关漏洞\u001B[m\n", v)
		case 110:
			fmt.Printf("%d 端口打开,POP3,\u001B[36m爆破、劫持\u001B[m\n", v)
		case 139:
			fmt.Printf("%d 端口打开,samba,可尝试\u001B[36m爆破、未授权、远程代码执行\u001B[m\n", v)
		case 143:
			fmt.Printf("%d 端口打开,imap,\u001B[36m爆破\u001B[m\n", v)
		case 161, 162:
			fmt.Printf("%d 端口打开,snmp,\u001B[36m爆破、搜集目标内网信息、弱口令\u001B[m\n", v)
		case 389:
			fmt.Printf("%d 端口打开,ldap,\u001B[36m注入,未授权、弱口令\u001B[m\n", v)
		case 445:
			fmt.Printf("%d 端口打开,smb,\u001B[36m弱口令,ms17_010、ms_08067\u001B[m\n", v)
		case 512, 513, 514:
			fmt.Printf("%d 端口打开,Linux Rexec,\u001B[36m爆破、Rlogin登录\u001B[m\n", v)
		case 873:
			fmt.Printf("%d 端口打开,Rsync,\u001B[36m未授权,文件上传\u001B[m\n", v)
		case 1025:
			fmt.Printf("%d 端口打开,rpc,\u001B[36mnfs匿名访问\u001B[m\n", v)
		case 1080:
			fmt.Printf("%d 端口打开,socket,\u001B[36m爆破\u001B[m\n", v)
		case 1099:
			fmt.Printf("%d 端口打开,java rmi,\u001B[36m远程命令执行\u001B[m\n", v)
		case 1352:
			fmt.Printf("%d 端口打开,lotus domino邮件,\u001B[36m弱口令,信息泄露、爆破\u001B[m\n", v)
		case 1433:
			fmt.Printf("%d 端口打开,mssql,\u001B[36m注入,提权、爆破\u001B[m\n", v)
		case 1521:
			fmt.Printf("%d 端口打开,Oracle,\u001B[36mTNS爆破、注入、反弹shell\u001B[m\n", v)
		case 2049:
			fmt.Printf("%d 端口打开,nfs,\u001B[36m配置不当\u001B[m\n", v)
		case 2082, 2083:
			fmt.Printf("%d 端口打开,cpanel主机管理系统,\u001B[36m弱口令爆破\u001B[m\n", v)
		case 2181:
			fmt.Printf("%d 端口打开,zookeeper,\u001B[36m未授权\u001B[m\n", v)
		case 2222:
			fmt.Printf("%d 端口打开,da虚拟主机管理系统,\u001B[36m弱口令爆破\u001B[m\n", v)
		case 2375:
			fmt.Printf("%d 端口打开,docker remote api,\u001B[36m未授权访问\u001B[m\n", v)
		case 2601, 2604:
			fmt.Printf("%d 端口打开,zebra路由器,\u001B[36m默认密码zebra\u001B[m\n", v)
		case 3128:
			fmt.Printf("%d 端口打开,squid代理,\u001B[36m未授权直接内网漫游\u001B[m\n", v)
		case 3306:
			fmt.Printf("%d 端口打开,mysql,\u001B[36m注入、提权、爆破\u001B[m\n", v)
		case 3690:
			fmt.Printf("%d 端口打开,svn服务,\u001B[36msvn泄露、未授权访问\u001B[m\n", v)
		case 3389:
			fmt.Printf("%d 端口打开,rdp,\u001B[36m爆破\u001B[m\n", v)
		case 4440:
			fmt.Printf("%d 端口打开,rundeck,\u001B[36mweb\u001B[m\n", v)
		case 4848:
			fmt.Printf("%d 端口打开,glass fish,\u001B[36m弱口令、爆破、认证绕过\u001B[m\n", v)
		case 5000:
			fmt.Printf("%d 端口打开,sysbase/db2数据库,\u001B[36m爆破、注入、提权\u001B[m\n", v)
		case 5432:
			fmt.Printf("%d 端口打开,postgresql,\u001B[36m爆破、注入、弱口令、缓冲区溢出\u001B[m\n", v)
		case 5632:
			fmt.Printf("%d 端口打开,pyanywhere,\u001B[36m抓密码、代码执行\u001B[m\n", v)
		case 5900, 5901, 5902:
			fmt.Printf("%d 端口打开,vnc,\u001B[36m弱口令爆破、认证绕过\u001B[m\n", v)
		case 5984:
			fmt.Printf("%d 端口打开,couchdb,\u001B[36mhttp://xxx.5984/_utils/\u001B[m\n", v)
		case 6379:
			fmt.Printf("%d 端口打开,redis,\u001B[36m未授权访问、弱口令\u001B[m\n", v)
		case 7001, 7002:
			fmt.Printf("%d 端口打开,weblogic控制台,\u001B[36mJava反序列化、弱口令\u001B[m\n", v)
		case 7778:
			fmt.Printf("%d 端口打开,kloxo,\u001B[36m主机控制面板登录\u001B[m\n", v)
		case 8069:
			fmt.Printf("%d 端口打开,zabbix,\u001B[36msql注入、远程执行\u001B[m\n", v)
		case 8080, 8089:
			fmt.Printf("%d 端口打开,jboss、resin、jetty、jenkins,\u001B[36m反序列化、控制台弱口令、put文件上传\u001B[m\n", v)
		case 8083, 8086:
			fmt.Printf("%d 端口打开,influxdb、vestacp,\u001B[36m未授权\u001B[m\n", v)
		case 8161:
			fmt.Printf("%d 端口打开,activemq,\u001B[36m弱口令、写文件\u001B[m\n", v)
		case 8888:
			fmt.Printf("%d 端口打开,amh、lumanager,\u001B[36m主机管理系统默认端口说明\u001B[m\n", v)
		case 9000:
			fmt.Printf("%d 端口打开,fastcgi,\u001B[36m远程命令执行\u001B[m\n", v)
		case 9200, 9300:
			fmt.Printf("%d 端口打开,elastic search,\u001B[36m远程代码执行\u001B[m\n", v)
		case 9090, 9043:
			fmt.Printf("%d 端口打开,websphere控制台,\u001B[36m爆破、Java反序列化、弱口令\u001B[m\n", v)
		case 10000:
			fmt.Printf("%d 端口打开,virtualmin、webmin-web控制面板,\u001B[36m弱口令\u001B[m\n", v)
		case 11211:
			fmt.Printf("%d 端口打开,memcache服务,\u001B[36m未授权、内存泄漏\u001B[m\n", v)
		case 27017, 27018:
			fmt.Printf("%d 端口打开,mongodb,\u001B[36m爆破、未授权\u001B[m\n", v)
		case 50000:
			fmt.Printf("%d 端口打开,sap management,\u001B[36m远程命令执行\u001B[m\n", v)
		case 50060, 50030:
			fmt.Printf("%d 端口打开,hadoop,\u001B[36m未授权访问\u001B[m\n", v)
		}
	}
	fmt.Println("所有开放的端口为:", results)
	t := time.Now().Sub(start)
	fmt.Printf("扫描总耗时:%v", t)
}
