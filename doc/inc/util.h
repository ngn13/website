#pragma once
#include <ctorm/ctorm.h>
#include <stdbool.h>
#include <stdint.h>
#include <ctype.h>

typedef struct {
  char    *content;
  uint64_t size;
} util_file_t;

#define util_compare_name(n1, n2) (strncmp(n1, n2, NAME_MAX) == 0)
#define util_toupper(str)                                                                                              \
  for (char *c = str; *c != 0; c++)                                                                                    \
  *c = toupper(*c)
#define util_tolower(str)                                                                                              \
  for (char *c = str; *c != 0; c++)                                                                                    \
  *c = tolower(*c)
uint64_t     util_endswith(char *str, char *suf);
void         util_send(ctorm_res_t *res, uint16_t code, cJSON *json);
bool         util_dir_contains(char *dir, const char *file);
util_file_t *util_file_load(char *path);
void         util_file_free(util_file_t *file);
