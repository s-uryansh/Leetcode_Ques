#include<stdio.h>
#include<stdlib.h>
#include<stdbool.h>
bool isPalindrome(int x);

bool isPalindrome(int x) {
     if (x<0){
        return false;
    }// if x is -ve it will disturb the check as "-" will be obstacle for finding palindrome
    long long re = 0; // new variable which will be having reversed x to check if x is palindrome or not
    int actual = x;

    while(x > 0){
        re = re*10 + x % 10;
        x /= 10; 
    }
    //re is continously multiplied by 10 and added x % 10 till x goes below 10 this is loop to reverse the number x
    return actual == re;   // comparing both of them to find if x is palindrome or not

    if(actual == re){
        printf("true");
    }else{
        printf("false");
    }  

}

int main(){

}