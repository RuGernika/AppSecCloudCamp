from flask import Flask, request
import subprocess

app = Flask(__name__)

@app.route("/dns")
def dns_lookup():
    hostname = request.values.get('hostname')
    if not hostname:
        return"Bad Request", 400
    cmd = 'nslookup ' + hostname
    try:
        output = subprocess.check_output(cmd, shell=True, text=True)
return output
except subprocess.CalledProcessError as e:
    return "Error execute:"+
str(e), 500    
if __name__ == "__main__":
    app.run(debug=True)
