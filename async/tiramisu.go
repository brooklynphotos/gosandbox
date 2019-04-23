package main

import (
	"fmt"
	"strings"
	"time"
)

type IngredientChannel chan string

func makeTiramisu() {
	outCh := make(IngredientChannel)
	cheeseCh := make(IngredientChannel)
	eggsCh := make(IngredientChannel)
	beatFatCount := 2
	beatFatCh := make(IngredientChannel, beatFatCount)
	go whisk(eggsCh, 1, beatFatCh)
	go beat(cheeseCh, 1, beatFatCh)
	go beat(beatFatCh, beatFatCount, outCh)
	eggsCh <- "4 eggs"
	close(eggsCh)
	cheeseCh <- "cream cheese"
	close(cheeseCh)
	fmt.Println("Done making tiramisu: " + <-outCh)
}

func whisk(ingredientChannel IngredientChannel, ingredientCount int, outputCh IngredientChannel) {
	ingredients := getAllIngredients(ingredientChannel, ingredientCount)
	ingStr := strings.Join(ingredients, ",")
	fmt.Println("Starting to whisk " + ingStr)
	time.Sleep(700 * time.Millisecond)
	fmt.Println("Done whisking " + ingStr)
	outputCh <- "whisked " + ingStr
}

func beat(ingredientChannel IngredientChannel, ingredientCount int, outputCh IngredientChannel) {
	ingredients := getAllIngredients(ingredientChannel, ingredientCount)
	ingStr := strings.Join(ingredients, ",")
	fmt.Println("Start beating " + ingStr)
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Done beatng " + ingStr)
	outputCh <- "beaten " + ingStr
}

func getAllIngredients(ingredientChannel IngredientChannel, ingredientCount int) []string {
	ings := []string{}
	c := 1
	for ing := range ingredientChannel {
		fmt.Println("Ingredient " + ing + " is ready!")
		ings = append(ings, ing)
		if c == ingredientCount {
			break
		}
		c++
	}
	return ings
}
