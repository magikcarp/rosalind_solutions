#include <stdio.h>
#include <string.h>

#define MAX_NT 1024

int hamming_distance(char* s1, char* s2) {
    int hd = 0;

    unsigned long len_s1 = strlen(s1);
    unsigned long len_s2 = strlen(s2);
    if (len_s1 != len_s2) return -1;
    
    for (size_t i = 0; i < len_s1; i++) {
        if (s1[i] != s2[i]) {
            hd++;
        }
    }

    return hd;
}

int main() {
    char first_seq[MAX_NT];
    char second_seq[MAX_NT];

    if (!fgets(first_seq, MAX_NT, stdin)) {
        fprintf(stderr, "First sequence not received.\n");
        return(1);
    }
    if (!fgets(second_seq, MAX_NT, stdin)){
        fprintf(stderr, "Second sequence not received.\n");
        return(1);
    }

    // TODO check if HD calculation was a success
    printf("%d\n", hamming_distance(first_seq, second_seq));
}
