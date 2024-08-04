const task = () => {
  let result = 0;
  for (let i = 0; i < 1000; i++) {
    result += i;
  }
};

const cpuTest = (durationSeconds) => {
  const startTime = Date.now();
  const duration = durationSeconds * 1000;
  let cpuUsage = [];

  while (Date.now() - startTime < duration) {
    const startUsage = process.cpuUsage();
    task();
    const endUsage = process.cpuUsage(startUsage);
    const cpuPercent =
      ((endUsage.user + endUsage.system) / 1000 / durationSeconds) * 100;
    cpuUsage.push(cpuPercent);
  }

  return cpuUsage;
};

const duration = 10; // Duration in seconds
const cpuUsage = cpuTest(duration);
const avgCpuUsage = cpuUsage.reduce((a, b) => a + b, 0) / cpuUsage.length;
console.log(`Average CPU usage: ${avgCpuUsage.toFixed(10)}%`);
