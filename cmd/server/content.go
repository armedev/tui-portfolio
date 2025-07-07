package server

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func (m *PortfolioModel) renderAbout() string {
	var content strings.Builder

	// Get data from loader or use fallback
	personal := m.dataLoader.GetPersonalInfo()
	asciiArt := m.dataLoader.GetAsciiArt()

	// Render ASCII art
	if asciiArt != nil && asciiArt.Logo != "" {
		content.WriteString(m.styles.AsciiArt.Render(asciiArt.Logo))
	} else {
		// Fallback ASCII art
		fallbackArt := `
    ╭ ─────────────────────────────────────────────────────────────────── ╮
    │                     PORTFOLIO TERMINAL                              │
    ╰ ─────────────────────────────────────────────────────────────────── ╯
`
		content.WriteString(m.styles.AsciiArt.Render(fallbackArt))
	}
	content.WriteString("\n\n")

	content.WriteString(m.styles.SectionTitle.Render("👋 About Me"))

	// Build about content from data
	var about strings.Builder

	if personal != nil {
		// Personal introduction
		about.WriteString(fmt.Sprintf("\nHi! I am %s", personal.Name))
		if personal.Nickname != "" {
			about.WriteString(fmt.Sprintf(" but you can call me %s.", personal.Nickname))
		}
		about.WriteString("\n")

		if personal.About.Intro != "" {
			about.WriteString(personal.About.Intro)
			about.WriteString("\n\n")
		}

		// What I Do section
		about.WriteString("🎯 What I Do:\n")
		if personal.About.WhatIDo != "" {
			about.WriteString("\t" + personal.About.WhatIDo + "\n\n")
		}

		// Background section
		if len(personal.About.Background) > 0 {
			about.WriteString("💻 Background:\n")
			for _, bg := range personal.About.Background {
				about.WriteString("\t• " + bg + "\n")
			}
			about.WriteString("\n")
		}

		// Philosophy section
		if personal.About.Philosophy != "" {
			about.WriteString("🌱 Always Learning:\n")
			about.WriteString("\t" + personal.About.Philosophy + "\n")
		}
	} else {
		// Fallback content if no data is loaded
		about.WriteString(`
Welcome to my interactive portfolio terminal!

🎯 What I Do:
	I'm a passionate developer who loves building innovative solutions
	and exploring new technologies.

💻 Background:
	• Full-stack developer with modern web technologies
	• Experience with both frontend and backend development
	• Always learning and adapting to new challenges

🌱 Always Learning:
	I believe in continuous learning and staying curious about 
	emerging technologies in the ever-evolving world of software development.
`)
	}

	content.WriteString(m.styles.ContentText.Render(about.String()))

	// Add random tech facts
	content.WriteString("\n")

	fact := m.dataLoader.GetRandomTechFact(m.animationTick / 50)
	content.WriteString(m.styles.FactBox.Render("💡 " + fact))

	// Add real-time clock
	content.WriteString("\n\n")
	now := time.Now()
	timeStr := now.Format("15:04:05 MST")
	dateStr := now.Format("Monday, January 2, 2006")

	content.WriteString(m.styles.LiveValue.Render(timeStr, m.styles.LiveSubtitle.Render(dateStr)))
	content.WriteString("\n\n")

	return content.String()
}

func (m *PortfolioModel) renderExperience() string {
	var content strings.Builder

	content.WriteString(m.styles.SectionTitle.Render("💼 Professional Experience"))
	content.WriteString("\n\n")

	experiences := m.dataLoader.GetExperiences()

	if len(experiences) == 0 {
		// Fallback content
		content.WriteString(m.styles.ContentText.Render("No experience data available. Please check the data file."))
		return content.String()
	}

	for i, exp := range experiences {
		if i > 0 {
			content.WriteString("\n")
		}

		header := fmt.Sprintf("%s @ %s", exp.Title, exp.Company)
		content.WriteString(m.styles.ExperienceTitle.Render(header))
		content.WriteString("\n")

		meta := fmt.Sprintf("📅 %s • 📍 %s", exp.Period, exp.Location)
		if exp.Current {
			meta += " • 🟢 Current"
		}
		content.WriteString(m.styles.ExperienceMeta.Render(meta))
		content.WriteString("\n\n")

		for _, detail := range exp.Details {
			content.WriteString(m.styles.ExperienceDetail.Render("  • " + detail))
			content.WriteString("\n")
		}

		// Add technology tags
		if len(exp.Technologies) > 0 {
			techStr := "Tech: " + strings.Join(exp.Technologies, ", ")
			content.WriteString(m.styles.ExperienceDetail.Render("  " + techStr))
			content.WriteString("\n")
		}
		content.WriteString("\n")
	}

	return content.String()
}

