import sys


# This function will watch our program while it runs
def monitor_var(frame, event, arg):
    print(event, frame.f_lineno)
    if event == "call":
        frame.f_trace = monitor_var
    if event == "line":
        print("  ", frame.f_locals)
        if "target_var" in frame.f_locals:
            frame.f_locals["target_var"] += 1
        #     print(f"Hey, target_var just changed to {frame.f_locals['target_var']}!")


def test_function():
    target_var = 0
    for i in range(3):
        target_var += i
        print("Doing something...")


if __name__ == "__main__":
    sys.settrace(monitor_var)
    test_function()
