import { useEffect, useState } from "react";
import "./App.css";
import axios from "axios";

type todo = {
  id: number;
  title: string;
};

function App() {
  const [todo, setTodo] = useState<todo[]>([]);
  const [newTodo, setNewTodo] = useState("");

  useEffect(() => {
    getTodos();
  }, []);

  const getTodos = async () => {
    const todos = await axios.get("http://localhost:8080/todos");
    if (!todos) {
      console.log("not found");
    } else {
      setTodo(todos.data);
    }
  };

  return (
    <div>
      <input
        type="text"
        value={newTodo}
        onChange={(e) => setNewTodo(e.target.value)}
      />
      <button
        onClick={async () => {
          if (newTodo) {
            const postData = {
              title: newTodo,
            };

            try {
              const response = await axios.post(
                "http://localhost:8080/todos",
                postData
              );
            } catch (error) {
              console.error("Error posting data:", error);
            }
            setNewTodo("");
            getTodos();
          }
        }}
      >
        送信
      </button>
      {todo.map((e) => (
        <div key={e.id}>
          <p>{e.id}</p>
          <p>{e.title}</p>
          <button
            onClick={async () => {
              const updateData = {
                id: e.id,
                title: "update",
              };
              try {
                const response = await axios.put(
                  "http://localhost:8080/todos" + e.id,
                  updateData
                );
              } catch (error) {
                console.error("Error updating data:", error);
              }
              getTodos();
            }}
          >
            更新しちゃうよ〜〜
          </button>
          <button
            onClick={async () => {
              try {
                const response = await axios.delete(
                  "http://localhost:8080/todos" + e.id
                );
              } catch (error) {
                console.error("Error deleting data:", error);
              }
              getTodos();
            }}
          >
            削除ぉ！！！！
          </button>
        </div>
      ))}
    </div>
  );
}

export default App;
