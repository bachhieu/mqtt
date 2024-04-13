package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	loop = 50000
	mess = 1000
)

// Hàm callback khi kết nối thành công
func onConnect(client mqtt.Client) {
	// for a := 0; a < loop/10; a++ {
	// 	token := client.Publish("volio/dol/test", 2, true, fmt.Sprintf("test message %v", i))
	// 	token.Wait()
	// }
	// Đăng ký các chủ đề bạn muốn nhận thông điệp từ đây

	// time.Sleep(5 * time.Second)
	// client.Unsubscribe("test")
	// fmt.Println("-----------------------")
	// client.Subscribe("phone-tracker/topic/noti/5653a8e114929c2a-c04a13", 0, func(c mqtt.Client, m mqtt.Message) {
	// 	fmt.Println("phone-tracker/topic/noti/5653a8e114929c2a-c04a13 ", string(m.Payload()))
	// })
}

// Hàm callback khi nhận được thông điệp
func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	// client.Subscribe("test", 0, func(c mqtt.Client, m mqtt.Message) {
	// 	fmt.Println("1:: ", string(m.Payload()))
	// 	client.Publish("test1", 0, true, "test")
	// })
	fmt.Printf("Received message: %s from topic: %s\n", message.Payload(), message.Topic())
}

// Hàm callback khi xảy ra lỗi kết nối
func onConnectionLost(client mqtt.Client, err error) {
	fmt.Printf("Connection lost: %v- %v\n", i, err)
}

var i = 1

func main() {
	// var wg sync.WaitGroup
	// LOOP:
	// 	client, err := connect()

	// 	if err != nil {
	// 		fmt.Printf("Connect to mqtt error:: %v", err)
	// 		goto LOOP
	// 	}

	// ticker := time.NewTicker(1 * time.Millisecond)
	// quit := make(chan struct{})
	// go func() {
	// 	for {
	// 		select {
	// 		case <-ticker.C:
	// 			// do stuff
	// 			client.Publish(fmt.Sprintf("test/%v", time.Now().UnixMilli()), 2, true, "test")
	// 		case <-quit:
	// 			ticker.Stop()
	// 			return
	// 		}
	// 	}
	// }()
	for a := 0; a < loop; a++ {
		// wg.Add(1)
		time.Sleep(10 * time.Millisecond)
		go connect(a)
	}
	// wg.Wait()
	// go func() {
	// 	for {
	// 		select {
	// 		case <-ticker.C:
	// 			// do stuff
	// 			connect()
	// 			i++
	// 		case <-quit:
	// 			ticker.Stop()
	// 			return
	// 		}
	// 	}
	// }()

	// Ngừng chương trình khi nhận được tín hiệu INT hoặc TERM
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c

	// Ngắt kết nối MQTT khi kết thúc chương trình
	// client.Disconnect(250)
}

func connect(i int) (mqtt.Client, error) {
	// Tạo một client MQTT
	clientOptions := mqtt.NewClientOptions().AddBroker("mqtt://ex1.mq.volio.vn:1883")
	clientOptions.ConnectTimeout = 100 * time.Second
	clientOptions.SetClientID(fmt.Sprintf("hieubv-ben%v", i))
	clientOptions.SetUsername("hieubv")
	clientOptions.SetPassword("Volio@123")

	// Gán các hàm callback cho client MQTT
	clientOptions.OnConnect = onConnect
	clientOptions.DefaultPublishHandler = onMessageReceived
	clientOptions.OnConnectionLost = onConnectionLost

	// Tạo MQTT client
	client := mqtt.NewClient(clientOptions)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	fmt.Printf("Connected to MQTT broker : %v \n", i)

	// for a := 0; a < mess; a++ {
	// 	time.Sleep(100 * time.Millisecond)
	// 	token := client.Publish("volio/dol/test", 2, true, fmt.Sprintf("test message %v", i))
	// 	token.Wait()
	// }

	return client, nil
}
