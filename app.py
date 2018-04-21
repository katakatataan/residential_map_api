import json
import requests

from flask import Flask, jsonify, request
app = Flask(__name__)

@app.route("/pref", methods=['GET'])
def pref():
    r = requests.request('GET', 'https://opendata.resas-portal.go.jp/api/v1/prefectures', headers={'X-API-KEY': 'QEO0NqHC4Iqpln3y4kxzzORVq5myL66QDklWzrDL'})
    return jsonify(r.json())

@app.route("/city", methods=['GET'])
def city():
    r = requests.request('GET', 'https://opendata.resas-portal.go.jp/api/v1/cities', headers={'X-API-KEY': 'QEO0NqHC4Iqpln3y4kxzzORVq5myL66QDklWzrDL'})
    return jsonify(r.json())

if __name__ == "__main__":
    app.run(host='0.0.0.0',port=5000,debug=True)
