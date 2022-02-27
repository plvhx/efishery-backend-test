import os


def get_cache_dir():
    return os.getcwd() + "/../cache"


def get_auth_cache_file():
    return get_cache_dir() + "/auth.json"


def get_currency_cache_file():
    return get_cache_dir() + "/currency.json"


def get_data_cache_file():
    return get_cache_dir() + "/data.json"


def get_jwt_validation_password():
    handler = open(os.getcwd() + "/../.jwt_pass", "rb")
    return handler.read(1024)
