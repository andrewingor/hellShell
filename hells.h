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
    { return "Hell bullshit happend"; }
};//class bell
/*
 *
 *
*/
template<
      class CharT
    , class Traits = char_traits<CharT>
>
class shell: public basic_iostream<CharT, Traits> {
     public:
     shell (void) {}
     ~shell (void) {}
 };//class shell
 } //namespace hell
#endif