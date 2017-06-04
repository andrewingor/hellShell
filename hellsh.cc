#ifndef Win64
#define Win64

#include <iostream>
#include "hells.h"

using namespace std;
using namespace hell;

static string revision("$Id$");

int main (int argc, char** argv)
{
try {
    // args (argc, argv);
    shell Hell;
    return 0;
} catch (bell& he11) {
    cout << he11.what() << endl;
    return he11.result;
}}
#endif