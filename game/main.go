package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"

	"github.com/fatih/color"
)


type card struct {
  text string
  color string
}

var cards []card

var userScore int

func main() {
  initGame()
  startGame()
  displayScore()
  checkRetry()
}

func displayScore() {
  fmt.Println("----------")
  fmt.Println(userScore, " points, Hype!")
  fmt.Println("----------")
}

func initGame() {
  userScore = 0
  var COLORS = []string{"red", "blue", "green", "yellow", "white"}
  cards = makeCards(200, COLORS)
  fmt.Println("----------")
  fmt.Println("UPPER CARD[Only Meaning of Text]")
  fmt.Println("DOWN CARD[Only Color of Text]")
  fmt.Println("If the above condition met, press 'y', if not, press 'n'")
  fmt.Println("START GAME")
}

func checkRetry() {
  fmt.Println("You can retry with type 'y'")
  retry := bufio.NewScanner(os.Stdin)
  retry.Scan()
  if retry.Text() == "y" {
    main()
  }
}

func startGame() {
  for i := 0; i < 100; i += 2 {
    var answer bool
    displayQuestion(cards[i], cards[i + 1], &answer);

    fmt.Println("if upperCard.text = downCard.color, press 'y', if not, press 'n'")
    input := bufio.NewScanner(os.Stdin)
    input.Scan()
    var userAnswer = input.Text() == "y"
    fmt.Println(userAnswer, answer)

    if userAnswer == answer {
      userScore++
    }
  }
}

func drawCard(card card) {
  var t = card.text
  var c = card.color

  if c == "red" {
    color.Red(t)
  } else if c == "blue" {
    color.Blue(t)
  } else if c == "green" {
    color.Green(t)
  } else if c == "yellow" {
    color.Yellow(t)
  } else if c == "white" {
    color.White(t)
  }
}

func displayQuestion(upperCard card, downCard card, answer *bool) {
  drawCard(upperCard)
  fmt.Println("----------")
  drawCard(downCard)
  println("")
  if upperCard.text == downCard.color {
    *answer = true
  } else {
    *answer = false
  }
  fmt.Println("")
}

func makeCards(length int, colors []string) []card {
  var cards []card
  rand.Seed(10)
  for i := 0; i < length; i++ {
    var randomIndex1 = rand.Intn(len(colors))
    var randomIndex2 = rand.Intn(len(colors))
    cards = append(cards, card{text: colors[randomIndex2], color: colors[randomIndex1]})
  }

  return cards
}
