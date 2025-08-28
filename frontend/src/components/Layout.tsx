import Sidebar from "@/components/sidebar";
import { cn } from "@/lib/utils";
import type React from "react";
import { useState } from "react";

interface LayoutProps {
  children: React.ReactNode;
}

const Layout: React.FC<LayoutProps> = ({ children }) => {
  const [isSidebarCollapsed, setIsSidebarCollapsed] = useState(false);

  const toggleSidebar = () => {
    setIsSidebarCollapsed(!isSidebarCollapsed);
  };

  return (
    <div className="min-h-screen bg-neutral-50 dark:bg-neutral-950">
      <div className="flex">
        <Sidebar
          isCollapsed={isSidebarCollapsed}
          toggleSidebar={toggleSidebar}
        />
        <main
          className={cn(
            "flex-1 min-h-screen h-full overflow-y-auto transition-all duration-300 ease-in-out",
            isSidebarCollapsed ? "ml-20" : "ml-64"
          )}
        >
          {children}
        </main>
      </div>
    </div>
  );
};

export default Layout;
