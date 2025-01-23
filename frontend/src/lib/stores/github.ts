import { writable } from 'svelte/store';

const API_BASE_URL = 'http://localhost:8080';

interface GitHubStats {
  profile: any;
  repositories: any[];
  languages: any[];
  contributions: any[];
  loading: boolean;
  error: string | null;
}

export const githubStats = writable<GitHubStats>({
  profile: null,
  repositories: [],
  languages: [],
  contributions: [],
  loading: false,
  error: null
});

export async function fetchGitHubStats() {
  try {
    const [profileRes, reposRes, langsRes, contribsRes] = await Promise.all([
      fetch(`${API_BASE_URL}/profile`),
      fetch(`${API_BASE_URL}/repositories`),
      fetch(`${API_BASE_URL}/languages`),
      fetch(`${API_BASE_URL}/contributions`)
    ]);

    const [profile, repositories, languages, contributions] = await Promise.all([
      profileRes.json(),
      reposRes.json(),
      langsRes.json(),
      contribsRes.json()
    ]);

    githubStats.set({
      profile,
      repositories,
      languages,
      contributions,
      loading: false,
      error: null
    });
  } catch (error) {
    console.error('Error fetching GitHub stats:', error);
  }
}

// WebSocket connection for real-time updates
const ws = new WebSocket(`ws://localhost:8080/ws`);

ws.onmessage = (event) => {
  const data = JSON.parse(event.data);
  githubStats.update(stats => ({
    ...stats,
    repositories: stats.repositories.map(repo => 
      repo.id === data.repository.id ? data.repository : repo
    )
  }));
};
