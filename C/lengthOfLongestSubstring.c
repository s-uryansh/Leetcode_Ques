#include<stdio.h>
#include<stdlib.h>


int includes(char *str , char c);
void empty(char *str);

int lengthOfLongestSubstring(char* str) {
    int max =0;
    int start =0;

    char store[100];

    int i ;

    for(int i = start ; str[i] != '\0' ; i++ ){
         if(includes(store , str[i])){
            int diff = i-start;
            if(diff > max){
                max = diff;
                }
            start += 1;
            i =  start - 1;
            empty(store);
            continue;
            
        }
        store[i] = str[i];
    }
    if(max == 0){
        max = i - start;
        printf("%d" , max);
        }
}


int includes(char *str, char c){
    int i;
    for(i =0 ; i< 100 ; i++){
        if(*(str + i) == c)
        return 1;
    }
    return 0;
}
void empty(char *str){
    int i;
    for(i = 0 ; i<100 ; i++){
        str[i] = '\0';
    }
}

int main(){
    
}