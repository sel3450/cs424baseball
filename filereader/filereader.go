package filereader

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sel3450/cs424baseball/baseballplayer"
)

func Read() {
	file, ferr := os.Open("C:/Users/Sierra Laney/Desktop/baseballdudes.txt")
	if ferr != nil {
		panic(ferr)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		t := &baseballplayer.BaseballPlayer{}
		line := scanner.Text()
		split := strings.Split(line, " ")
		t.SetFirstName(split[0])
		t.SetLastName(split[1])
		j, err := strconv.Atoi(split[2])
		if err != nil {
			panic(err)
		}
		t.SetPlateAppearance(j)
		k, err := strconv.Atoi(split[3])
		if err != nil {
			panic(err)
		}
		t.SetAtBats(k)
		l, err := strconv.Atoi(split[4])
		if err != nil {
			panic(err)
		}
		t.SetSingles(l)
		m, err := strconv.Atoi(split[5])
		if err != nil {
			panic(err)
		}
		t.SetDoubles(m)
		n, err := strconv.Atoi(split[6])
		if err != nil {
			panic(err)
		}
		t.SetTriples(n)
		o, err := strconv.Atoi(split[7])
		if err != nil {
			panic(err)
		}
		t.SetHomeruns(o)
		p, err := strconv.Atoi(split[8])
		if err != nil {
			panic(err)
		}
		t.SetWalks(p)
		q, err := strconv.Atoi(split[9])
		if err != nil {
			panic(err)
		}
		t.SetHitByPitch(q)
		//t.ComputeBattingAverage()
		t.ComputeSlugging()
		t.ComputeOnBase()
		t.ComputeOPS()
		//fmt.Println(t.BattingAverage)
		//fmt.Println(t.Slugging)
		//fmt.Println(t.OnbasePercentage)
		//fmt.Println(t.OPS)
	}
}

func Hello() {
	fmt.Println("hi")
}
