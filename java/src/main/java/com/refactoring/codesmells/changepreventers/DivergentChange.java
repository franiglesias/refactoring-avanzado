package com.refactoring.codesmells.changepreventers;

import java.util.*;
import java.util.stream.Collectors;

/**
 * Code Smell: Divergent Change [Cambio divergente]
 * ProfileManager maneja múltiples razones de cambio:
 * - Gestión de usuarios (registro, actualización)
 * - Exportación de datos (JSON, CSV)
 * - Notificaciones (email)
 *
 * Cada nuevo formato o canal de comunicación requiere modificar esta clase.
 */
public class DivergentChange {

    public static void main(String[] args) {
        ProfileManager pm = new ProfileManager();
        User user = new User("1", "Ana García", "ana@example.com");

        String result = demoDivergentChange(pm, user);
        System.out.println(result);
    }

    public static class User {
        String id;
        String name;
        String email;

        public User(String id, String name, String email) {
            this.id = id;
            this.name = name;
            this.email = email;
        }
    }

    public static class ProfileManager {
        private Map<String, User> store = new HashMap<>();

        public void register(User user) {
            if (user.name == null || user.name.trim().isEmpty()) {
                throw new IllegalArgumentException("invalid name");
            }
            if (!user.email.contains("@")) {
                throw new IllegalArgumentException("invalid email");
            }
            store.put(user.id, user);
        }

        public void updateEmail(String id, String newEmail) {
            if (!newEmail.contains("@")) {
                throw new IllegalArgumentException("invalid email");
            }
            User u = store.get(id);
            if (u == null) {
                throw new IllegalArgumentException("not found");
            }
            u.email = newEmail;
            store.put(id, u);
        }

        public String exportAsJson() {
            List<String> users = store.values().stream()
                .map(u -> String.format("{\"id\":\"%s\",\"name\":\"%s\",\"email\":\"%s\"}",
                    u.id, u.name, u.email))
                .collect(Collectors.toList());
            return "[" + String.join(",", users) + "]";
        }

        public String exportAsCsv() {
            List<String> rows = new ArrayList<>();
            rows.add("id,name,email");
            for (User u : store.values()) {
                rows.add(String.format("%s,%s,%s", u.id, u.name, u.email));
            }
            return String.join("\n", rows);
        }

        public String sendWelcomeEmail(User user) {
            return String.format("Welcome %s! Sent to %s", user.name, user.email);
        }
    }

    public static String demoDivergentChange(ProfileManager pm, User u) {
        pm.register(u);
        pm.updateEmail(u.id, u.email);
        return pm.exportAsJson();
    }
}
