package server

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Data structures for portfolio data
type PersonalInfo struct {
	Name     string  `json:"name"`
	Nickname string  `json:"nickname"`
	Title    string  `json:"title"`
	Location string  `json:"location"`
	Timezone string  `json:"timezone"`
	About    About   `json:"about"`
	Contact  Contact `json:"contact"`
}

type About struct {
	Intro      string   `json:"intro"`
	WhatIDo    string   `json:"whatIDo"`
	Background []string `json:"background"`
	Philosophy string   `json:"philosophy"`
}

type Contact struct {
	Email            string   `json:"email"`
	GitHub           string   `json:"github"`
	LinkedIn         string   `json:"linkedin"`
	Portfolio        string   `json:"portfolio"`
	PreferredContact string   `json:"preferredContact"`
	ResponseTime     string   `json:"responseTime"`
	AvailableFor     []string `json:"availableFor"`
	Specializations  []string `json:"specializations"`
}

type Experience struct {
	Title        string   `json:"title"`
	Company      string   `json:"company"`
	Period       string   `json:"period"`
	Location     string   `json:"location"`
	Type         string   `json:"type"`
	Current      bool     `json:"current"`
	Details      []string `json:"details"`
	Technologies []string `json:"technologies"`
}

type Skill struct {
	Name       string `json:"name"`
	Percentage int    `json:"percentage"`
	Experience string `json:"experience"`
}

type AsciiArt struct {
	Logo    string `json:"logo"`
	Contact string `json:"contact"`
}

type PortfolioData struct {
	Personal    PersonalInfo       `json:"personal"`
	Experiences []Experience       `json:"experiences"`
	Skills      map[string][]Skill `json:"skills"`
	TechFacts   []string           `json:"techFacts"`
	AsciiArt    AsciiArt           `json:"asciiArt"`
}

// DataLoader handles loading and caching portfolio data
type DataLoader struct {
	data     *PortfolioData
	dataPath string
}

// NewDataLoader creates a new data loader instance
func NewDataLoader(dataPath string) *DataLoader {
	return &DataLoader{
		dataPath: dataPath,
	}
}

// LoadData loads portfolio data from JSON file
func (dl *DataLoader) LoadData() error {
	// Get the absolute path for the data file
	absPath, err := filepath.Abs(dl.dataPath)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	// Check if file exists
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return fmt.Errorf("data file not found at: %s", absPath)
	}

	// Read the JSON file
	data, err := os.ReadFile(absPath)
	if err != nil {
		return fmt.Errorf("failed to read data file: %w", err)
	}

	// Parse JSON
	var portfolioData PortfolioData
	if err := json.Unmarshal(data, &portfolioData); err != nil {
		return fmt.Errorf("failed to parse JSON data: %w", err)
	}

	dl.data = &portfolioData
	return nil
}

// GetData returns the loaded portfolio data
func (dl *DataLoader) GetData() *PortfolioData {
	return dl.data
}

// IsLoaded checks if data has been loaded
func (dl *DataLoader) IsLoaded() bool {
	return dl.data != nil
}

// GetPersonalInfo returns personal information
func (dl *DataLoader) GetPersonalInfo() *PersonalInfo {
	if dl.data == nil {
		return nil
	}
	return &dl.data.Personal
}

// GetExperiences returns all experiences
func (dl *DataLoader) GetExperiences() []Experience {
	if dl.data == nil {
		return nil
	}
	return dl.data.Experiences
}

// GetCurrentExperience returns current experience (if any)
func (dl *DataLoader) GetCurrentExperience() *Experience {
	if dl.data == nil {
		return nil
	}

	for _, exp := range dl.data.Experiences {
		if exp.Current {
			return &exp
		}
	}
	return nil
}

// GetSkills returns all skills organized by category
func (dl *DataLoader) GetSkills() map[string][]Skill {
	if dl.data == nil {
		return nil
	}
	return dl.data.Skills
}

// GetSkillsByCategory returns skills for a specific category
func (dl *DataLoader) GetSkillsByCategory(category string) []Skill {
	if dl.data == nil {
		return nil
	}

	skills, exists := dl.data.Skills[category]
	if !exists {
		return nil
	}
	return skills
}

// GetTechFacts returns all tech facts
func (dl *DataLoader) GetTechFacts() []string {
	if dl.data == nil {
		return nil
	}
	return dl.data.TechFacts
}

// GetRandomTechFact returns a random tech fact based on index
func (dl *DataLoader) GetRandomTechFact(index int) string {
	facts := dl.GetTechFacts()
	if len(facts) == 0 {
		return "Loading awesome tech facts..."
	}
	return facts[index%len(facts)]
}

// GetAsciiArt returns ASCII art
func (dl *DataLoader) GetAsciiArt() *AsciiArt {
	if dl.data == nil {
		return nil
	}
	return &dl.data.AsciiArt
}

// GetContact returns contact information
func (dl *DataLoader) GetContact() *Contact {
	if dl.data == nil {
		return nil
	}
	return &dl.data.Personal.Contact
}

// ReloadData reloads data from file (useful for hot-reloading during development)
func (dl *DataLoader) ReloadData() error {
	return dl.LoadData()
}

// ValidateData performs basic validation on loaded data
func (dl *DataLoader) ValidateData() error {
	if dl.data == nil {
		return fmt.Errorf("no data loaded")
	}

	// Validate personal info
	if dl.data.Personal.Name == "" {
		return fmt.Errorf("personal name is required")
	}

	if dl.data.Personal.Contact.Email == "" {
		return fmt.Errorf("contact email is required")
	}

	// Validate experiences
	if len(dl.data.Experiences) == 0 {
		return fmt.Errorf("at least one experience is required")
	}

	// Validate skills
	if len(dl.data.Skills) == 0 {
		return fmt.Errorf("at least one skill category is required")
	}

	return nil
}
