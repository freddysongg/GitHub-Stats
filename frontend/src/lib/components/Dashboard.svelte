<script lang="ts">
  import { githubStats } from '../stores/github';
  import { Line } from 'svelte-chartjs';
  import { format } from 'date-fns';

  $: profile = $githubStats.profile;
  $: repositories = $githubStats.repositories;
</script>

<div class="p-6">
  {#if $githubStats.loading}
    <div class="flex justify-center items-center h-64">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-gray-900"></div>
    </div>
  {:else if profile}
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <div class="bg-white dark:bg-gray-800 rounded-lg shadow p-6">
        <h3 class="text-lg font-semibold mb-2">Repositories</h3>
        <p class="text-3xl font-bold">{profile.public_repos}</p>
      </div>
      
      <div class="bg-white dark:bg-gray-800 rounded-lg shadow p-6">
        <h3 class="text-lg font-semibold mb-2">Followers</h3>
        <p class="text-3xl font-bold">{profile.followers}</p>
      </div>
      
      <div class="bg-white dark:bg-gray-800 rounded-lg shadow p-6">
        <h3 class="text-lg font-semibold mb-2">Following</h3>
        <p class="text-3xl font-bold">{profile.following}</p>
      </div>
      
      <div class="bg-white dark:bg-gray-800 rounded-lg shadow p-6">
        <h3 class="text-lg font-semibold mb-2">Gists</h3>
        <p class="text-3xl font-bold">{profile.public_gists}</p>
      </div>
    </div>

    <div class="mt-8">
      <h2 class="text-2xl font-bold mb-4">Recent Repositories</h2>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        {#each repositories.slice(0, 6) as repo}
          <div class="bg-white dark:bg-gray-800 rounded-lg shadow p-6">
            <h3 class="text-xl font-semibold mb-2">{repo.name}</h3>
            <p class="text-gray-600 dark:text-gray-400 mb-4">{repo.description || 'No description'}</p>
            <div class="flex justify-between text-sm">
              <span>⭐ {repo.stargazers_count}</span>
              <span>🔄 {repo.forks_count}</span>
              <span>👁️ {repo.watchers_count}</span>
            </div>
          </div>
        {/each}
      </div>
    </div>
  {:else if $githubStats.error}
    <div class="text-red-500">
      Error: {$githubStats.error}
    </div>
  {/if}
</div>
