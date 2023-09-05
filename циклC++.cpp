// циклC++.cpp : Этот файл содержит функцию "main". Здесь начинается и заканчивается выполнение программы.
//

#include <iostream>

int main() {
    setlocale(LC_ALL,"RU");
    std::cout << "Enter the number from which the summation will begin \n";
    int a = 0, v = 1,c =0 ,d = 1,y=0;
    std::cin >> a;
        std::cout << "Enter the number from which the summation will end \n";
    std::cin >> y;
    c = a;
    d = v;
    while (v <= y) {
        a += v;
        ++v;
    }
    std::cout << " The sum of the numbers from "<<c<<" tо "<<d<<" they're a rivne: " << a<<std::endl;
        return 0;
}


