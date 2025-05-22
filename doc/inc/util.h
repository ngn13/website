#pragma once
#include <ctorm/ctorm.h>
#include <stdbool.h>
#include <stdint.h>
#include <ctype.h>

#define util_toupper(str)                                                      \
  for (char *c = str; *c != 0; c++)                                            \
  *c = toupper(*c)
uint64_t util_endswith(char *str, char *suf);
void     util_send(ctorm_res_t *res, uint16_t code, cJSON *json);
