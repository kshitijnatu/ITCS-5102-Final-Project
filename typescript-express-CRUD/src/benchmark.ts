/**
 * Load-test client for the Express CRUD server.
 *
 * Run the server first:
 *   npm run build && npm start
 *
 * Then:
 *   npm run benchmark
 *
 * Environment variables:
 *   BENCH_URL          (default: http://localhost:3000/students)
 *   BENCH_REQUESTS     (default: 10000)
 *   BENCH_CONCURRENCY  (default: 50)
 */

const DEFAULT_URL = "http://localhost:3000/students";
const DEFAULT_REQUESTS = 10_000;
const DEFAULT_CONCURRENCY = 50;

function envInt(key: string, fallback: number): number {
  const value = process.env[key];
  if (value && /^\d+$/.test(value) && parseInt(value, 10) > 0) {
    return parseInt(value, 10);
  }
  return fallback;
}

function percentile(sorted: number[], p: number): number {
  if (sorted.length === 0) return 0;
  const idx = Math.min(Math.floor((p * sorted.length) / 100), sorted.length - 1);
  return sorted[idx];
}

async function fetchOnce(url: string): Promise<{ elapsed: number; ok: boolean }> {
  const start = performance.now();
  try {
    const res = await fetch(url);
    await res.arrayBuffer();
    return { elapsed: performance.now() - start, ok: res.ok };
  } catch {
    return { elapsed: performance.now() - start, ok: false };
  }
}

async function runBatch(
  url: string,
  count: number
): Promise<{ elapsed: number; ok: boolean }[]> {
  return Promise.all(Array.from({ length: count }, () => fetchOnce(url)));
}

async function main(): Promise<void> {
  const url = process.env.BENCH_URL ?? DEFAULT_URL;
  const total = envInt("BENCH_REQUESTS", DEFAULT_REQUESTS);
  const concurrency = envInt("BENCH_CONCURRENCY", DEFAULT_CONCURRENCY);

  console.log("TypeScript benchmark client");
  console.log(`  Target:      ${url}`);
  console.log(`  Requests:    ${total}`);
  console.log(`  Concurrency: ${concurrency}\n`);

  try {
    const warmup = await fetch(url);
    if (!warmup.ok) throw new Error(`HTTP ${warmup.status}`);
  } catch (err) {
    console.error(`Server not reachable at ${url}:`, err);
    process.exit(1);
  }

  const latencies: number[] = [];
  let okCount = 0;
  let failed = 0;

  const start = performance.now();
  const batches = Math.ceil(total / concurrency);
  let sent = 0;

  for (let b = 0; b < batches; b++) {
    const batchSize = Math.min(concurrency, total - sent);
    const results = await runBatch(url, batchSize);
    for (const result of results) {
      latencies.push(result.elapsed);
      if (result.ok) okCount++;
      else failed++;
    }
    sent += batchSize;
  }
  const duration = (performance.now() - start) / 1000;

  latencies.sort((a, b) => a - b);
  const rps = total / duration;
  const avg = latencies.reduce((a, b) => a + b, 0) / latencies.length;

  console.log("Results");
  console.log("-------");
  console.log("Language:     TypeScript (Express client)");
  console.log(`URL:          ${url}`);
  console.log(`Duration:     ${duration.toFixed(3)}s`);
  console.log(`Requests/sec: ${rps.toFixed(2)}`);
  console.log(`Successful:   ${okCount}`);
  console.log(`Failed:       ${failed}`);
  console.log(`Avg latency:  ${avg.toFixed(2)}ms`);
  console.log(`p50 latency:  ${percentile(latencies, 50).toFixed(2)}ms`);
  console.log(`p95 latency:  ${percentile(latencies, 95).toFixed(2)}ms`);
  console.log(`p99 latency:  ${percentile(latencies, 99).toFixed(2)}ms`);
}

main();
