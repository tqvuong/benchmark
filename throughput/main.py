import time

def task():
    sum(i for i in range(1000))

def throughput_test(duration_seconds):
    start_time = time.time()
    end_time = start_time + duration_seconds
    count = 0

    while time.time() < end_time:
        task()
        count += 1

    return count

if __name__ == "__main__":
    duration = 10
    operations = throughput_test(duration)
    print(f"Operations completed in {duration} seconds: {operations}")
