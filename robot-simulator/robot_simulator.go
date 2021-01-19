package robot

const N = Dir(0)
const E = Dir(1)
const S = Dir(2)
const W = Dir(3)

// STEP1

func (d Dir) String() string {
	switch d {
	case N:
		return "North"
	case E:
		return "East"
	case W:
		return "West"
	case S:
		return "South"
	}
	panic("invalid dir")
}

func Left() {
	Step1Robot.Dir.turnLeft()
}

func (d *Dir) turnLeft() {
	*d--
	if *d < 0 {
		*d = W
	}
}

func (d *Dir) turnRight() {
	*d = (*d + 1) % 4
}

func Right() {
	Step1Robot.Dir.turnRight()
}

func Advance() {
	to := move(Pos{
		Easting:  RU(Step1Robot.X),
		Northing: RU(Step1Robot.Y),
	}, Step1Robot.Dir)
	Step1Robot.Y = int(to.Northing)
	Step1Robot.X = int(to.Easting)
}

func move(pos Pos, dir Dir) Pos {
	out := pos
	switch dir {
	case N:
		out.Northing++
	case E:
		out.Easting++
	case W:
		out.Easting--
	case S:
		out.Northing--
	}
	return out
}

// STEP2

type Action Command

func StartRobot(cmd chan Command, act chan Action) {
	for i := range cmd {
		act <- Action(i)
	}
	close(act)
}

func Room(rect Rect, robot Step2Robot, act chan Action, rep chan Step2Robot) {
	robot.Easting = rect.Min.Easting
	robot.Northing = rect.Min.Northing

	for i := range act {
		switch i {
		case 'A':
			newPos := move(robot.Pos, robot.Dir)
			if isValidForRoom(rect, newPos) {
				robot.Pos = newPos
			}
		case 'R':
			robot.Dir.turnRight()
		case 'L':
			robot.Dir.turnLeft()
		}
	}
	rep <- robot
}

func isValidForRoom(rect Rect, pos Pos) bool {
	return pos.Easting <= rect.Max.Easting &&
		pos.Easting >= rect.Min.Easting &&
		pos.Northing <= rect.Max.Northing &&
		pos.Northing >= rect.Min.Northing
}

// STEP3
type Action3 struct {
	name string
	cmd  Command
}

var validCommand = map[int32]bool{
	'R': true,
	'L': true,
	'A': true,
}

func Room3(rect Rect, robots []Step3Robot, act chan Action3, rep chan []Step3Robot, log chan string) {
	running := 0
	theRobots := map[string]int{}
	var tmpRobots []Step3Robot
	for idx, j := range robots {
		running++
		if j.Name == "" {
			log <- "no name"
			continue
		}
		if _, exists := theRobots[j.Name]; exists {
			log <- "already exists"
			continue
		}
		if isValidForRoom(rect, j.Pos) && isPositionAvailable(tmpRobots, j.Pos) {
			tmpRobots = append(tmpRobots, j)
			theRobots[j.Name] = idx
		} else {
			log <- "bump"
		}
	}
	for a := range act {
		robotIdx := theRobots[a.name]
		switch a.cmd {
		case 'Q':
			running--
		case 'A':
			newPos := move(robots[robotIdx].Pos, robots[robotIdx].Dir)

			if isValidForRoom(rect, newPos) && isPositionAvailable(robots, newPos) {
				robots[robotIdx].Pos = newPos
			} else {
				log <- "invalid"
			}
		case 'R':
			robots[robotIdx].Dir.turnRight()
		case 'L':
			robots[robotIdx].Dir.turnLeft()
		}
		if running <= 0 {
			close(act)
		}
	}

	rep <- robots
}

func StartRobot3(name string, scr string, act chan Action3, log chan string) {
	for _, i := range scr {
		if _, isValid := validCommand[i]; !isValid {
			log <- "bad command"
			break
		}
		act <- Action3{
			name: name,
			cmd:  Command(i),
		}
	}
	act <- Action3{
		name: name,
		cmd:  'Q', // Signal the robot doesn't have more action
	}
}

func isPositionAvailable(robots []Step3Robot, pos Pos) bool {
	for _, robot := range robots {
		if robot.Pos == pos {
			return false
		}
	}

	return true
}
