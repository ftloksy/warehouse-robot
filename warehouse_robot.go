
/*
____This is GoLang____

the answer is:

1. Robot A: ball 	Robot B: ball,		 All Step:  NEE
2. Robot A: hat, ball 	Robot B: box, hat,	 All Step:  NNSSE
3. Robot A: ball, ball 	Robot B: hat, ball,	 All Step:  ENE

*/


package main

import (
	"fmt"
)

type location struct {
	x int
	y int
}

type grid struct {
	l location
	content byte
}

// Calculates the distance between two locations (robot and target)
func distance (robot location, target location) (dis location) {
	dis.x = target.x - robot.x
	dis.y = target.y - robot.y
	return
}

/* 
A_rl is A Robot's location.
B_rl is B Robot's location.
A_WP is A Robot collection thing.
B_WP is B Robot collection thing.
*/

// Finds the common target point for both robots based on their current locations and the items to collect
func whereTpoint (A_rl location, B_rl location, A_WP []location, B_WP []location)(t_point location ) {
	for i := 0; i < len(A_WP) ; i++ {
		var A_d location
		var B_d location

		A_d = distance(A_rl, A_WP[i])
		for j := 0; j < len(B_WP); j++ {
			B_d = distance(B_rl, B_WP[j])
			if A_d == B_d {
				t_point = B_d
				break
			}
		}
		if A_d == B_d {
			break
		}
	}
	return
}

// Generates the steps needed to walk from the robot's current location to the target location
func walk_toTarget(robot location, dis location)(walkStep map[string][]location) {
	var walkStep_xf []location
	var walkStep_yf []location
	var walk_xf string
	var walk_yf string
	var n_x int
	var n_y int

	walkStep = make(map[string][]location)

	if dis.x > 0 {
		for ix := robot.x + 1 ; ix <= robot.x + dis.x ; ix ++ {
			walkStep_xf = append(walkStep_xf, location{x: ix, y: robot.y})
			walk_xf = walk_xf + "E"
		}
	}
	if dis.x < 0 {
		for ix := robot.x - 1 ; ix >= robot.x + dis.x ; ix -- {
			walkStep_xf = append(walkStep_xf, location{x: ix, y: robot.y})
			walk_xf = walk_xf + "W"
		}
	}

	if len(walkStep_xf) - 1 >= 0 {
		n_x = walkStep_xf[len(walkStep_xf) - 1].x
	} else {
		n_x = robot.x
	}

	if dis.y > 0 {
		for iy := robot.y + 1 ; iy <= robot.y + dis.y ; iy ++ {
			walkStep_xf = append(walkStep_xf, location{x: n_x, y: iy} )
			walk_xf = walk_xf + "N"
		}
	}
	if dis.y < 0 {
		for iy := robot.y - 1 ; iy >= robot.y + dis.y ; iy -- {
			walkStep_xf = append(walkStep_xf, location{x: n_x, y: iy} )
			walk_xf = walk_xf + "S"
		}
	}

	walkStep[walk_xf] = walkStep_xf

	if dis.y > 0 {
		for iy := robot.y + 1 ; iy <= robot.y + dis.y ; iy ++ {
			walkStep_yf = append(walkStep_yf, location{x: robot.x , y: iy} )
			walk_yf = walk_yf + "N"
		}
	}
	if dis.y < 0 {
		for iy := robot.y - 1 ; iy >= robot.y + dis.y ; iy -- {
			walkStep_yf = append(walkStep_yf, location{x: robot.x , y: iy} )
			walk_yf = walk_yf + "S"
		}
	}

	if len(walkStep_yf) - 1 >= 0 {
		n_y = walkStep_yf[len(walkStep_yf) - 1].y
	} else {
		n_y = robot.y
	}

	if dis.x > 0 {
		for ix := robot.x + 1 ; ix <= robot.x + dis.x ; ix ++ {
			walkStep_yf = append(walkStep_yf, location{x: ix, y: n_y})
			walk_yf = walk_yf + "E"
		}
	}
	if dis.x < 0 {
		for ix := robot.x - 1 ; ix >= robot.x + dis.x ; ix -- {
			walkStep_yf = append(walkStep_yf, location{x: ix, y: n_y})
			walk_yf = walk_yf + "S"
		}
	}

	walkStep[walk_yf] = walkStep_yf
	return
}

// Checks if all elements of Tester exist in Pool
func allinPool(Pool []location, Tester []location)(iH bool) {
	var result int
	result = 0

        for _,element := range Pool{
	    for _,t_element := range Tester{
		if element == t_element {
			result += 1
			break
		}
	    }
        }
        if result == len(Tester){
	    iH = true
        }
	return
}

