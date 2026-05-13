package com.refactoring.parallelchange.sproutchange;

import org.junit.jupiter.api.Test;

import java.util.List;

import static org.assertj.core.api.Assertions.assertThat;

class SproutChangeTest {

    private static final List<SproutChange.CartItem> CART = List.of(
            new SproutChange.CartItem("p1", 10, 2, "general"),
            new SproutChange.CartItem("b1", 20, 1, "books"),
            new SproutChange.CartItem("f1", 15, 4, "food")
    );

    private double executeSubject(List<SproutChange.CartItem> cart, SproutChange.Region region) {
        return SproutChange.calculateTotal(cart, region);
    }

    @Test
    void shouldCalculateTotalForEU() {
        assertThat(executeSubject(CART, SproutChange.Region.EU)).isEqualTo(104.0);
    }

    @Test
    void shouldCalculateTotalForUS() {
        assertThat(executeSubject(CART, SproutChange.Region.US)).isEqualTo(107.0);
    }
}
