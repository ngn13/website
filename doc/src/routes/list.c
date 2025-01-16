#include <linux/limits.h>
#include <ctorm/ctorm.h>

#include <dirent.h>
#include <string.h>
#include <stdio.h>

#include "util.h"
#include "config.h"

void route_list(ctorm_req_t *req, ctorm_res_t *res) {
  config_t *conf     = REQ_LOCAL("config");
  char     *docs_dir = config_get(conf, "docs_dir");

  cJSON *array = NULL, *json = NULL;
  DIR   *docs_dir_fd = NULL;

  if (NULL == (array = cJSON_CreateArray())) {
    ctorm_fail("failed to create cJSON array");
    return util_send(res, 500, NULL);
  }

  if (NULL == (docs_dir_fd = opendir(docs_dir))) {
    ctorm_fail("failed to open the docs dir (%s): %s", docs_dir, ctorm_geterror());
    return util_send(res, 500, NULL);
  }

  char           doc_path[PATH_MAX + 1], doc_name[NAME_MAX + 1];
  util_file_t   *doc_file = NULL;
  struct dirent *doc      = NULL;
  uint64_t       ext_indx = 0;

  while (NULL != (doc = readdir(docs_dir_fd))) {
    if ((ext_indx = util_endswith(doc->d_name, ".json")) == 0)
      continue;

    snprintf(doc_path, sizeof(doc_path), "%s/%s", docs_dir, doc->d_name);

    if (NULL == (doc_file = util_file_load(doc_path))) {
      ctorm_fail("failed to load the JSON file: %s", doc_path);
      goto next;
    }

    if (NULL == (json = cJSON_Parse(doc_file->content))) {
      ctorm_fail("failed to parse the JSON file: %s", doc_path);
      goto next;
    }

    strcpy(doc_name, doc->d_name);
    util_tolower(doc_name);
    doc_name[ext_indx] = 0;

    cJSON_AddStringToObject(json, "name", doc_name);
    cJSON_AddItemToArray(array, json);

  next:
    if (NULL != doc_file)
      util_file_free(doc_file);
    doc_file = NULL;
  }

  closedir(docs_dir_fd);

  if (NULL == (json = cJSON_CreateObject())) {
    ctorm_fail("failed to create cJSON object");
    return util_send(res, 500, NULL);
  }

  cJSON_AddItemToObject(json, "list", array);
  util_send(res, 200, json);
  cJSON_free(json);
}
