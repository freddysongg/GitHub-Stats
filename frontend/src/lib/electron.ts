interface ElectronAPI {
  getGithubToken: () => Promise<string>;
  setGithubToken: (token: string) => Promise<void>;
}

declare global {
  interface Window {
    electronAPI: ElectronAPI;
  }
}

// Mock implementation for development
const mockElectronAPI: ElectronAPI = {
  getGithubToken: async () => {
    return localStorage.getItem('githubToken') || '';
  },
  setGithubToken: async (token: string) => {
    localStorage.setItem('githubToken', token);
  }
};

// Use the real Electron API in production, mock in development
export const electronAPI = window.electronAPI || mockElectronAPI;