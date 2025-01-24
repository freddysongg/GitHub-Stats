<script lang="ts">
  import { onMount } from 'svelte';
  import * as d3 from 'd3';
  import type { ContributionDay } from '../services/github';

  export let data: ContributionDay[] = [];

  onMount(() => {
    if (data.length) {
      createGraph();
    }
  });

  let svg: SVGElement;
  const margin = { top: 20, right: 20, bottom: 20, left: 20 };
  const cellSize = 10;
  const cellPadding = 2;

  $: width = (cellSize + cellPadding) * 53 + margin.left + margin.right;
  $: height = (cellSize + cellPadding) * 7 + margin.top + margin.bottom;

  function createGraph() {
    if (!data.length) return;

    // Clear previous graph
    d3.select(svg).selectAll('*').remove();

    const colorScale = d3
      .scaleQuantize<string>()
      .domain([0, d3.max(data, (d) => d.count) || 10])
      .range([
        'var(--contribution-color-0)',
        'var(--contribution-color-1)',
        'var(--contribution-color-2)',
        'var(--contribution-color-3)',
        'var(--contribution-color-4)',
      ]);

    const dateExtent = d3.extent(data, (d) => new Date(d.date)) as [Date, Date];
    const weeks = d3.timeWeeks(dateExtent[0], dateExtent[1]);

    const chart = d3
      .select(svg)
      .attr('width', width)
      .attr('height', height)
      .append('g')
      .attr('transform', `translate(${margin.left}, ${margin.top})`);

    const days = chart
      .selectAll('.day')
      .data(data)
      .enter()
      .append('rect')
      .attr('class', 'day')
      .attr('width', cellSize)
      .attr('height', cellSize)
      .attr('x', (d) => {
        const date = new Date(d.date);
        const week = Math.floor(
          (date.getTime() - dateExtent[0].getTime()) / (7 * 24 * 60 * 60 * 1000)
        );
        return week * (cellSize + cellPadding);
      })
      .attr('y', (d) => {
        const date = new Date(d.date);
        return date.getDay() * (cellSize + cellPadding);
      })
      .attr('rx', 2)
      .attr('fill', (d) => colorScale(d.count))
      .attr('data-date', (d) => d.date)
      .attr('data-count', (d) => d.count);

    // Add tooltips
    days.append('title').text((d) => `${d.date}: ${d.count} contributions`);
  }

  $: if (data) {
    createGraph();
  }
</script>

<div class="contribution-graph overflow-x-auto">
  <svg bind:this={svg} class="contribution-graph__svg"></svg>
</div>

<style>
  .contribution-graph {
    @apply w-full;
  }

  .contribution-graph__svg {
    @apply min-w-full;
  }
</style>
