package main

import (
	"bufio"
	"fmt"
	"github.com/sumitasok/battleship"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Hi!")

	file, err := os.Open(os.Args[1])
	defer file.Close()

	if err != nil {
		panic(err.Error())
	}

	m := 0
	s := 0
	p1 := ""
	p2 := ""
	attackCount := 0
	strickBy1 := ""
	strickBy2 := ""

	i := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		switch i {
		case 1:
			m, err = strconv.Atoi(scanner.Text())
		case 2:
			s, err = strconv.Atoi(scanner.Text())
		case 3:
			p1 = scanner.Text()
		case 4:
			p2 = scanner.Text()
		case 5:
			attackCount, err = strconv.Atoi(scanner.Text())
		case 6:
			strickBy1 = scanner.Text()
		case 7:
			strickBy2 = scanner.Text()
		default:
			fmt.Printf("trailing content found in the file.")
		}
		i += 1
	}

	bg1, err := battleship.NewBattleGround(m, s, p1)
	if err != nil {
		panic(err)
	}

	print("P1:\n" + bg1.String())

	battleship.Attack(attackCount, bg1, strickBy2)

	fileP1Status, err := os.Create(os.Args[1] + "-P1-status")
	defer fileP1Status.Close()
	if err != nil {
		panic(err)
	}
	print("P1 Attacked:\n" + bg1.String())

	// bg1 status write to file
	w := bufio.NewWriter(fileP1Status)
	w.WriteString(bg1.String())
	w.Flush()
	// bg1 write to io

	bg2, err := battleship.NewBattleGround(m, s, p2)
	if err != nil {
		panic(err)
	}

	print("P2:\n" + bg2.String())

	fileP2Status, err := os.Create(os.Args[1] + "-P2-status")
	defer fileP2Status.Close()
	if err != nil {
		panic(err)
	}

	battleship.Attack(attackCount, bg2, strickBy1)
	// bg2 status write to file
	w2 := bufio.NewWriter(fileP2Status)
	w2.WriteString(bg2.String())
	w2.Flush()
	// bg2 status write to io
	print("P2 Attacked:\n" + bg2.String())
}
