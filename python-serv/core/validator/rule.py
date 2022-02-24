from abc import ABC


class AbstractRule(ABC):
    pass


class Rule(AbstractRule):
    def __init__(self, name, data_type, required=True):
        self.name = name
        self.type = data_type
        self.required = required

    def get_name(self):
        return self.name

    def get_type(self):
        return self.type

    def is_required(self):
        return self.required
