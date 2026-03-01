package com.refactoring.codesmells.oopabusers;

public class RefusedBequest {

    public static void main(String[] args) {
        demoRefusedBequest(true);
        demoRefusedBequest(false);
    }

    public static class BaseController {
        public void start() {
            System.out.println("starting");
        }

        public void stop() {
            System.out.println("stopping");
        }

        public void reset() {
            System.out.println("resetting");
        }
    }

    public static class ReadOnlyController extends BaseController {
        @Override
        public void start() {
            // No hace nada - rechaza la herencia
        }

        @Override
        public void stop() {
            // No hace nada - rechaza la herencia
        }
    }

    public static void demoRefusedBequest(boolean readonly) {
        BaseController controller = readonly ? new ReadOnlyController() : new BaseController();
        controller.start();
        controller.stop();
    }
}
