# Todo-cli
A simple and lightweight command-line ToDo List application written in Go — add, list, complete, and delete your tasks with ease, all from your terminal.

---

## 📌 Features

- Add new tasks
- View task list
- Mark tasks as completed
- Delete tasks
- Local storage with JSON
- Easy to use, no external dependencies

---

## 🚀 Getting Started

Follow these steps to get `todo-cli` running on your local machine.

---

### 1. 📥 Clone the Repository

First, clone the project using Git, or download the ZIP manually.

```bash
git clone https://github.com/h1ntz0/Todo-cli.git
cd todo-cli
```

### 2. 🔧 Install Dependencies
Make sure Go is installed, then run:
```bash
go mod tidy
```

### 3. Install Application
Use this format to run this code from terminal:
```bash
go run main.go <command> [your arguments]
```
Examples:
```bash
go run main.go add "read a book"
go run main.go list
go run main.go done 1
go run main.go delete 1
```

And then done.
