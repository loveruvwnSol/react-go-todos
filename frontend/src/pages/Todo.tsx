import { useEffect, useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";

type todo = {
  id: number;
  title: string;
};

export const Todo = () => {
  const navigate = useNavigate();
  const [todo, setTodo] = useState<todo[]>([]);
  const [newTodo, setNewTodo] = useState("");
  const [user, setUser] = useState<any>([]);

  useEffect(() => {
    if (!localStorage.getItem("ACCESS_KEY")) {
      navigate("/signin");
    } else {
      getCurrentUser();
      getTodos();
    }
  }, [navigate]);

  const getTodos = async () => {
    const todos = await axios.get("http://localhost:8080/todos");
    if (!todos) {
      console.log("not found");
    } else {
      setTodo(todos.data);
    }
  };

  const getCurrentUser = async () => {
    const token = localStorage.getItem("ACCESS_KEY");
    axios
      .get("http://localhost:8080/user", {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })
      .then((res) => {
        setUser(res.data);
      })
      .catch((error) => {
        alert(error);
      });
  };

  return (
    <div>
      <p>{user.name}でログイン中</p>
      <input
        type="text"
        placeholder="今日のタスク"
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
              setTodo(response.data);
            } catch (error) {
              console.error("Error posting data:", error);
            }
            setNewTodo("");
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
                setTodo(response.data);
              } catch (error) {
                console.error("Error updating data:", error);
              }
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
                setTodo(response.data);
              } catch (error) {
                console.error("Error deleting data:", error);
              }
            }}
          >
            削除ぉ！！！！
          </button>
        </div>
      ))}
      <button
        onClick={() => {
          localStorage.removeItem("ACCESS_KEY");
          navigate("/signin");
        }}
      >
        ログアウト
      </button>
    </div>
  );
};
