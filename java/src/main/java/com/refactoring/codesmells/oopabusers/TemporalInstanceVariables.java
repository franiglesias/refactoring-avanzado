package com.refactoring.codesmells.oopabusers;

import java.util.ArrayList;
import java.util.List;

/**
 * Code Smell: Temporal Instance Variables [Variables de instancia temporales]
 * PizzaOrder usa variables de instancia (size, toppings, address) que solo son
 * válidas durante un período específico del flujo de trabajo.
 *
 * Estas variables deben ser parte de un objeto de estado separado o parámetros.
 */
public class TemporalInstanceVariables {

    public static void main(String[] args) {
        String result = demoPizzaOrder();
        System.out.println(result);
    }

    public static class PizzaOrder {
        private String size;
        private List<String> toppings = new ArrayList<>();
        private String address;

        public void start(String size) {
            this.size = size;
            this.toppings = new ArrayList<>();
            this.address = null;
        }

        public void addTopping(String topping) {
            if (this.size == null) {
                return;
            }
            this.toppings.add(topping);
        }

        public void setDeliveryAddress(String address) {
            this.address = address;
        }

        public String place() {
            String sizeStr = this.size != null ? this.size : "?";
            String addressStr = this.address != null ? this.address : "UNKNOWN";
            String summary = String.format("Pizza %s to %s with [%s]",
                sizeStr, addressStr, String.join(", ", this.toppings));

            // Reset temporal state
            this.size = null;
            this.address = null;
            this.toppings = new ArrayList<>();

            return summary;
        }
    }

    public static String demoPizzaOrder() {
        PizzaOrder o = new PizzaOrder();
        o.start("L");
        o.addTopping("pepperoni");
        o.addTopping("mushroom");
        o.setDeliveryAddress("123 Main St");
        return o.place();
    }
}
