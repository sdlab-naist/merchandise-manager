import os
from flask import Flask, request, jsonify, send_from_directory
from json import dumps, loads
import mysql.connector
import json

with open('config_db.json') as json_file:
    data = json.load(json_file)
    print(data)

mydb = mysql.connector.connect(
  host=data['Host'],
  user=data['Username'],
  passwd=data['Password']
)


app = Flask(__name__, static_folder=".", template_folder=".")

@app.after_request
def add_headers(response):
    response.headers.add('Access-Control-Allow-Origin', '*')
    response.headers.add('Access-Control-Allow-Headers', 'Content-Type,Authorization, data')
    return response

@app.route("/")
def home():
    return 'Kundjanasith Thonglek . . .'

@app.route("/verify/<temp>/<username>", methods=['GET'])
def verify(temp,username):
    mycursor = mydb.cursor()
    mycursor.execute("SELECT * FROM pbl.Users WHERE TempPassword='"+temp+"'")
    myresult = mycursor.fetchall()
    if len(myresult) == 0:
        return "Authentication is invalid"
    else:
        for x in myresult:
            print(x[1])
            if x[1] == username:
                sql = "UPDATE pbl.Users SET Status = 'Active' WHERE TempPassword='"+temp+"' AND Username='"+username+"'"
                mycursor.execute(sql)
                mydb.commit()
                return "Authentication is verified"
    return "Authentication is invalid"

if __name__ == "__main__":
	app.run(host='0.0.0.0',port=13135)