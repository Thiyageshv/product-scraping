from flask import Flask
import sys 


class App():
	def __init__(self):
		app = Flask(__name__)
		@app.route("/")
