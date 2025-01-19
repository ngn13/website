#include "routes.h"
#include "util.h"

void route_notfound(ctorm_req_t *req, ctorm_res_t *res) {
  return util_send(res, 404, NULL);
}
