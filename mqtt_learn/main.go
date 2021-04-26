package main

import (
    "fmt"
    mqtt "github.com/eclipse/paho.mqtt.golang"
    "log"
    "os"
    "time"
)

func sub(client mqtt.Client, topic string) {
    token := client.Subscribe(topic, 1, nil)
    token.Wait()
    log.Printf("Subscribed to topic %s", topic)
}
func public(client mqtt.Client, topic string) {
    for i := 0;i < 10;i ++ {
        payload := fmt.Sprintf("Paylod data[%d]", i)
        token := client.Publish(topic, 0, false, payload)
        token.Wait()
        time.Sleep(1 * time.Second)
    }
}

func main() {
    args := os.Args
    if len(args) < 4 {
        return
    }
    clienId := args[1]
    subTopic := args[2]
    pubTopic := args[3]
    
    var broker = "xa"
    var port = 1883
    opts := mqtt.NewClientOptions()
    opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
    opts.SetClientID(clienId)
    opts.SetUsername("emqx")
    opts.SetPassword("public")
    opts.SetDefaultPublishHandler(func(client mqtt.Client, message mqtt.Message) {
        log.Printf("Recieve message : %s from topic: %s\n", message.Payload(), message.Topic())
    })
    opts.OnConnect = func(client mqtt.Client) {
        log.Println("Connected")
    }
    
    opts.OnConnectionLost = func(client mqtt.Client, err error) {
        log.Printf("Connect lost: %v", err)
    }
    
    client := mqtt.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }
    log.Println("start???")
    sub(client, subTopic)
    public(client, pubTopic)
    
    client.Disconnect(250)
}
