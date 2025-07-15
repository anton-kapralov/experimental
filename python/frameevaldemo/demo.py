import sys

from types import FrameType
from typing import Any


def trace(frame: FrameType, event: str, arg: Any):
    print(
        f"[trace] {frame.f_code.co_name} {event} ({frame.f_code.co_filename}:{frame.f_lineno}) {frame.f_locals}"
    )
    
    if frame.f_code.co_name == "sum" and event == "call":
        return mul(arg)

    if event == "call":
        frame.f_trace = trace


def foo() -> int:
    target_var = 0
    n = 3
    for i in range(n):
        target_var = sum(target_var, i)
        print(f"{i}/{n}")
    return target_var


def sum(v: int, i: int) -> int:
    return v + i


def mul(v: int, i: int) -> int:
    return v * i


if __name__ == "__main__":
    sys.settrace(trace)
    print(foo())
