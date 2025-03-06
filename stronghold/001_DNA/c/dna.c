#include <stdio.h>

#define MAX_NT 1000

struct BaseCount {
    int a;
    int c;
    int g;
    int t;
};

int main() {
    char input[MAX_NT];
    struct BaseCount bc;

    fgets(input, MAX_NT, stdin);

    for (int i = 0; input[i] != 0; i++) {
        switch (input[i])
        {
        case 'A':
            bc.a += 1;
            break;
        case 'C':
            bc.c += 1;
            break;
        case 'G':
            bc.g += 1;
            break;
        case 'T':
            bc.t += 1;
            break;
        default:
            break;
        }
    }

    printf("%d %d %d %d\n", bc.a, bc.c, bc.g, bc.t);
}
