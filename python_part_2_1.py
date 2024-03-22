from flask import Flask, request
from jinja2 import Template

# Некоректная передача имени приложения name, __name__
app = Flask(__name__)

# Нет обработки ошибок. 
@app.route("/page")
def page():
    name = request.values.get('name')
    age = request.values.get('age', 'unknown')
    output = Template('Hello ' + name + '! Your age is ' + age + '.').render()
return output

if __name__== "__main__":
    app.run(debug=True)
