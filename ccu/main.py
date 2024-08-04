import time
import concurrent.futures

def task(n):
    sum(i for i in range(1000 * n))

def concurrency_test(num_tasks, num_workers):
    with concurrent.futures.ThreadPoolExecutor(max_workers=num_workers) as executor:
        futures = [executor.submit(task, i) for i in range(num_tasks)]
        start_time = time.time()
        concurrent.futures.wait(futures)
        end_time = time.time()

    return end_time - start_time

if __name__ == "__main__":
    num_tasks = 10000
    num_workers = 10
    duration = concurrency_test(num_tasks, num_workers)
    print(f"Time taken to complete {num_tasks} tasks with {num_workers} workers: {duration:.2f} seconds")
