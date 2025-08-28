import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { Badge } from "@/components/ui/badge";
import { Info } from "lucide-react";
import React from "react";
import { useForm } from "react-hook-form";
import { toast } from "sonner";
import { CreateServer } from "../../wailsjs/go/main/App";
import { useSettings } from "@/contexts/SettingsContext";

interface MockFormData {
  name: string;
  description: string;
  endpoint: string;
  method: string;
  requestHeaders: string;
  requestBody: string;
  responseStatus: number;
  responseHeaders: string;
  responseBody: string;
  addressAssigned?: { port: number };
  latency: number;
  requestQuery: string;
  status: string;
}

const CreateMock: React.FC = () => {
  const { settings, loading: settingsLoading } = useSettings();
  const form = useForm<MockFormData>({
    defaultValues: {
      name: "",
      description: "",
      endpoint: "/api/example",
      method: "GET",
      requestHeaders: "{}",
      requestBody: "{}",
      responseStatus: 200,
      responseHeaders: '{"Content-Type": "application/json"}',
      responseBody: '{"message": "Hello World"}',
      addressAssigned: { port: settings?.defaultUnifiedPort || 8080 },
      latency: 0,
      requestQuery: "{}",
      status: "active",
    },
  });

  // Update port when settings change
  React.useEffect(() => {
    if (settings && !settings.allowDedicatedPorts) {
      form.setValue("addressAssigned.port", settings.defaultUnifiedPort);
    }
  }, [settings, form]);

  const onSubmit = async (data: MockFormData) => {
    try {
      console.log("Mock data:", data);
      const response = await CreateServer(data);
      console.log("Create server response:", response);

      if (response.success) {
        toast("Mock created successfully!");
        form.reset(); // Reset form on success
      } else {
        toast(`Failed to create mock: ${response.message}`);
      }
    } catch (error) {
      console.error("Error creating server:", error);
      toast("An unexpected error occurred while creating the mock");
    }
  };

  return (
    <div className="w-full max-w-4xl mx-auto p-6">
      <div className="mb-6">
        <h1 className="text-2xl font-bold text-foreground mb-2">
          Create Mock Endpoint
        </h1>
        <p className="text-muted-foreground">
          Define a mock API endpoint with request and response specifications.
        </p>
      </div>

      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-6">
          <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
            {/* Basic Information */}
            <Card className="p-6">
              <h2 className="text-lg font-semibold mb-4">Basic Information</h2>

              <FormField
                control={form.control}
                name="name"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Mock Name</FormLabel>
                    <FormControl>
                      <Input placeholder="e.g., User API Mock" {...field} />
                    </FormControl>
                    <FormDescription>
                      A descriptive name for your mock
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="description"
                render={({ field }) => (
                  <FormItem className="mt-4">
                    <FormLabel>Description</FormLabel>
                    <FormControl>
                      <Textarea
                        placeholder="Describe what this mock is used for..."
                        className="min-h-[80px]"
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name="addressAssigned.port"
                render={({ field }) => {
                  const isUnifiedMode = !settings?.allowDedicatedPorts;
                  const unifiedPort = settings?.defaultUnifiedPort || 8080;
                  
                  return (
                    <FormItem>
                      <FormLabel>PORT</FormLabel>
                      <FormControl>
                        {isUnifiedMode ? (
                          <div className="flex items-center space-x-2">
                            <Input
                              type="number"
                              value={unifiedPort}
                              disabled
                              className="flex-1"
                            />
                            <Badge variant="secondary" className="bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-200">
                              <Info className="h-3 w-3 mr-1" />
                              Unified
                            </Badge>
                          </div>
                        ) : (
                          <Input
                            type="number"
                            placeholder="e.g., 8080"
                            value={field.value}
                            onChange={(e) =>
                              field.onChange(Number(e.target.value))
                            }
                          />
                        )}
                      </FormControl>
                      <FormDescription>
                        {isUnifiedMode 
                          ? "All mock endpoints share this unified port with different paths"
                          : "The dedicated port that this server will listen on"
                        }
                      </FormDescription>
                      <FormMessage />
                    </FormItem>
                  );
                }}
              />

              <div className="grid grid-cols-2 gap-4 mt-4">
                <FormField
                  control={form.control}
                  name="endpoint"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Endpoint</FormLabel>
                      <FormControl>
                        <Input placeholder="/api/users" {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />

                <FormField
                  control={form.control}
                  name="method"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>HTTP Method</FormLabel>
                      <FormControl>
                        <select
                          className="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                          {...field}
                        >
                          <option value="GET">GET</option>
                          <option value="POST">POST</option>
                          <option value="PUT">PUT</option>
                          <option value="DELETE">DELETE</option>
                          <option value="PATCH">PATCH</option>
                        </select>
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </div>
            </Card>

            {/* Response Configuration */}
            <Card className="p-6">
              <h2 className="text-lg font-semibold mb-4">
                Response Configuration
              </h2>

              <FormField
                control={form.control}
                name="responseStatus"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Response Status Code</FormLabel>
                    <FormControl>
                      <Input
                        type="number"
                        placeholder="200"
                        {...field}
                        onChange={(e) =>
                          field.onChange(Number.parseInt(e.target.value))
                        }
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="responseHeaders"
                render={({ field }) => (
                  <FormItem className="mt-4">
                    <FormLabel>Response Headers (JSON)</FormLabel>
                    <FormControl>
                      <Textarea
                        placeholder='{"Content-Type": "application/json"}'
                        className="min-h-[80px] font-mono text-sm"
                        {...field}
                      />
                    </FormControl>
                    <FormDescription>
                      Headers to return with the response
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="responseBody"
                render={({ field }) => (
                  <FormItem className="mt-4">
                    <FormLabel>Response Body (JSON)</FormLabel>
                    <FormControl>
                      <Textarea
                        placeholder='{"message": "Hello World", "status": "success"}'
                        className="min-h-[120px] font-mono text-sm"
                        {...field}
                      />
                    </FormControl>
                    <FormDescription>
                      The response data to return
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="latency"
                render={({ field }) => (
                  <FormItem className="mt-4">
                    <FormLabel>Latency (ms)</FormLabel>
                    <FormControl>
                      <Input
                        type="number"
                        placeholder="0"
                        {...field}
                        onChange={(e) =>
                          field.onChange(Number.parseInt(e.target.value))
                        }
                      />
                    </FormControl>
                    <FormDescription>
                      The simulated network latency in milliseconds
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="status"
                render={({ field }) => (
                  <FormItem className="mt-4">
                    <FormLabel>Status</FormLabel>
                    <FormControl>
                      <select
                        className="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                        {...field}
                      >
                        <option value="active">Active</option>
                        <option value="inactive">Inactive</option>
                      </select>
                    </FormControl>
                    <FormDescription>
                      The status of the mock endpoint
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </Card>
          </div>

          {/* Request Configuration */}
          <Card className="p-6">
            <h2 className="text-lg font-semibold mb-4">
              Request Configuration (Optional)
            </h2>
            <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
              <FormField
                control={form.control}
                name="requestHeaders"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Expected Request Headers (JSON)</FormLabel>
                    <FormControl>
                      <Textarea
                        placeholder='{"Authorization": "Bearer token"}'
                        className="min-h-[100px] font-mono text-sm"
                        {...field}
                      />
                    </FormControl>
                    <FormDescription>
                      Headers that the request should contain
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="requestBody"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Expected Request Body (JSON)</FormLabel>
                    <FormControl>
                      <Textarea
                        placeholder='{"name": "John Doe", "email": "john@example.com"}'
                        className="min-h-[100px] font-mono text-sm"
                        {...field}
                      />
                    </FormControl>
                    <FormDescription>
                      Expected structure of the request body
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="requestQuery"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Expected Request query (JSON)</FormLabel>
                    <FormControl>
                      <Textarea
                        placeholder='{"id": "123"}'
                        className="min-h-[100px] font-mono text-sm"
                        {...field}
                      />
                    </FormControl>
                    <FormDescription>
                      Expected query in the request
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </div>
          </Card>

          {/* Submit Button */}
          <div className="flex justify-end space-x-4">
            <Button
              type="button"
              variant="outline"
              onClick={() => form.reset()}
            >
              Reset Form
            </Button>
            <Button
              type="submit"
              className="bg-primary text-primary-foreground hover:bg-primary/90"
            >
              Create Mock Endpoint
            </Button>
          </div>
        </form>
      </Form>
    </div>
  );
};

export default CreateMock;
