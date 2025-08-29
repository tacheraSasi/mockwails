import { useState } from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Badge } from "@/components/ui/badge";
import { cn } from "@/lib/utils";
import ReloadIcon from "@/components/reload-icon";

const mockRequests = [
  {
    id: 1,
    method: "GET",
    path: "/api/users",
    status: 200,
    request: {
      headers: { Authorization: "Bearer ..." },
      body: "",
    },
    response: {
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify([{ id: 1, name: "John Doe" }], null, 2),
    },
  },
  {
    id: 2,
    method: "POST",
    path: "/api/products",
    status: 201,
    request: {
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ name: "New Product" }, null, 2),
    },
    response: {
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ id: 100, name: "New Product" }, null, 2),
    },
  },
];

const statusColor = (status: number) => {
  if (status >= 200 && status < 300) return "bg-green-500/20 text-green-700";
  if (status >= 400 && status < 500) return "bg-yellow-500/20 text-yellow-700";
  if (status >= 500) return "bg-red-500/20 text-red-700";
  return "bg-gray-500/20 text-gray-700";
};

const RequestInspector = () => {
  const [selectedRequest, setSelectedRequest] = useState(mockRequests[0]);

  return (
    <div className="flex h-[calc(100vh-4rem)] p-4 gap-4">
      {/* Left Pane: Request List */}
      <div className="w-1/3 border rounded-lg p-2 overflow-y-auto bg-card">
        <div className="flex p-2 justify-between">
          <span className="flex text-lg font-semibold">Requests</span>
          <span className="flex"><ReloadIcon /></span>
        </div>

        <ul className="space-y-1">
          {mockRequests.map((req) => (
            <li key={req.id}>
              <button
                onClick={() => setSelectedRequest(req)}
                className={cn(
                  "w-full text-left p-3 rounded-md transition-colors flex justify-between items-center",
                  selectedRequest.id === req.id
                    ? "bg-accent text-accent-foreground"
                    : "hover:bg-muted"
                )}
              >
                <div>
                  <span className="font-bold mr-2">{req.method}</span>
                  <span className="text-sm text-muted-foreground">
                    {req.path}
                  </span>
                </div>
                <Badge
                  variant="secondary"
                  className={cn("px-2 py-1 text-xs", statusColor(req.status))}
                >
                  {req.status}
                </Badge>
              </button>
            </li>
          ))}
        </ul>
      </div>

      {/* Right Pane: Request Details */}
      <div className="w-2/3 overflow-y-auto">
        {selectedRequest && (
          <Card className="h-full flex flex-col">
            <CardHeader>
              <CardTitle className="flex items-center justify-between">
                <span>Request Details</span>
                <Badge
                  variant="secondary"
                  className={statusColor(selectedRequest.status)}
                >
                  {selectedRequest.status}
                </Badge>
              </CardTitle>
            </CardHeader>
            <CardContent className="flex-1">
              <Tabs defaultValue="request" className="h-full flex flex-col">
                <TabsList className="mb-2">
                  <TabsTrigger value="request">Request</TabsTrigger>
                  <TabsTrigger value="response">Response</TabsTrigger>
                </TabsList>

                <TabsContent
                  value="request"
                  className="flex-1 overflow-y-auto space-y-2"
                >
                  <h3 className="font-semibold text-sm">Headers</h3>
                  <pre className="bg-neutral-100 dark:bg-neutral-800 p-3 rounded-md text-xs font-mono overflow-x-auto">
                    {JSON.stringify(selectedRequest.request.headers, null, 2)}
                  </pre>
                  <h3 className="font-semibold text-sm">Body</h3>
                  <pre className="bg-neutral-100 dark:bg-neutral-800 p-3 rounded-md text-xs font-mono overflow-x-auto">
                    {selectedRequest.request.body || "<empty>"}
                  </pre>
                </TabsContent>

                <TabsContent
                  value="response"
                  className="flex-1 overflow-y-auto space-y-2"
                >
                  <h3 className="font-semibold text-sm">Headers</h3>
                  <pre className="bg-neutral-100 dark:bg-neutral-800 p-3 rounded-md text-xs font-mono overflow-x-auto">
                    {JSON.stringify(selectedRequest.response.headers, null, 2)}
                  </pre>
                  <h3 className="font-semibold text-sm">Body</h3>
                  <pre className="bg-neutral-100 dark:bg-neutral-800 p-3 rounded-md text-xs font-mono overflow-x-auto">
                    {selectedRequest.response.body}
                  </pre>
                </TabsContent>
              </Tabs>
            </CardContent>
          </Card>
        )}
      </div>
    </div>
  );
};

export default RequestInspector;
