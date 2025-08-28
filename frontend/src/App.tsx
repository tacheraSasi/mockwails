import Layout from "@/components/Layout";
import PageRouter from "@/components/PageRouter";
import { ThemeProvider } from "@/components/theme-provider";
import { Toaster } from "@/components/ui/sonner";
import { NavigationProvider } from "@/contexts/NavigationContext";
import { SettingsProvider } from "@/contexts/SettingsContext";

function App() {
  return (
    <ThemeProvider defaultTheme="system" storageKey="mockwails-theme">
      <SettingsProvider>
        <NavigationProvider>
          <Layout>
            <Toaster />
            <PageRouter />
          </Layout>
        </NavigationProvider>
      </SettingsProvider>
    </ThemeProvider>
  );
}

export default App;
