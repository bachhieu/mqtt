package main

// import (
// 	"context"
// 	"fmt"
// 	"net/url"
// 	"os"
// 	"os/signal"
// 	"strconv"
// 	"syscall"
// 	"time"

// 	"github.com/eclipse/paho.golang/autopaho"
// 	"github.com/eclipse/paho.golang/paho"
// )

// const (
// 	clientID = "hieubv" // Change this to something random if using a public test server
// 	topic    = "#"
// )

// func main() {
// 	// App will run until cancelled by user (e.g. ctrl-c)
// 	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
// 	defer stop()

// 	// We will connect to the Eclipse test server (note that you may see messages that other users publish)
// 	u, err := url.Parse("mqtt://draw:draw@localhost:1883")
// 	if err != nil {
// 		panic(err)
// 	}

// 	pass, _ := u.User.Password()
// 	fmt.Println("pass:", pass, "user:", u.User.Username())
// 	fmt.Println("-------------", u.Host)
// 	cliCfg := autopaho.ClientConfig{
// 		ServerUrls: []*url.URL{u},
// 		KeepAlive:  20, // Keepalive message should be sent every 20 seconds
// 		// CleanStartOnInitialConnection defaults to false. Setting this to true will clear the session on the first connection.
// 		CleanStartOnInitialConnection: true,
// 		// SessionExpiryInterval - Seconds that a session will survive after disconnection.
// 		// It is important to set this because otherwise, any queued messages will be lost if the connection drops and
// 		// the server will not queue messages while it is down. The specific setting will depend upon your needs
// 		// (60 = 1 minute, 3600 = 1 hour, 86400 = one day, 0xFFFFFFFE = 136 years, 0xFFFFFFFF = don't expire)
// 		SessionExpiryInterval: 0xFFFFFFFF,
// 		OnConnectionUp: func(cm *autopaho.ConnectionManager, connAck *paho.Connack) {
// 			fmt.Println("mqtt connection up")
// 			// Subscribing in the OnConnectionUp callback is recommended (ensures the subscription is reestablished if
// 			// the connection drops)
// 			if _, err := cm.Subscribe(context.Background(), &paho.Subscribe{
// 				Subscriptions: []paho.SubscribeOptions{
// 					{Topic: topic, QoS: 1},
// 				},
// 			}); err != nil {
// 				fmt.Printf("failed to subscribe (%s). This is likely to mean no messages will be received.", err)
// 			}
// 			fmt.Println("mqtt subscription made")
// 		},
// 		OnConnectError: func(err error) { fmt.Printf("error whilst attempting connection: %s\n", err) },
// 		// eclipse/paho.golang/paho provides base mqtt functionality, the below config will be passed in for each connection
// 		ClientConfig: paho.ClientConfig{
// 			// If you are using QOS 1/2, then it's important to specify a client id (which must be unique)
// 			ClientID: clientID,
// 			// OnPublishReceived is a slice of functions that will be called when a message is received.
// 			// You can write the function(s) yourself or use the supplied Router
// 			OnPublishReceived: []func(paho.PublishReceived) (bool, error){
// 				func(pr paho.PublishReceived) (bool, error) {
// 					fmt.Printf("received message on topic %s; body: %s (retain: %t)\n", pr.Packet.Topic, pr.Packet.Payload, pr.Packet.Retain)
// 					return true, nil
// 				},
// 			},
// 			OnClientError: func(err error) { fmt.Printf("client error: %s\n", err) },
// 			OnServerDisconnect: func(d *paho.Disconnect) {
// 				if d.Properties != nil {
// 					fmt.Printf("server requested disconnect: %s\n", d.Properties.ReasonString)
// 				} else {
// 					fmt.Printf("server requested disconnect; reason code: %d\n", d.ReasonCode)
// 				}
// 			},
// 		},
// 	}

// 	c, err := autopaho.NewConnection(ctx, cliCfg) // starts process; will reconnect until context cancelled
// 	if err != nil {
// 		panic(err)
// 	}
// 	// Wait for the connection to come up
// 	if err = c.AwaitConnection(ctx); err != nil {
// 		panic(err)
// 	}

