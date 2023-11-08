import React from "react";
import ReactDOM from "react-dom/client";
import HelloPage from "./components/Hello";
import AddTaskPage from "./components/AddTask";

const App = () => {
  return (
    <div>
      <HelloPage />
      <AddTaskPage />
    </div>
  );
};

const domContainer = document.querySelector("#app");
const root = ReactDOM.createRoot(domContainer!);
root.render(<App />);
