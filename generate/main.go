package main

import (
	"flag"
	"fmt"
	"os"
	"resume/generate/ui"

	tea "charm.land/bubbletea/v2"
)

func main() {
	resumePath := flag.String("resume", "resume.json", "Path to the resume.json to use as a template")

	flag.Parse()

	resume, err := FindAndReadResume(*resumePath)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		fmt.Printf("Could not find the resume.json file! Searched: \n - %s\n - data/%s\n", *resumePath, *resumePath)
		return
	}

	p := tea.NewProgram(ui.InitialModel(resume))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
