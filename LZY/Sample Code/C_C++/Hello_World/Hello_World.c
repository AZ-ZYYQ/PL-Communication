#include<stdio.h>
// see if the return type of main() function in C can be void ?
#if 0
void main()
{
    printf("Print with void main()");
#else
int main(void)
{
    printf("Print with int main()");
#endif
}
//Conclusion: without return statement, the return type of main() function can be either void or int