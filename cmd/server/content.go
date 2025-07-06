package server

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func (m *PortfolioModel) renderAbout() string {
	var content strings.Builder

	asciiArt := `
    ╭ ─────────────────────────────────────────────────────────────────── ╮
                                                                           
    │                                                                     │
    │      █████╗ ██████╗ ███╗   ███╗███████╗██████╗ ███████╗██╗   ██╗    │
    │     ██╔══██╗██╔══██╗████╗ ████║██╔════╝██╔══██╗██╔════╝██║   ██║    │
    ⚡    ███████║██████╔╝██╔████╔██║█████╗  ██║  ██║█████╗  ██║   ██║    ⚡
    │     ██╔══██║██╔══██╗██║╚██╔╝██║██╔══╝  ██║  ██║██╔══╝  ╚██╗ ██╔╝    │
    │     ██║  ██║██║  ██║██║ ╚═╝ ██║███████╗██████╔╝███████╗ ╚████╔╝     │
    │     ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝╚═════╝ ╚══════╝  ╚═══╝      │
    │                                                                     │
                                                                          
    ╰ ─────────────────────────────────────────────────────────────────── ╯
`
	content.WriteString(m.styles.AsciiArt.Render(asciiArt))
	content.WriteString("\n\n")

	content.WriteString(m.styles.SectionTitle.Render("👋 About Me"))

	about := `
Hi! I am ABDUL HAMEED but you can call me Armedev. 
I am a tech enthusiast with an awesome skillset based in India. 

🎯 What I Do:
	I like $(COMPUTERS) and their C00l $tacks. I'm passionate about exploring new technologies, 
	building innovative projects, and continuously expanding my knowledge in the ever-evolving world of technology.

💻 Background:
	• Information Science graduate with a strong foundation in computer systems
	• Tech enthusiast who loves diving deep into different technology stacks
	• Based in India, contributing to the global tech community
	• Always eager to learn and adapt to new technological challenges

🌱 Always Learning:
	I believe in continuous learning and staying curious about emerging technologies. 
	The world of computing fascinates me, and I enjoy exploring everything from 
	low-level systems to modern development frameworks.
	`

	content.WriteString(m.styles.ContentText.Render(about))

	// Add random tech facts
	content.WriteString("\n")

	facts := []string{
		"The first computer bug was an actual bug found in 1947",
		"The term 'debugging' was coined by Grace Hopper",
		"There are more possible chess games than atoms in the universe",
		"The first computer programmer was Ada Lovelace in 1843",
		"Linux powers 96.3% of the world's top 1 million web servers",
		"Go was created at Google by Rob Pike, Ken Thompson, and Robert Griesemer",
		"The first version of Git was written in just 2 weeks",
		"PostgreSQL is older than MySQL by 5 years",
		"The word 'robot' comes from the Czech word 'robota' meaning work",
		"The @ symbol was used in emails for the first time in 1971",
	}

	fact := facts[m.animationTick/50%len(facts)]
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

	experiences := []struct {
		title    string
		company  string
		period   string
		location string
		details  []string
	}{
		{
			title:    "Software Developer",
			company:  "Tursio",
			period:   "Jun 2025 - Present",
			location: "Bengaluru, India (On-site)",
			details: []string{
				"Currently working as a Full-time Software Developer",
				"Building scalable software solutions and contributing to product development",
				"Working with modern technology stacks and development practices",
				"Collaborating with cross-functional teams to deliver high-quality software",
				"Tech: Modern web technologies and cloud platforms",
			},
		},
		{
			title:    "Software Developer",
			company:  "Gida Technologies",
			period:   "September 2023 - May 2025",
			location: "Bengaluru, India (On-site)",
			details: []string{
				"Built and maintained full-stack web applications using Next.js and NestJS",
				"Developed multiple products: AgeEasyByAntara (Max group), Ergo Self-Help Portal (HDFC), Convenex Portal (HDFC)",
				"Created and integrated RESTful APIs for efficient data management and user authentication",
				"Utilized Next.js features like SSR and SSG to optimize application performance and SEO",
				"Tech: Next.js, NestJS, TypeScript, REST APIs, SSR, SSG",
			},
		},
		{
			title:    "Full-Stack Developer Intern",
			company:  "BurdenOff Consultancy Services",
			period:   "Feb 2023 - Jun 2023",
			location: "Remote",
			details: []string{
				"Designed and implemented a payment model to support seamless transactions",
				"Introduced adapter architecture to ensure flexibility and reduce reliance on single payment provider",
				"Integrated Stripe for payment processing with webhooks and 2-way verification",
				"Developed PaymentMethod model for secure payment storage and recurring payments",
				"Tech: Stripe API, Payment Architecture, Webhooks, Security Implementation",
			},
		},
		{
			title:    "Full-Stack Developer Intern",
			company:  "BurdenOff Consultancy Services",
			period:   "Jun 2022 - Dec 2022",
			location: "Remote",
			details: []string{
				"Worked on Payment, Notification, Billing/Account, Wallet, Store, and Product modules",
				"Designed type-safe, clean model structure to enhance security and prevent vulnerabilities",
				"Added and enhanced features using GraphQL, TypeScript, and ArangoDB",
				"Ensured efficient and scalable functionality across all modules",
				"Tech: GraphQL, TypeScript, ArangoDB, Security Architecture",
			},
		},
		{
			title:    "React Developer Intern",
			company:  "NETART-INDIA",
			period:   "Jun 2021 - Jan 2022",
			location: "Remote",
			details: []string{
				"Built R&D dashboard using React and FireCMS for generating SEO reports",
				"Implemented scheduler to prevent data capture clashes",
				"Automated data capture process with UIvision and App Script API",
				"Reduced manual workload significantly through automation",
				"Tech: React, FireCMS, UIvision, Google App Script, SEO Tools",
			},
		},
	}

	for i, exp := range experiences {
		if i > 0 {
			content.WriteString("\n")
		}

		header := fmt.Sprintf("%s @ %s", exp.title, exp.company)
		content.WriteString(m.styles.ExperienceTitle.Render(header))
		content.WriteString("\n")

		meta := fmt.Sprintf("📅 %s • 📍 %s", exp.period, exp.location)
		content.WriteString(m.styles.ExperienceMeta.Render(meta))
		content.WriteString("\n\n")

		for _, detail := range exp.details {
			content.WriteString(m.styles.ExperienceDetail.Render("  • " + detail))
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

	skillCategories := map[string][]Skill{
		"💻 Programming Languages": {
			{"TypeScript", 95, "Expert"},
			{"JavaScript", 90, "Advanced"},
			{"Rust", 85, "Advanced"},
			{"Golang", 80, "Intermediate"},
			{"C++", 75, "Intermediate"},
		},
		"🚀 Frontend Frameworks": {
			{"React.js", 95, "Expert"},
			{"Next.js", 95, "Expert"},
			{"Solid.js", 80, "Advanced"},
			{"HTML/CSS", 90, "Expert"},
		},
		"⚡ Backend & APIs": {
			{"NestJS", 90, "Advanced"},
			{"Node.js", 85, "Advanced"},
			{"REST APIs", 95, "Expert"},
			{"GraphQL", 90, "Advanced"},
			{"gRPC", 80, "Intermediate"},
			{"Actix-web", 75, "Intermediate"},
		},
		"🗃️ Databases": {
			{"PostgreSQL", 90, "Advanced"},
			{"MongoDB", 80, "Intermediate"},
			{"ArangoDB", 85, "Advanced"},
		},
		"🔧 Tools & DevOps": {
			{"Docker", 85, "Advanced"},
			{"Kubernetes", 80, "Intermediate"},
			{"Webpack", 80, "Intermediate"},
			{"esbuild", 75, "Intermediate"},
		},
		"📊 Data Formats & Protocols": {
			{"Protobuf", 80, "Advanced"},
			{"JSON", 95, "Expert"},
			{"WebSockets", 85, "Advanced"},
		},
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

type Skill struct {
	Name       string
	Percentage int
	Experience string
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

	contact := `Ready to collaborate or discuss exciting opportunities? I'm always open to 
interesting conversations about technology and new projects!

📧 Email:     armedev@protonmail.com
🐙 GitHub:    github.com/armedev
💼 LinkedIn:  linkedin.com/in/abdul-hameed-armedev
🌐 Portfolio: arme.dev

🌍 Location:  Bangalore, India
🕐 Timezone:  IST (UTC+5:30)

💬 Preferred contact method: Email or LinkedIn
⚡ Response time: Usually within 24 hours

Feel free to reach out for:
• Full-stack development opportunities
• Technical discussions and collaboration
• Open source contributions
• Consulting and freelance projects

🚀 Specializations:
• Microservices architecture and system design
• Payment systems and financial technology
• Modern web development with React/Next.js
• Backend development with NestJS and GraphQL
• Database design and optimization
`

	content.WriteString(m.styles.ContentText.Render(contact))

	// Add ASCII art
	content.WriteString("\n\n")
	ascii := `
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
	content.WriteString(m.styles.AsciiArt.Render(ascii))

	return content.String()
}
