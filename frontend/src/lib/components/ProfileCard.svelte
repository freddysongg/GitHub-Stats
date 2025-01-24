<script lang="ts">
  import type { GithubUser } from '../services/github';

  export let user: GithubUser | null = null;
</script>

{#if user}
  <div class="card">
    <div class="flex items-center space-x-4">
      <img src={user.avatar_url} alt={user.login} class="w-16 h-16 rounded-full" />
      <div>
        <h2 class="text-xl font-bold">{user.name || user.login}</h2>
        {#if user.bio}
          <p class="text-github-600 dark:text-github-300">{user.bio}</p>
        {/if}
      </div>
    </div>

    <div class="grid grid-cols-3 gap-4 mt-6">
      <div class="text-center">
        <div class="text-2xl font-bold">{user.public_repos}</div>
        <div class="text-sm text-github-600 dark:text-github-300">Repositories</div>
      </div>
      <div class="text-center">
        <div class="text-2xl font-bold">{user.followers}</div>
        <div class="text-sm text-github-600 dark:text-github-300">Followers</div>
      </div>
      <div class="text-center">
        <div class="text-2xl font-bold">{user.following}</div>
        <div class="text-sm text-github-600 dark:text-github-300">Following</div>
      </div>
    </div>
  </div>
{/if}
