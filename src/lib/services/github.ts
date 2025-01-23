import { Octokit } from '@octokit/rest';

const octokit = new Octokit();

export interface GithubUser {
  login: string;
  name: string | null;
  avatar_url: string;
  bio: string | null;
  public_repos: number;
  followers: number;
  following: number;
}

export interface ContributionDay {
  date: string;
  count: number;
}

export async function getUser(username: string): Promise<GithubUser> {
  const { data } = await octokit.users.getByUsername({ username });
  return data;
}

export async function getContributions(username: string): Promise<ContributionDay[]> {
  // Since GitHub's API doesn't provide contribution data directly,
  // we'll simulate it for now with random data for the last year
  const contributions: ContributionDay[] = [];
  const today = new Date();

  for (let i = 365; i >= 0; i--) {
    const date = new Date(today);
    date.setDate(date.getDate() - i);
    contributions.push({
      date: date.toISOString().split('T')[0],
      count: Math.floor(Math.random() * 10),
    });
  }

  return contributions;
}

export async function getRepositories(username: string) {
  const { data } = await octokit.repos.listForUser({
    username,
    sort: 'updated',
    per_page: 100,
  });
  return data;
}
