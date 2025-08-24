import { ModeToggle } from "@/components/mode-toggle";
import { ThemeProvider } from "@/components/theme-provider";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Card } from "@/components/ui/card";
import { AspectRatio } from "@/components/ui/aspect-ratio";
import Logo from "./assets/images/logo-universal.png";
import { useState } from "react";
import { Toaster } from "@/components/ui/sonner"
import { toast } from "sonner"
import type * as React from "react";
import { Greet } from "../wailsjs/go/main/App.js";

function App() {
  const [resultText, setResultText] = useState("Please enter your name below");
  const [name, setName] = useState("");
  const updateName = (e: React.ChangeEvent<HTMLInputElement>) => setName(e.target.value);
  const updateResultText = (result: string) => setResultText(result);

  function greet() {
    Greet(name).then(updateResultText).then(() => toast("Greeted!"));
  }

  return (
    <div id="App">
      <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
       
        <Toaster />
        <div className="flex align-center justify-center mt-2">
          <Card className="p-4">
            <div className="flex align-end justify-end">
              <ModeToggle />
            </div>
            <div className="flex align-center justify-center mb-12">
              <AspectRatio ratio={16 / 9}>
                <img src={Logo} alt="Logo" />
              </AspectRatio>
            </div>
            <div className="text-md font-bold text-center">
              {resultText}
            </div>
            <Input
              id="name"
              onChange={updateName}
              autoComplete="off"
              name="input"
              type="text"
              placeholder="Enter your name"
              className="w-[20rem]"
            />
            <Button variant="outline" onClick={greet}>
              Greet
            </Button>
          </Card>
        </div>
      </ThemeProvider>
    </div>
  );
}

export default App;
