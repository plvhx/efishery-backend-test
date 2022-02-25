import hashlib
from core.exception import InvalidHashAlgorithmError

SHA1 = 1
SHA224 = 2
SHA256 = 3
SHA384 = 4
SHA512 = 5

_hash_to_impl_map = {
    SHA1: hashlib.sha1(),
    SHA224: hashlib.sha224(),
    SHA256: hashlib.sha256(),
    SHA384: hashlib.sha384(),
    SHA512: hashlib.sha512(),
}


def calculate_hash(str, algorithm=SHA1):
    try:
        hobj = _hash_to_impl_map[algorithm]
    except KeyError as e:
        raise InvalidHashAlgorithmError("Invalid hashing algorithm")

    hobj.update(bytes(str.encode("ascii")))
    return hobj.hexdigest()
