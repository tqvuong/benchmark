import time

from memory_profiler import memory_usage


def task():
    sum(i for i in range(1000))

def memory_test():
    mem_usage = memory_usage((run_tasks,), interval=0.1)
    max_mem = max(mem_usage)
    total_mem = sum(mem_usage)
    return max_mem, total_mem

def run_tasks():
    start_time = time.time()
    end_time = start_time + 10
    while time.time() < end_time:
        task()

if __name__ == "__main__":
    before_mem = memory_usage()[0]
    max_mem, total_mem = memory_test()
    after_mem = memory_usage()[0]
    print(f"Memory usage before test (MB): {before_mem}")
    print(f"Memory usage after test (MB): {after_mem}")
    print(f"Max memory usage (MB): {max_mem}")
    print(f"Total memory usage (MB): {total_mem}")

# Memory usage before test (MB): 21.84375
# Memory usage after test (MB): 22.703125
# Max memory usage (MB): 22.703125
# Total memory usage (MB): 2292.8125