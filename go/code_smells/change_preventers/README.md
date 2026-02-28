# Divergent change

## Definición

Una clase tiene múltiples razones para cambiar, lo que normalmente indica que se ocupa de muchas responsabilidades que se deberían separar en clases especialistas más pequeñas.

## Ejemplo

ProfileManager maneja validación, persistencia, exportación y envío de emails—múltiples razones para cambiar concentradas en una sola clase.

```go
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ProfileManager struct {
	store map[string]User
}

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

func (pm *ProfileManager) ExportAsCSV() string {
	rows := []string{"id,name,email"}
	for _, u := range pm.store {
		rows = append(rows, fmt.Sprintf("%s,%s,%s", u.ID, u.Name, u.Email))
	}
	return strings.Join(rows, "\n")
}

func (pm *ProfileManager) SendWelcomeEmail(user User) string {
	return fmt.Sprintf("Welcome %s! Sent to %s", user.Name, user.Email)
}
```

## Ejercicio

Añade un número de teléfono con validación, inclúyelo en las exportaciones y envía un SMS.

## Problemas que encontrarás

Tocarás validación, almacenamiento, ExportAsJSON/CSV y mensajería en un solo lugar, demostrando cómo un cambio fuerza ediciones en responsabilidades no relacionadas.
