import { useNavigation } from "@/contexts/NavigationContext";
import CreateMock from "@/pages/CreateMock";
import ListMock from "@/pages/ListMock";
import RequestInspector from "@/pages/RequestInspector";
import type React from "react";

const PageRouter: React.FC = () => {
  const { currentPage } = useNavigation();

  switch (currentPage) {
    case "create-mock":
      return <CreateMock />;
    case "list-mock":
      return <ListMock />;
    case "request-inspector":
      return <RequestInspector />;
    case "settings":
      return (
        <div className="w-full max-w-4xl mx-auto p-6">
          <div className="text-center">
            <h1 className="text-2xl font-bold text-foreground mb-2">Settings</h1>
            <p className="text-muted-foreground">Settings page coming soon...</p>
          </div>
        </div>
      );
    default:
      return <CreateMock />;
  }
};

export default PageRouter;
