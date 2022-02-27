import json
import os
import core.util as util


def cache_key_exists(file, key):
    handler = open(file, "rb")
    buf = handler.read(1024)
    data = json.loads(buf)

    handler.close()

    try:
        data["data"][key]
    except KeyError as e:
        return False

    return True


def put_cache_file(file, key, obj):
    if cache_key_exists(file, key):
        return

    handler = open(file, "r+")
    buf = handler.read(1024)
    data = json.loads(buf)

    data["data"][key] = obj

    handler.seek(0, os.SEEK_SET)
    handler.write(json.dumps(data))
    handler.close()


def update_cache_file(file, key, obj):
    if not cache_key_exists(file, key):
        return

    handler = open(file, "r+")
    buf = handler.read(1024)
    data = json.loads(buf)

    data["data"][key] = obj

    handler.seek(0, os.SEEK_SET)
    handler.write(json.dumps(data))
    handler.close()


def get_cache_data(file):
    handler = open(file, "rb")
    buf = handler.read(1024)
    data = json.loads(buf)

    handler.close()
    return data


def get_auth_cache_data():
    return get_cache_data(util.get_auth_cache_file())


def get_currency_cache_data():
    return get_cache_data(util.get_currency_cache_file())


def get_regular_cache_data():
    return get_cache_data(util.get_data_cache_file())
