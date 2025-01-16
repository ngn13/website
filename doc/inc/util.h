#pragma once
#include <ctorm/ctorm.h>
#include <stdbool.h>
#include <stdint.h>
#include <ctype.h>

typedef struct {
  char    *content;
  uint64_t size;
} util_file_t;

#define util_toupper(str)                                                                                              \
  for (char *c = str; *c != 0; c++)                                                                                    \
  *c = toupper(*c)
uint64_t     util_endswith(char *str, char *suf);
void         util_send(ctorm_res_t *res, uint16_t code, cJSON *json);
util_file_t *util_file_load(int dirfd, char *path);
void         util_file_free(util_file_t *file);
bool         util_parse_doc_name(char *name, char **lang, const char *ext);
