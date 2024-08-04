import time
import psutil

def task():
    sum(i for i in range(1000))

def cpu_test(duration_seconds):
    cpu_usage = []
    start_time = time.time()
    end_time = start_time + duration_seconds

    while time.time() < end_time:
        task()
        cpu_percent = psutil.cpu_percent(interval=None)
        cpu_usage.append(cpu_percent)

    return cpu_usage

if __name__ == "__main__":
    duration = 10
    cpu_usage = cpu_test(duration)
    avg_cpu_usage = sum(cpu_usage) / len(cpu_usage)
    print(f"Average CPU usage: {avg_cpu_usage:.10f}%")
