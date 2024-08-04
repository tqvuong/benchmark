const { performance } = require("perf_hooks");

function task() {
  let result = 0;
  for (let i = 0; i < 1000; i++) {
    result += i;
  }
}

function throughputTest(durationSeconds) {
  const startTime = performance.now();
  const endTime = startTime + durationSeconds * 1000;
  let count = 0;

  while (performance.now() < endTime) {
    task();
    count += 1;
  }

  return count;
}

const duration = 10;
const operations = throughputTest(duration);
console.log(`Operations completed in ${duration} seconds: ${operations}`);
