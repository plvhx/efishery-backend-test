from core.exception import KeyNotFoundError

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
				raise KeyNotFoundError("Mandatory key '%s' not found in current object." % self.rules[i].get_name())

		return True
