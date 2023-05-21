#include <iostream>
#include <math.h>

auto main() -> int {
  unsigned int result = 0;

  for (int i = 0; i < 1000; ++i) {
    if (i % 3 == 0 && i % 5 == 0) {
      result += i;
    }
  }

  printf("Result %d\n", result);
  return 0;
}
