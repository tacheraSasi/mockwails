import Sidebar from "@/components/sidebar";
import type React from "react";

interface LayoutProps {
  children: React.ReactNode;
}

const Layout: React.FC<LayoutProps> = ({ children }) => {
  return (
    <div className="flex min-h-screen bg-neutral-50 dark:bg-neutral-950">
      <Sidebar />
      <main className="flex-1 overflow-y-auto">{children}</main>
    </div>
  );
};

export default Layout;
