package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nsf/termbox-go"
)

const screenMinWidth = 80
const screenMinHeight = 22

var m MapData

func main() {

	// Initialize termbox library
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	// Check that the screen is big enough for us
	width, height := termbox.Size()
	if width < screenMinWidth || height < screenMinHeight {
		termbox.Close()
		log.Fatal("Error - screen W/H ", width, "/", height, " is less than ", screenMinWidth, "/", screenMinHeight)
	}
	termbox.SetOutputMode(termbox.OutputNormal)

	m, _ = loadMap(0)
	for i := 0; i < 20; i++ {
		for j := 0; j < 79; j++ {
			//			fmt.Printf("%c", m.field[i][j])
		}
		fmt.Println()
	}
	//	fmt.Println("Lad is at ", m.lad)
	//	fmt.Println("Dispensers at ", m.dispensers)

	drawMap()
	termbox.Flush()

	// Make termbox generate events on a channel for keypresses
	keystroke := make(chan termbox.Event)
	go func() {
		for {
			keystroke <- termbox.PollEvent()
		}
	}()

	// Set up a channel for the game frame rate
	ticker := time.NewTicker(200 * time.Millisecond)

gameloop:
	for {
		select {
		case <-ticker.C:
			m.score++
			drawStatusline()
			termbox.Flush()
		case ev := <-keystroke:
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
				break gameloop
			}
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyArrowLeft {
			}
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyArrowRight {
			}
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyArrowUp {
			}
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyArrowDown {
			}
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeySpace {
			}
		}

	}

}

//
//
//
func drawMap() {
	for y := 0; y < len(m.field); y++ {
		for x := 0; x < len(m.field[0]); x++ {
			termbox.SetCell(x, y, rune(m.field[y][x]), termbox.ColorDefault, termbox.ColorDefault)
		}
	}
	drawStatusline()
}

//
//
//
func drawStatusline() {
	//	Lads    5     Level    1     Score     000                 Bonus time    3400
	status := fmt.Sprintf("Lads   %2d     Level   %2d     Score    %04d                 Bonus time    %4d", m.ladsRemaining, m.level, m.score, m.bonustime)
	printXY(0, 20, termbox.ColorDefault, termbox.ColorDefault, status)
}

//
//
//
func printXY(x int, y int, fgcolor termbox.Attribute, bgcolor termbox.Attribute, txt string) {
	for i, elem := range txt {
		termbox.SetCell(x+i, y, rune(elem), fgcolor, bgcolor)
	}
}
