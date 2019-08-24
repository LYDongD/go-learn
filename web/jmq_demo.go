package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

var count = uint32(0)
var total = uint32(100000)

var z0 = " 吃了没，您吶?"
var z3 = " 嗨！吃饱了溜溜弯儿。"
var z5 = " 回头去给老太太请安！"
var l1 = " 刚吃。"
var l2 = " 您这，嘛去？"
var l4 = " 有空家里坐坐啊。"

var liWriteLock sync.Mutex
var zhangWriteLock sync.Mutex

type RequestResponse struct {
	Serial  uint32 //消息序列号
	Payload string //消息内容
}

//发送消息： 1 序列化协议 2 通信协议 3 线程安全
func writeTo(r *RequestResponse, conn *net.TCPConn, lock *sync.Mutex) {
	lock.Lock()
	defer lock.Unlock()

	//通信协议，报文序列化金额组装
	payloadBytes := []byte(r.Payload) //报文byts
	serialBytes := make([]byte, 4)    //序列号bytes
	binary.BigEndian.PutUint32(serialBytes, r.Serial)
	length := uint32(len(payloadBytes) + len(serialBytes))
	lengthBytes := make([]byte, 4) //长度bytes
	binary.BigEndian.PutUint32(lengthBytes, length)

	conn.Write(lengthBytes)
	conn.Write(serialBytes)
	conn.Write(payloadBytes)
}

//读消息： 1 反序列化协议 2 应用协议，断句
func readFrom(conn *net.TCPConn) (*RequestResponse, error) {
	ret := &RequestResponse{}
	//根据通信协议解析报文，先读长度，序列号，然后才是内容
	//use buffer to read
	buf := make([]byte, 4)
	if _, err := io.ReadFull(conn, buf); err != nil {
		return nil, fmt.Errorf("读长度故障: %s", err.Error())
	}
	length := binary.BigEndian.Uint32(buf)

	if _, err := io.ReadFull(conn, buf); err != nil {
		return nil, fmt.Errorf("读序列号故障: %s", err.Error())
	}
	ret.Serial = binary.BigEndian.Uint32(buf)

	payloadBytes := make([]byte, length-4)
	if _, err := io.ReadFull(conn, payloadBytes); err != nil {
		return nil, fmt.Errorf("读payload故障：%s", err.Error())
	}
	ret.Payload = string(payloadBytes)

	return ret, nil
}

//启动client
func startClient() {
	//建立连接
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)

	defer conn.Close()
	var wg sync.WaitGroup
	wg.Add(1)
	go liDaYeListen(conn, &wg)
	go liDaYeSay(conn)
	wg.Wait()
}

func startServer() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	fmt.Println("张大爷在胡同口等着")
	for {
		conn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("碰见一个李大爷：" + conn.RemoteAddr().String())
		go zhangDaYeListen(conn)
		go zhangDaYeSay(conn)
	}
}

func zhangDaYeListen(conn *net.TCPConn) {
	//听1w次
	for count < total {
		r, err := readFrom(conn)
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		if r.Payload == l2 {
			go writeTo(&RequestResponse{r.Serial, z3}, conn, &zhangWriteLock)
		} else if r.Payload == l4 {
			go writeTo(&RequestResponse{r.Serial, z5}, conn, &zhangWriteLock)
		} else if r.Payload == l1 {

		} else {
			fmt.Println("张大爷听不懂:" + r.Payload)
			break
		}
	}
}

//说1w次，每次说的序列化不同
func zhangDaYeSay(conn *net.TCPConn) {
	nextSerial := uint32(0)
	for i := uint32(0); i < total; i++ {
		writeTo(&RequestResponse{nextSerial, z0}, conn, &zhangWriteLock)
		nextSerial++
	}
}

func liDaYeListen(conn *net.TCPConn, wg *sync.WaitGroup) {
	defer wg.Done()
	//听1w次
	for count < total {
		r, err := readFrom(conn)
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		if r.Payload == z0 {
			writeTo(&RequestResponse{r.Serial, l1}, conn, &liWriteLock)
		} else if r.Payload == z3 {

		} else if r.Payload == z5 {
			count++

		} else {
			fmt.Println("李大爷听不懂:" + r.Payload)
			break
		}
	}
}

func liDaYeSay(conn *net.TCPConn) {
	nextSerial := uint32(0)
	for i := uint32(0); i < total; i++ {
		writeTo(&RequestResponse{nextSerial, l2}, conn, &liWriteLock)
		nextSerial++
		writeTo(&RequestResponse{nextSerial, l4}, conn, &liWriteLock)
		nextSerial++
	}
}
func main() {
	go startServer()
	time.Sleep(time.Second)
	t1 := time.Now()
	startClient()
	elapsed := time.Since(t1)
	fmt.Println("耗时:", elapsed)
}
