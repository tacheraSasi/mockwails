import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Edit, Play, Search, Square, Trash2 } from "lucide-react";
import type React from "react";
import { useEffect, useState } from "react";
import { GetAllServers } from "../../wailsjs/go/main/App";
import { db } from "../../wailsjs/go/models";
import { getMethodColor } from "@/lib/utils";


const ListMock: React.FC = () => {
  const [searchTerm, setSearchTerm] = useState("");
  const [endpoints, setEndpoints] = useState<db.Server[]>([]);

  useEffect(() => {
    const fetchData = async () => {
      const allServers = await GetAllServers();
      setEndpoints(allServers);
      console.log("ALL SERVERS", allServers);
    };

    fetchData();
  }, []);

  const filteredEndpoints = endpoints.filter(
    (endpoint) =>
      endpoint.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
      endpoint.endpoint.toLowerCase().includes(searchTerm.toLowerCase()) ||
      endpoint.description.toLowerCase().includes(searchTerm.toLowerCase()),
  );

  const toggleStatus = (id: number) => {
    // setEndpoints(
    //   endpoints.map((endpoint) =>
    //     endpoint.id === id
    //       ? { ...endpoint, status: endpoint.status === "active" ? "inactive" : "active" }
    //       : endpoint,
    //   ),
    // );
  };

  const deleteMock = (id: number) => {
    setEndpoints(endpoints.filter((endpoint) => endpoint.id !== id));
  };



  return (
    <div className="w-full max-w-6xl mx-auto p-6">
      <div className="mb-6">
        <h1 className="text-2xl font-bold text-foreground mb-2">Mock Endpoints</h1>
        <p className="text-muted-foreground">
          Manage your mock API endpoints. You can enable, disable, edit, or delete existing mocks.
        </p>
      </div>

      {/* Search and Filters */}
      <Card className="p-4 mb-6">
        <div className="flex flex-col sm:flex-row gap-4 items-center">
          <div className="relative flex-1">
            <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 text-muted-foreground h-4 w-4" />
            <Input
              placeholder="Search mocks by name, endpoint, or description..."
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
              className="pl-10"
            />
          </div>
          <div className="flex gap-2">
            <Button variant="outline" size="sm">
              All ({endpoints.length})
            </Button>
            {/* <Button variant="outline" size="sm">
              Active ({endpoints.filter((e) => e.status === "active").length})
            </Button>
            <Button variant="outline" size="sm">
              Inactive ({endpoints.filter((e) => e.status === "inactive").length})
            </Button> */}
          </div>
        </div>
      </Card>

      {/* Mock Endpoints List */}
      <div className="space-y-4">
        {filteredEndpoints.length === 0 ? (
          <Card className="p-8 text-center">
            <div className="text-muted-foreground">
              {searchTerm
                ? "No mocks found matching your search."
                : "No mock endpoints created yet."}
            </div>
            {!searchTerm && (
              <Button className="mt-4" variant="outline">
                Create Your First Mock
              </Button>
            )}
          </Card>
        ) : (
          filteredEndpoints.map((endpoint) => (
            <Card key={endpoint.id} className="p-6 hover:shadow-lg transition-shadow">
              <div className="flex items-start justify-between">
                <div className="flex-1">
                  <div className="flex items-center gap-3 mb-2">
                    <h3 className="text-lg font-semibold text-foreground">{endpoint.name}</h3>
                    <span
                      className={`px-2 py-1 rounded-full text-xs font-medium ${getMethodColor(endpoint.method)}`}
                    >
                      {endpoint.method}
                    </span>
                    <span
                      className={`px-2 py-1 rounded-full text-xs font-medium ${
                        endpoint.status === "active"
                          ? "bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-200"
                          : "bg-gray-100 text-gray-800 dark:bg-gray-900 dark:text-gray-200"
                      }`}
                    >
                      {endpoint.status}
                    </span>
                  </div>

                  <p className="text-muted-foreground mb-3">{endpoint.description}</p>

                  <div className="flex items-center gap-6 text-sm text-muted-foreground">
                    <div>
                      <span className="font-medium">Endpoint:</span>
                      <code className="ml-1 px-2 py-1 bg-muted rounded font-mono">
                        {endpoint.endpoint}
                      </code>
                    </div>
                    <div>
                      <span className="font-medium">Status Code:</span> {endpoint.responseStatus}
                    </div>
                    <div>
                      <span className="font-medium">Created:</span> {endpoint.createdAt}
                    </div>
                  </div>
                </div>

                <div className="flex items-center gap-2 ml-4">
                  <Button
                    variant="outline"
                    size="sm"
                    onClick={() => toggleStatus(endpoint.id)}
                    className="flex items-center gap-1"
                  >
                    {endpoint.status === "active" ? (
                      <>
                        <Square className="h-3 w-3" />
                        Stop
                      </>
                    ) : (
                      <>
                        <Play className="h-3 w-3" />
                        Start
                      </>
                    )}
                  </Button>

                  <Button variant="outline" size="sm" className="flex items-center gap-1">
                    <Edit className="h-3 w-3" />
                    Edit
                  </Button>

                  <Button
                    variant="outline"
                    size="sm"
                    onClick={() => deleteMock(endpoint.id)}
                    className="flex items-center gap-1 text-destructive hover:text-destructive"
                  >
                    <Trash2 className="h-3 w-3" />
                    Delete
                  </Button>
                </div>
              </div>
            </Card>
          ))
        )}
      </div>
    </div>
  );
};

export default ListMock;
