#ifndef HELL
#define HELL

#include <iostream>
#include <exception>
using namespace std;

namespace hell {
/*
 *
 *
*/
class bell: public exception {
    public:
    int result;
    virtual const char* what() const throw()
    { return "Hell shit happend"; }
};
/*
 *
 *
*/
class shell {

     public:
     
     shell (void) {
         cout   << "-=<The_$hell_from_the_Hell>=-" << endl
                << "Type '?' for help" << endl;
     }
     ~shell (void) {}
     string& version (void) {
         static string v("0.0.1");
         return v;
     }
 };
 } //namespace hell
#endif