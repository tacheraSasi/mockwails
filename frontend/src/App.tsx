import { ModeToggle } from "@/components/mode-toggle";
import { ThemeProvider } from "@/components/theme-provider";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Card } from "@/components/ui/card";
import { AspectRatio } from "@/components/ui/aspect-ratio";
import Logo from "./assets/images/logo-universal.png";
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
          <div className="mb-10 text-center">
            <h1 className="text-3xl font-extrabold tracking-tight mb-2 text-blue-700 dark:text-blue-400">
              Welcome to mockWAILS
            </h1>
            <p className="text-lg text-gray-600 dark:text-gray-300 mb-4">
              A simple, beautiful mock server creator for rapid API prototyping.
            </p>
          </div>
          <Card className="p-8 shadow-xl border border-blue-100 dark:border-blue-900 bg-white/80 dark:bg-gray-900/80">
            <div className="flex flex-col items-center mb-8">
              <AspectRatio ratio={16 / 9} className="w-32 mb-4">
                <img src={Logo} alt="Logo" className="rounded shadow" />
              </AspectRatio>
              <div className="text-xl font-semibold text-center mb-2">
                Create a Mock Endpoint
              </div>
              <div className="text-sm text-gray-500 dark:text-gray-400 mb-4 text-center max-w-xs">
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
                className="w-full max-w-md"
              >
                Create Mock
              </Button>
              <div className="text-md font-bold text-center mt-4 text-blue-700 dark:text-blue-400">
                {resultText}
              </div>
            </div>
          </Card>
        </ThemeProvider>
      </div>
    </Layout>
  );
}

export default App;
