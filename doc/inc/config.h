#pragma once

#include <stdbool.h>
#include <stdint.h>

typedef struct option {
  const char *name;
  char       *value;
  bool        required;
} option_t;

typedef struct config {
  option_t *options;
  int32_t   count;
} config_t;

int32_t config_load(config_t *conf);
char   *config_get(config_t *conf, const char *name);
