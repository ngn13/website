#include <ctorm/all.h>
#include <stdlib.h>

#include "routes.h"
#include "config.h"

int main() {
  config_t     conf;
  app_config_t config;

  if (config_load(&conf) < 0)
    return EXIT_FAILURE;

  app_config_new(&config);
  config.disable_logging = true;

  app_t *app = app_new(&config);
  GET(app, "/read", GET_read);

  if (!app_run(app, config_get(&conf, "host")))
    error("failed to start the app: %s", app_geterror());

  app_free(app);
  return EXIT_SUCCESS;
}
