package queueManager

import (
  "github.com/streadway/amqp"
  "fmt"
)

type QueueManager struct {
    Request *Queue
    Response *Queue
    Uri string        
    Connection *amqp.Connection
}

func (queuemanager *QueueManager) Init() bool {
    //init producer & consumer queues
    fmt.Println("Initializing the QueueManager")
    queuemanager.Uri = "amqp://guest:guest@localhost:5672/"
    conn, _ := amqp.Dial(queuemanager.Uri)
    
    queuemanager.Connection = conn

    fmt.Println("Initializing the request queue....")
    queuemanager.Request = new(Queue)
    queuemanager.Request.setconnection(queuemanager.Connection)
 
    fmt.Println("Initializing the response queue....")
    queuemanager.Response = new(Queue)
    queuemanager.Response.setconnection(queuemanager.Connection)

    return true     
}

func (queuemanager *QueueManager) reinit() bool {

    return true
}

func (queuemanager *QueueManager) senddatarequest() bool {

    return true
}

func (queuemanager *QueueManager) receivedatarequest() {

}

func (queuemanager *QueueManager) senddataresponse() bool {
    return true
}

func (queuemanager *QueueManager) receivedataresponse() {


} 
