<script lang="ts">
  import { theme } from '../stores/theme';
  import { onMount, createEventDispatcher } from 'svelte';

  export let loading = false;
  const dispatch = createEventDispatcher<{
    search: string;
  }>();

  onMount(() => {
    theme.initialize();
  });

  let username = '';

  function handleSubmit() {
    if (username && !loading) {
      dispatch('search', username);
    }
  }
</script>

<header class="card mb-6">
  <div class="flex items-center justify-between">
    <div class="flex items-center space-x-4">
      <h1 class="text-2xl font-bold">GitHub Stats Dashboard</h1>
      <button
        on:click={() => theme.toggleTheme()}
        class="p-2 rounded-full hover:bg-github-100 dark:hover:bg-github-700 transition-colors"
      >
        {#if $theme === 'light'}
          🌙
        {:else}
          ☀️
        {/if}
      </button>
    </div>

    <form on:submit|preventDefault={handleSubmit} class="flex space-x-2">
      <input
        type="text"
        bind:value={username}
        placeholder="Enter GitHub username"
        class="px-4 py-2 rounded-mac border border-github-200 dark:border-github-600 bg-white dark:bg-github-700 focus:outline-none focus:ring-2 focus:ring-github-500"
      />
      <button
        type="submit"
        disabled={loading}
        class="px-4 py-2 bg-github-900 dark:bg-github-100 text-white dark:text-github-900 rounded-mac hover:opacity-90 transition-opacity disabled:opacity-50"
      >
        {loading ? 'Loading...' : 'Search'}
      </button>
    </form>
  </div>
</header>
