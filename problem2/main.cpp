#include <iostream>
#include <math.h>

int algorithm_v2() {
	int counter = 0;
	auto Fn = [](double n) -> double {
    // Binet
		return (std::pow((1.0 + std::sqrt(5)), n) - std::pow((1.0 - std::sqrt(5)),n)) / (std::pow(2, n) * std::sqrt(5)); //CALCULO DE FIBO
	};
	for(unsigned long x = 1;; ++x) {
		if(Fn(x*3) > 4000000) break;
		counter += Fn(x * 3);
	}
	return counter;
}

auto main() -> int {
	int result = algorithm_v2();
	std::cout << result << std::endl;
}
