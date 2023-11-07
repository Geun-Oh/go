import React from "react";
import ReactDOM from "react-dom/client";
import HelloPage from "./components/Hello";

const App = () => {
  return (
    <div>
      <HelloPage />
    </div>
  );
};

const domContainer = document.querySelector("#app");
const root = ReactDOM.createRoot(domContainer!);
root.render(<App />);
