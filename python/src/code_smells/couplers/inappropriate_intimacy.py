"""
Inappropriate Intimacy Code Smell Example

Two classes are too tightly coupled. They access each other's internal data
and methods excessively, breaking encapsulation.

Refactoring suggestion: Move Method, Move Field, Change Bidirectional Association to Unidirectional
"""


class Budget:
    """Budget class with validation logic"""

    def __init__(self, amount: float) -> None:
        if amount < 900:
            raise ValueError('Budget too low')
        if amount > 3000:
            raise ValueError('Budget too high')
        self.amount = amount

    def raise_budget(self, amount: float) -> None:
        if self.amount + amount > 3000:
            raise ValueError('Budget exceeded')
        self.amount += amount


class Team:
    """Team knows too much about Manager's internals"""

    def __init__(self, name: str, budget: float) -> None:
        self._budget = Budget(budget)
        self._name = name
        self._manager: 'Manager | None' = None

    def assign_manager(self, manager: 'Manager') -> None:
        self._manager = manager
        manager.assign_team(self)  # Bidirectional coupling

    def raise_budget(self, amount: float) -> None:
        self._budget = Budget(self._budget.amount + amount)

    def rename(self, new_name: str) -> None:
        self._name = new_name


class Manager:
    """Manager manipulates Team's internals directly"""

    def __init__(self, name: str) -> None:
        self._name = name
        self._team: Team | None = None

    def assign_team(self, team: Team) -> None:
        if self._team:
            raise ValueError('Team already assigned')
        self._team = team

    def raise_team_budget(self, amount: float) -> None:
        """Directly manipulating team's budget"""
        if self._team:
            self._team.raise_budget(amount)

    def rename_team(self, new_name: str) -> None:
        """Directly manipulating team's name"""
        if self._team:
            self._team.rename(new_name)


def demo_inappropriate_intimacy() -> Team:
    team = Team('Core', 1000)
    manager = Manager('Alice')
    manager.assign_team(Team('Frontend', 2000))
    team.assign_manager(manager)
    manager.raise_team_budget(200)
    manager.rename_team('Platform')
    return team
