package main

import "fmt"

func main() {

	// since both struct satisfies writer interface the can be stored in varibale for a common type writer
	// this is called polymorphic behaviour
	var cw Writer = ConsoleWriter{}
	var nw Writer = NetworkWriter{}
	cw.Write([]byte("Write this to console"))
	nw.Write([]byte("Write this to network"))

	var allrounderPlayer AllRounder = Player{} // this will only work if all methods in Player are value reciver for pointer receiver methods pass &Player{}

	allrounderPlayer.Batting()
	allrounderPlayer.Bowling()

	//TYPE CONVERSION

	player, ok := allrounderPlayer.(Player)

	if !ok {
		fmt.Println("Error while converting type")
	} else {
		fmt.Printf("%T", player)

	}

	//BEST PRACTICES
	/*
      break interface in multiple smaller once than one big interface
      design function to accept interfaces and not concrete types
	*/
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct {
}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

type NetworkWriter struct {
}

func (cw NetworkWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

//compose multiple interfaces

type Player struct {
}

func (p Player) Bowling() {
	fmt.Println("Bowling")
}
func (p Player) Batting() {
	fmt.Println("Batting")
}

type Bowler interface {
	Bowling()
}

type Batsman interface {
	Batting()
}

type AllRounder interface {
	Bowler
	Batsman
}
