from __future__ import annotations

import os
import statistics
import sys
import time
import urllib.error
import urllib.request
from concurrent.futures import ThreadPoolExecutor, as_completed


def env_int(key: str, default: int) -> int:
    value = os.environ.get(key)
    if value and value.isdigit() and int(value) > 0:
        return int(value)
    return default


def fetch(url: str) -> tuple[float, bool]:
    start = time.perf_counter()
    try:
        with urllib.request.urlopen(url, timeout=30) as resp:
            resp.read()
            ok = 200 <= resp.status < 300
    except (urllib.error.URLError, TimeoutError):
        ok = False
    elapsed = time.perf_counter() - start
    return elapsed, ok


def percentile(sorted_values: list[float], p: int) -> float:
    if not sorted_values:
        return 0.0
    idx = min((p * len(sorted_values)) // 100, len(sorted_values) - 1)
    return sorted_values[idx]


def main() -> None:
    url = os.environ.get("BENCH_URL", "http://localhost:8000/students")
    total = env_int("BENCH_REQUESTS", 10000)
    concurrency = env_int("BENCH_CONCURRENCY", 50)

    print("Python benchmark client")
    print(f"  Target:      {url}")
    print(f"  Requests:    {total}")
    print(f"  Concurrency: {concurrency}\n")

    try:
        urllib.request.urlopen(url, timeout=5).read()
    except Exception as exc:
        print(f"Server not reachable at {url}: {exc}", file=sys.stderr)
        sys.exit(1)

    latencies: list[float] = []
    ok_count = 0
    failed = 0

    start = time.perf_counter()
    with ThreadPoolExecutor(max_workers=concurrency) as pool:
        futures = [pool.submit(fetch, url) for _ in range(total)]
        for future in as_completed(futures):
            elapsed, ok = future.result()
            latencies.append(elapsed)
            if ok:
                ok_count += 1
            else:
                failed += 1
    duration = time.perf_counter() - start

    latencies.sort()
    rps = total / duration if duration > 0 else 0

    print("Results")
    print("-------")
    print("Language: Python (FastAPI client)")
    print(f"URL: {url}")
    print(f"Duration: {duration:.3f}s")
    print(f"Requests/sec: {rps:.2f}")
    print(f"Successful: {ok_count}")
    print(f"Failed: {failed}")
    print(f"Avg latency: {statistics.mean(latencies) * 1000:.2f}ms")
    print(f"p50 latency: {percentile(latencies, 50) * 1000:.2f}ms")
    print(f"p95 latency: {percentile(latencies, 95) * 1000:.2f}ms")
    print(f"p99 latency: {percentile(latencies, 99) * 1000:.2f}ms")


if __name__ == "__main__":
    main()
