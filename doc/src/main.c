#include <ctorm/app.h>
#include <ctorm/ctorm.h>
#include <stdlib.h>

#include "config.h"
#include "routes.h"

int main() {
  ctorm_app_t   *app = NULL;
  ctorm_config_t app_config;

  config_t conf;
  char    *host = NULL;

  if (!config_load(&conf))
    return EXIT_FAILURE;

  if (NULL == (host = config_get(&conf, "host"))) {
    ctorm_fail("failed to get the host configuration");
    return EXIT_FAILURE;
  }

  // initialize the app config
  ctorm_config_new(&app_config);
  app_config.disable_logging = true;

  // create a new app
  app = ctorm_app_new(&app_config);

  // middlewares
  ALL(app, "/*", route_cors);
  ALL(app, "/*/*", route_cors);

  // routes
  GET(app, "/list", route_list);
  GET(app, "/get/:name", route_get);

  ctorm_app_default(app, route_notfound);
  ctorm_app_local(app, "config", &conf);

  ctorm_info("starting the web server on %s", host);

  if (!ctorm_app_run(app, host))
    ctorm_fail("failed to start the app: %s", ctorm_error());

  ctorm_app_free(app);
  return EXIT_SUCCESS;
}
