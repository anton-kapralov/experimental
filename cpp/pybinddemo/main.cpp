#include <iostream>

#include <pybind11/embed.h>

namespace py = pybind11;

const char *script = R"(
import sys
print(sys.executable)
print()

for path in sys.path:
    print(path)
print()

import os
print(os.__file__)
print()

import numpy as np
)";

int main() {
    py::scoped_interpreter guard{};

    py::exec(script);

    return 0;
}
