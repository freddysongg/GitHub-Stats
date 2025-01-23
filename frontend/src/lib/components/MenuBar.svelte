<script lang="ts">
  import { githubStats } from '../stores/github';
  
  $: profile = $githubStats.profile;
  $: repositories = $githubStats.repositories;
</script>

<div class="w-80 bg-white dark:bg-gray-800 rounded-lg shadow-lg overflow-hidden">
  {#if $githubStats.loading}
    <div class="flex justify-center items-center h-32">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-gray-900"></div>
    </div>
  {:else if profile}
    <div class="p-4">
      <div class="flex items-center space-x-4">
        <img 
          src={profile.avatar_url} 
          alt="Profile" 
          class="w-12 h-12 rounded-full"
        />
        <div>
          <h3 class="font-semibold">{profile.name || profile.login}</h3>
          <p class="text-sm text-gray-600 dark:text-gray-400">{profile.bio || ''}</p>
        </div>
      </div>
      
      <div class="mt-4 grid grid-cols-2 gap-4">
        <div class="text-center">
          <p class="text-lg font-semibold">{profile.public_repos}</p>
          <p class="text-sm text-gray-600 dark:text-gray-400">Repositories</p>
        </div>
        <div class="text-center">
          <p class="text-lg font-semibold">{profile.followers}</p>
          <p class="text-sm text-gray-600 dark:text-gray-400">Followers</p>
        </div>
      </div>
    </div>
    
    <div class="border-t dark:border-gray-700">
      <div class="p-4">
        <h4 class="text-sm font-semibold mb-2">Recent Activity</h4>
        {#each repositories.slice(0, 3) as repo}
          <div class="mb-2">
            <p class="text-sm font-medium">{repo.name}</p>
            <p class="text-xs text-gray-600 dark:text-gray-400">
              ⭐ {repo.stargazers_count} · 🔄 {repo.forks_count}
            </p>
          </div>
        {/each}
      </div>
    </div>
  {:else if $githubStats.error}
    <div class="p-4 text-red-500 text-sm">
      Error: {$githubStats.error}
    </div>
  {/if}
</div>