/*** 
a_sl is Robot A start location
b_sl is Robot B start location
a_c is Robot A collect item location list.
b_c is Robot B collect item location list.
nllp is nullPool , it is location list.
     In here, don't have anything.
a_fl is Robot A final location
b_fl is Robot B final location
abstep is they ( Robot A and Robot B ) howto walk to collect item.
***/

// Chooses the path for both robots based on their starting locations, collected items, and the null pool
func choicePath (a_sl location, b_sl location, a_c []location, b_c []location, nllp []location) (abstep string, a_fl location, b_fl location) {

    /* Robot A: ball 	Robot B: ball */
	var t_point location
	t_point = whereTpoint(a_sl, b_sl, a_c, b_c)
	a_ws := walk_toTarget(a_sl, t_point)
	b_ws := walk_toTarget(b_sl, t_point)
	for k, v := range a_ws {
		if allinPool(nllp, v[0:(len(v)-1)]) {
			if allinPool(nllp, b_ws[k][0:(len(v)-1)]){
				abstep = k
				a_fl = a_ws[k][len(a_ws[k]) - 1]
				b_fl = b_ws[k][len(a_ws[k]) - 1]
			} else {
				abstep = "MISSION FALSE"
			}
		}
	}
	return
}

func main() {
	/* Set t byte is hat */
	/* Set l byte is ball */
	/* Set x byte is box */
	/* Set N byte is NULL */

	t := byte("t"[0])
	l := byte("l"[0])
	x := byte("x"[0])
	N := byte("L"[0])

	/* create three Pools and record what goods are there. */

	var boxPool []location
	var ballPool []location
	var hatPool []location
	var nullPool []location

	var A_robot location
	var B_robot location

	A_robot = location{2, 2}
	B_robot = location{0, 0}

	wareHouse := [5][5]grid{
		{
			grid{l: location{0,0}, content: N},
			grid{l: location{1,0}, content: t},
			grid{l: location{2,0}, content: N},
			grid{l: location{3,0}, content: N},
			grid{l: location{4,0}, content: x},
	        },
		{
			grid{l: location{0,1}, content: N},
			grid{l: location{1,1}, content: N},
			grid{l: location{2,1}, content: l},
			grid{l: location{3,1}, content: N},
			grid{l: location{4,1}, content: N},
	        },
		{
			grid{l: location{0,2}, content: x},
			grid{l: location{1,2}, content: N},
			grid{l: location{2,2}, content: N},
			grid{l: location{3,2}, content: l},
			grid{l: location{4,2}, content: N},
	        },
		{
			grid{l: location{0,3}, content: N},
			grid{l: location{1,3}, content: N},
			grid{l: location{2,3}, content: N},
			grid{l: location{3,3}, content: N},
			grid{l: location{4,3}, content: l},
	        },
		{
			grid{l: location{0,4}, content: N},
			grid{l: location{1,4}, content: N},
			grid{l: location{2,4}, content: t},
			grid{l: location{3,4}, content: x},
			grid{l: location{4,4}, content: N},
	        },
	}

	for i := len(wareHouse) - 1 ; i >= 0; i -- {
		for j := 0; j < len(wareHouse[i]); j ++ {
			switch wareHouse[i][j].content {
			case byte("t"[0]):
				hatPool = append(hatPool, wareHouse[i][j].l) 
			case byte("l"[0]):
				ballPool = append(ballPool, wareHouse[i][j].l)
			case byte("x"[0]):
				boxPool = append(boxPool, wareHouse[i][j].l)
			case byte("L"[0]):
				nullPool = append(nullPool, wareHouse[i][j].l)
			}
		}
	}

    /* 1. 	Robot A: ball 	Robot B: ball */
	as, _, _ := choicePath(A_robot, B_robot, ballPool, ballPool, nullPool)
	fmt.Println("1. Robot A: ball 	Robot B: ball,\t\t All Step: ", as)

    /*2.  Robot A: hat, ball 	Robot B: box, hat */
	as, af, bf := choicePath(A_robot, B_robot, hatPool, boxPool, nullPool)
	as2, _, _ := choicePath(af, bf, ballPool, hatPool, nullPool)
	fmt.Println("2. Robot A: hat, ball 	Robot B: box, hat,\t All Step: ", as + as2)

    /*3.  Robot A: ball, ball 	Robot B: hat, ball */
	as, af, bf = choicePath(A_robot, B_robot, ballPool, hatPool, nullPool)
	as2, _, _ = choicePath(af, bf, ballPool, ballPool, nullPool)
	fmt.Println("3. Robot A: ball, ball 	Robot B: hat, ball,\t All Step: ", as + as2)
}
