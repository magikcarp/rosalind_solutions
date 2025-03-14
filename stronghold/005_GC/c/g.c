#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_HEADER 128
#define MAX_SEQ 1024

typedef struct {
  char header[MAX_HEADER];
  char seq[MAX_SEQ];
  float gc_content;
} FastaRecord;

void calc_gc_content(FastaRecord *fa) {
  int gc = 0, len = 0;
  for (char *p = fa->seq; *p; p++) {
    if (*p == 'C' || *p == 'G') {
      gc++;
    }
    len++;
  }
  fa->gc_content = (len > 0) ? (100.0 * gc / len) : 0.0;
}

void remove_first_char(char *str) {
  if (str != NULL && str[0] != '\0') {
    memmove(str, str + 1, strlen(str));
  }
}

int main() {
  char line[MAX_SEQ];
  FastaRecord sequences[10];
  int seq_count = 0;

  while (fgets(line, sizeof(line), stdin)) {
    if (line[0] == '>') {
      if (seq_count > 0) {
        calc_gc_content(&sequences[seq_count - 1]);
      }
      remove_first_char(line);
      line[strlen(line)-1] = '\0';
      strcpy(sequences[seq_count].header, line);
      sequences[seq_count].seq[0] = '\0';
      seq_count++;
    } else {
      line[strcspn(line, "\n")] = 0;
      strcat(sequences[seq_count - 1].seq, line);
    }
  }
  if (seq_count > 0) {
    calc_gc_content(&sequences[seq_count - 1]);
  }

  FastaRecord max_gc;
  for (int i = 0; i < seq_count; i++) {
    if (sequences[i].gc_content > max_gc.gc_content) {
      max_gc = sequences[i];
    }
  }

  printf("%s\n%.5f\n", max_gc.header, max_gc.gc_content);
  return 0;
}
