from bottle import route, run

@route('/')
def hello():
    return "hello world"

run(port=8080, debug=True)
