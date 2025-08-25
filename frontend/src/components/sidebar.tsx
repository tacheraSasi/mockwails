import React from "react";
import { Home, Settings, User, Server, List } from "lucide-react";
import { ModeToggle } from "@/components/mode-toggle";
import Logo from "../assets/images/logo-universal.png";

const Sidebar: React.FC = () => {
  return (
    <div className="flex h-screen">
      {/* Sidebar */}
      <aside className="w-64 bg-gradient-to-b from-white via-gray-50 to-gray-200 dark:from-gray-900 dark:via-gray-950 dark:to-gray-900 text-gray-900 dark:text-gray-100 shadow-xl flex flex-col border-r border-gray-200 dark:border-gray-800">
        {/* Logo & App Name */}
        <div className="flex items-center gap-3 p-6 border-b dark:border-gray-800">
          <img
            src={Logo}
            alt="MockWAILS Logo"
            className="w-10 h-10 rounded shadow"
          />
          <span className="text-xl font-extrabold tracking-tight">
            mock<span className="text-blue-600 dark:text-blue-400">WAILS</span>
          </span>
        </div>
        {/* Nav */}
        <nav className="flex-1 p-6 space-y-2">
          <a
            href="#"
            className="flex items-center gap-3 p-2 rounded-lg font-medium hover:bg-blue-100 dark:hover:bg-blue-900 transition-colors"
          >
            <Server size={20} /> Create Mock
          </a>
          <a
            href="#"
            className="flex items-center gap-3 p-2 rounded-lg font-medium hover:bg-blue-100 dark:hover:bg-blue-900 transition-colors"
          >
            <List size={20} /> Mocks List
          </a>
          <a
            href="#"
            className="flex items-center gap-3 p-2 rounded-lg font-medium hover:bg-blue-100 dark:hover:bg-blue-900 transition-colors"
          >
            <Settings size={20} /> Settings
          </a>
        </nav>
        {/* Footer: Mode Toggle */}
        <div className="p-4 border-t dark:border-gray-800 flex justify-end">
          <ModeToggle />
        </div>
      </aside>
    </div>
  );
};

export default Sidebar;
