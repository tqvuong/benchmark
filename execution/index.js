const { performance } = require("perf_hooks");

function task() {}

const startTime = performance.now();
for (let i = 0; i < 1000; i++) {
  task();
}
const endTime = performance.now();

console.log(`Execution Time: ${endTime - startTime} milliseconds`);
