/*
Sierra Laney
CS424-01
10/5/2021
Purpose: Baseball Player Calculator that ranks based off of OPS
and reads data from user-inputted file.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
Holds the computed and read values for a baseball player.
*/
type BaseballPlayer struct {
	firstName        string  // string of baseball player's first name from text file
	lastName         string  // string of baseball player's last name from text file
	plateAppearance  int     // int of baseball player's plate appearance from text file
	atBats           int     // int of baseball player's at bats from text file
	singles          int     // int of baseball player's singles from text file
	doubles          int     // int of baseball player's doubles from text file
	triples          int     // int of baseball player's triples from text file
	homeRuns         int     // int of baseball player's homeruns from text file
	walks            int     // int of baseball player's walks from text file
	hitByPitch       int     // int of baseball player's hit by pitch from text file
	battingAverage   float64 // float of sum of hits divided by at bats, computed
	slugging         float64 // float of the total # of bases divided by at bats, computed
	onbasePercentage float64 // float of the sum of all hits, walks, hbp divided by plate appearances, computed
	ops              float64 // float of the sum of on-base percentage and slugging percentage, computed
}

/*
Sets the string to a baseball player's object's first name field.
Arguments: string
*/
func (player *BaseballPlayer) SetFirstName(firstname string) {
	player.firstName = firstname
}

/*
Sets the string to a baseball player's object's last name field.
Arguments: string
*/
func (player *BaseballPlayer) SetLastName(lastname string) {
	player.lastName = lastname
}

/*
Sets the int to a baseball player's object's plate appearance field.
Arguments: int
*/
func (player *BaseballPlayer) SetPlateAppearance(plateappearance int) {
	player.plateAppearance = plateappearance
}

/*
Sets the int to a baseball player's object's at bats field.
Arguments: int
*/
func (player *BaseballPlayer) SetAtBats(atbats int) {
	player.atBats = atbats
}

/*
Sets the int to a baseball player's object's singles field.
Arguments: int
*/
func (player *BaseballPlayer) SetSingles(singles int) {
	player.singles = singles
}

/*
Sets the int to a baseball player's object's doubles field.
Arguments: int
*/
func (player *BaseballPlayer) SetDoubles(doubles int) {
	player.doubles = doubles
}

/*
Sets the int to a baseball player's object's triples field.
Arguments: int
*/
func (player *BaseballPlayer) SetTriples(triples int) {
	player.triples = triples
}

/*
Sets the int to a baseball player's object's home runs field.
Arguments: int
*/
func (player *BaseballPlayer) SetHomeruns(homeruns int) {
	player.homeRuns = homeruns
}

/*
Sets the int to a baseball player's object's walks field.
Arguments: int
*/
func (player *BaseballPlayer) SetWalks(walks int) {
	player.walks = walks
}

/*
Sets the int to a baseball player's object's hits by pitch field.
Arguments: int
*/
func (player *BaseballPlayer) SetHitByPitch(hbp int) {
	player.hitByPitch = hbp
}

/*
Computes the batting average based off of baseball player's singles, doubles,
triples, homeruns divided by at bats.
Note: ints are converted to float64 to not lose precision.
*/
func (player *BaseballPlayer) ComputeBattingAverage() {
	var battingaverage float64 = (float64(player.singles) + float64(player.doubles) +
		float64(player.triples) + float64(player.homeRuns)) / float64(player.atBats)
	player.battingAverage = battingaverage
}

/*
Computes the total # of bases divided by at bats.
Note: ints are converted to float64 to not lose precision.
*/
func (player *BaseballPlayer) ComputeSlugging() {
	var slugging float64 = (float64(player.singles) + (2 * float64(player.doubles)) +
		(3 * float64(player.triples)) + (4 * float64(player.homeRuns))) / float64(player.atBats)
	player.slugging = slugging
}

/*
Computes the sum of all hits, walks and hit-by-pitch divided by the number of plate appearances.
Note: ints are converted to float64 to not lose precision.
*/
func (player *BaseballPlayer) ComputeOnBase() {
	var onbasepercentage float64 = (float64(player.singles) + float64(player.doubles) +
		float64(player.triples) + float64(player.homeRuns) + float64(player.walks) +
		float64(player.hitByPitch)) / float64(player.plateAppearance)
	player.onbasePercentage = onbasepercentage
}

