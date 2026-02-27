from dataclasses import dataclass
from typing import Optional


@dataclass
class User:
    id: str
    name: str
    email: str
    first_name: Optional[str] = None
    last_name: Optional[str] = None


# --- Consumers of User.name ---

def format_greeting(user: User) -> str:
    return f"Hello, {user.name}!"


def format_email_header(user: User) -> str:
    return f"From: {user.name} <{user.email}>"


def format_display_name(user: User) -> str:
    return f"{user.name} ({user.id})"


def build_user_summary(users: list[User]) -> str:
    return '\n'.join(f"- {u.name}" for u in users)
