from fastapi import FastAPI, Form, Request
from fastapi.responses import RedirectResponse
from fastapi.templating import Jinja2Templates

app = FastAPI()
templates = Jinja2Templates(directory="templates")

todos = []
next_id = 1


@app.get("/")
def index(request: Request):
    return templates.TemplateResponse(request, "index.html", {"todos": todos})


@app.post("/add")
def add(text: str = Form(...)):
    global next_id
    todos.append({"id": next_id, "text": text})
    next_id += 1
    return RedirectResponse("/", status_code=303)


@app.post("/delete/{id}")
def delete(id: int):
    global todos
    todos = [t for t in todos if t["id"] != id]
    return RedirectResponse("/", status_code=303)


if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)
