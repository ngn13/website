#pragma once
#include <stdint.h>

typedef struct {
  char   *content;
  int64_t size;
} file_t;

file_t *file_load(int dirfd, char *path);
void    file_free(file_t *file);
