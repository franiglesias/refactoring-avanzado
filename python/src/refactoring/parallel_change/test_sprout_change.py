import pytest
from .sprout_change import CartItem, Region, calculate_total


class TestSproutChange:
    """Sprout Change (Parallel Change)"""

    @pytest.fixture
    def cart(self) -> list[CartItem]:
        return [
            CartItem(id='p1', price=10, qty=2, category='general'),
            CartItem(id='b1', price=20, qty=1, category='books'),
            CartItem(id='f1', price=15, qty=4, category='food'),
        ]

    def execute_subject(self, cart: list[CartItem], region: Region) -> float:
        return calculate_total(cart, region)

    def test_should_calculate_the_total_for_eu(self, cart):
        """should calculate the total for EU"""
        assert self.execute_subject(cart, 'EU') == 104

    def test_should_calculate_the_total_for_us(self, cart):
        """should calculate the total for US"""
        assert self.execute_subject(cart, 'US') == 107
