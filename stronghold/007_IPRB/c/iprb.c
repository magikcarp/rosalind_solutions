#include <stdio.h>

float mendel(int k, int m, int n) {
  float h = (float) k;
  float i = (float) m;
  float j = (float) n;
  float total = h + i + j;
  
  float two_homo_rec = (j / total) * ((j - 1.0) / (total - 1.0));
  float two_het = (i / total) * ((i - 1) / (total - 1));
  float het_rec = ((j / total) * (i / (total - 1))) + ((i / total) * (j / (total - 1)));
  float recessive_prob = two_homo_rec + (two_het * 0.25) + (het_rec * 0.5);

  return 1.0 - recessive_prob;
}

int main() {
  int k;
  int m;
  int n;

  if (scanf("%d %d %d", &k, &m, &n) != 3) {
    fprintf(stderr, "Three integers not provided.\n");
    return 1;
  }

  printf("%f\n", mendel(k, m, n));
  return 0; 
}
