import { ModeToggle } from "@/components/mode-toggle";
import { useNavigation } from "@/contexts/NavigationContext";
import { Home, List, Server, Settings, User } from "lucide-react";
import type React from "react";
import Logo from "../assets/images/logo-universal.png";

const Sidebar: React.FC = () => {
  const { currentPage, setCurrentPage } = useNavigation();

  const navItems = [
    { id: "create-mock" as const, icon: Server, label: "Create Mock" },
    { id: "list-mock" as const, icon: List, label: "Mocks List" },
    { id: "settings" as const, icon: Settings, label: "Settings" },
  ];

  return (
    <div className="flex h-screen">
      {/* Sidebar */}
      <aside className="w-64 bg-background text-foreground shadow-xl flex flex-col border-r border-border">
        {/* Logo & App Name */}
        <div className="flex items-center gap-3 p-6 border-b border-border">
          <img src={Logo} alt="MockWAILS Logo" className="w-10 h-10 rounded shadow" />
          <span className="text-xl font-extrabold tracking-tight">
            mock<span className="text-primary">WAILS</span>
          </span>
        </div>
        {/* Nav */}
        <nav className="flex-1 p-6 space-y-2">
          {navItems.map(({ id, icon: Icon, label }) => (
            <button
              type="button"
              key={id}
              onClick={() => setCurrentPage(id)}
              className={`w-full flex items-center gap-3 p-3 rounded-lg font-medium transition-all duration-200 ${
                currentPage === id
                  ? "bg-primary text-primary-foreground shadow-md"
                  : "text-muted-foreground hover:bg-accent hover:text-accent-foreground"
              }`}
            >
              <Icon size={20} />
              {label}
            </button>
          ))}
        </nav>
        {/* Footer: Mode Toggle */}
        <div className="p-4 border-t border-border flex justify-end">
          <ModeToggle />
        </div>
      </aside>
    </div>
  );
};

export default Sidebar;
