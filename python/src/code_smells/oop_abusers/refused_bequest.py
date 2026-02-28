"""
Refused Bequest Code Smell Example

A subclass inherits methods from a parent class but doesn't use them
or overrides them with empty implementations.

This indicates incorrect use of inheritance - the subclass doesn't need
all the parent's functionality.

Refactoring suggestion: Replace Inheritance with Delegation, Push Down Method
"""


class BaseController:
    """Base controller with full functionality"""

    def start(self) -> None:
        print('starting')

    def stop(self) -> None:
        print('stopping')

    def reset(self) -> None:
        print('resetting')


class ReadOnlyController(BaseController):
    """
    Refused bequest: inherits start() and stop() but refuses to use them.

    This class overrides methods with empty implementations because it
    doesn't need that functionality. This suggests inheritance is not
    the right relationship here.
    """

    def start(self) -> None:
        pass  # Refuses to implement

    def stop(self) -> None:
        pass  # Refuses to implement


def demo_refused_bequest(readonly: bool) -> None:
    controller: BaseController = ReadOnlyController() if readonly else BaseController()
    controller.start()
    controller.stop()
