<script lang="ts">
  import { electronAPI } from '../electron';
  import { fetchGitHubStats } from '../stores/github';

  let token = '';
  let username = '';
  let saving = false;
  let error = '';

  async function saveSettings() {
    saving = true;
    error = '';

    try {
      await electronAPI.setGithubToken(token);
      await fetchGitHubStats();
      window.location.reload();
    } catch (e) {
      error = e.message;
    } finally {
      saving = false;
    }
  }
</script>

<div class="flex items-center justify-center min-h-screen bg-gray-50 dark:bg-gray-900">
  <div class="w-full max-w-md p-8 bg-white dark:bg-gray-800 rounded-lg shadow-lg">
    <h2 class="text-2xl font-bold mb-6">GitHub Settings</h2>
    
    <form on:submit|preventDefault={saveSettings}>
      <div class="mb-4">
        <label class="block text-sm font-medium mb-2" for="token">
          GitHub Personal Access Token
        </label>
        <input
          type="password"
          id="token"
          bind:value={token}
          class="w-full p-2 border rounded-md dark:bg-gray-700 dark:border-gray-600"
          placeholder="ghp_..."
          required
        />
      </div>

      <div class="mb-6">
        <label class="block text-sm font-medium mb-2" for="username">
          GitHub Username
        </label>
        <input
          type="text"
          id="username"
          bind:value={username}
          class="w-full p-2 border rounded-md dark:bg-gray-700 dark:border-gray-600"
          placeholder="your-username"
          required
        />
      </div>

      {#if error}
        <div class="mb-4 text-red-500 text-sm">
          {error}
        </div>
      {/if}

      <button
        type="submit"
        class="w-full bg-blue-500 text-white py-2 px-4 rounded-md hover:bg-blue-600 disabled:opacity-50"
        disabled={saving}
      >
        {saving ? 'Saving...' : 'Save Settings'}
      </button>
    </form>
  </div>
</div>
