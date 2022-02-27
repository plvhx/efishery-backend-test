from flask import jsonify, request
from core.exception import InvalidBearerTokenError, JwtDecodeError

import jwt

"""
10x status code.
"""
HTTP_CONTINUE = 100
HTTP_SWITCHING_PROTOCOLS = 101
HTTP_PROCESSING_WEBDAV = 102
HTTP_EARLY_HINTS = 103

"""
20x status code.
"""
HTTP_OK = 200
HTTP_CREATED = 201
HTTP_ACCEPTED = 202
HTTP_NON_AUTHORITATIVE_INFORMATION = 203
HTTP_NO_CONTENT = 204
HTTP_RESET_CONTENT = 205
HTTP_PARTIAL_CONTENT = 206
HTTP_MULTI_STATUS_WEBDAV = 207
HTTP_ALREADY_REPORTED_WEBDAV = 208
HTTP_IM_USED = 226

"""
30x status code
"""
HTTP_MULTIPLE_CHOICE = 300
HTTP_MOVED_PERMANENTLY = 301
HTTP_FOUND = 302
HTTP_SEE_OTHER = 303
HTTP_NOT_MODIFIED = 304
HTTP_USE_PROXY = 305
HTTP_UNUSED = 306
HTTP_TEMPORARY_REDIRECT = 307
HTTP_PERMANENT_REDIRECT = 308

"""
40x status code
"""
HTTP_BAD_REQUEST = 400
HTTP_UNAUTHORIZED = 401
HTTP_PAYMENT_REQUIRED = 402
HTTP_FORBIDDEN = 403
HTTP_NOT_FOUND = 404
HTTP_METHOD_NOT_ALLOWED = 405
HTTP_NOT_ACCEPTABLE = 406
HTTP_PROXY_AUTHENTICATION_REQUIRED = 407
HTTP_REQUEST_TIMEOUT = 408
HTTP_CONFLICT = 409
HTTP_GONE = 410
HTTP_LENGTH_REQUIRED = 411
HTTP_PRECONDITION_FAILED = 412
HTTP_PAYLOAD_TOO_LARGE = 413
HTTP_URI_TOO_LONG = 414
HTTP_UNSUPPORTED_MEDIA_TYPE = 415
HTTP_RANGE_NOT_SATISFIABLE = 416
HTTP_EXPECTATION_FAILED = 417
HTTP_IM_A_TEAPOT = 418
HTTP_MISDIRECTED_REQUEST = 421
HTTP_UNPROCESSABLE_ENTITY = 422
HTTP_LOCKED = 423
HTTP_FAILED_DEPENDENCY = 424
HTTP_TOO_EARLY = 425
HTTP_UPGRADE_REQUIRED = 426
HTTP_PRECONDITION_REQUIRED = 428
HTTP_TOO_MANY_REQUESTS = 429
HTTP_REQUEST_HEADER_FIELDS_TOO_LARGE = 431
HTTP_UNAVAILABLE_FOR_LEGAL_REASONS = 451

"""
50x status code
"""
HTTP_INTERNAL_SERVER_ERROR = 500
HTTP_NOT_IMPLEMENTED = 501
HTTP_BAD_GATEWAY = 502
HTTP_SERVICE_UNAVAILABLE = 503
HTTP_GATEWAY_TIMEOUT = 504
HTTP_VERSION_NOT_SUPPORTED = 505
HTTP_VARIANT_ALSO_NEGOTIATES = 506
HTTP_INSUFFICIENT_STORAGE_WEBDAV = 507
HTTP_LOOP_DETECTED_WEBDAV = 508
HTTP_NOT_EXTENDED = 510
HTTP_NETWORK_AUTHENTICATION_REQUIRED = 511

