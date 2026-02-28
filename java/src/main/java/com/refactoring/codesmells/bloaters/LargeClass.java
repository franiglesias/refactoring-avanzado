package com.refactoring.codesmells.bloaters;

import java.util.ArrayList;
import java.util.Date;
import java.util.List;

/**
 * Code smell: Large Class [Clase grande].
 * UserAccount acumula muchas responsabilidades no relacionadas
 * como autenticación, perfil, notificaciones y gestión de administración,
 * lo que dificulta el cambio.
 *
 * Ejercicio: Añade autenticación de dos factores (2FA) y preferencias de notificación.
 *
 * Tocarás autenticación, estado y notificaciones en una clase inflada,
 * aumentando la probabilidad de romper comportamiento no relacionado.
 */
public class LargeClass {

    public static void main(String[] args) {
        UserAccount user = new UserAccount("Ana", "ana@example.com", "password123", false);
        user.login("password123");
        user.updateEmail("ana.new@example.com");
        user.addNotification("Bienvenida al sistema");
        System.out.println("Notificaciones: " + user.getNotifications());
    }

    public static class UserAccount {
        private String name;
        private String email;
        private String password;
        private Date lastLogin;
        private int loginAttempts = 0;
        private List<String> notifications = new ArrayList<>();
        private boolean isAdmin;

        public UserAccount(String name, String email, String password, boolean isAdmin) {
            this.name = name;
            this.email = email;
            this.password = password;
            this.lastLogin = new Date();
            this.isAdmin = isAdmin;
        }

        // --- Autenticación ---
        public boolean login(String password) {
            if (this.password.equals(password)) {
                this.lastLogin = new Date();
                this.loginAttempts = 0;
                System.out.println("Inicio de sesión exitoso");
                return true;
            } else {
                this.loginAttempts++;
                System.out.println("Contraseña incorrecta");
                return false;
            }
        }

        public void resetPassword(String newPassword) {
            this.password = newPassword;
            System.out.println("Contraseña actualizada");
        }

        // --- Perfil ---
        public void updateEmail(String newEmail) {
            this.email = newEmail;
            System.out.println("Correo actualizado");
        }

        public void updateName(String newName) {
            this.name = newName;
            System.out.println("Nombre actualizado");
        }

        // --- Notificaciones ---
        public void addNotification(String message) {
            this.notifications.add(message);
        }

        public List<String> getNotifications() {
            return notifications;
        }

        public void clearNotifications() {
            this.notifications.clear();
        }

        // --- Administración ---
        public void promoteToAdmin() {
            this.isAdmin = true;
        }

        public void revokeAdmin() {
            this.isAdmin = false;
        }
    }
}
