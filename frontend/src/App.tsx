import { ThemeProvider } from "@/components/theme-provider";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Card } from "@/components/ui/card";
import { useState } from "react";
import { Toaster } from "@/components/ui/sonner";
import { toast } from "sonner";
import type * as React from "react";
import { Greet } from "../wailsjs/go/main/App.js";
import Layout from "@/components/Layout";

function App() {
  const [resultText, setResultText] = useState("Please enter your name below");
  const [name, setName] = useState("");
  const updateName = (e: React.ChangeEvent<HTMLInputElement>) =>
    setName(e.target.value);
  const updateResultText = (result: string) => setResultText(result);

  function greet() {
    Greet(name)
      .then(updateResultText)
      .then(() => toast("Greeted!"));
  }

  return (
    <Layout>
      <div id="App" className="w-full max-w-xl">
        <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
          <Toaster />
          <Card className="p-8 shadow-xl border border-black dark:border-black/30 bg-white/80 dark:bg-neutral-900/80">
            <div className="flex flex-col items-center mb-8">
              <div className="text-xl font-semibold text-center mb-2 text-black">
                Create a Mock Endpoint
              </div>
              <div className="text-sm text-neutral-500 dark:text-neutral-400 mb-4 text-center max-w-xs">
                Enter a name to get started. More options coming soon!
              </div>
            </div>
            <div className="flex flex-col gap-4 items-center">
              <Input
                id="name"
                onChange={updateName}
                autoComplete="off"
                name="input"
                type="text"
                placeholder="Enter mock name"
                className="w-full max-w-md"
              />
              <Button
                variant="outline"
                onClick={greet}
                className="w-full max-w-md border-black text-black hover:bg-black/30"
              >
                Create Mock
              </Button>
            </div>
          </Card>
        </ThemeProvider>
      </div>
    </Layout>
  );
}

export default App;
