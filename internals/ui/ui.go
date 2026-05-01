package ui

import (
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/olekukonko/tablewriter"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205"))

	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("42"))

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("196"))

	mutedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240"))
)

// ---------- UTILS ----------

func nullSafe(s string) string {
	if s == "" || s == "0" {
		return "-"
	}
	return s
}

// ---------- BASIC MESSAGES ----------

func Title(msg string) {
	fmt.Println(titleStyle.Render(msg))
}

func Success(msg string) {
	fmt.Println(successStyle.Render("✔ " + msg))
}

func Error(err error) {
	fmt.Println(errorStyle.Render("✖ " + fmt.Sprint(err)))
}

func Info(msg string) {
	fmt.Println(mutedStyle.Render(msg))
}

// ---------- APPOINTMENTS TABLE ----------

type AppointmentRow struct {
	ID           int64
	PatientName  string
	ScheduledFor string
	Status       string
}

func RenderAppointments(rows []AppointmentRow) {
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"ID", "Patient", "Scheduled", "Status"})

	for _, r := range rows {
		table.Append([]string{
			strconv.Itoa(int(r.ID)),
			r.PatientName,
			r.ScheduledFor,
			colorStatus(r.Status),
		})
	}

	table.Render()
}

// ---------- PATIENTS ----------

type PatientRow struct {
	ID          int64
	Name        string
	Email       string
	PhoneNumber uint
}

func RenderPatients(rows []PatientRow) {
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"ID", "Name", "Email", "Phone"})

	for _, r := range rows {
		table.Append([]string{
			strconv.Itoa(int(r.ID)),
			r.Name,
			nullSafe(r.Email),
			nullSafe(strconv.Itoa(int(r.PhoneNumber))),
		})
	}

	table.Render()
}

// ---------- VISITS ----------

type VisitRow struct {
	ID           int64
	PatientName  string
	ScheduledFor string
	Notes        string
}

func RenderVisits(rows []VisitRow) {
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"ID", "Patient", "Scheduled", "Notes"})

	for _, r := range rows {
		table.Append([]string{
			strconv.Itoa(int(r.ID)),
			r.PatientName,
			r.ScheduledFor,
			nullSafe(r.Notes),
		})
	}

	table.Render()
}

// ---------- PATIENT HISTORY ----------

type PatientHistoryView struct {
	PatientName string

	TotalAppointments int
	Completed         int
	Cancelled         int
	NoShow            int

	Appointments []AppointmentRow
	Visits       []VisitRow
}

func RenderPatientHistory(h PatientHistoryView) {
	Title("Patient History")

	// Header
	fmt.Println()
	fmt.Println("Patient: " + h.PatientName)
	fmt.Println()

	// Summary
	fmt.Println("Summary")
	fmt.Println("-------")
	fmt.Printf("Total Appointments: %d\n", h.TotalAppointments)
	fmt.Printf("Completed: %s\n", successStyle.Render(strconv.Itoa(h.Completed)))
	fmt.Printf("Canceled: %s\n", mutedStyle.Render(strconv.Itoa(h.Cancelled)))
	fmt.Printf("No Show: %s\n", errorStyle.Render(strconv.Itoa(h.NoShow)))
	fmt.Println()

	// Appointments
	if len(h.Appointments) > 0 {
		fmt.Println("Appointments")
		fmt.Println("------------")
		RenderAppointments(h.Appointments)
		fmt.Println()
	}

	// Visits
	if len(h.Visits) > 0 {
		fmt.Println("Visits")
		fmt.Println("------")
		RenderVisits(h.Visits)
		fmt.Println()
	}
}

// ---------- STATUS COLORS ----------

func colorStatus(status string) string {
	switch status {
	case "COMPLETED":
		return successStyle.Render(status)
	case "CANCELED":
		return mutedStyle.Render(status)
	case "NO_SHOW":
		return errorStyle.Render(status)
	default:
		return status
	}
}
