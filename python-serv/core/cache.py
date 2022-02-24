import json


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


def put_cache_file(file, key, rbuf):
    if cache_key_exists(file, key):
        return

    handler = open(file, "rb")
    buf = handler.read(1024)
    data = json.loads(buf)

    handler.close()

    data["data"][key] = rbuf


def update_cache_file(file, key, rbuf):
    if not cache_key_exists(file, key):
        return

    handler = open(file, "rb")
    buf = handler.read(1024)
    data = json.loads(buf)

    handler.close()

    data["data"][key] = rbuf
