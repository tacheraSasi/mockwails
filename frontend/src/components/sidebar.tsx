import { ModeToggle } from "@/components/mode-toggle";
import { useNavigation } from "@/contexts/NavigationContext";
import { ChevronLeft, List, Search, Server, Settings } from "lucide-react";
import type React from "react";
import Logo from "../assets/images/logo-universal.png";
import { cn } from "@/lib/utils";

interface SidebarProps {
  isCollapsed: boolean;
  toggleSidebar: () => void;
}

const Sidebar: React.FC<SidebarProps> = ({
  isCollapsed,
  toggleSidebar,
}) => {
  const { currentPage, setCurrentPage } = useNavigation();

  const navItems = [
    { id: "create-mock" as const, icon: Server, label: "Create Mock" },
    { id: "list-mock" as const, icon: List, label: "Mocks List" },
    {
      id: "request-inspector" as const,
      icon: Search,
      label: "Request Inspector",
    },
    { id: "settings" as const, icon: Settings, label: "Settings" },
  ];

  return (
    <div className="flex h-screen fixed">
      {/* Sidebar */}
      <aside
        className={cn(
          "bg-background text-foreground shadow-xl flex flex-col border-r border-border transition-all duration-300 ease-in-out",
          isCollapsed ? "w-22" : "w-64"
        )}
      >
        {/* Logo & App Name */}
        <div className="flex items-center m-auto gap-3 p-3 border-b border-border h-[4rem]">
          <img
            src={Logo}
            alt="MockWAILS Logo"
            className="w-12 h-12 rounded shadow"
          />
          <span
            className={cn(
              "text-xl font-extrabold tracking-tight transition-opacity duration-300",
              isCollapsed && "opacity-0"
            )}
          >
            mock<span className="text-primary">WAILS</span>
          </span>
        </div>
        {/* Nav */}
        <nav className="flex-1 p-3 space-y-2">
          {navItems.map(({ id, icon: Icon, label }) => (
            <button
              type="button"
              key={id}
              onClick={() => setCurrentPage(id)}
              className={cn(
                "w-full flex items-center gap-3 p-3 rounded-lg font-medium transition-all duration-200",
                currentPage === id
                  ? "bg-primary text-primary-foreground shadow-md"
                  : "text-muted-foreground hover:bg-accent hover:text-accent-foreground",
                isCollapsed && "justify-center"
              )}
            >
              <Icon size={20} />
              <span
                className={cn(
                  "transition-opacity",
                  isCollapsed && "opacity-0 hidden"
                )}
              >
                {label}
              </span>
            </button>
          ))}
        </nav>

        {/* Footer: Collapse & Mode Toggle */}
        <div className="p-4 border-t border-border flex items-center justify-between">
          <button
            onClick={toggleSidebar}
            className={`p-2 ${cn(isCollapsed && "m-auto")} rounded-lg hover:bg-accent`}
          >
            <ChevronLeft
              size={20}
              className={cn(
                "transition-transform duration-300",
                isCollapsed && "rotate-180"
              )}
            />
          </button>
          <div className={cn(isCollapsed && "hidden")}>
            <ModeToggle />
          </div>
        </div>
      </aside>
    </div>
  );
};

export default Sidebar;
