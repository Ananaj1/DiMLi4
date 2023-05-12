// projekt2.cpp : Этот файл содержит функцию "main". Здесь начинается и заканчивается выполнение программы.
//

#include <iostream>

int main()
{
    setlocale(LC_ALL,"RU");
    std::cout << "введите 2 числа"<<std::endl;
    int v = 0,v2 = 0;
    std::cin >> v >> v2;
    std::cout << "При сложении "<<v<<" И "<<v2<<" получим "<<v+v2<<std::endl <<" A при умножении получим "<<v*v2;
    return 0;
}


