import React, { createContext, useContext, useState, useEffect, ReactNode } from 'react';
import { GetSettings, UpdateSettings } from '../../wailsjs/go/main/App';

interface Settings {
  id?: number;
  allowDedicatedPorts: boolean;
  defaultUnifiedPort: number;
  createdAt?: string;
  updatedAt?: string;
}

interface SettingsContextType {
  settings: Settings | null;
  loading: boolean;
  updateSettings: (newSettings: Partial<Settings>) => Promise<boolean>;
  refreshSettings: () => Promise<void>;
}

const SettingsContext = createContext<SettingsContextType | undefined>(undefined);

export const useSettings = (): SettingsContextType => {
  const context = useContext(SettingsContext);
  if (!context) {
    throw new Error('useSettings must be used within a SettingsProvider');
  }
  return context;
};

interface SettingsProviderProps {
  children: ReactNode;
}

export const SettingsProvider: React.FC<SettingsProviderProps> = ({ children }) => {
  const [settings, setSettings] = useState<Settings | null>(null);
  const [loading, setLoading] = useState(true);

  const refreshSettings = async () => {
    try {
      setLoading(true);
      const response = await GetSettings();
      if (response.success && response.data) {
        setSettings(response.data);
      }
    } catch (error) {
      console.error('Failed to fetch settings:', error);
    } finally {
      setLoading(false);
    }
  };

  const updateSettings = async (newSettings: Partial<Settings>): Promise<boolean> => {
    try {
      const updatedSettings = { ...settings, ...newSettings };
      const response = await UpdateSettings(updatedSettings);
      if (response.success && response.data) {
        setSettings(response.data);
        return true;
      }
      return false;
    } catch (error) {
      console.error('Failed to update settings:', error);
      return false;
    }
  };

  useEffect(() => {
    refreshSettings();
  }, []);

  return (
    <SettingsContext.Provider
      value={{
        settings,
        loading,
        updateSettings,
        refreshSettings,
      }}
    >
      {children}
    </SettingsContext.Provider>
  );
};