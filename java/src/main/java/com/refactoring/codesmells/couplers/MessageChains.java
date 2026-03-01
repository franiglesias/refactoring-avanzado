package com.refactoring.codesmells.couplers;

public class MessageChains {

    public static void main(String[] args) {
        Root root = new Root(new Level1(new Level2(42)));
        int value = readDeep(root);
        System.out.println("Valor profundo: " + value);
    }

    public static class Level2 {
        private int value;

        public Level2(int value) {
            this.value = value;
        }

        public int getValue() {
            return value;
        }
    }

    public static class Level1 {
        private Level2 next;

        public Level1(Level2 next) {
            this.next = next;
        }

        public Level2 getNext() {
            return next;
        }
    }

    public static class Root {
        private Level1 next;

        public Root(Level1 next) {
            this.next = next;
        }

        public Level1 getNext() {
            return next;
        }
    }

    public static int readDeep(Root root) {
        return root.getNext().getNext().getValue();
    }
}
