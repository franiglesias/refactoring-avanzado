package com.refactoring.calisthenics;

/**
 * Object Calisthenics: Don't Use Abbreviations
 * Las abreviaturas dificultan la comprensión del código.
 * Los nombres deben ser completos y expresivos.
 *
 * Clase C con variables u, p, s, e y método cnx() - imposible de entender.
 */
public class NoAbbreviations {

    public static void main(String[] args) {
        C config = new C("admin", "secret", "localhost", "prod");
        System.out.println("Connection: " + config.cnx());
    }

    public static class C {
        String u;  // ¿username? ¿url? ¿user?
        String p;  // ¿password? ¿port? ¿path?
        String s;  // ¿server? ¿service? ¿schema?
        String e;  // ¿environment? ¿endpoint?

        public C(String u, String p, String s, String e) {
            this.u = u;
            this.p = p;
            this.s = s;
            this.e = e;
        }

        public String cnx() {  // ¿connection? ¿context?
            return String.format("%s:%s@%s/%s", u, p, s, e);
        }
    }
}
