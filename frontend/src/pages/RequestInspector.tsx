import { useState } from "react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { cn } from "@/lib/utils";

// Mock data for demonstration
const mockRequests = [
  {
    id: 1,
    method: "GET",
    path: "/api/users",
    status: 200,
    request: {
      headers: { "Authorization": "Bearer ..." },
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

const RequestInspector = () => {
  const [selectedRequest, setSelectedRequest] = useState(mockRequests[0]);

  return (
    <div className="flex h-[calc(100vh-4rem)] p-4 gap-4">
      {/* Left Pane: Request List */}
      <div className="w-1/3 border rounded-lg p-2 overflow-y-auto">
        <h2 className="text-lg font-semibold p-2">Requests</h2>
        <ul>
          {mockRequests.map((req) => (
            <li key={req.id}>
              <button
                onClick={() => setSelectedRequest(req)}
                className={cn(
                  "w-full text-left p-2 rounded-md",
                  selectedRequest.id === req.id && "bg-accent"
                )}
              >
                <div className="flex justify-between">
                  <span>
                    <span className="font-bold">{req.method}</span> {req.path}
                  </span>
                  <span
                    className={cn(
                      "font-bold",
                      req.status >= 400 ? "text-red-500" : "text-green-500"
                    )}
                  >
                    {req.status}
                  </span>
                </div>
              </button>
            </li>
          ))}
        </ul>
      </div>

      {/* Right Pane: Request Details */}
      <div className="w-2/3 overflow-y-auto">
        {selectedRequest && (
          <Card>
            <CardHeader>
              <CardTitle>Request Details</CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              <div>
                <h3 className="font-semibold">Request</h3>
                <pre className="bg-neutral-100 dark:bg-neutral-800 p-2 rounded-md mt-1 text-sm">
                  {`Headers: ${JSON.stringify(selectedRequest.request.headers, null, 2)}
Body: ${selectedRequest.request.body}`}
                </pre>
              </div>
              <div>
                <h3 className="font-semibold">Response</h3>
                <pre className="bg-neutral-100 dark:bg-neutral-800 p-2 rounded-md mt-1 text-sm">
                  {`Headers: ${JSON.stringify(selectedRequest.response.headers, null, 2)}
Body: ${selectedRequest.response.body}`}
                </pre>
              </div>
            </CardContent>
          </Card>
        )}
      </div>
    </div>
  );
};

export default RequestInspector;
