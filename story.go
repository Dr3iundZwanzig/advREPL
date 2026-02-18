package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Story struct {
	Chapter      int           `json:"chapter"`
	ChapterName  string        `json:"chapterName"`
	ChapterSteps []ChapterStep `json:"ChapterSteps"`
}

type ChapterStep struct {
	StepNo        int            `json:"stepNo"`
	MainString    string         `json:"mainString"`
	HasChoice     bool           `json:"hasChoice"`
	NextStep      *int           `json:"nextStep"`                // nil when absent/null
	TriggerChoice []ChoiceOption `json:"triggerChoice,omitempty"` // present only if hasChoice=true
}

type ChoiceOption struct {
	ChoiceText        string  `json:"choiceText"`
	ChoiceNextStep    int     `json:"choiceNextStep"`
	ChoiceRequirement *string `json:"choiceRequirement"` // nil when null
}

func loadStory(fileName string) Story {
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(fmt.Errorf("Error opening file: %v", err))
	}

	var story Story
	err = json.Unmarshal(file, &story)
	if err != nil {
		log.Fatal(fmt.Errorf("Error creating struct: %v", err))
	}
	return story
}
