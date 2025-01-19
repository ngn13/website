#pragma once
#include <ctorm/ctorm.h>

void route_cors(ctorm_req_t *req, ctorm_res_t *res);
void route_list(ctorm_req_t *req, ctorm_res_t *res);
void route_get(ctorm_req_t *req, ctorm_res_t *res);
void route_notfound(ctorm_req_t *req, ctorm_res_t *res);
