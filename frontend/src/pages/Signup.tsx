import axios from "axios";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

export const Signup = () => {
  const navigate = useNavigate();
  const [name, setName] = useState<string>("");
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");

  useEffect(() => {
    if (localStorage.getItem("ACCESS_KEY")) {
      navigate("/");
    }
  }, [navigate]);
  return (
    <div>
      <h1>SignUp</h1>
      <div>
        <input
          type="text"
          placeholder="name"
          value={name}
          onChange={(e) => setName(e.target.value)}
        ></input>
        <input
          type="email"
          placeholder="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        ></input>
        <input
          type="password"
          placeholder="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        ></input>
        <button
          onClick={async () => {
            if (name && email && password) {
              const newUser = {
                name: name,
                email: email,
                password: password,
              };
              try {
                const res = await axios.post(
                  "http://localhost:8080/signup",
                  newUser
                );
                alert("作成しました");
                navigate("/signin");
              } catch (error) {
                console.error("error", error);
              }
            }
          }}
        >
          アカウント作成
        </button>
        <a href="/signin" onClick={() => navigate("/signin")}>
          SignIn
        </a>
      </div>
    </div>
  );
};
