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
	Hello! I'm a passionate Senior Software Engineer with 8+ years of experience building scalable, 
high-performance systems. I specialize in backend development with Go, cloud architecture, and 
modern web technologies.

🎯 What I Do:
• Design and implement distributed microservices architectures
• Build robust APIs and backend systems that handle millions of requests
• Optimize database performance and implement efficient caching strategies
• Lead technical teams and mentor junior developers
• Contribute to open-source projects and tech communities

🌱 Always Learning:
I believe in continuous learning and staying up-to-date with the latest technologies. 
Currently exploring Kubernetes, serverless architectures, and advanced Go patterns.
	`

	content.WriteString(m.styles.ContentText.Render(about))

	// Add some live stats
	content.WriteString("\n\n")
	content.WriteString(m.styles.StatsBox.Render(m.renderLiveStats()))

	// Add ASCII art at the bottom

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
			title:    "Senior Software Engineer",
			company:  "TechCorp Inc.",
			period:   "2022 - Present",
			location: "San Francisco, CA (Remote)",
			details: []string{
				"Led development of microservices architecture serving 50M+ users",
				"Reduced API response times by 40% through optimization and caching",
				"Mentored team of 5 junior developers and conducted code reviews",
				"Implemented CI/CD pipelines reducing deployment time by 60%",
				"Tech: Go, Kubernetes, PostgreSQL, Redis, AWS",
			},
		},
		{
			title:    "Full Stack Developer",
			company:  "StartupXYZ",
			period:   "2020 - 2022",
			location: "New York, NY",
			details: []string{
				"Built entire backend infrastructure from scratch for fintech product",
				"Developed real-time payment processing system handling $10M+ monthly",
				"Created React dashboard for internal tools and customer analytics",
				"Implemented security best practices and SOC2 compliance",
				"Tech: Go, React, TypeScript, PostgreSQL, Docker",
			},
		},
		{
			title:    "Software Engineer",
			company:  "Enterprise Solutions Ltd.",
			period:   "2018 - 2020",
			location: "Austin, TX",
			details: []string{
				"Developed REST APIs and backend services for enterprise clients",
				"Optimized database queries improving performance by 3x",
				"Collaborated with frontend teams to deliver seamless user experiences",
				"Participated in on-call rotation and incident response",
				"Tech: Java, Spring Boot, MySQL, AWS, Jenkins",
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
		"🔧 Languages": {
			{"Go", 95, "5+ years"},
			{"TypeScript", 90, "4+ years"},
			{"Python", 85, "3+ years"},
			{"Java", 80, "3+ years"},
			{"Rust", 70, "Learning"},
		},
		"🏗️ Backend": {
			{"Microservices", 95, "Expert"},
			{"REST APIs", 95, "Expert"},
			{"GraphQL", 85, "Advanced"},
			{"gRPC", 80, "Advanced"},
			{"WebSockets", 75, "Intermediate"},
		},
		"🗄️ Databases": {
			{"PostgreSQL", 90, "Advanced"},
			{"Redis", 85, "Advanced"},
			{"MongoDB", 80, "Intermediate"},
			{"ElasticSearch", 75, "Intermediate"},
		},
		"☁️ Cloud & DevOps": {
			{"AWS", 90, "Advanced"},
			{"Docker", 95, "Expert"},
			{"Kubernetes", 85, "Advanced"},
			{"Terraform", 80, "Intermediate"},
			{"CI/CD", 90, "Advanced"},
		},
		"🎨 Frontend": {
			{"React", 85, "Advanced"},
			{"Next.js", 80, "Advanced"},
			{"Vue.js", 75, "Intermediate"},
			{"HTML/CSS", 90, "Expert"},
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
		animatedPercentage = skill.Percentage + pulse
		if animatedPercentage > 100 {
			animatedPercentage = 100
		}
		if animatedPercentage < 0 {
			animatedPercentage = 0
		}
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

func (m *PortfolioModel) renderProjects() string {
	var content strings.Builder

	content.WriteString(m.styles.SectionTitle.Render("🚀 Featured Projects"))
	content.WriteString("\n\n")

	projects := []struct {
		name        string
		description string
		tech        []string
		highlights  []string
		github      string
		demo        string
	}{
		{
			name:        "DistributedChat",
			description: "Real-time chat application with horizontal scaling capabilities",
			tech:        []string{"Go", "WebSockets", "Redis", "PostgreSQL", "React"},
			highlights: []string{
				"Supports 10,000+ concurrent connections",
				"Message delivery in <50ms",
				"Auto-scaling based on load",
				"End-to-end encryption",
			},
			github: "github.com/johndoe/distributed-chat",
			demo:   "chat.johndoe.dev",
		},
		{
			name:        "MetricsCollector",
			description: "High-performance metrics collection and visualization platform",
			tech:        []string{"Go", "InfluxDB", "Grafana", "Kubernetes"},
			highlights: []string{
				"Processes 1M+ metrics/second",
				"Custom query language",
				"Real-time alerting system",
				"99.99% uptime SLA",
			},
			github: "github.com/johndoe/metrics-collector",
			demo:   "metrics.johndoe.dev",
		},
		{
			name:        "PaymentGateway",
			description: "Secure payment processing microservice with fraud detection",
			tech:        []string{"Go", "PostgreSQL", "Redis", "AWS Lambda"},
			highlights: []string{
				"PCI DSS compliant",
				"ML-based fraud detection",
				"Support for 20+ payment methods",
				"99.9% success rate",
			},
			github: "github.com/johndoe/payment-gateway",
			demo:   "Private repository",
		},
	}

	for i, project := range projects {
		if i > 0 {
			content.WriteString("\n")
		}

		content.WriteString(m.styles.ProjectTitle.Render(project.name))
		content.WriteString("\n")
		content.WriteString(m.styles.ProjectDescription.Render(project.description))
		content.WriteString("\n\n")

		// Tech stack
		content.WriteString(m.styles.ProjectLabel.Render("🛠️ Tech Stack: "))
		content.WriteString(strings.Join(project.tech, " • "))
		content.WriteString("\n\n")

		// Highlights
		content.WriteString(m.styles.ProjectLabel.Render("✨ Highlights:"))
		content.WriteString("\n")
		for _, highlight := range project.highlights {
			content.WriteString(fmt.Sprintf("  • %s\n", highlight))
		}
		content.WriteString("\n")

		// Links
		content.WriteString(m.styles.ProjectLabel.Render("🔗 Links: "))
		content.WriteString(fmt.Sprintf("GitHub: %s", project.github))
		if project.demo != "Private repository" {
			content.WriteString(fmt.Sprintf(" • Demo: %s", project.demo))
		}
		content.WriteString("\n\n")
	}

	return content.String()
}

func (m *PortfolioModel) renderContact() string {
	var content strings.Builder

	content.WriteString(m.styles.SectionTitle.Render("📞 Get In Touch"))
	content.WriteString("\n\n")

	contact := `Ready to collaborate or just want to chat about technology? I'm always open to 
