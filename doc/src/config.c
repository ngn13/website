#include <ctorm/log.h>
#include <stdlib.h>
#include <string.h>
#include <errno.h>
#include "config.h"

option_t options[] = {
    {"host", "0.0.0.0:7003", true }, // host the server should listen on
    {NULL,   NULL,           false},
};

int32_t config_load(config_t *conf) {
  bzero(conf, sizeof(*conf));

  char *value   = NULL;
  conf->options = options;

  for (option_t *opt = conf->options; opt->name != NULL; opt++) {
    if ((value = getenv(opt->name)) != NULL)
      opt->value = value;

    if (opt->required && *opt->value == 0) {
      error("please specify a value for the required config option %s", opt->name);
      errno = EFAULT;
      return -1;
    }

    conf->count++;
  }

  return 0;
}

char *config_get(config_t *conf, const char *name) {
  for (int32_t i = 0; i < conf->count; i++)
    if (strcmp(conf->options[i].name, name) == 0)
      return conf->options[i].value;
  return NULL;
}
