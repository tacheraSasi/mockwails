import { useNavigation } from "@/contexts/NavigationContext";
import CreateMock from "@/pages/CreateMock";
import ListMock from "@/pages/ListMock";
import RequestInspector from "@/pages/RequestInspector";
import Settings from "@/pages/Settings";
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
      return <Settings />;
    default:
      return <CreateMock />;
  }
};

export default PageRouter;
