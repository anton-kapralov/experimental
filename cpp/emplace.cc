#include <iostream>
#include <deque>

class Point {
public:
  Point(int x, int y) {
    this->x = x;
    this->y = y;
  }

public:
  int x;
  int y;
};

int main() {
  std::deque<Point> dq;
  dq.emplace_back(42, 34);

  for (auto it = dq.begin(); it != dq.end(); ++it) {
    std::cout << "(" << it->x << ", " << it->y << ")" << std::endl;
  }

  return 0;
}

