package queueManager

import (
     "github.com/streadway/amqp"
     "fmt"
)

type Queue struct {
    Channel *amqp.Channel    
    Name string
    Durable bool
    Deletewhenused bool
    Exclusive bool
    Nowait bool
    Arguments *amqp.Table
    QueueReturn amqp.Queue
}

func (Q *Queue) init(conn *amqp.Connection) {
}

func (Q *Queue) setconnection(conn *amqp.Connection) {
    ch, _ := conn.Channel()
    Q.Channel = ch
    Q.Arguments = nil
    Q.Durable = false
    Q.Deletewhenused = false
    Q.Exclusive = false
    Q.Nowait = false
    q, _ := Q.Channel.QueueDeclare(
                Q.Name,
                Q.Durable,
                Q.Deletewhenused,
                Q.Exclusive,
                Q.Nowait,
                nil,
    )
    Q.QueueReturn = q
    fmt.Printf("LOG: Queue has been Initialized : %s", Q.QueueReturn.Name) 
 
}

func (Q *Queue) close() {


}

func (Q *Queue) send(payload []byte) {
    _ = Q.Channel.Publish(
                  "",
                  Q.QueueReturn.Name,
                  false,
                  false,
                  amqp.Publishing{
                      Body:[]byte(payload),
                  })
}

func (Q *Queue) receive() {


}
