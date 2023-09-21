import pytest

from python.pytestdemo import pytest_demo


@pytest.mark.parametrize(
    "a, b, c, want_x1, want_x2",
    [
        (1, -7, 15, None, None),
        (1, -4, 4, 2, 2),
        (1, -7, 10, 2, 5),
    ],
)
def test_roots_of_quadratic_equation(a, b, c, want_x1, want_x2):
    x1, x2 = pytest_demo.roots_of_quadratic_equation(a, b, c)
    assert x1 == want_x1
    assert x2 == want_x2
