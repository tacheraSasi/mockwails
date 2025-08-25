import Sidebar from "@/components/sidebar";
import type React from "react";

interface LayoutProps {
  children: React.ReactNode;
}

const Layout: React.FC<LayoutProps> = ({ children }) => {
  return (
    <div className="min-h-screen bg-neutral-50 dark:bg-neutral-950">
      <div className="flex">
        <Sidebar />
        <main className="flex-1 ml-64 min-h-screen h-full overflow-y-auto">
          {children}
        </main>
      </div>
    </div>
  );
};

export default Layout;
