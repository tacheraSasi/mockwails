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
import type React from "react";
import { useForm } from "react-hook-form";
import { toast } from "sonner";
import { CreateServer, GetAllServers } from "../../wailsjs/go/main/App";

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
}

const CreateMock: React.FC = () => {
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
    },
  });

  const onSubmit = async (data: MockFormData) => {
    console.log("Mock data:", data);
    await CreateServer(data);
    toast("Mock created successfully!");
  };

  return (
    <div className="w-full max-w-4xl mx-auto p-6">
      <div className="mb-6">
        <h1 className="text-2xl font-bold text-foreground mb-2">Create Mock Endpoint</h1>
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
                    <FormDescription>A descriptive name for your mock</FormDescription>
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
              <h2 className="text-lg font-semibold mb-4">Response Configuration</h2>

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
                        onChange={(e) => field.onChange(Number.parseInt(e.target.value))}
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
                    <FormDescription>Headers to return with the response</FormDescription>
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
                    <FormDescription>The response data to return</FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </Card>
          </div>

          {/* Request Configuration */}
          <Card className="p-6">
            <h2 className="text-lg font-semibold mb-4">Request Configuration (Optional)</h2>
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
                    <FormDescription>Headers that the request should contain</FormDescription>
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
                    <FormDescription>Expected structure of the request body</FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </div>
          </Card>

          {/* Submit Button */}
          <div className="flex justify-end space-x-4">
            <Button type="button" variant="outline" onClick={() => form.reset()}>
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
