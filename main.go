package main

import (
  "fmt"
  "os"
)
func main() {
	a := [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for {
		printGame(a)
		play(&a, "X")
    if checkWinner(a) == 1 {
      fmt.Println("X Won!!!")
      os.Exit(0)
    }

		printGame(a)
		play(&a, "O")
    if checkWinner(a) == 1 {
      fmt.Println("O Won!!!")
      os.Exit(0)
    }
	}
}

func printGame(a [9]string) {
	for i := 0; i < 7; i += 3 {
		fmt.Printf("%s | %s | %s\n", a[i], a[i+1], a[i+2])
	}
}

func play(a *[9]string, player string) {
	var choice int

	fmt.Print("\r=> ")
	fmt.Scanf("%d", &choice)

	if choice < 1 || choice > 9 {
		play(a, player)
	} else if a[choice-1] == "X" || a[choice-1] == "O" {
		fmt.Println("Already taken!")
    play(a, player)
	} else {
    a[choice-1] = player
	}
}

func checkWinner(a [9]string) int{
  for i := 0; i < 3; i += 3 {
    if a[i] == a[i+1] && a[i+1] == a[i+2] {return 1}
    if a[i] == a[i+3] && a[i+3] == a[i+6] {return 1}
  }
  if a[0] == a[4] && a[4] == a[8] {return 1}
  if a[2] == a[4] && a[4] == a[6] {return 1}

  return 0
}
