#pragma once

#include <stdbool.h>
#include <stdint.h>

typedef struct {
#define OPT_NAME_MAX 20
  char  name[20]; // option name
  char *value;    // option value
  bool  required; // is the option required (does it need to have a value)
} option_t;

typedef struct config {
  option_t *options;
  int32_t   count;
} config_t;

bool  config_load(config_t *conf);
char *config_get(config_t *conf, const char *name);
