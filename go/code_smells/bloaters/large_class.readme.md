# Large class

Clase grande.

## Definición

Una clase contiene muchas propiedades, muchos métodos o muchas líneas de código, acumulando muchas responsabilidades no relacionadas o que pueden responder a necesidades diferentes.

## Ejemplo

```go
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
```

## Ejercicio

Añade soporte para autenticación de dos factores (2FA) y preferencias de notificación.

## Problemas que encontrarás

Tocarás autenticación, estado y notificaciones en una clase inflada, aumentando la probabilidad de romper comportamiento no relacionado.
