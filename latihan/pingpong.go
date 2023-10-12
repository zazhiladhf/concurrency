package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
PINGPONG APPS
2 player => 2 goroutine

kondisi kalah : saat flag/counter/random number, habis dibagi 11 (random % 11 == 0 THAN break)

Contoh :
Player A = Hit 1 // counter 8
Player B = Hit 2 // counter 3
Player A = Hit 3 // counter 24
Player B = Hit 4 // counter 33

Player B kalah, total hit : 4, kalah di nomor 33

Player A = Hit 1 // counter 8
Player B = Hit 2 // counter 9
Player A = Hit 3 // counter 22

Player A kalah, total hit : 3, kalah di nomor 22

Requirement :
- Struct Player {
Name string
Hit int
}

a := Player{Name: A, Hit:0}
b := Player{Name: B, Hit:0}

a.Hit++
b.Hit++

*/

/*
   struct untuk korek
   hits        => untuk proses buka tutup korek api
   lastPlayer  => player yang sedang memegang korek
*/
type Player struct {
	Name	string
	Hit		int
}

// break point
const BreakPoint = 11

// var counter int
// var userToken map[string]time.Time

func main() {
	// ball := make(chan int)
	player := make(chan *Player)
	done := make(chan *Player)

	players := []string{"Zazhil", "Budi"}

	for _, p := range players {
		go play(p, player, done)
	}

    // initialize channel kosong
	player <-new(Player)

	// tambahin function finish
	finish(done)

    // proses pengeluaran nilai
	// <-player
}

func play(name string, player, done chan *Player) {
	randomizer := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	min := 1
	max := 100

	// lakukan perulangan terus menerus
	for {
		select {
		// akan di eksekusi jika ada data yg dikirim ke channel korek
		case p := <-player:
			// mengambil angka random
			v := randomizer.Intn(max-min) + min
			time.Sleep(500 * time.Millisecond)
			// melakukan proses buka tutup korek
			p.Name = name
			p.Hit++

			fmt.Println("Player: ", p.Name, ", Hit: ", p.Hit, " got value", v)

            // proses pengiriman korek antar player
			if v%BreakPoint == 0 {
                // jika oke, maka akan mengirim value ke channel `done`
				done <- p

                // return akan memberhentikan perulangan
				return
			}

			player <- p
		}
	}
}

func finish(done chan *Player){
    for {
		select {
        // jika ada data yang masuk pada channel done,
        // maka game selesai
		case d := <-done:
			fmt.Println(d.Name, "kalah pada hit ke", d.Hit)
			return
		}
	}
}

// 		// userToken["token"] = time.Now()
// 		// counter++

// 		time.Sleep(1 * time.Second)
// 		// player 1 hit ball to player 2
// 		val := <-ball
// 		log.Println("Player :", name, ", hit 1, got value", val)
// 		if val%11 == 0 {
// 			log.Println("Player :", name, "fail in value", val)
// 			done <- true
// 			break
// 		}

// 		// proses / teknik pukulan player
// 		// val++
// 		val = rand.Intn(100-1) + 1

// 		log.Println("player", name, "hit the ball to another player with value", val)
// 		// ball will be delivered to player 2
// 		ball <- val
// 	}
// }