/*
Computes the sum of on-base percentage and slugging percentage.
Note: ints are converted to float64 to not lose precision.
*/
func (player *BaseballPlayer) ComputeOPS() {
	var tempOps float64 = player.slugging + player.onbasePercentage
	player.ops = tempOps
}

/*
Reads a text file from a user inputted string, creates a baseball player object, sets the fields of a
baseball player object, computes the specific fields, appends it to a baseball player array, and sorts
the baseball player array based on ops.
Arguments: string
*/
func ReadAndCompute(path string) []BaseballPlayer {
	var bbpArray []BaseballPlayer
	file, ferr := os.Open(path) // attempt to open the file path
	if ferr != nil {            // if error was found, abort the function
		panic(ferr)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var t BaseballPlayer              // create a new baseball player
		line := scanner.Text()            // grab a line
		split := strings.Split(line, " ") // split based off of a space
		t.SetFirstName(split[0])
		t.SetLastName(split[1])
		j, err := strconv.Atoi(split[2])
		if err != nil { // if error was found, abort the function
			panic(err)
		}
		t.SetPlateAppearance(j)
		k, err := strconv.Atoi(split[3])
		if err != nil { // if error was found, abort the function
			panic(err)
		}
		t.SetAtBats(k)
		l, err := strconv.Atoi(split[4])
		if err != nil { // if error was found, abort the function
			panic(err)
		}
		t.SetSingles(l)
		m, err := strconv.Atoi(split[5])
		if err != nil { // if error was found, abort the function
			panic(err)
		}
		t.SetDoubles(m)
		n, err := strconv.Atoi(split[6])
		if err != nil { // if error was found, abort the function
			panic(err)
		}
		t.SetTriples(n)
		o, err := strconv.Atoi(split[7])
		if err != nil { // if error was found, abort the function
			panic(err)
		}
		t.SetHomeruns(o)
		p, err := strconv.Atoi(split[8])
		if err != nil { // if error was found, abort the function
			panic(err)
		}
		t.SetWalks(p)
		q, err := strconv.Atoi(split[9])
		if err != nil { // if error was found, abort the function
			panic(err)
		}
		t.SetHitByPitch(q)
		t.ComputeBattingAverage()
		t.ComputeSlugging()
		t.ComputeOnBase()
		t.ComputeOPS()

		bbpArray = append(bbpArray, t) // add to the array
	}
	sort.Slice(bbpArray, func(p, q int) bool {
		return bbpArray[p].ops > bbpArray[q].ops // sort the ops in ascending order
	})

	return bbpArray // return the sorted array
}

/*
Format the report, and prints the report.
Arguments: array
*/
func Report(bbArray []BaseballPlayer) {

	fmt.Println("\nBASEBALL TEAM REPORT ---", len(bbArray), "PLAYERS FOUND IN FILE")
	fmt.Println("")
	fmt.Println("  PLAYER NAME\t    BAT AVG\tSLUGGING\tONBASE%\t\tOPS")
	fmt.Println("----------------------------------------------------------------------")
	for _, t := range bbArray {
		name := t.lastName + ", " + t.firstName
		fmt.Printf("%-20s", name)
		fmt.Printf("%.4f", t.battingAverage)
		fmt.Printf("\t")
		fmt.Printf("%.4f", t.slugging)
		fmt.Printf("\t\t")
		fmt.Printf("%.4f", t.onbasePercentage)
		fmt.Printf("\t\t")
		fmt.Printf("%.4f", t.ops)
		fmt.Println("")
	}
}

/*
Runs the main program.
*/
func main() {
	// Below is the user prompt.
	fmt.Println("Welcome to the player stats calculator. Please provide a valid file path",
		"to a player text file, and I will return a report based on the highest OPS rankings: ")
	// C:/Users/Sierra Laney/Desktop/baseballdudes.txt
	reader := bufio.NewReader(os.Stdin)
	filePath, err := reader.ReadString('\n')
	if err != nil { // if error was found, abort the function
		panic(err)
	}

	filePath = strings.TrimSpace(filePath) // trims the space off of the path

	baseballplayers := ReadAndCompute(filePath) // reads and computes based on the file path
	Report(baseballplayers)                     // prints off the baseball player reports
}