// 	ticker := time.NewTicker(time.Second)
// 	msgCount := 0
// 	defer ticker.Stop()
// 	for {
// 		select {
// 		case <-ticker.C:
// 			msgCount++
// 			// Publish a test message (use PublishViaQueue if you don't want to wait for a response)
// 			if _, err = c.Publish(ctx, &paho.Publish{
// 				QoS:     1,
// 				Topic:   topic,
// 				Payload: []byte("TestMessage: " + strconv.Itoa(msgCount)),
// 			}); err != nil {
// 				if ctx.Err() == nil {
// 					panic(err) // Publish will exit when context cancelled or if something went wrong
// 				}
// 			}
// 			continue
// 		case <-ctx.Done():
// 		}
// 		break
// 	}

// 	fmt.Println("signal caught - exiting")
// 	<-c.Done() // Wait for clean shutdown (cancelling the context triggered the shutdown)
// }

// // func connect(i int) (mqtt.Client, error) {
// // 	// Tạo một client MQTT
// // 	// clientOptions := mqtt.NewClientOptions().AddBroker("mqtt://ex1.mq.volio.vn:1883")
// // 	clientOptions := mqtt.NewClientOptions().AddBroker("mqtt://ex-test.mq.volio.vn:1884")
// // 	// clientOptions := mqtt.NewClientOptions().AddBroker("mqtt://192.168.1.51:1883")
// // 	clientOptions.ConnectTimeout = 10 * time.Second
// // 	clientOptions.SetClientID(fmt.Sprintf("hieubv-%v", i))
// // 	clientOptions.SetUsername("hieubv")
// // 	clientOptions.SetPassword("Volio@123")
// // 	clientOptions.KeepAlive = 1
// // 	// Gán các hàm callback cho client MQTT
// // 	clientOptions.OnConnect = onConnect
// // 	clientOptions.DefaultPublishHandler = onMessageReceived
// // 	clientOptions.OnConnectionLost = onConnectionLost

// // 	// Tạo MQTT client
// // 	client := mqtt.NewClient(clientOptions)

// // 	if token := client.Connect(); token.Wait() && token.Error() != nil {
// // 		return nil, token.Error()
// // 	}
// // 	fmt.Printf("Connected to MQTT broker : %v \n", i)

// // 	timeStart := time.Now().Unix()
// // 	for a := 0; a <= mess; a++ {
// // 		// time.Sleep(10 * time.Millisecond)

// // 		start := time.Now().Unix()
// // 		// token := client.Publish(fmt.Sprintf("volio/dol/test/%v", i), 2, true, fmt.Sprintf("send message index : %v - time start : %v -  %vs", a, start, start-timeStart))
// // 		// token := client.Publish(fmt.Sprintf("volio/dol/test/%v", i), 1, true, fmt.Sprintf("send message index : %v - time start : %v -  %vs", a, start, start-timeStart))
// // 		// token := client.Publish(fmt.Sprintf("volio/dol/test/%v", i), 0, true, fmt.Sprintf("send message index : %v - time start : %v -  %vs", a, start, start-timeStart))
// // 		// token := client.Publish(fmt.Sprintf("volio/dol/test/%v", i), 0, true, fmt.Sprintf("timestamp: %v - total:%v-------::::%s", time.Now().Local(), start-timeStart, m))
// // 		// token := client.Publish(fmt.Sprintf("volio/dol/test/%v", 1), 1, true, fmt.Sprintf("timestamp: %v - total:%v-------::::%s", time.Now().Local(), start-timeStart, m))
// // 		// client.Publish(fmt.Sprintf("volio/dol/test/%v", i), 2, true, fmt.Sprintf("timestamp: %v - total:%v-------::::%s", time.Now().Local(), start-timeStart, m))
// // 		token := client.Publish(fmt.Sprintf("volio/dol/test/%v", i), 2, true, fmt.Sprintf("timestamp: %v - total:%v-------::::%s", time.Now().Local(), start-timeStart, m))
// // 		token.Wait()
// // 		// go pushMess(a, timeStart, client)
// // 		if a == mess {
// // 			fmt.Println("--------------------------------")
// // 			fmt.Println("send message index ", i, ":", start, "-", start-timeStart)
// // 			fmt.Println("************************")
// // 		}
// // 	}

// // 	return client, nil
// // }
