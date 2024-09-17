package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()


	// Helper function to add a new line
	addLine := func(height float64) {
		pdf.Ln(height)
	}

	// Helper function to add a section title
	addSectionTitle := func(title string) {
		pdf.SetFont("Arial", "B", 12)
		pdf.Cell(0, 10, title)
		pdf.Ln(5)
		pdf.Line(10, pdf.GetY(), 200, pdf.GetY())
		pdf.Ln(5)
	}

	// Header
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, "John Doe")
	addLine(5)

	pdf.SetFont("Arial", "", 10)
	pdf.Cell(0, 5, "Email: john.doe@email.com | Phone: (123) 456-7890")
	addLine(5)
	pdf.Cell(0, 5, "Location: New York, NY | LinkedIn: linkedin.com/in/johndoe")
	addLine(5)
	pdf.Cell(0, 5, "GitHub: github.com/johndoe")
	addLine(10)

	// Summary
	addSectionTitle("Summary")
	pdf.SetFont("Arial", "", 10)
	pdf.MultiCell(0, 5, "Experienced software engineer with 5+ years of expertise in full-stack development, specializing in scalable web applications and cloud technologies. Proven track record of delivering high-quality code and leading successful projects.", "", "", false)
	addLine(5)

	// Skills
	addSectionTitle("Skills")
	pdf.SetFont("Arial", "", 10)
	pdf.MultiCell(0, 5, "Languages: Java, Python, JavaScript, TypeScript, SQL\nTechnologies: React, Node.js, Spring Boot, Docker, Kubernetes, AWS, Git\nMethodologies: Agile, Scrum, TDD, CI/CD", "", "", false)
	addLine(5)

	// Work Experience
	addSectionTitle("Work Experience")

	pdf.SetFont("Arial", "B", 11)
	pdf.Cell(0, 5, "Senior Software Engineer - TechCorp Inc.")
	addLine(5)
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(0, 5, "June 2020 - Present")
	addLine(5)
	pdf.SetFont("Arial", "", 10)
	pdf.MultiCell(0, 5, "• Led the development of a microservices-based e-commerce platform, improving scalability by 200%\n• Implemented CI/CD pipelines, reducing deployment time by 70%\n• Mentored junior developers and conducted code reviews to ensure best practices", "", "", false)
	addLine(5)

	pdf.SetFont("Arial", "B", 11)
	pdf.Cell(0, 5, "Software Engineer - InnoSoft Solutions")
	addLine(5)
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(0, 5, "August 2017 - May 2020")
	addLine(5)
	pdf.SetFont("Arial", "", 10)
	pdf.MultiCell(0, 5, "• Developed and maintained RESTful APIs for mobile applications\n• Optimized database queries, improving application performance by 40%\n• Collaborated with UX designers to implement responsive web designs", "", "", false)
	addLine(5)

	// Projects
	addSectionTitle("Projects")
	pdf.SetFont("Arial", "B", 11)
	pdf.Cell(0, 5, "Task Management App")
	addLine(5)
	pdf.SetFont("Arial", "", 10)
	pdf.MultiCell(0, 5, "A full-stack web application built with React, Node.js, and MongoDB\n• Implemented real-time updates using WebSockets\n• Integrated OAuth 2.0 for secure authentication", "", "", false)
	addLine(5)

	// Education
	addSectionTitle("Education")
	pdf.SetFont("Arial", "B", 11)
	pdf.Cell(0, 5, "Bachelor of Science in Computer Science")
	addLine(5)
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(0, 5, "University of Technology, Graduated May 2017")

	err := pdf.OutputFileAndClose("software_engineer_resume.pdf")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
