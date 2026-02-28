"""
Divergent Change Code Smell Example

A class suffers changes for multiple unrelated reasons.
When you need to modify different aspects, you end up changing the same class.
This violates the Single Responsibility Principle.

Refactoring suggestion: Extract Class
"""
from typing import TypedDict


class User(TypedDict):
    id: str
    name: str
    email: str


class ProfileManager:
    """
    This class handles too many responsibilities:
    - User validation and registration
    - Email updates
    - Data export (JSON, CSV)
    - Email notifications

    Each responsibility is a different reason to change.
    """

    def __init__(self) -> None:
        self._store: dict[str, User] = {}

    def register(self, user: User) -> None:
        if not user['name'].strip():
            raise ValueError('invalid name')
        if '@' not in user['email']:
            raise ValueError('invalid email')
        self._store[user['id']] = user

    def update_email(self, user_id: str, new_email: str) -> None:
        if '@' not in new_email:
            raise ValueError('invalid email')
        user = self._store.get(user_id)
        if not user:
            raise ValueError('not found')
        user['email'] = new_email
        self._store[user_id] = user

    def export_as_json(self) -> str:
        import json
        return json.dumps(list(self._store.values()))

    def export_as_csv(self) -> str:
        rows = ['id,name,email']
        for user in self._store.values():
            rows.append(f"{user['id']},{user['name']},{user['email']}")
        return '\n'.join(rows)

    def send_welcome_email(self, user: User) -> str:
        return f"Welcome {user['name']}! Sent to {user['email']}"


def demo_divergent_change(pm: ProfileManager, user: User) -> str:
    pm.register(user)
    pm.update_email(user['id'], user['email'])
    return pm.export_as_json()
