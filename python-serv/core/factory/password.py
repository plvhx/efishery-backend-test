import random
import string
from core.factory.abstract import AbstractFactory


class PasswordFactory(AbstractFactory):
    @staticmethod
    def generate(length=4):
        password = ""
        source = string.ascii_lowercase + string.ascii_uppercase + string.digits

        for i in range(length):
            password += random.choice(source)

        return password