interesting conversations and new opportunities!

📧 Email:     john.doe@email.com
🐙 GitHub:    github.com/johndoe
💼 LinkedIn:  linkedin.com/in/johndoe
🐦 Twitter:   @johndoe_dev
🌐 Website:   johndoe.dev
📱 Phone:     +1 (555) 123-4567

🌍 Location:  San Francisco, CA (Open to remote work)
🕐 Timezone:  PST (UTC-8)

💬 Preferred contact method: Email or LinkedIn
⚡ Response time: Usually within 24 hours

Feel free to reach out for:
• Technical discussions and collaboration
• Speaking opportunities at conferences/meetups
• Open source contributions
• Consulting and freelance projects
• Career opportunities`

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

func (m *PortfolioModel) renderLiveDemo() string {
	var content strings.Builder

	content.WriteString(m.styles.SectionTitle.Render("🎮 Live Interactive Demo"))
	content.WriteString("\n\n")

	// Real-time clock
	now := time.Now()
	timeStr := now.Format("15:04:05 MST")
	dateStr := now.Format("Monday, January 2, 2006")

	content.WriteString(m.styles.LiveTitle.Render("🕐 Real-time Clock"))
	content.WriteString("\n")
	content.WriteString(m.styles.LiveValue.Render(timeStr))
	content.WriteString("\n")
	content.WriteString(m.styles.LiveSubtitle.Render(dateStr))
	content.WriteString("\n\n")

	// Animation demo
	content.WriteString(m.styles.LiveTitle.Render("🎬 Animation Demo"))
	content.WriteString("\n")

	// Rotating progress bar
	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	frame := frames[m.animationTick%len(frames)]
	progress := (m.animationTick % 50) * 2

	progressBar := strings.Repeat("█", progress/10) + strings.Repeat("░", 10-progress/10)
	content.WriteString(fmt.Sprintf("%s Processing... %s %d%%\n", frame, progressBar, progress))
	content.WriteString("\n")

	// System stats simulation
	content.WriteString(m.styles.LiveTitle.Render("📊 System Performance"))
	content.WriteString("\n")

	// Simulate realistic metrics
	baseCPU := 25 + int(15*math.Sin(float64(m.animationTick)*0.05))
	baseMem := 55 + int(10*math.Sin(float64(m.animationTick)*0.03))
	baseConnections := 1200 + int(300*math.Sin(float64(m.animationTick)*0.02))
	baseRequests := 750 + int(200*math.Sin(float64(m.animationTick)*0.04))

	stats := fmt.Sprintf(`CPU Usage:    %d%% %s
