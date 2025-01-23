const { contextBridge, ipcRenderer } = require('electron');

contextBridge.exposeInMainWorld('electronAPI', {
  getGithubToken: () => ipcRenderer.invoke('get-github-token'),
  setGithubToken: (token) => ipcRenderer.invoke('set-github-token', token)
});