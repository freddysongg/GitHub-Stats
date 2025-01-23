<script lang="ts">
  import { onMount } from 'svelte';
  import { githubStats, fetchGitHubStats } from './lib/stores/github';
  import Dashboard from './lib/components/Dashboard.svelte';
  import MenuBar from './lib/components/MenuBar.svelte';
  import Settings from './lib/components/Settings.svelte';
  import { electronAPI } from './lib/electron';

  let view = 'dashboard';
  let token: string;

  onMount(async () => {
    token = await electronAPI.getGithubToken();
    if (token) {
      await fetchGitHubStats();
    }
  });
</script>

<main class="h-screen bg-gray-50 dark:bg-gray-900 text-gray-900 dark:text-gray-100">
  {#if !token}
    <Settings />
  {:else if view === 'dashboard'}
    <Dashboard />
  {:else}
    <MenuBar />
  {/if}
</main>

<style>
  :global(body) {
    @apply m-0 p-0;
  }
</style>
