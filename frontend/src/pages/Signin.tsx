import axios from "axios";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

export const Signin = () => {
  const navigate = useNavigate();
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");

  useEffect(() => {
    if (localStorage.getItem("ACCESS_KEY")) {
      navigate("/");
    }
  }, [navigate]);
  return (
    <div>
      <h1>SignIn</h1>
      <div>
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
            const user = {
              email: email,
              password: password,
            };
            try {
              const res = await axios.post(
                "http://localhost:8080/signin",
                user
              );
              localStorage.setItem("ACCESS_KEY", res.data);
              navigate("/");
            } catch (error) {
              console.error("error", error);
            }
          }}
        >
          ログイン
        </button>
        <a href="/signup" onClick={() => navigate("/signup")}>SignUp</a>
      </div>
    </div>
  );
};
