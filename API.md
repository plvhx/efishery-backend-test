```
{
	"openapi": "3.0.0",
	"info": {
		"version": "1.0.0",
		"title": "eFishery Backend App Test",
		"license": {
			"name": "BSD-3-License"
		}
	},
	"servers": [
		{
			"url": "http://localhost:5000"
		},
		{
			"url": "http://localhost:9000"
		}
	],
	"paths": {
		"/auth/register": {
			"post": {
				"summary": "Register",
				"operationId": "register",
				"tags": [
					"auth"
				],
				"requestBody": {
					"content": {
						"application/json": {
							"schema": {
								"type": "object",
								"properties": {
									"name": {
										"type": "string"
									},
									"phone": {
										"type": "string"
									},
									"role": {
										"type": "string"
									}
								}
							}
						}
					},
					"required": true
				},
				"responses": {
					"201": {
						"description": "Created",
						"content": {
							"application/json": {
								"schema": {
									"$ref": "#/components/schemas/Auth"
								}
							}
						}
					},
					"405": {
						"description": "Method Not Allowed",
						"content": {
							"application/json": {
								"schema": {
									"type": "object",
									"properties": {
										"message": {
											"type": "string",
											"description": "Error message"
										},
										"code": {
											"type": "integer",
											"description": "HTTP response code"
										}
									}
								}
							}
						}
					},
					"415": {
						"description": "Unsupported Media Type",
						"content": {
							"application/json": {
								"schema": {
									"type": "object",
									"properties": {
										"message": {
											"type": "string",
											"description": "Error message"
										},
										"code": {
											"type": "integer",
											"description": "HTTP response code"
										}
									}
								}
							}
						}
					},
					"422": {
						"description": "Unprocessable Entity",
						"content": {
							"application/json": {
								"schema": {
									"type": "object",
									"properties": {
										"message": {
											"type": "string",
											"description": "Error message"
										},
										"code": {
											"type": "integer",
											"description": "HTTP response code"
										}
									}
								}
							}
						}
					},
					"500": {
						"description": "Internal Server Error",
						"content": {
							"application/json": {
								"schema": {
									"type": "object",
									"properties": {
										"message": {
											"type": "string",
											"description": "Error message"
										},
										"code": {
											"type": "integer",
											"description": "HTTP response code"
										}
									}
								}
							}
						}
					}
				}
			}
		},
		"/auth/token": {
			"post": {
				"summary": "Generate JWT Token",
				"operationId": "generateToken",
				"tags": [
					"auth"
				],
				"requestBody": {
					"content": {
						"application/json": {
							"schema": {
								"type": "object",
								"properties": {
									"phone": {
										"type": "string"
									},
									"password": {
										"type": "string"
									}
								}
							}
						}
					},
					"required": true
				},
				"responses": {
					"200": {
						"description": "Ok",
						"content": {
							"application/json": {
								"schema": {
									"type": "object",
									"properties": {
										"token": {
											"type": "string"
										}
									}
								}
							}
						}
					},
					"405": {
						"description": "Method Not Allowed",
						"content": {
							"application/json": {
								"schema": {
									"type": "object",
									"properties": {
										"message": {
											"type": "string",
											"description": "Error message"
										},
										"code": {
											"type": "integer",
											"description": "HTTP response code"
										}
									}
								}
							}
						}
					},
					"415": {
						"description": "Unsupported Media Type",
						"content": {
							"application/json": {
								"schema": {
									"type": "object",
									"properties": {
										"message": {
											"type": "string",
											"description": "Error message"
										},
										"code": {
											"type": "integer",
											"description": "HTTP response code"
										}
									}
								}
							}
						}
					},
					"422": {
						"description": "Unsupported Media Type",
						"content": {
							"application/json": {
								"schema": {
									"type": "object",
									"properties": {
										"message": {
											"type": "string",
											"description": "Error message"
										},
										"code": {
											"type": "integer",
											"description": "HTTP response code"
										}
									}
								}
							}
						}
					},
					"500": {
						"description": "Internal Server Error",
						"content": {
							"application/json": {
								"schema": {
									"type": "object",
									"properties": {
										"message": {
											"type": "string",
											"description": "Error message"
										},
										"code": {
											"type": "integer",
											"description": "HTTP response code"
										}
									}
								}
							}
						}
					}	
				}
			}
		},
		"/auth/parse": {
			"post": {
				"summary": "Parse Authentication Token",
				"operationId": "parseToken",
				"tags": [
					"auth"
				],
				"requestBody": {
					"content": {
						"application/json": {
							"schema": {
								"type": "object",
								"properties": {
									"token": {
										"type": "string"
									}
								}
							}
						}
					},
					"required": true
				},
				"responses": {
					"200": {
						"description": "Created",
						"content": {
							"application/json": {
								"schema": {
									"type": "object",
									"properties": {
										"name": {
											"type": "string"
										},
										"phone": {
											"type": "string"
										},
										"role": {
											"type": "string"
										},
										"timestamp": {
											"type": "integer"
										}
									}
								}
							}
						}
					},
					"405": {
						"description": "Method Not Allowed",
						"content": {
							"application/json": {
								"schema": {
									"type": "object",
									"properties": {
										"message": {
											"type": "string",
											"description": "Error message"
										},
										"code": {
											"type": "integer",
											"description": "HTTP response code"
										}
									}
								}
							}
						}
					},
					"415": {
						"description": "Unsupported Media Type",
						"content": {
							"application/json": {
								"schema": {
									"type": "object",
									"properties": {
										"message": {
											"type": "string",
											"description": "Error message"
										},
										"code": {
											"type": "integer",
											"description": "HTTP response code"
										}
									}
								}
							}
						}
					},
					"422": {
						"description": "Unprocessable Entity",
						"content": {
							"application/json": {
								"schema": {
									"type": "object",
									"properties": {
										"message": {
											"type": "string",
											"description": "Error message"
										},
										"code": {
											"type": "integer",
											"description": "HTTP response code"
										}
									}
								}
							}
						}
					},
					"400": {
						"description": "Bad Request",
						"content": {
							"application/json": {
								"schema": {
									"type": "object",
									"properties": {
										"message": {
											"type": "string",
											"description": "Error message"
										},
										"code": {
											"type": "integer",
											"description": "HTTP response code"
										}
									}
								}
							}
						}
					}
				}
			}
		},
		"/fetch": {
			"get": {
				"summary": "Fetch All Resources",
				"operationId": "fetchAll",
				"tags": [
					"fetch"
				],
				"responses": {
					"200": {
						"description": "Created",
						"content": {
							"application/json": {
								"schema": {
									"type": "array",
									"items": {
										"type": "object",
										"properties": {
											"uuid": {
												"type": "string"
											},
											"komoditas": {
												"type": "string"
											},
											"area_provinsi": {
												"type": "string"
											},
											"area_kota": {
												"type": "string"
											},
											"size": {
												"type": "string"
											},
											"price": {
												"type": "string"
											},
											"price_usd": {
												"type": "float"
											},
											"tgl_parsed": {
												"type": "string"
											},
											"timestamp": {
												"type": "string"
											}
										}
									}
								}
							}
						}
					},
					"401": {
						"description": "Unauthorized",
						"content": {
							"application/json": {
								"schema": {
									"type": "object",
									"properties": {
										"message": {
											"type": "string",
											"description": "Error message"
										},
										"code": {
											"type": "integer",
											"description": "HTTP response code"
										}
									}
								}
							}
						}
					},
					"500": {
						"description": "Internal Server Error",
						"content": {
							"application/json": {
								"schema": {
									"type": "object",
									"properties": {
										"message": {
											"type": "string",
											"description": "Error message"
										},
										"code": {
											"type": "integer",
											"description": "HTTP response code"
										}
									}
								}
							}
						}
					}
				}
			}
		}
	},
	"components": {
		"schemas": {
			"Auth": {
				"type": "object",
				"properties": {
					"name": {
						"type": "string",
						"required": true
					},
					"phone": {
						"type": "string",
						"required": true
					},
					"role": {
						"type": "string",
						"required": true
					},
					"password": {
						"type": "string",
						"required": false
					},
					"timestamp": {
						"type": "integer",
						"required": false
					}
				}
			}
		}
	}
}
```
