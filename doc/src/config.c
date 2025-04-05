#include <ctorm/ctorm.h>

#include <stdlib.h>
#include <string.h>

#include <stdio.h>
#include <errno.h>

#include "config.h"
#include "util.h"

option_t options[] = {
    // name      value           requied
    {"host",     "0.0.0.0:7003", true }, // host the server should listen on
    {"docs_dir", "./docs",       true }, // documentation directory
    {"",         NULL,           false},
};

bool config_load(config_t *conf) {
  bzero(conf, sizeof(*conf));

  char name_env[OPT_NAME_MAX + 10], name_copy[OPT_NAME_MAX], *value = NULL;
  conf->options = options;

  for (option_t *opt = conf->options; opt->value != NULL; opt++, conf->count++) {
    strcpy(name_copy, opt->name);
    util_toupper(name_copy);
    snprintf(name_env, sizeof(name_env), "WEBSITE_%s", name_copy);

    if ((value = getenv(name_env)) != NULL)
      opt->value = value;

    if (*opt->value == 0)
      opt->value = NULL;

    if (!opt->required || NULL != opt->value)
      continue;

    ctorm_fail("please specify a value for the required config option: %s (%s)", opt->name, name_env);
    errno = EFAULT;
    return false;
  }

  return true;
}

char *config_get(config_t *conf, const char *name) {
  for (int32_t i = 0; i < conf->count; i++)
    if (strcmp(conf->options[i].name, name) == 0)
      return conf->options[i].value;
  return NULL;
}