_status_code_map = {
    HTTP_CONTINUE: "Continue",
    HTTP_SWITCHING_PROTOCOLS: "Switching Protocols",
    HTTP_PROCESSING_WEBDAV: "Processing (WebDAV)",
    HTTP_EARLY_HINTS: "Early Hints",
    HTTP_OK: "Ok",
    HTTP_CREATED: "Created",
    HTTP_ACCEPTED: "Accepted",
    HTTP_NON_AUTHORITATIVE_INFORMATION: "Non Authoritative Information",
    HTTP_NO_CONTENT: "No Content",
    HTTP_RESET_CONTENT: "Reset Content",
    HTTP_PARTIAL_CONTENT: "Partial Content",
    HTTP_MULTI_STATUS_WEBDAV: "Multi Status (WebDAV)",
    HTTP_ALREADY_REPORTED_WEBDAV: "Already Reported (WebDAV)",
    HTTP_IM_USED: "I'm Used",
    HTTP_MULTIPLE_CHOICE: "Multiple Choice",
    HTTP_MOVED_PERMANENTLY: "Moved Permanently",
    HTTP_FOUND: "Found",
    HTTP_SEE_OTHER: "See Other",
    HTTP_NOT_MODIFIED: "Not Modified",
    HTTP_USE_PROXY: "Use Proxy",
    HTTP_UNUSED: "Unused",
    HTTP_TEMPORARY_REDIRECT: "Temporary Redirect",
    HTTP_PERMANENT_REDIRECT: "Permanent Redirect",
    HTTP_BAD_REQUEST: "Bad Request",
    HTTP_UNAUTHORIZED: "Unauthorized",
    HTTP_PAYMENT_REQUIRED: "Payment Required",
    HTTP_FORBIDDEN: "Forbidden",
    HTTP_NOT_FOUND: "Not Found",
    HTTP_METHOD_NOT_ALLOWED: "Method Not Allowed",
    HTTP_NOT_ACCEPTABLE: "Not Acceptable",
    HTTP_PROXY_AUTHENTICATION_REQUIRED: "Proxy Authentication Required",
    HTTP_REQUEST_TIMEOUT: "Request Timeout",
    HTTP_CONFLICT: "Conflict",
    HTTP_GONE: "Gone",
    HTTP_LENGTH_REQUIRED: "Length Required",
    HTTP_PRECONDITION_FAILED: "Precondition Failed",
    HTTP_PAYLOAD_TOO_LARGE: "Payload Too Large",
    HTTP_URI_TOO_LONG: "URI Too Long",
    HTTP_UNSUPPORTED_MEDIA_TYPE: "Unsupported Media Type",
    HTTP_RANGE_NOT_SATISFIABLE: "Range Not Satisfiable",
    HTTP_EXPECTATION_FAILED: "Expectation Failed",
    HTTP_IM_A_TEAPOT: "I'm A Teapot",
    HTTP_MISDIRECTED_REQUEST: "Misdirected Request",
    HTTP_UNPROCESSABLE_ENTITY: "Unprocessable Entity",
    HTTP_LOCKED: "Locked",
    HTTP_FAILED_DEPENDENCY: "Failed Dependency",
    HTTP_TOO_EARLY: "Too Early",
    HTTP_UPGRADE_REQUIRED: "Upgrade Required",
    HTTP_PRECONDITION_REQUIRED: "Precondition Required",
    HTTP_TOO_MANY_REQUESTS: "Too Many Requests",
    HTTP_REQUEST_HEADER_FIELDS_TOO_LARGE: "Request Header Fields Too Large",
    HTTP_UNAVAILABLE_FOR_LEGAL_REASONS: "Unavailable For Legal Reasons",
    HTTP_INTERNAL_SERVER_ERROR: "Internal Server Error",
    HTTP_NOT_IMPLEMENTED: "Not Implemented",
    HTTP_BAD_GATEWAY: "Bad Gateway",
    HTTP_SERVICE_UNAVAILABLE: "Service Unavailable",
    HTTP_GATEWAY_TIMEOUT: "Gateway Timeout",
    HTTP_VERSION_NOT_SUPPORTED: "HTTP Version Not Supported",
    HTTP_VARIANT_ALSO_NEGOTIATES: "Variant Also Negotiates",
    HTTP_INSUFFICIENT_STORAGE_WEBDAV: "Insufficient Storage (WebDAV)",
    HTTP_LOOP_DETECTED_WEBDAV: "Loop Detected (WebDAV)",
    HTTP_NOT_EXTENDED: "Not Extended",
    HTTP_NETWORK_AUTHENTICATION_REQUIRED: "Network Authentication Required",
}


def get_response_reason(code=HTTP_OK):
    return _status_code_map[code]


def handle_request_exception(exception, code=HTTP_BAD_REQUEST):
    return jsonify({"message": str(exception), "code": code}), code


def is_json():
    if request.content_type != "application/json":
        return False
    return True


def has_authorization():
    try:
        request.headers["Authorization"]
    except KeyError as e:
        return False

    return True


def parse_bearer_token(password):
    bearer = request.headers["Authorization"]
    splitted = bearer.split(" ")

    if splitted[0] != "Bearer":
        raise InvalidBearerTokenError("Bearer token must be prefixed by 'Bearer'.")

    try:
        claims = jwt.decode(splitted[1], password, {"verify_signature": True})
    except jwt.exceptions.DecodeError as e:
        raise JwtDecodeError(str(e))

    return claims
