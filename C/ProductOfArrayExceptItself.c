#include<stdio.h>
#include<stdlib.h>
int* productExceptSelf(int* nums , int numsSize , int* returnSize);

//  * Note: The returned array must be malloced, assume caller calls free().
//  */


int* productExceptSelf(int* nums, int numsSize, int* returnSize) {

    int originalArrMulti =1;


    int zeroCount = 0;
    for(int i = 0; i < numsSize; i++) {
        if (nums[i] != 0)
            originalArrMulti *= nums[i];
        else
            zeroCount++;
    }
    
    int* multipliedArr = (int *)malloc(numsSize * sizeof(int));

    if (multipliedArr == NULL) {
            *returnSize = 0;
            return NULL;
        }
   


    if(zeroCount > 1){
        //array with more than one zero
        for(int i =0 ; i < numsSize ; i++){
            multipliedArr[i] = 0; //if more than one zero whole array will be marked as zero
        }
    }else{
        if(zeroCount == 1){
            //array with only one zero
         for (int i = 0; i < numsSize; i++) {
            
            multipliedArr[i] = (nums[i] == 0) ? originalArrMulti : 0;
            //case with only one zero ,all elements will be marked as zero expect the one whose index is 0
        }
        }else{
            // Array with no zero
            for(int i = 0 ; i <numsSize ; i++){
                multipliedArr[i] = originalArrMulti / nums[i];
                //in case of array with no zero normal loop will be executed
            }
        }
    }
    

    *returnSize = numsSize;
    return multipliedArr;

    free(multipliedArr);
    return NULL;
    }


int main(){
    
}