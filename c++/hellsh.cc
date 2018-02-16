#include <iostream>
#include "hells.h"

using namespace std;
using namespace hell;

static const string revision("$Id$");

int main (int argc, char** argv)
{
try {
    // args (argc, argv);
    cout   << "Hell$hell Id:" << revision << endl
           << "Type '?' for help" << endl;
    shell<string> Hell; //(ip,port)
    string input;
    while (cin >> input) {
        //preprocessing
        cout << Hell << input << flush;     
    }
  return 0;
//    
} catch (bell& he11) {
    cout << he11.what() << endl;
    return he11.result;
}}