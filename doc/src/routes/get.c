#include <linux/limits.h>
#include <ctorm/ctorm.h>

#include <string.h>
#include <stdlib.h>
#include <stdio.h>

#include "routes.h"
#include "config.h"
#include "util.h"

void route_get(ctorm_req_t *req, ctorm_res_t *res) {
  config_t   *conf     = REQ_LOCAL("config");
  char       *docs_dir = config_get(conf, "docs_dir");
  const char *name     = REQ_PARAM("name");

  if (NULL == name) {
    ctorm_fail("doc name is not specified");
    return util_send(res, 500, NULL);
  }

  if (strlen(name) > NAME_MAX - 6)
    return util_send(res, 404, NULL);

  char         full_path[PATH_MAX + 1], full_name[NAME_MAX + 1];
  util_file_t *doc_file = NULL, *json_file = NULL;
  cJSON       *doc_json = NULL;

  // read the doc markdown
  snprintf(full_name, sizeof(full_name), "%s.md", name);
  snprintf(full_path, sizeof(full_path), "%s/%s", docs_dir, full_name);

  if (!util_dir_contains(docs_dir, full_name)) {
    util_send(res, 404, NULL);
    goto end;
  }

  if (NULL == (doc_file = util_file_load(full_path))) {
    ctorm_fail("failed to load file: %s", full_path);
    util_send(res, 500, NULL);
    goto end;
  }

  // read the doc JSON
  snprintf(full_name, sizeof(full_name), "%s.json", name);
  snprintf(full_path, sizeof(full_path), "%s/%s", docs_dir, full_name);

  if (!util_dir_contains(docs_dir, full_name)) {
    util_send(res, 404, NULL);
    goto end;
  }

  if (NULL == (json_file = util_file_load(full_path))) {
    ctorm_fail("failed to load file: %s", full_path);
    util_send(res, 500, NULL);
    goto end;
  }

  // parse the doc JSON
  if (NULL == (doc_json = cJSON_Parse(json_file->content))) {
    ctorm_fail("failed to parse file: %s", full_path);
    util_send(res, 500, NULL);
    goto end;
  }

  cJSON_AddStringToObject(doc_json, "content", doc_file->content);
  util_send(res, 200, doc_json);

end:
  if (NULL != doc_file)
    util_file_free(doc_file);

  if (NULL != json_file)
    util_file_free(json_file);

  if (NULL != doc_json)
    cJSON_free(doc_json);
}