Memory:       %d%% %s  
Connections:  %d active
Uptime:       %s
Requests/sec: %d`,
		baseCPU, m.renderMiniBar(baseCPU, 20),
		baseMem, m.renderMiniBar(baseMem, 20),
		baseConnections,
		time.Since(m.startTime).Truncate(time.Second),
		baseRequests,
	)

	content.WriteString(m.styles.StatsBox.Render(stats))
	content.WriteString("\n\n")

	// Interactive features
	content.WriteString(m.styles.LiveTitle.Render("🎆 Interactive Features"))
	content.WriteString("\n")

	interactiveContent := `Try these interactive features:

• Press 'x' to trigger particle explosions
• Press 'e' to toggle visual effects on/off
• Use Tab/Shift+Tab to navigate sections
• Press 'h' for detailed help

Effects Status: `

	if m.effectsEnabled {
		interactiveContent += "✅ ENABLED"
	} else {
		interactiveContent += "❌ DISABLED"
	}

	if len(m.particles) > 0 {
		interactiveContent += fmt.Sprintf("\n\nActive Particles: %d 🎆", len(m.particles))
		interactiveContent += "\nParticles are currently exploding!"
	} else {
		interactiveContent += "\n\nPress 'x' to create some fireworks! 🎇"
	}

	content.WriteString(m.styles.ContentText.Render(interactiveContent))
	content.WriteString("\n\n")

	// Tech facts
	content.WriteString(m.styles.LiveTitle.Render("🎲 Random Tech Facts"))
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

	return content.String()
}

func (m *PortfolioModel) renderLiveStats() string {
	uptime := time.Since(m.startTime)

	// Animated view counter
	baseViews := 150 + (m.animationTick % 50)
	connections := 2834 + (m.animationTick % 100)

	stats := fmt.Sprintf(`📈 Portfolio Statistics:
  • Views today: %d
  • Total connections: %d
  • Session uptime: %s
  • Server location: San Francisco, CA
  • Effects: %s
  • Last updated: %s`,
		baseViews,
		connections,
		uptime.Truncate(time.Second),
		func() string {
			if m.effectsEnabled {
				return "ENABLED"
			} else {
				return "DISABLED"
			}
		}(),
		time.Now().Format("15:04:05"),
	)

	return stats
}

func (m *PortfolioModel) renderMiniBar(value, width int) string {
	if value > 100 {
		value = 100
	}
	if value < 0 {
		value = 0
	}

	filled := int(float64(width) * float64(value) / 100.0)
	return "[" + strings.Repeat("█", filled) + strings.Repeat("░", width-filled) + "]"
}
