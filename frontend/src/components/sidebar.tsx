import React, { useState, useEffect } from "react";
import { Sun, Moon, Home, Settings, User } from "lucide-react";

const Sidebar: React.FC = () => {
  const [darkMode, setDarkMode] = useState(false);

  useEffect(() => {
    if (darkMode) {
      document.documentElement.classList.add("dark");
    } else {
      document.documentElement.classList.remove("dark");
    }
  }, [darkMode]);

  return (
    <div className="flex h-screen">
      {/* Sidebar */}
      <aside className="w-64 bg-white dark:bg-gray-900 text-gray-900 dark:text-gray-100 shadow-lg flex flex-col">
        {/* Header */}
        <div className="flex items-center justify-between p-4 border-b dark:border-gray-700">
          <h1 className="text-lg font-bold">My App</h1>
          <button
            onClick={() => setDarkMode(!darkMode)}
            className="p-2 rounded-full hover:bg-gray-200 dark:hover:bg-gray-800"
          >
            {darkMode ? <Sun size={20} /> : <Moon size={20} />}
          </button>
        </div>

        {/* Nav */}
        <nav className="flex-1 p-4 space-y-2">
          <a
            href="#"
            className="flex items-center gap-3 p-2 rounded-lg hover:bg-gray-200 dark:hover:bg-gray-800"
          >
            <Home size={20} /> Home
          </a>
          <a
            href="#"
            className="flex items-center gap-3 p-2 rounded-lg hover:bg-gray-200 dark:hover:bg-gray-800"
          >
            <User size={20} /> Profile
          </a>
          <a
            href="#"
            className="flex items-center gap-3 p-2 rounded-lg hover:bg-gray-200 dark:hover:bg-gray-800"
          >
            <Settings size={20} /> Settings
          </a>
        </nav>
      </aside>
    </div>
  );
};

export default Sidebar;
