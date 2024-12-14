import { BrowserRouter as Router, Routes, Route} from "react-router-dom"
import { ToastContainer } from "react-toastify"
import 'react-toastify/dist/ReactToastify.css';
import Home from "./pages/Home";
import Login from "./pages/Login";
import Gallery from "./pages/Gallery";
import Error from "./pages/Error";

function App() {

  return (
    <Router>
      <ToastContainer position="bottom-right" autoClose={3000} />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/login" element={<Login />} />
        <Route path="/gallery" element={<Gallery />} />
        <Route path="*" element={<Error />} />
      </Routes>
    </Router>
  )
}

export default App
