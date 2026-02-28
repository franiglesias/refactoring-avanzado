package dispensables

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Code smell: Data Class [Clase de datos].
// UserRecord es una clase que solo almacena datos, mientras que la validación
// y la lógica de creación residen en UserService.

// Ejercicio: Implementa reglas de dominio adicionales, como requerir verificación de email
// o restringir el registro a ciertos dominios (ej. company.com).

// Tendrás que modificar múltiples servicios y lugares que manipulan UserRecord.
// Esto demuestra cómo separar el comportamiento de los datos provoca que cambios simples
// se dispersen ampliamente por el código (Shotgun Surgery).

type UserRecord struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
}

type UserService struct{}

func (s *UserService) CreateUser(name, email string) (*UserRecord, error) {
	if !contains(email, "@") {
		return nil, fmt.Errorf("Invalid email")
	}

	return &UserRecord{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}, nil
}

func (s *UserService) UpdateUserEmail(user *UserRecord, newEmail string) error {
	if !contains(newEmail, "@") {
		return fmt.Errorf("Invalid email")
	}
	user.Email = newEmail
	return nil
}

type UserReportGenerator struct{}

func (g *UserReportGenerator) GenerateUserSummary(user *UserRecord) string {
	return fmt.Sprintf("User %s (%s) created on %s",
		user.Name,
		user.Email,
		user.CreatedAt.Format("2006-01-02"))
}

func DemoDataClass() (string, error) {
	service := &UserService{}
	report := &UserReportGenerator{}
	user, err := service.CreateUser("Lina", "lina@example.com")
	if err != nil {
		return "", err
	}
	err = service.UpdateUserEmail(user, "lina+news@example.com")
	if err != nil {
		return "", err
	}
	return report.GenerateUserSummary(user), nil
}

func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
