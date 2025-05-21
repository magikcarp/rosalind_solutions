#include <stdio.h>
#include <string.h>
#include <stdint.h>

unsigned long int rabbit_repro(unsigned long int n, unsigned long int k) {
  unsigned long int months[n];
  months[0] = 1;
  months[1] = 1;

  for (int i = 2; i < n; i++) {
    months[i] = months[i-1] + (months[i-2] * k);
  }

  return months[n-1];
}

int main() {
  int n;
  int k;
  
  if (scanf("%d %d", &n, &k) != 2) {
    fprintf(stderr, "[ERROR] Expected two integers\n");
    return 1;
  }
  
  printf("%lu\n", rabbit_repro((unsigned long int) n, (unsigned long int) k));
  return 0;
}
