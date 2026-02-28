"""
Data Class Code Smell Example

A class that only contains fields and getters/setters without any behavior.
Other classes manipulate the data class's data directly, violating encapsulation.

Refactoring suggestion: Move Method, Encapsulate Field
"""
from datetime import datetime
from uuid import uuid4


class UserRecord:
    """
    Data class: only stores data, no behavior.

    All logic that operates on this data lives in other classes,
    making the design anemic and non-object-oriented.
    """

    def __init__(self, user_id: str, name: str, email: str, created_at: datetime) -> None:
        self.id = user_id
        self.name = name
        self.email = email
        self.created_at = created_at


class UserService:
    """Service class that manipulates UserRecord data"""

    def create_user(self, name: str, email: str) -> UserRecord:
        if '@' not in email:
            raise ValueError('Invalid email')

        return UserRecord(str(uuid4()), name, email, datetime.now())

    def update_user_email(self, user: UserRecord, new_email: str) -> None:
        if '@' not in new_email:
            raise ValueError('Invalid email')
        user.email = new_email


class UserReportGenerator:
    """Another class manipulating UserRecord data"""

    def generate_user_summary(self, user: UserRecord) -> str:
        return f"User {user.name} ({user.email}) created on {user.created_at.strftime('%Y-%m-%d')}"


def demo_data_class() -> str:
    service = UserService()
    report = UserReportGenerator()
    user = service.create_user('Lina', 'lina@example.com')
    service.update_user_email(user, 'lina+news@example.com')
    return report.generate_user_summary(user)
