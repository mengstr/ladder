package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

const screenMinWidth = 80
const screenMinHeight = 22

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

var m MapData
var lad Actor
var stones [5]Actor
var dispensers []XY
var ss = 0

func main() {
	InitLogger(ioutil.Discard, os.Stderr, os.Stderr, os.Stderr)
	Info.Println("Ladder started")

	rand.Seed(int64(time.Now().Nanosecond()))

	// Initialize termbox library and check if the screen dimensions are big enough
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	width, height := termbox.Size()
	if width < screenMinWidth || height < screenMinHeight {
		termbox.Close()
		log.Fatal("Error - screen W/H ", width, "/", height, " is less than ", screenMinWidth, "/", screenMinHeight, "\x10\x0d")
	}
	defer termbox.Close()
	termbox.SetOutputMode(termbox.OutputNormal)

	// Load a level ahs show it
	var ladXY XY
	m, ladXY, dispensers, _ = LoadMap(0)
	InitActor(&lad, LAD, ladXY)
	for i := 0; i < len(stones); i++ {
		InitActor(&stones[i], STONE, dispensers[0])
	}

	// Show the initial map on screen
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
	ticker := time.NewTicker(100 * time.Millisecond)

gameloop:
	for {
		drawStatusline()
		select {
		case <-ticker.C:
			// Test if in SingleStep mode and act accordingly
			if ss == 1 {
				continue
			}
			if ss == 2 {
				ss = 1
			}

			for i := 0; i < 5; i++ {
				if (stones[i].Dir == PENDING) && (rand.Intn(10) == 0) {
					stones[i].DirRequest = FALLING
				}
			}

			// Move the lad according to players whishes
			lad.Ch = '@'
			termbox.SetCell(lad.X, lad.Y, rune(m.Field[lad.Y][lad.X]), termbox.ColorDefault, termbox.ColorDefault)
			MoveActor(&lad, m)
			termbox.SetCell(lad.X, lad.Y, rune(lad.Ch), termbox.ColorDefault, termbox.ColorDefault)

			// Move the stone(s)
			for i := 0; i < 5; i++ {
				stones[i].Ch = byte(48 + i)
				termbox.SetCell(stones[i].X, stones[i].Y, rune(m.Field[stones[i].Y][stones[i].X]), termbox.ColorDefault, termbox.ColorDefault)
				MoveActor(&stones[i], m)
				// Don't draw the stone if it's pending
				if stones[i].Dir != PENDING {
					termbox.SetCell(stones[i].X, stones[i].Y, rune(stones[i].Ch), termbox.ColorDefault, termbox.ColorDefault)
				}
			}
			drawStatusline()
			termbox.Flush()
		case ev := <-keystroke:
			if ev.Type == termbox.EventKey {
				if ev.Key == termbox.KeyEsc {
					break gameloop
				}
				if ev.Key == termbox.KeyEnter {
					ss = 2
				} else if ev.Key == termbox.KeyArrowLeft {
					lad.DirRequest = LEFT
				} else if ev.Key == termbox.KeyArrowRight {
					lad.DirRequest = RIGHT
				} else if ev.Key == termbox.KeyArrowUp {
					lad.DirRequest = UP
				} else if ev.Key == termbox.KeyArrowDown {
					lad.DirRequest = DOWN
				} else if ev.Key == termbox.KeySpace {
					lad.DirRequest = JUMP
				} else {
					lad.DirRequest = STOPPED
				}
			}
		}

	}

}

// //
// //
// //
// func move(o Coords, dir int, reqdir int) Coords {
// 	// 	if m.laddir == STOPPED && m.ladrequest == STOPPED {
// 	// 	return
// 	// }
// 	// xd := 0
// 	// yd := 0
// 	// ladchar := 'g'

// 	// if m.laddir == STOPPED && m.ladrequest != STOPPED {
// 	// 	newcell := m.field[m.lad.y+dirdata[m.ladrequest].y][m.lad.x+dirdata[m.ladrequest].x]
// 	// 	newcellbelow := m.field[m.lad.y+dirdata[m.ladrequest].y+1][m.lad.x+dirdata[m.ladrequest].x]
// 	// 	m.ladrequest == LEFT || m.laddir = RIGHT{}
// 	// }

// 	// if m.laddir == LEFT || m.laddir == RIGHT {

// 	// }

// 	// // Restore screen where the lad currently are
// 	// termbox.SetCell(m.lad.x, m.lad.y, rune(m.field[m.lad.y][m.lad.x]), termbox.ColorDefault, termbox.ColorDefault)

// 	// // Get the value of the next position of the lad
// 	// newcell := m.field[m.lad.y+yd][m.lad.x+xd]

// 	// // Move the lad to it's new position
// 	// m.lad.y += yd
// 	// m.lad.x = +xd

// 	// if m.laddir == 1 || m.laddir == 2 {
// 	// 	if bytes.IndexByte([]byte(" H"), newcell) != -1 {
// 	// 		m.lad.x = m.lad.x + xd
// 	// 	} else {
// 	// 		m.laddir = 0
// 	// 	}
// 	// }

// }

//
//
//
func drawMap() {
	for y := 0; y < len(m.Field); y++ {
		for x := 0; x < len(m.Field[0]); x++ {
			termbox.SetCell(x, y, rune(m.Field[y][x]), termbox.ColorDefault, termbox.ColorDefault)
		}
	}
	drawStatusline()
}

//
//
//
func drawStatusline() {
	//	Lads    5     Level    1     Score     000                 Bonus time    3400
	status := fmt.Sprintf("Lads   %2d     Level   %2d     Score    %04d                 Bonus time    %4d", m.LadsRemaining, m.Level, m.Score, m.Bonustime)
	printXY(0, 20, termbox.ColorDefault, termbox.ColorDefault, status)
	status = fmt.Sprintf("Lad    Dir=%s, DirRequest=%s      ", lad.Dir, lad.DirRequest)
	printXY(0, 21, termbox.ColorDefault, termbox.ColorDefault, status)
	status = fmt.Sprintf("Stone0 Dir=%s, DirRequest=%s  x=%d y=%d    ", stones[0].Dir, stones[0].DirRequest, stones[0].X, stones[0].Y)
	printXY(0, 22, termbox.ColorDefault, termbox.ColorDefault, status)
}

//
//
//
func printXY(x int, y int, fgcolor termbox.Attribute, bgcolor termbox.Attribute, txt string) {
	for i, elem := range txt {
		termbox.SetCell(x+i, y, rune(elem), fgcolor, bgcolor)
	}
}

//
//
//
func InitLogger(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
