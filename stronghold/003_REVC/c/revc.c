#include <stdio.h>
#include <string.h>

#define MAX_NT 1000

char complement(char base) {
    switch (base)
    {
    case 'A':
        return 'T';
    case 'C':
        return 'G';
    case 'G':
        return 'C';
    case 'T':
        return 'A';
    default:
        return 'N';
    }
}

void reverse(char *str) {
    int l = 0;
    int r = strlen(str) - 1;
    char tmp;

    while (l < r) {
        tmp = str[l];
        str[l] = str[r];
        str[r] = tmp;

        l++;
        r--;
    }
}

int main() {
    char in_seq[MAX_NT], out_seq[MAX_NT];
    fgets(in_seq, MAX_NT, stdin);
    in_seq[strcspn(in_seq, "\n")] = 0;
    
    strcpy(out_seq, in_seq);
    reverse(out_seq);

    for (int i = 0; i < strlen(out_seq); i++) {
        out_seq[i] = complement(out_seq[i]);
    }

    printf("%s\n", out_seq);

    return 0;
}