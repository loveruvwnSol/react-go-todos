import { BrowserRouter, Route, Routes, useNavigate } from "react-router-dom";
import { Signin } from "./pages/Signin";
import { Todo } from "./pages/Todo";
import { Signup } from "./pages/Signup";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/signin" element={<Signin />} />
        <Route path="/signup" element={<Signup />} />
        <Route path="/" element={<Todo />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
