import React from "react";
import ReactDOM from "react-dom/client";

const App = () => {
  return <div>application</div>;
};

const domContainer = document.querySelector("#app");
const root = ReactDOM.createRoot(domContainer!);
root.render(<App />);