func (m *PortfolioModel) renderSkills() string {
	var content strings.Builder

	content.WriteString(m.styles.SectionTitle.Render("🛠️ Technical Skills"))
	content.WriteString("\n\n")

	skillCategories := m.dataLoader.GetSkills()

	if len(skillCategories) == 0 {
		// Fallback content
		content.WriteString(m.styles.ContentText.Render("No skills data available. Please check the data file."))
		return content.String()
	}

	for category, skills := range skillCategories {
		content.WriteString(m.styles.SkillCategory.Render(category))
		content.WriteString("\n")

		for _, skill := range skills {
			content.WriteString(m.renderSkillBar(skill))
			content.WriteString("\n")
		}
		content.WriteString("\n")
	}

	return content.String()
}

func (m *PortfolioModel) renderSkillBar(skill Skill) string {
	barWidth := 30

	// Add subtle animation to skill bars
	animatedPercentage := skill.Percentage
	if m.effectsEnabled {
		// Gentle pulsing effect
		pulse := int(3 * math.Sin(float64(m.animationTick+skill.Percentage)*0.05))
		animatedPercentage = max(min(skill.Percentage+pulse, 100), 0)
	}

	filled := int(float64(barWidth) * float64(animatedPercentage) / 100.0)
	bar := strings.Repeat("█", filled) + strings.Repeat("░", barWidth-filled)

	skillLine := fmt.Sprintf("  %-15s %s %3d%% (%s)",
		skill.Name,
		m.styles.SkillBar.Render(bar),
		skill.Percentage,
		skill.Experience,
	)

	return skillLine
}

func (m *PortfolioModel) renderContact() string {
	var content strings.Builder

	content.WriteString(m.styles.SectionTitle.Render("📞 Get In Touch"))
	content.WriteString("\n\n")

	contact := m.dataLoader.GetContact()
	personal := m.dataLoader.GetPersonalInfo()
	asciiArt := m.dataLoader.GetAsciiArt()

	if contact == nil {
		// Fallback content
		content.WriteString(m.styles.ContentText.Render("No contact data available. Please check the data file."))
		return content.String()
	}

	var contactContent strings.Builder

	contactContent.WriteString("Ready to collaborate or discuss exciting opportunities? I'm always open to\n")
	contactContent.WriteString("interesting conversations about technology and new projects!\n\n")

	// Contact information
	contactContent.WriteString(fmt.Sprintf("📧 Email:     %s\n", contact.Email))
	contactContent.WriteString(fmt.Sprintf("🐙 GitHub:    %s\n", contact.GitHub))
	contactContent.WriteString(fmt.Sprintf("💼 LinkedIn:  %s\n", contact.LinkedIn))
	contactContent.WriteString(fmt.Sprintf("🌐 Portfolio: %s\n\n", contact.Portfolio))

	// Location and timezone
	if personal != nil {
		contactContent.WriteString(fmt.Sprintf("🌍 Location:  %s\n", personal.Location))
		contactContent.WriteString(fmt.Sprintf("🕐 Timezone:  %s\n\n", personal.Timezone))
	}

	// Contact preferences
	contactContent.WriteString(fmt.Sprintf("💬 Preferred contact method: %s\n", contact.PreferredContact))
	contactContent.WriteString(fmt.Sprintf("⚡ Response time: %s\n\n", contact.ResponseTime))

	// Available for
	if len(contact.AvailableFor) > 0 {
		contactContent.WriteString("Feel free to reach out for:\n")
		for _, item := range contact.AvailableFor {
			contactContent.WriteString("• " + item + "\n")
		}
		contactContent.WriteString("\n")
	}

	// Specializations
	if len(contact.Specializations) > 0 {
		contactContent.WriteString("🚀 Specializations:\n")
		for _, spec := range contact.Specializations {
			contactContent.WriteString("• " + spec + "\n")
		}
	}

	content.WriteString(m.styles.ContentText.Render(contactContent.String()))

	// Add ASCII art
	content.WriteString("\n\n")
	if asciiArt != nil && asciiArt.Contact != "" {
		content.WriteString(m.styles.AsciiArt.Render(asciiArt.Contact))
	} else {
		// Fallback ASCII art
		fallbackArt := `
    ╭─────────────────────╮
    │  Let's build cool   │
    │  stuff together! 🚀 │
    ╰─────────────────────╯
         │
         ▼
       ┌─────────┐
       │ ( ◕‿◕ ) │
       └─────────┘
`
		content.WriteString(m.styles.AsciiArt.Render(fallbackArt))
	}

	return content.String()
}
