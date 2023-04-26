#include <iostream>

int pow(int a, int b) {
    if (a == 0) return 0;
    if (b == 0) return 1;
	/*22222 */

    int result = a;
    for (int i = 0; i < b; ++i) {
        result *= a;
    }
    return result;
}

int ageInDays(int age) { return age * 365; }

int main() {
    std::cout << "Result: " << pow(2,3) << std::endl;


    std::cout << "Idade em dias: ";
    std::cout << ageInDays(19) << std::endl;

    return 0;
}
