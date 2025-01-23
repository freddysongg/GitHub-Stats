<script lang="ts">
  import Header from './lib/components/Header.svelte';
  import ProfileCard from './lib/components/ProfileCard.svelte';
  import ContributionGraph from './lib/components/ContributionGraph.svelte';
  import {
    getUser,
    getContributions,
    type GithubUser,
    type ContributionDay,
  } from './lib/services/github';

  let user: GithubUser | null = null;
  let contributions: ContributionDay[] = [];
  let error: string | null = null;
  let loading = false;

  async function handleSearch(event: CustomEvent<string>) {
    const username = event.detail;
    loading = true;
    error = null;

    try {
      const [userData, contributionsData] = await Promise.all([
        getUser(username),
        getContributions(username),
      ]);

      user = userData;
      contributions = contributionsData;
    } catch (e) {
      error = e instanceof Error ? e.message : 'Failed to fetch user data';
    } finally {
      loading = false;
    }
  }
</script>

<main class="container mx-auto px-4 py-8 min-h-screen">
  <Header on:search={handleSearch} {loading} />

  {#if error}
    <div class="card bg-red-50 dark:bg-red-900 text-red-900 dark:text-red-50">
      {error}
    </div>
  {/if}

  {#if user}
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-6">
      <div class="lg:col-span-3">
        <ProfileCard {user} />
      </div>
    </div>

    <div class="card">
      <h2 class="text-xl font-bold mb-4">Contribution Activity</h2>
      <ContributionGraph data={contributions} />
    </div>
  {:else if !loading}
    <div class="card">
      <h2 class="text-xl font-bold mb-4">Welcome to GitHub Stats Dashboard</h2>
      <p class="text-github-600 dark:text-github-300">
        Enter a GitHub username above to view their stats
      </p>
    </div>
  {/if}
</main>
