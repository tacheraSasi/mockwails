import React from "react";
import { Home, Settings, User, Server, List } from "lucide-react";
import { ModeToggle } from "@/components/mode-toggle";
import Logo from "../assets/images/logo-universal.png";

const Sidebar: React.FC = () => {
  return (
    <div className="flex h-screen">
      {/* Sidebar */}
      <aside className="w-64  text-neutral-900 dark:text-neutral-100 shadow-xl flex flex-col border-r border-neutral-200 dark:border-neutral-800">
        {/* Logo & App Name */}
        <div className="flex items-center gap-3 p-6 border-b dark:border-neutral-800">
          <img
            src={Logo}
            alt="MockWAILS Logo"
            className="w-10 h-10 rounded shadow"
          />
          <span className="text-xl font-extrabold tracking-tight">
            mock<span className="text-black">WAILS</span>
          </span>
        </div>
        {/* Nav */}
        <nav className="flex-1 p-6 space-y-2">
          <a
            href="#"
            className="flex items-center gap-3 p-2 rounded-lg font-medium hover:bg-black dark:hover:bg-black/20 transition-colors"
          >
            <Server size={20} /> Create Mock
          </a>
          <a
            href="#"
            className="flex items-center gap-3 p-2 rounded-lg font-medium hover:bg-black dark:hover:bg-black/20 transition-colors"
          >
            <List size={20} /> Mocks List
          </a>
          <a
            href="#"
            className="flex items-center gap-3 p-2 rounded-lg font-medium hover:bg-black dark:hover:bg-black/20 transition-colors"
          >
            <Settings size={20} /> Settings
          </a>
        </nav>
        {/* Footer: Mode Toggle */}
        <div className="p-4 border-t dark:border-neutral-800 flex justify-end">
          <ModeToggle />
        </div>
      </aside>
    </div>
  );
};

export default Sidebar;
