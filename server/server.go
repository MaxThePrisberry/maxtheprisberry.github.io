package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
	"log"
	"helpers"
//	"errors"
)

const gamePort = ":1234"

var upgrader = websocket.Upgrader{}
var games []Game //Slice of current games open or running
var pprofs map[string]PlayerProfile //Slice of current players

type PlayerProfile struct {
	PName string
	Connection *websocket.Conn
}

func sendPlayerPacket(pid string) (err error) {
	message := append(append(append(append([]byte("{\"PID\":\""),[]byte(pid)...),[]byte("\",\"Name\":\"")...),[]byte(pprofs[pid].PName)...),[]byte("\"}")...)
	fmt.Printf("Your message is: %v\n", string(message))
	err = pprofs[pid].Connection.WriteMessage(websocket.TextMessage,message)
	return
}

type Game struct {
	P1PID, P2PID, Turn string
	P1Checkers, P2Checkers [12][3]int
}

func newConnection(w http.ResponseWriter, r *http.Request) {
	//Enact websocket handshakes and get connection object
	upgrader.CheckOrigin = func(r *http.Request) bool {return true} //Insecure: permits cross-site forgeries
	conn, err := upgrader.Upgrade(w, r, nil)//"conn" should be the first result here
	if err != nil {
		fmt.Println(err)
		return
	}

	//Make new new PID and player profile
	pid := helpers.RandString(10)
	for _, found := pprofs[pid]; found; _, found = pprofs[pid] {
		pid = helpers.RandString(10)
	}
	pprofs[pid] = PlayerProfile{Connection:conn}

	//Send new PID to computer in a player info packet
	if err := sendPlayerPacket(pid); err != nil {
		fmt.Println(err)
		return
	}

	//Call newGame()
	go newGame(pid)
}

func newGame(pid string) {
	//Wait for new player info packet to signify that the computer is ready to be put into a game

	//Call findGame()
}

func findGame(pid string) {
	//Search through the slice "games" for games with only one player. If found, add player and then call runGame() and return

	//Open a new Game and append it to the slice "games"
}

func runGame(gameIndex int) {
	game := games[gameIndex]
	//Send the opposing player's info packet to each player

	for len(game.P1Checkers) > 0 && len(game.P2Checkers) > 0 {
		//Check if the connection for who's turn it is is still active. If not, remove all players' checkers and send the packet to the opposing player. Then remove the game, call newGame() for active player, and return

		//Call PlayerMove(). If "false", remove all players' checkers and send the packet to the opposing player. Then remove the game, call newGame() for active player, and return

		//Change turn in game

	}
	//Send both players a board state packet to let them know the game is over

	//Call newGame() for both players
}

func PlayerMove(game *Game, pid string) bool {
	for {
		//Send the player with pid a game state packet, signifying that it's their turn

		//Wait for move packet. If error occurs (a.k.a. connection cut), return "false"

		//Check if move packet is legal. If legal, update players' checkers states in game and return "true"
	}
}

func main() {
	//Have the newConnection function handle all connections
	http.HandleFunc("/", newConnection)
	//Start listening at gamePort
	log.Fatal(http.ListenAndServe(gamePort, nil))
}
