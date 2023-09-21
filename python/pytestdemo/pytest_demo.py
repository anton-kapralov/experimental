import math


def roots_of_quadratic_equation(a: float, b: float, c: float) -> (float, float):
    d = b ** 2 - 4 * a * c
    if d < 0:
        return None, None
    x1 = (-b - math.sqrt(d)) / 2 * a
    x2 = (-b + math.sqrt(d)) / 2 * a
    return x1, x2
