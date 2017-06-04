#ifndef Win64
#define Win64

#include <iostream>
#include "hellsh"

using namespace std;
using namespace hell;

int main (int argc, char** argv)
{
try {
    // args (argc, argv);
    shell Hell;
    cout << Hell.version() << endl;
    return 0;
} catch (bell& he11) {
    cout << he11.what() << endl;
    return he11.result;
}}
#endif