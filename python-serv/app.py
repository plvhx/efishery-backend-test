from flask import Flask, request, jsonify
from core.factory.password import PasswordFactory
from core.validator.rule import Rule
from core.validator.validator import Validator

import jwt
import os
import time
import core.cache as cache
import core.hash as hashgen
import core.http as http
import core.resource as resource
import core.util as util
import core.validator.type as data_type

app = Flask(__name__)
password = util.get_jwt_validation_password()


@app.route("/auth/register", methods=["POST"])
def register():
    if not http.is_json():
        return (
            jsonify(
                {
                    "message": "Content-Type must be 'application/json'.",
                    "code": http.HTTP_UNSUPPORTED_MEDIA_TYPE,
                }
            ),
            http.HTTP_UNSUPPORTED_MEDIA_TYPE,
        )

    content = request.json
    validator = Validator().set_rules(
        [
            Rule("name", data_type.STRING),
            Rule("phone", data_type.STRING),
            Rule("role", data_type.STRING),
        ]
    )

    try:
        validator.validate(content)
    except Exception as e:
        return http.handle_request_exception(e, http.HTTP_UNPROCESSABLE_ENTITY)

    password = PasswordFactory.generate()
    rbuf = {
        "name": content["name"],
        "phone": content["phone"],
        "role": content["role"],
        "password": password,
    }

    cache.put_cache_file(
        util.get_auth_cache_file(),
        hashgen.calculate_hash(content["name"] + content["phone"] + content["role"]),
        rbuf,
    )

    return jsonify(rbuf), http.HTTP_CREATED


@app.route("/auth/token", methods=["POST"])
def auth():
    if not http.is_json():
        return (
            jsonify(
                {
                    "message": "Content-Type must be 'application/json'.",
                    "code": http.HTTP_UNSUPPORTED_MEDIA_TYPE,
                }
            ),
            http.HTTP_UNSUPPORTED_MEDIA_TYPE,
        )

    content = request.json
    validator = Validator().set_rules(
        [Rule("phone", data_type.STRING), Rule("password", data_type.STRING)]
    )

    try:
        validator.validate(content)
    except Exception as e:
        return http.handle_request_exception(e, http.HTTP_UNPROCESSABLE_ENTITY)

    auths = cache.get_auth_cache_data()
    match = None

    for auth in auths["data"].items():
        if (
            content["phone"] == auth[1]["phone"]
            and content["password"] == auth[1]["password"]
        ):
            match = {
                "name": auth[1]["name"],
                "phone": auth[1]["phone"],
                "role": auth[1]["role"],
                "timestamp": int(time.time()) + 3600,
            }

            break

    if match == None:
        return (
            jsonify(
                {
                    "message": "Either phone or password did not match.",
                    "code": http.HTTP_NOT_FOUND,
                }
            ),
            http.HTTP_NOT_FOUND,
        )

    return (
        jsonify(
            {
                "access_token": jwt.encode(match, password, algorithm="HS256").decode(
                    "utf-8"
                )
            }
        ),
        http.HTTP_OK,
    )


@app.route("/auth/parse", methods=["POST"])
def parse_auth_token():
    if not http.is_json():
        return (
            jsonify(
                {
                    "message": "Content-Type must be 'application/json'.",
                    "code": http.HTTP_UNSUPPORTED_MEDIA_TYPE,
                }
            ),
            http.HTTP_UNSUPPORTED_MEDIA_TYPE,
        )

    content = request.json
    validator = Validator().add_rule(Rule("token", data_type.STRING))

    try:
        validator.validate(content)
    except Exception as e:
        return http.handle_request_exception(e, http.HTTP_UNPROCESSABLE_ENTITY)

    try:
        claims = jwt.decode(content["token"], password, {"verify_signature": True})
    except jwt.exceptions.DecodeError as e:
        return http.handle_request_exception(e, http.HTTP_UNPROCESSABLE_ENTITY)

    return jsonify(claims), http.HTTP_OK


@app.route("/fetch", methods=["GET"])
def get_resource():
    if not http.has_authorization():
        return (
            jsonify(
                {
                    "message": "Authorization header must be provided.",
                    "code": http.HTTP_BAD_REQUEST,
                }
            ),
            http.HTTP_BAD_REQUEST,
        )

    try:
        http.parse_bearer_token(password)
    except Exception as e:
        return http.handle_request_exception(e, http.HTTP_UNAUTHORIZED)

    return jsonify(resource.fetch_resources()), http.HTTP_OK

if __name__ == '__main__':
    port = int(os.environ.get('PORT', 5000))
    app.run(debug=False, host='0.0.0.0', port=port)
