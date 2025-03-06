#include <stdio.h>

#define MAX_NT 1000

int main() {
    char seq[MAX_NT];
    fgets(seq, MAX_NT, stdin);

    for (int i = 0; seq[i] != 0; i++) {
        if (seq[i] == 'T') {
            seq[i] = 'U';
        }
    }

    printf("%s\n", seq);
    return 0;
}
