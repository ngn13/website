#include <ctorm/ctorm.h>

void route_cors(ctorm_req_t *req, ctorm_res_t *res) {
  RES_SET("Access-Control-Allow-Origin", "*");
  RES_SET("Access-Control-Allow-Headers",
      "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, "
      "Authorization, accept, origin, Cache-Control, "
      "X-Requested-With");
  RES_SET("Access-Control-Allow-Methods", "PUT, DELETE, GET");
}
