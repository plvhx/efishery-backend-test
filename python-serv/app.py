from flask import Flask, request
from core.factory.password import PasswordFactory
from core.validator.rule import Rule
from core.validator.validator import Validator

import os
import core.http as http
import core.util as util
import core.validator.type as data_type

app = Flask(__name__)

@app.route('/register', methods=['POST'])
def register():
	content   = request.json
	validator = Validator().set_rules([
		Rule('name', data_type.STRING),
		Rule('phone', data_type.STRING),
		Rule('role', data_type.STRING)
	])

	try:
		validator.validate(content)
	except BaseException as e:
		return http.handle_request_exception(e, http.HTTP_UNPROCESSABLE_ENTITY)

	password = PasswordFactory.generate()

@app.route("/login")
def login():
	return "This is /login"

@app.route("/token/parse")
def parse_token():
	return open(util.get_auth_cache_file()).read(1024)
