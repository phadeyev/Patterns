package main

import (
	"fmt"
	"log"
)

func main() {

	user1 := NewUser("jkulvich")
	user2 := NewUser("vasya")

	ch1 := NewChannel()
	ch1.Subscribe(user1)
	ch1.Subscribe(user2)

	ch2 := NewChannel()
	ch2.Subscribe(user2)

	pub := NewPublisher()
	pub.AddChannel("test", ch1)
	pub.AddChannel("test2", ch2)

	if err := pub.Send("HELLO TO ALL CHANNELS!"); err != nil {
		log.Fatalf("can't send: %s", err)
	}

	if err := pub.Send("HELLO TO CH1!", "test"); err != nil {
		log.Fatalf("can't send: %s", err)
	}

	if err := pub.Send("HELLO FROM CH2", "test2"); err != nil {
		log.Fatalf("can't send: %s", err)
	}

	if err := pub.RemoveChannel("test"); err != nil {
		fmt.Println(err)
	}

	if err := pub.RemoveChannel("test"); err != nil {
		fmt.Println(err)
	}

}
