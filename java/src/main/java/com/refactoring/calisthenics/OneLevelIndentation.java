package com.refactoring.calisthenics;

import java.util.List;
import java.util.Map;

/**
 * Example demonstrating multiple levels of indentation.
 * Rule: Only one level of indentation per method.
 */
public class OneLevelIndentation {

    public static void processOrders(List<Map<String, Object>> orders) {
        for (Map<String, Object> order : orders) {
            if ("pending".equals(order.get("status"))) {
                if ((Double) order.get("total") > 0) {
                    if (order.get("customer") != null) {
                        @SuppressWarnings("unchecked")
                        Map<String, Object> customer = (Map<String, Object>) order.get("customer");
                        if (customer.get("email") != null && !customer.get("email").toString().isEmpty()) {
                            System.out.printf("Processing order for %s%n", customer.get("email"));
                        }
                    }
                }
            }
        }
    }

    /*
     * Exercise: Refactor this code to have only one level of indentation
     * Techniques to consider:
     * 1. Extract methods
     * 2. Early returns
     * 3. Guard clauses
     * 4. Replace nested conditionals with boolean expressions
     */
}
