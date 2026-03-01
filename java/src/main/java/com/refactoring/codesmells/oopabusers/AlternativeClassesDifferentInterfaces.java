package com.refactoring.codesmells.oopabusers;

public class AlternativeClassesDifferentInterfaces {

    public static void main(String[] args) {
        useAltClasses("logger", "Hello from logger");
        useAltClasses("writer", "Hello from writer");
    }

    public static class TextLogger {
        public void log(String message) {
            System.out.println("[text] " + message);
        }
    }

    public static class MessageWriter {
        public void write(String entry) {
            System.out.println("[text] " + entry);
        }
    }

    public static void useAltClasses(String choice, String msg) {
        if ("logger".equals(choice)) {
            new TextLogger().log(msg);
        } else {
            new MessageWriter().write(msg);
        }
    }
}
