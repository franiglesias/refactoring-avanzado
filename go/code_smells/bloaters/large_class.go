package bloaters

import (
	"fmt"
	"time"
)

// Code smell: Large Class [Clase grande].
// UserAccount acumula muchas responsabilidades no relacionadas
// como autenticación, perfil, notificaciones y gestión de administración,
// lo que dificulta el cambio.

// Ejercicio: Añade autenticación de dos factores (2FA) y preferencias de notificación.

// Tocarás autenticación, estado y notificaciones en una clase inflada,
// aumentando la probabilidad de romper comportamiento no relacionado.

type UserAccount struct {
	name          string
	email         string
	password      string
	lastLogin     time.Time
	loginAttempts int
	notifications []string
	isAdmin       bool
}

func NewUserAccount(name, email, password string, isAdmin bool) *UserAccount {
	return &UserAccount{
		name:          name,
		email:         email,
		password:      password,
		lastLogin:     time.Now(),
		loginAttempts: 0,
		notifications: []string{},
		isAdmin:       isAdmin,
	}
}

// --- Autenticación ---

func (u *UserAccount) Login(password string) bool {
	if u.password == password {
		u.lastLogin = time.Now()
		u.loginAttempts = 0
		fmt.Println("Inicio de sesión exitoso")
		return true
	}
	u.loginAttempts++
	fmt.Println("Contraseña incorrecta")
	return false
}

func (u *UserAccount) ResetPassword(newPassword string) {
	u.password = newPassword
	fmt.Println("Contraseña actualizada")
}

// --- Perfil ---

func (u *UserAccount) UpdateEmail(newEmail string) {
	u.email = newEmail
	fmt.Println("Correo actualizado")
}

func (u *UserAccount) UpdateName(newName string) {
	u.name = newName
	fmt.Println("Nombre actualizado")
}

// --- Notificaciones ---

func (u *UserAccount) AddNotification(message string) {
	u.notifications = append(u.notifications, message)
}

func (u *UserAccount) GetNotifications() []string {
	return u.notifications
}

func (u *UserAccount) ClearNotifications() {
	u.notifications = []string{}
}

// --- Administración ---

func (u *UserAccount) PromoteToAdmin() {
	u.isAdmin = true
}

func (u *UserAccount) RevokeAdmin() {
	u.isAdmin = false
}
