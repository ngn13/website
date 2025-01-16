#include <cjson/cJSON.h>
#include <dirent.h>
#include <linux/limits.h>
#include <ctorm/ctorm.h>

#include <string.h>
#include <stdlib.h>
#include <stdio.h>

#include "routes.h"
#include "config.h"
#include "docs.h"

void route_get(ctorm_req_t *req, ctorm_res_t *res) {
  config_t *conf     = REQ_LOCAL("config");
  char     *doc_name = REQ_PARAM("name");
  char     *docs_dir = config_get(conf, "docs_dir"), *doc_data = NULL;
  cJSON    *json = NULL, *doc_json = NULL;
  docs_t    docs;

  if (NULL == doc_name) {
    ctorm_fail("documentation name not specified (how did that even happend)");
    util_send(res, 500, NULL);
    goto end;
  }

  if (!docs_init(&docs, docs_dir)) {
    ctorm_fail("docs_init failed: %s", ctorm_geterror());
    util_send(res, 500, NULL);
    goto end;
  }

  if (NULL == (json = cJSON_CreateObject())) {
    ctorm_fail("failed to create cJSON object");
    util_send(res, 500, NULL);
    goto end;
  }

  while (NULL != (doc_data = docs_next(&docs, doc_name, false))) {
    if (NULL == (doc_json = cJSON_Parse(doc_data))) {
      ctorm_fail("failed to parse JSON: %s (%s)", docs.name, docs.lang);
      continue;
    }

    cJSON_AddStringToObject(doc_json, "content", "");
    cJSON_AddItemToObject(json, docs.lang, doc_json);
  }

  if (NULL == doc_json) {
    util_send(res, 404, NULL);
    goto end;
  }

  docs_reset(&docs);

  while (NULL != (doc_data = docs_next(&docs, doc_name, true))) {
    if (NULL == (doc_json = cJSON_GetObjectItem(json, docs.lang)))
      continue;

    cJSON_DeleteItemFromObject(doc_json, "content");
    cJSON_AddStringToObject(doc_json, "content", doc_data);
  }

  util_send(res, 200, json);

end:
  docs_free(&docs);
  if (NULL != json)
    cJSON_Delete(json);
}
