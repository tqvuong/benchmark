const task = (n) => {
  let result = 0;
  for (let i = 0; i < 1000 * n; i++) {
    result += i;
  }

  return result;
};

const concurrencyTest = async (numTasks) => {
  const tasks = Array.from({ length: numTasks }, (_, i) => i);
  const worker = async (n) => task(n);

  const startTime = Date.now();
  await Promise.all(tasks.map(worker));
  const endTime = Date.now();

  return (endTime - startTime) / 1000;
};

(async () => {
  const numTasks = 10000;
  const duration = await concurrencyTest(numTasks);
  console.log(
    `Time taken to complete ${numTasks} tasks: ${duration.toFixed(2)} seconds`
  );
})();
