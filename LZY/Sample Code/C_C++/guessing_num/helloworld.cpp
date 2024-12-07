#include <iostream>
using namespace std;
#include <ctime>
int main(void)
{
    srand((unsigned int)time(NULL));

    int num = rand() % 100 + 1;
    int num1;
    while(1)
    {
        cout << "Please input a value: ";
        cin >> num1;
        if (num1 == num){
            cout << "You are right!!" << endl;
            break;
        }else if (num1 > num){
            cout << "The value you entered is too big, please try again" << endl;
        }else {
            cout << "The value you entered is too small, please try again" << endl;
        }

    }
    system("pause");
    return 0;
}