package com.refactoring.calisthenics;

public class NoAbbreviations {

    public static void main(String[] args) {
        C c = new C("admin", "secret", "localhost", "prod");
        System.out.println("Connection: " + c.cx());
    }

    public static class C {
        String u;
        String p;
        String s;
        String e;

        public C(String u, String p, String s, String e) {
            this.u = u;
            this.p = p;
            this.s = s;
            this.e = e;
        }

        public String cx() {
            return String.format("%s:%s@%s/%s", u, p, s, e);
        }
    }
}
