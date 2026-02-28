package change_preventers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// User represents a user in the system
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ProfileManager demonstrates divergent change code smell
// It handles validation, persistence, export, and email sending—multiple reasons to change
type ProfileManager struct {
	store map[string]User
}

// NewProfileManager creates a new ProfileManager
func NewProfileManager() *ProfileManager {
	return &ProfileManager{
		store: make(map[string]User),
	}
}

// Register registers a new user
func (pm *ProfileManager) Register(user User) error {
	if strings.TrimSpace(user.Name) == "" {
		return errors.New("invalid name")
	}
	if !strings.Contains(user.Email, "@") {
		return errors.New("invalid email")
	}
	pm.store[user.ID] = user
	return nil
}

// UpdateEmail updates a user's email
func (pm *ProfileManager) UpdateEmail(id string, newEmail string) error {
	if !strings.Contains(newEmail, "@") {
		return errors.New("invalid email")
	}
	u, exists := pm.store[id]
	if !exists {
		return errors.New("not found")
	}
	u.Email = newEmail
	pm.store[id] = u
	return nil
}

// ExportAsJSON exports users as JSON
func (pm *ProfileManager) ExportAsJSON() (string, error) {
	users := make([]User, 0, len(pm.store))
	for _, user := range pm.store {
		users = append(users, user)
	}
	data, err := json.Marshal(users)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ExportAsCSV exports users as CSV
func (pm *ProfileManager) ExportAsCSV() string {
	rows := []string{"id,name,email"}
	for _, u := range pm.store {
		rows = append(rows, fmt.Sprintf("%s,%s,%s", u.ID, u.Name, u.Email))
	}
	return strings.Join(rows, "\n")
}

// SendWelcomeEmail sends a welcome email to the user
func (pm *ProfileManager) SendWelcomeEmail(user User) string {
	return fmt.Sprintf("Welcome %s! Sent to %s", user.Name, user.Email)
}

// DemoDivergentChange demonstrates the divergent change smell
func DemoDivergentChange(pm *ProfileManager, u User) (string, error) {
	if err := pm.Register(u); err != nil {
		return "", err
	}
	if err := pm.UpdateEmail(u.ID, u.Email); err != nil {
		return "", err
	}
	return pm.ExportAsJSON()
}

// Exercise: Add a phone number with validation, include it in exports, and send an SMS.
//
// Problems you will encounter:
// You will touch validation, storage, ExportAsJSON/CSV, and messaging in one place,
// demonstrating how a change forces edits across unrelated responsibilities.
