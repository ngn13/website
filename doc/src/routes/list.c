#include <linux/limits.h>
#include <cjson/cJSON.h>
#include <ctorm/ctorm.h>

#include <dirent.h>
#include <stdio.h>

#include "routes.h"
#include "config.h"
#include "util.h"
#include "docs.h"

void route_list(ctorm_req_t *req, ctorm_res_t *res) {
  config_t *conf = REQ_LOCAL("config");
  char     *dir = config_get(conf, "dir"), *doc_data = NULL;
  cJSON    *array = NULL, *json = NULL, *doc_json = NULL;
  docs_t    docs;

  if (!docs_init(&docs, dir)) {
    ctorm_fail("docs_init failed: %s", ctorm_error());
    util_send(res, 500, NULL);
    goto end;
  }

  if (NULL == (json = cJSON_CreateObject())) {
    ctorm_fail("failed to create cJSON object");
    util_send(res, 500, NULL);
    goto end;
  }

  while (NULL != (doc_data = docs_next(&docs, NULL, false))) {
    if (NULL == (array = cJSON_GetObjectItem(json, docs.lang)) &&
        NULL == (array = cJSON_AddArrayToObject(json, docs.lang))) {
      ctorm_fail(
          "failed to create an array object for the language %s", docs.lang);
      continue;
    }

    if (NULL == (doc_json = cJSON_Parse(doc_data))) {
      ctorm_fail("failed to parse JSON: %s (%s)", docs.name, docs.lang);
      continue;
    }

    cJSON_AddStringToObject(doc_json, "name", docs.name);
    cJSON_AddItemToArray(array, doc_json);
  }

  util_send(res, 200, json);

end:
  docs_free(&docs);
  if (NULL != json)
    cJSON_Delete(json);
}
