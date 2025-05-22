#pragma once
#include <linux/limits.h>
#include <stdbool.h>
#include <dirent.h>

#include "file.h"

typedef struct {
  DIR    *dir;
  file_t *file;
  char    name[NAME_MAX + 1];
  char   *lang;
} docs_t;

bool  docs_init(docs_t *docs, char *dir);
char *docs_next(docs_t *docs, char *name, bool content);
void  docs_reset(docs_t *docs);
void  docs_free(docs_t *docs);
