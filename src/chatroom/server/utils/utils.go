package utils

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)


//将这些方法关联到结构体中

type Transfer struct {
	//分析应该有哪些字段
	Conn net.Conn
	Buf [8096]byte //传输时使用的缓冲

}

func (this *Transfer) ReadPkg()(mes message.Message,err error){
	fmt.Println("读取客户端发送的数据...")
	// conn.Read只有在未关闭的情况下，才会阻塞
	//如果客户端关系了conn，就不会阻塞
	_,err = this.Conn.Read(this.Buf[:4])
	if err!= nil{
		//err = errors.New("conn.Read header err")
		return
	}
	// 根据buf[:4]读到的长度转成uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])
	n,err := this.Conn.Read(this.Buf[:pkgLen]) //从套接字中读pkgLen个字节到buf中
	if n != int(pkgLen) || err != nil{
		//err = errors.New("conn.Read body err")
		return
	}

	//把pkgLen反序列化成->message.Message
	err = json.Unmarshal(this.Buf[:pkgLen],&mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=",err)
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte)(err error){
	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:4],pkgLen)

	//发送长度
	n,err := this.Conn.Write(this.Buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn,Write(bytes) fail",err)
		return
	}
	//发送data本身
	n,err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn,Write(bytes) fail",err)
		return
	}
	return
}