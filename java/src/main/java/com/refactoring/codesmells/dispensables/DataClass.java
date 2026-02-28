package com.refactoring.codesmells.dispensables;

import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.UUID;

/**
 * Code Smell: Data Class [Clase de datos]
 * UserRecord es una clase que solo contiene datos sin comportamiento propio.
 * Toda la lógica está en otras clases (UserService, UserReportGenerator).
 *
 * La clase UserRecord debería contener la lógica relacionada con usuarios.
 */
public class DataClass {

    public static void main(String[] args) {
        String summary = demoDataClass();
        System.out.println(summary);
    }

    public static class UserRecord {
        public String id;
        public String name;
        public String email;
        public Date createdAt;

        public UserRecord(String id, String name, String email, Date createdAt) {
            this.id = id;
            this.name = name;
            this.email = email;
            this.createdAt = createdAt;
        }
    }

    public static class UserService {
        public UserRecord createUser(String name, String email) {
            if (!email.contains("@")) {
                throw new IllegalArgumentException("Invalid email");
            }

            return new UserRecord(UUID.randomUUID().toString(), name, email, new Date());
        }

        public void updateUserEmail(UserRecord user, String newEmail) {
            if (!newEmail.contains("@")) {
                throw new IllegalArgumentException("Invalid email");
            }
            user.email = newEmail;
        }
    }

    public static class UserReportGenerator {
        public String generateUserSummary(UserRecord user) {
            SimpleDateFormat sdf = new SimpleDateFormat("dd/MM/yyyy");
            return String.format("User %s (%s) created on %s",
                user.name, user.email, sdf.format(user.createdAt));
        }
    }

    public static String demoDataClass() {
        UserService service = new UserService();
        UserReportGenerator report = new UserReportGenerator();
        UserRecord user = service.createUser("Lina", "lina@example.com");
        service.updateUserEmail(user, "lina+news@example.com");
        return report.generateUserSummary(user);
    }
}
