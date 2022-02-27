class KeyNotFoundError(Exception):
    pass


class InvalidHashAlgorithmError(Exception):
    pass


class InvalidBearerTokenError(Exception):
    pass


class JwtDecodeError(Exception):
    pass
