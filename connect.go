package main

// import (
// 	"fmt"
// 	"time"

// 	mqtt "github.com/eclipse/paho.mqtt.golang"
// )

// func connect(i int) (mqtt.Client, error) {
// 	// Tạo một client MQTT
// 	// clientOptions := mqtt.NewClientOptions().AddBroker("mqtt://host.docker.internal:1884")
// 	// clientOptions := mqtt.NewClientOptions().AddBroker("mqtt://ex1.mq.volio.vn:1883")
// 	// clientOptions := mqtt.NewClientOptions().AddBroker("mqtt://45.77.248.75:1883")
// 	clientOptions := mqtt.NewClientOptions().AddBroker("mqtt://45.77.248.75:1883")
// 	clientOptions.ConnectTimeout = 30 * time.Second
// 	clientOptions.SetClientID(fmt.Sprintf("hieubv-%v", i))
// 	clientOptions.SetUsername("hieubv")
// 	clientOptions.SetPassword("Volio@123")
// 	clientOptions.KeepAlive = 600
// 	clientOptions.CleanSession = true
// 	// Gán các hàm callback cho client MQTT
// 	clientOptions.OnConnect = onConnect
// 	clientOptions.DefaultPublishHandler = onMessageReceived
// 	clientOptions.OnConnectionLost = onConnectionLost

// 	// Tạo MQTT client
// 	client := mqtt.NewClient(clientOptions)

// 	if token := client.Connect(); token.Wait() && token.Error() != nil {
// 		return nil, token.Error()
// 	}
// 	fmt.Printf("Connected to MQTT broker : %v \n", i)

// 	timeStart := time.Now().Unix()
// 	for a := 0; a <= mess; a++ {
// 		// 	// time.Sleep(10 * time.Millisecond)

// 		start := time.Now().Unix()
// 		// 	// token := client.Publish(fmt.Sprintf("volio/dol/test/%v", i), 2, true, fmt.Sprintf("send message index : %v - time start : %v -  %vs", a, start, start-timeStart))
// 		// 	// token := client.Publish(fmt.Sprintf("volio/dol/test/%v", i), 1, true, fmt.Sprintf("send message index : %v - time start : %v -  %vs", a, start, start-timeStart))
// 		// 	// token := client.Publish(fmt.Sprintf("volio/dol/test/%v", i), 0, true, fmt.Sprintf("send message index : %v - time start : %v -  %vs", a, start, start-timeStart))
// 		// 	// token := client.Publish(fmt.Sprintf("volio/dol/test/%v", i), 0, true, fmt.Sprintf("timestamp: %v - total:%v-------::::%s", time.Now().Local(), start-timeStart, m))
// 		// 	// token := client.Publish(fmt.Sprintf("volio/dol/test/%v", 1), 1, true, fmt.Sprintf("timestamp: %v - total:%v-------::::%s", time.Now().Local(), start-timeStart, m))
// 		// 	// client.Publish(fmt.Sprintf("volio/dol/test/%v", i), 2, true, fmt.Sprintf("timestamp: %v - total:%v-------::::%s", time.Now().Local(), start-timeStart, m))
// 		startT := time.Now()
// 		token := client.Publish(fmt.Sprintf("volio/dol/test/%v", i), 2, true, fmt.Sprintf("timestamp: %v - total:%v-------::::%s", time.Now().Local(), start-timeStart, m))
// 		token.Wait()
// 		// go pushMess(a, timeStart, client)

// 		elapsed := time.Since(startT)
// 		if elapsed > time.Second {
// 			fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
// 			fmt.Println("                                                a")
// 			fmt.Println("                                                a")
// 			fmt.Printf("Thời gian thực thi: %s                   a\n", elapsed)
// 			fmt.Println("                                                a")
// 			fmt.Println("                                                a")
// 			fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
// 		}
// 		if a == mess {
// 			fmt.Println("--------------------------------")
// 			fmt.Println("send total message: ", mess, "index ", i, ":", start, "-", start-timeStart)
// 			fmt.Println("************************")
// 		}
// 	}

// 	return client, nil
// }
