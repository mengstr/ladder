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
	// If stopped and no request to move then just exit
	if a.Dir == STOPPED && a.DirRequest == STOPPED {
		return a
	}

	// If stopped and requst to do something, then try to do it
	if a.Dir == STOPPED && a.DirRequest != STOPPED {
		a.Dir = a.DirRequest
		a.DirRequest = STOPPED
	}

	// If going LEFT and RIGHT is requested
	if a.Dir == LEFT && a.DirRequest == RIGHT {
		a.Dir = a.DirRequest
		a.DirRequest = STOPPED
	}

	// If going RIGHT and LEFT is requested
	if a.Dir == RIGHT && a.DirRequest == LEFT {
		a.Dir = a.DirRequest
		a.DirRequest = STOPPED
	}

	// If moving to the left at edge then stop and exit
	if a.Dir == LEFT && a.X == 0 {
		a.Dir = STOPPED
		return a
	}

	// If falling then continue doing that until not in free space anymore
	if a.Dir == FALLING {
		if m.Field[a.Y+1][a.X] == ' ' {
			a.Y++
			return a
		}
		a.Dir = STOPPED
		return a
	}

	// If at a ladder and want to go up
	if a.DirRequest == UP && m.Field[a.Y][a.X] == 'H' {
		a.Dir = a.DirRequest
		a.DirRequest = STOPPED
		a.Y--
		return a
	}

	// If at a ladder and want to go down
	if a.DirRequest == DOWN && m.Field[a.Y][a.X] == 'H' {
		a.Dir = a.DirRequest
		a.DirRequest = STOPPED
		a.Y++
		return a
	}

	// Climb up until ladder is no more
	if a.Dir == UP {
		if m.Field[a.Y-1][a.X] == 'H' {
			a.Y--
			return a
		}
		a.Dir = STOPPED
		return a
	}

	// Climb up until ladder is no more
	if a.Dir == DOWN {
		if m.Field[a.Y+1][a.X] == 'H' {
			a.Y++
			return a
		}
		a.Dir = STOPPED
		return a
	}

	if a.Dir == LEFT && (m.Field[a.Y][a.X] == ' ' || m.Field[a.Y][a.X] == 'H' || m.Field[a.Y][a.X] == '*') {
		if m.Field[a.Y+1][a.X] == ' ' {
			a.Dir = FALLING
			a.Y++
			return a
		}
		a.X--
		return a
	}

	// If moving to the right at edge then stop and exit
	if a.Dir == RIGHT && a.X >= 78 {
		a.Dir = STOPPED
		return a
	}

	if a.Dir == RIGHT && (m.Field[a.Y][a.X] == ' ' || m.Field[a.Y][a.X] == 'H' || m.Field[a.Y][a.X] == '*') {
		if m.Field[a.Y+1][a.X] == ' ' {
			a.Dir = FALLING
			a.Y++
			return a
		}
		a.X++
		return a
	}

	return a

}
