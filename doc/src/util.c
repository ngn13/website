#include <linux/limits.h>
#include <ctorm/ctorm.h>

#include <stdint.h>
#include <string.h>
#include <errno.h>

#include "util.h"

uint64_t util_endswith(char *str, char *suf) {
  if (NULL == str || NULL == suf) {
    errno = EINVAL;
    return false;
  }

  uint64_t sufl = strlen(suf);
  uint64_t strl = strlen(str);

  if (sufl > strl)
    return false;

  uint64_t indx = strl - sufl;

  return strncmp(str + indx, suf, sufl) == 0 ? indx : 0;
}

void util_send(ctorm_res_t *res, uint16_t code, cJSON *json) {
  if (NULL == json)
    json = cJSON_CreateObject();
  const char *error = "";

  switch (code) {
  case 404:
    error = "not found";
    break;

  case 400:
    error = "bad request";
    break;

  case 500:
    error = "internal server error";
    break;
  }

  if (*error != 0)
    cJSON_AddStringToObject(json, "error", error);

  RES_JSON(json);
}
