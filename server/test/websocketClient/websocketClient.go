package main

import (
	"crypto/tls"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func main() {
	// WebSocket 服务器地址
	//u := url.URL{Scheme: "ws", Host: "appapidev.yeebok.cn", Path: "/api/ws"}
	//log.Printf("Connecting to %s", u.String())
	//Authorization := ""
	// 创建一个 HTTP 请求头
	header := http.Header{}
	//fmt.Println("请输入用户身份令牌")
	//_, err := fmt.Scanln(&Authorization)
	//if err != nil{
	//	panic(err)
	//}
	// 司机端身份
	header.Add("Authorization", "58bf31b7a0cc9d7035dfd3bd3ae033cf5ea6ae74a45b0dafdf02ab63019c62544a5cfab30baf0e1b97ca6b5deff494cb82abfff63b93bda6126aea7d47839b07819b17e08596a527c533364fd1c2129e183c11ad908583efb974a963e8e1345f747442c0e6388f066145f5f619d25177b701fcfb9e8e0072b7879e19423986e0601d10ba83b5fc8095d54838021eb112013212b781785a9198210bc1e131a869faddefdf243c92dceeb71dd9b2aa842b04833ccf605fabeff41cabd24deb4843e2e747601c9b0a14ec3c713d890106ebe19769f840a2a3913ff72ecb10534c41070119bfadfcb7e508eb060faf0a00aa25e2bc7c755297d43cf1145e7591e9eca4f419f1860f172403b1518e16f7864a1c0a5ed2ba28a14b259f49dcbd14ba17efde90bf061069222d7fefc5309fcd1e709e1c8da4221f443ea57f866f63a35446322dd3f0addc8d4362ad7115af769e1d4e25bc3d8e6eaa68db02b26a880120cfdace4ad74ad608ebb3a117a2c4fd53658714a82e470196a2dc0393414e51545b2ba28c1b95ad5b160209e5d88721da3a36776fa559d4c03bed8b40d6c3dad8c59b1e486e95c91aef3ca7c4f7234d21adf151a1437400103a4b87b4e7ae245f9a88efa3899876dfebfb72d1ecc8b3e35150f76b81675c8ccd223b4e5ebfd8e65c129e56e9b014e267e0c2b16be90f01117d4e9181bef64a29ecaf7d453e2439aac0f25a3a226bb844008ac93bf794368e1f22eae3e3c1b525083c20c33620a2d123e412673a5edc742e49901b51d47755f5fe6e39e81f032fab0bab514236cbb7704efa8a66ab5a65050c11fa1aef96f14146b5ac4cc8b59e08f848c66d7ac5fdb82c1dc91fd2505f7e4417cc0e7bbe18551fc117a8e8ff48859ff03eb609bd3f4c2e4bdb4a4fec30225fb23b537a6a8bd1b26a42dc1dbf1d982fe2f46c3b3506a1df3c40b506d0641568e1831d555c85fe90c97cae8408e2d65be56e4ca8faa25dde0c857d1021776320272d60a4c3fdf6f94e6fbaf4cc44c42aba7e41fe5f5e6d9d4c9c515abd1dc68a6527c9da36aaeea3ad894821463c5200e4eb4c3ee3142b9e7c45fef2d8fc91330d3bc4b96a5cefefac98d0013c2f9a3b4b864b414d41b4437ce023f556e2146c62acccb873037a6ee390560bc880f949d5f491ae74ddb71b840c68c0b6d7a1d998654b9737a0bfc45a478a8f6655ad723f03a13b9cc74ad1aac26be1f39352741276174d92")
	header.Add("Connection", "upgrade")
	header.Add("websocket", "Upgrade")
	// 乘客端身份
	//header.Add("Authorization", "4e2ef20f5c9d1d232811314ee0845751044bf25ec44f9912b7c1f234d21743644513d500ef553b707cdded0c05f4558ee7faaaeb89400bfddcd75612c5b8477be944bc46e997ae6bf9a321d6a5b9f712282839a008e8df7876149ccddc2e2582af762d2e856c59e0c6dcd9894d4cc4bdaecc6f8f02aa20643bd8c2eede60ad709eac29c77469a71728d612204cac2162a21940ea5de4446c14737767a3481bafc50e682fcf440ef2adf24e4cf7ba58b6e9d647d6548ff904ccc0e0b2fdb2e331228f3462212d1607e8948db9931628d327b834983b57ef886123c388aad955a4635c97f59eade1c5b97265db11674e8c3b08ba9ae501c058a2ff830afbf28ac987415ff4cc7eff57387e3bf8b3d546154f1c2b636a55933426e497f746daa8e3aedff78d0fcbceb30d0753c586d5c796c90f728b6c91c2eabe53ba4fec0d8e8ba4b088ca578b8cae323ba5b12db1e7b15f07cc0376d6f11310e282464c350a8fa20ad75a0f08402460036cbe7f656a63211602ce7439b609cf76cd4aeb2130ff177cd4bedcfaa34068933300b95bdfc6af1fe93568fa5290691407c3235d2cf23fa897ecb5ecad1930a7910a22a99bfc86357df3e19070152612523d3c22a2117040bcd47d06be19de7e63f1657d22b9ab9a805de2baed37db98da4e2b7b40cfcb18d714963f9b878bea6e0539d3d4db50e82375a08e542e24c9f3769fe48eaeb535ed0e97f3682d0b3d958a74ffc534fc1aa72e15cb4a3fc587488a0d1ebd28bc8dbe40b8a5c1f4b5834851148b0f48a4508323ab755850b96d61dd416795c9fc72b1b548e788858f47b9bdfb65382d2ccb8abeccf9fb0ec9b198b1618f603ce3784b71a6a0c3fd8ebe94874852e5efda0b0bf776235d19e08fa044a34be55870f2192748e9c7b3b4b882ebd3e64d776bdadef72b8046fa8ff4e76a4401c9b0526e83f9262e477910f68a0177b52ead7b3468ec36cd97d0ed696c880cbaf778fcb1042ea961942faa8000c5d7fa1737a9536bb47ad8bbfe6a418a349b00ddeae5b8f2fa0599b508a6bd9d054650b8dd9709150980d44e953460a93a758d9078e62bb8740dc0b717092a357c760789c2")
	//header.Add("Custom-Header", "Custom-Value")
	dialer := websocket.Dialer{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	// 连接到 WebSocket 服务器
	conn, resp, err := dialer.Dial("wss://appapi.apt11.net/api/ws", header)
	if err != nil {
		if resp != nil {
			log.Printf("HTTP status: %s", resp.Status)
		}
		log.Fatalf("Failed to connect to WebSocket server: %v", err)
	}
	defer conn.Close()
	fmt.Println("connect success")
	// 发送消息到服务器
	//message := []byte("Hello, WebSocket Server!")
	//err = conn.WriteMessage(websocket.TextMessage, message)
	//if err != nil {
	//	log.Fatalf("Failed to send message: %v", err)
	//}
	//fmt.Printf("Sent message: %s\n", string(message))
	for {
		// 接收服务器的响应
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Fatalf("Failed to receive message: %v", err)
		}
		fmt.Printf("Received message: %s\n", string(p))
	}
}
