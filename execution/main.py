import time

def task():
    pass

start_time = time.time()
for _ in range(1000):
    task()
end_time = time.time()

print(f"Execution Time: {end_time - start_time} seconds")

