package main

// The Actor struct holds info about an actor I.E the Lad or a Rock
type Actor struct {
	Type       int
	Y          int
	X          int
	Ch         byte
	Dir        Action
	DirRequest Action
}

// The Action constans define what the Actor currently is, or is requested to be, doing
type Action int

const (
	STOPPED Action = iota
	LEFT
	RIGHT
	UP
	DOWN
	FALLING
	JUMP
)

//
//
//
func MoveActor(a Actor, m MapData) Actor {

loopAgain: // If just started falling we need to retest all conditions

	// If stopped or going left or going right and requst to do something else, then try to do it
	if (a.Dir == STOPPED && a.DirRequest != STOPPED) ||
		(a.Dir == LEFT && a.DirRequest == RIGHT) ||
		(a.Dir == RIGHT && a.DirRequest == LEFT) {
		a.Dir = a.DirRequest
		a.DirRequest = STOPPED
	}

	// If at the edige of the playfield stop and exit
	if (a.Dir == LEFT && a.X == 0) ||
		(a.Dir == RIGHT && a.X >= 78) {
		a.Dir = STOPPED
	}

	// If at a ladder and want to go up
	if (a.DirRequest == UP && m.Field[a.Y][a.X] == 'H') ||
		(a.DirRequest == DOWN && m.Field[a.Y][a.X] == 'H') {
		a.Dir = a.DirRequest
		a.DirRequest = STOPPED
	}

	// If falling then continue doing that until not in free space anymore,
	// then continue the previous direction (if any)
	if a.Dir == FALLING {
		if m.Field[a.Y+1][a.X] == ' ' {
			a.Y++
			return a
		}
		a.Dir = a.DirRequest
	}

	// Climb up until ladder is no more
	if a.Dir == UP {
		if m.Field[a.Y-1][a.X] == 'H' {
			a.Y--
		} else {
			a.Dir = STOPPED
		}
	}

	// Climb down until ladder is no more
	if a.Dir == DOWN {
		if m.Field[a.Y+1][a.X] == 'H' {
			a.Y++
		} else {
			a.Dir = STOPPED
		}
	}

	if a.Dir == LEFT && (m.Field[a.Y][a.X] == ' ' || m.Field[a.Y][a.X] == 'H' || m.Field[a.Y][a.X] == '*') {
		// Stepped out into the void?, the start falling, but remember the previous direction
		if m.Field[a.Y+1][a.X] == ' ' {
			a.DirRequest = a.Dir
			a.Dir = FALLING
			goto loopAgain
		}
		a.X--
	}

	if a.Dir == RIGHT && (m.Field[a.Y][a.X] == ' ' || m.Field[a.Y][a.X] == 'H' || m.Field[a.Y][a.X] == '*') {
		// Stepped out into the void?, the start falling, but remember the previous direction
		if m.Field[a.Y+1][a.X] == ' ' {
			a.DirRequest = a.Dir
			a.Dir = FALLING
			goto loopAgain
		}
		a.X++
	}

	return a
}
