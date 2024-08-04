function task() {
  let result = 0;
  for (let i = 0; i < 1000; i++) {
    result += i;
  }
}

function memoryTest() {
  const startTime = Date.now();
  const duration = 10 * 1000;
  let memUsage = [];
  let totalMem = 0;

  const beforeMem = process.memoryUsage().heapUsed / 1024 / 1024;

  while (Date.now() - startTime < duration) {
    task();
    const heapUsed = process.memoryUsage().heapUsed / 1024 / 1024;
    memUsage.push(heapUsed);
    totalMem += heapUsed;
  }

  const afterMem = process.memoryUsage().heapUsed / 1024 / 1024;

  let maxMem = 0;
  for (let i = 0; i < memUsage.length; i++) {
    if (memUsage[i] > maxMem) {
      maxMem = memUsage[i];
    }
  }

  return { beforeMem, afterMem, maxMem, totalMem };
}

const { beforeMem, afterMem, maxMem, totalMem } = memoryTest();
console.log(`Memory usage before tasks (MB): ${beforeMem}`);
console.log(`Memory usage after tasks (MB): ${afterMem}`);
console.log(`Max memory usage (MB): ${maxMem}`);
console.log(`Total memory usage (MB): ${totalMem}`);
