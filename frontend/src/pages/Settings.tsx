import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { Switch } from "@/components/ui/switch";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { Info, Save, RotateCcw } from "lucide-react";
import React, { useState } from "react";
import { useSettings } from "@/contexts/SettingsContext";
import { toast } from "sonner";

const Settings: React.FC = () => {
  const { settings, loading, updateSettings, refreshSettings } = useSettings();
  const [formData, setFormData] = useState({
    allowDedicatedPorts: settings?.allowDedicatedPorts || false,
    defaultUnifiedPort: settings?.defaultUnifiedPort || 8080,
  });
  const [saving, setSaving] = useState(false);

  React.useEffect(() => {
    if (settings) {
      setFormData({
        allowDedicatedPorts: settings.allowDedicatedPorts,
        defaultUnifiedPort: settings.defaultUnifiedPort,
      });
    }
  }, [settings]);

  const handleSave = async () => {
    setSaving(true);
    try {
      // Validate port range
      if (formData.defaultUnifiedPort < 1 || formData.defaultUnifiedPort > 65535) {
        toast.error("Port must be between 1 and 65535");
        setSaving(false);
        return;
      }

      const success = await updateSettings(formData);
      if (success) {
        toast.success("Settings updated successfully");
      } else {
        toast.error("Failed to update settings");
      }
    } catch (error) {
      console.error("Error updating settings:", error);
      toast.error("Failed to update settings");
    } finally {
      setSaving(false);
    }
  };

  const handleReset = async () => {
    try {
      await refreshSettings();
      toast.success("Settings reset to current values");
    } catch (error) {
      console.error("Error refreshing settings:", error);
      toast.error("Failed to refresh settings");
    }
  };

  const hasChanges = 
    settings &&
    (formData.allowDedicatedPorts !== settings.allowDedicatedPorts ||
     formData.defaultUnifiedPort !== settings.defaultUnifiedPort);

  if (loading) {
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="max-w-2xl mx-auto">
          <div className="animate-pulse">
            <div className="h-8 bg-gray-300 rounded mb-4"></div>
            <div className="h-64 bg-gray-300 rounded"></div>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="max-w-2xl mx-auto">
        <h1 className="text-3xl font-bold mb-6">Settings</h1>
        
        <Card className="p-6">
          <div className="space-y-6">
            <div>
              <h2 className="text-xl font-semibold mb-4">Port Management</h2>
              <Alert className="mb-4">
                <Info className="h-4 w-4" />
                <AlertDescription>
                  Changes to these settings will only affect new mock endpoints. 
                  Existing endpoints will keep their current port assignments.
                </AlertDescription>
              </Alert>
            </div>

            <div className="space-y-4">
              <div className="flex items-center space-x-2">
                <Switch
                  id="allowDedicatedPorts"
                  checked={formData.allowDedicatedPorts}
                  onCheckedChange={(checked) =>
                    setFormData(prev => ({ ...prev, allowDedicatedPorts: checked }))
                  }
                />
                <Label htmlFor="allowDedicatedPorts" className="text-sm font-medium">
                  Allow dedicated ports for each endpoint
                </Label>
              </div>
              <p className="text-sm text-muted-foreground ml-6">
                When enabled, each mock endpoint can have its own port. 
                When disabled, all endpoints share the unified port below.
              </p>
            </div>

            <div className="space-y-2">
              <Label htmlFor="defaultUnifiedPort" className="text-sm font-medium">
                Default Unified Port
              </Label>
              <Input
                id="defaultUnifiedPort"
                type="number"
                min="1"
                max="65535"
                value={formData.defaultUnifiedPort}
                onChange={(e) =>
                  setFormData(prev => ({ 
                    ...prev, 
                    defaultUnifiedPort: parseInt(e.target.value) || 8080 
                  }))
                }
                disabled={formData.allowDedicatedPorts}
                className="max-w-xs"
              />
              <p className="text-sm text-muted-foreground">
                {formData.allowDedicatedPorts 
                  ? "This port will be used as default when creating new endpoints"
                  : "All mock endpoints will share this port when unified mode is active"
                }
              </p>
            </div>

            <div className="flex items-center space-x-2 pt-4">
              <Button 
                onClick={handleSave} 
                disabled={!hasChanges || saving}
                className="flex items-center gap-2"
              >
                <Save className="h-4 w-4" />
                {saving ? "Saving..." : "Save Settings"}
              </Button>
              
              <Button 
                variant="outline" 
                onClick={handleReset}
                disabled={!hasChanges}
                className="flex items-center gap-2"
              >
                <RotateCcw className="h-4 w-4" />
                Reset
              </Button>
            </div>
          </div>
        </Card>

        <Card className="p-6 mt-6">
          <h3 className="text-lg font-semibold mb-4">How It Works</h3>
          <div className="space-y-4 text-sm text-muted-foreground">
            <div>
              <h4 className="font-medium text-foreground">Unified Mode (Recommended)</h4>
              <p>All mock endpoints share a single port with different paths. This is more efficient and easier to manage.</p>
              <p className="mt-1 font-mono text-xs">
                Example: http://localhost:8080/api/users, http://localhost:8080/api/products
              </p>
            </div>
            <div>
              <h4 className="font-medium text-foreground">Dedicated Mode</h4>
              <p>Each mock endpoint gets its own unique port. Useful when you need to simulate different services.</p>
              <p className="mt-1 font-mono text-xs">
                Example: http://localhost:8080/api/users, http://localhost:8081/api/products
              </p>
            </div>
          </div>
        </Card>
      </div>
    </div>
  );
};

export default Settings;