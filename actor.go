package main

//go:generate stringer -type=Action

// The Actor struct holds info about an actor I.E the Lad or a Rock
type Actor struct {
	Type       int
	Y          int
	X          int
	Ch         byte
	Dir        Action
	DirRequest Action
	JumpStep   int
}

type Action int

// The Action constans define what the Actor currently is, or is requested to be, doing
const (
	STOPPED Action = iota
	UP
	UPRIGHT
	RIGHT
	DOWNRIGHT
	DOWN
	DOWNLEFT
	LEFT
	UPLEFT
	FALLING
	JUMP // Generic jump set by keyhandler
	JUMPRIGHT
	JUMPUP
	JUMPLEFT
)

var jumpLeft = []Action{UPLEFT, UPLEFT, LEFT, LEFT, DOWNLEFT, DOWNLEFT}
var jumpRight = []Action{UPRIGHT, UPRIGHT, RIGHT, RIGHT, DOWNRIGHT, DOWNRIGHT}
var jumpUp = []Action{UP, UP, STOPPED, DOWN, DOWN}

var jumpPaths = map[Action][]Action{
	JUMPUP:    jumpUp,
	JUMPLEFT:  jumpLeft,
	JUMPRIGHT: jumpRight,
}

var dirs = map[Action]struct {
	x, y int
}{
	STOPPED:   {0, 0},
	UP:        {0, -1},
	UPRIGHT:   {1, -1},
	RIGHT:     {1, 0},
	DOWNRIGHT: {1, 1},
	DOWN:      {0, 1},
	DOWNLEFT:  {-1, 1},
	LEFT:      {-1, 0},
	UPLEFT:    {-1, -1},
	FALLING:   {0, 1},
	JUMP:      {0, 0},
	JUMPRIGHT: {0, 0},
	JUMPLEFT:  {0, 0},
	JUMPUP:    {0, 0},
}

//
// A moving jummp is UR/UR/R/R/DR/DR
// or                UL/UL/L/L/DL/DL
// A standing jump is U/U/-/D/D
//
// 	 ====================
//   ----234-----23------
//   ---1---5----14------
//   --0-----6---05------
//   ====================
//

//
//
//
func MoveActor(a Actor, m MapData) Actor {

loopAgain: // If just started falling we need to retest all conditions

	// If stopped or going left or going right and requst to do something else, then try to do it
	if (a.Dir == STOPPED && a.DirRequest == LEFT) ||
		(a.Dir == STOPPED && a.DirRequest == RIGHT) ||
		(a.Dir == LEFT && a.DirRequest == RIGHT) ||
		(a.Dir == RIGHT && a.DirRequest == LEFT) {
		a.Dir = a.DirRequest
		a.DirRequest = STOPPED

	}

	// Handle starting of jumps
	if a.DirRequest == JUMP {
		switch a.Dir {
		case STOPPED:
			a.DirRequest = a.Dir
			a.Dir = JUMPUP
			a.JumpStep = 0
		case LEFT:
			a.DirRequest = a.Dir
			a.Dir = JUMPLEFT
			a.JumpStep = 0
		case RIGHT:
			a.DirRequest = a.Dir
			a.Dir = JUMPRIGHT
			a.JumpStep = 0
		}
	}

	// Do the jumping
	if a.Dir == JUMPUP || a.Dir == JUMPLEFT || a.Dir == JUMPRIGHT {
		jd := jumpPaths[a.Dir][a.JumpStep]
		if m.Field[a.Y+dirs[jd].y][a.X+dirs[jd].x] == ' ' {
			a.X += dirs[jd].x
			a.Y += dirs[jd].y
			a.JumpStep++
			if a.JumpStep >= len(jumpPaths[a.Dir]) {
				a.Dir = a.DirRequest
			}
		} else {
			// If bumped into something try falling
			a.Dir = FALLING
		}
	}

	// Don't allow player to end up outside of the playfeild
	ClampToPlayfield(&a)

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

	// Don't allow player to end up outside of the playfeild
	ClampToPlayfield(&a)

	return a
}

//
// If walking or jumping of the playfield edges then set actor mode to FALLING
// and make sure the actor stays inside the playfield
//
func ClampToPlayfield(a *Actor) {
	if a.X <= 0 && (a.Dir == LEFT || a.Dir == JUMPLEFT) {
		a.X = 0
		a.Dir = FALLING
		a.DirRequest = STOPPED
	}

	if a.X >= 78 && (a.Dir == RIGHT || a.Dir == JUMPRIGHT) {
		a.X = 78
		a.Dir = FALLING
		a.DirRequest = STOPPED
	}

}
