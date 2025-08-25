import Layout from "@/components/Layout";
import PageRouter from "@/components/PageRouter";
import { ThemeProvider } from "@/components/theme-provider";
import { Toaster } from "@/components/ui/sonner";
import { NavigationProvider } from "@/contexts/NavigationContext";

function App() {
  return (
    <ThemeProvider defaultTheme="system" storageKey="vite-ui-theme">
      <NavigationProvider>
        <Layout>
          <Toaster />
          <PageRouter />
        </Layout>
      </NavigationProvider>
    </ThemeProvider>
  );
}

export default App;
