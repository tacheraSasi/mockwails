import React from "react";
import { createRoot } from "react-dom/client";
import "./global.css";
import App from "./App";

const container = document.getElementById("root");

if (!container) {
  throw new Error("Root element not found. Failed to render the application.");
}

const root = createRoot(container);

root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
);
