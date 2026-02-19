package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Story struct {
	Chapter      int           `json:"chapter"` // chapter number
	ChapterName  string        `json:"chapterName"`
	ChapterSteps []ChapterStep `json:"ChapterSteps"` // list of all the steps in the chapter
}

type ChapterStep struct {
	StepNo        int            `json:"stepNo"`
	MainString    string         `json:"mainString"` // story string that is shown to the player
	HasChoice     bool           `json:"hasChoice"`
	HasEvent      bool           `json:"hasEvent"`
	NextStep      *int           `json:"nextStep"`                // will be nil when hasChoice=true
	TriggerChoice []ChoiceOption `json:"triggerChoice,omitempty"` // present only if hasChoice=true
	Events        []Event        `json:"events,omitempty"`        // present only if hasEvent=true
}

type ChoiceOption struct {
	ChoiceText        string  `json:"choiceText"`        // name of the choice shown to the player
	ChoiceNextStep    int     `json:"choiceNextStep"`    // next step if player selects this choice
	ChoiceRequirement *string `json:"choiceRequirement"` // nil when null only present when requirements are needed
}

type Event struct {
	EventName        string  `json:"eventName"`
	EventDescription string  `json:"eventDescription"`
	EventRequirement *string `json:"eventRequirement"` // nil when null only present when requirements are needed
}

func loadStory(fileName string) Story {
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(fmt.Errorf("Error opening file: %v", err))
	}

	var story Story
	err = json.Unmarshal(file, &story)
	if err != nil {
		log.Fatal(fmt.Errorf("Error creating struct from %v: %v", fileName, err))
	}
	return story
}
