import pytest
from .expand_migrate_contract import (
    User,
    format_greeting,
    format_email_header,
    format_display_name,
    build_user_summary
)


class TestExpandMigrateContract:
    """Expand-Migrate-Contract (Parallel Change)"""

    @pytest.fixture
    def alice(self):
        return User(id='u-1', name='Alice Smith', email='alice@example.com')

    @pytest.fixture
    def bob(self):
        return User(id='u-2', name='Bob Jones', email='bob@example.com')

    def test_format_greeting(self, alice):
        """should greet the user by name"""
        assert format_greeting(alice) == 'Hello, Alice Smith!'

    def test_format_email_header(self, alice):
        """should format the email header with name and email"""
        assert format_email_header(alice) == 'From: Alice Smith <alice@example.com>'

    def test_format_display_name(self, alice):
        """should format name with id"""
        assert format_display_name(alice) == 'Alice Smith (u-1)'

    def test_build_user_summary(self, alice, bob):
        """should list all user names"""
        assert build_user_summary([alice, bob]) == '- Alice Smith\n- Bob Jones'

    def test_build_user_summary_empty(self):
        """should return empty string for empty list"""
        assert build_user_summary([]) == ''
