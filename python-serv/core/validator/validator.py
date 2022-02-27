from core.exception import KeyNotFoundError

import re
import core.validator.type as data_type


class Validator(object):
    def __init__(self):
        self.rules = []

    def add_rule(self, rule):
        self.rules.append(rule)
        return self

    def set_rules(self, rules):
        for i in range(len(rules)):
            self.add_rule(rules[i])
        return self

    def get_rules(self):
        return self.rules

    def validate(self, obj):
        for i in range(len(self.rules)):
            if not self.rules[i].is_required():
                obj[self.rules[i].get_name()] = None
                continue

            try:
                obj[self.rules[i].get_name()]
            except KeyError as e:
                raise KeyNotFoundError(
                    "Mandatory key '%s' not found in current object."
                    % self.rules[i].get_name()
                )

            if not self._type_check(self.rules[i], obj[self.rules[i].get_name()]):
                return self._parse_type_name(obj[self.rules[i].get_name()])

        return True

    def _type_check(self, rule, val):
        if rule.get_type() == data_type.INTEGER:
            return isinstance(val, int)
        elif rule.get_type() == data_type.STRING:
            return isinstance(val, str)
        elif rule.get_type() == data_type.FLOAT:
            return isinstance(val, float)
        elif rule.get_type() == data_type.BOOLEAN:
            return isinstance(val, bool)

    def _parse_type_name(self, val):
        name = str(type(val))
        matches = re.findall("^(?:<class ')(int|str|float|bool)(?:')$", name)

        return matches
