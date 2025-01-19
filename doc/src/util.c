#include <linux/limits.h>
#include <ctorm/ctorm.h>

#include <sys/mman.h>
#include <sys/stat.h>

#include <stdint.h>
#include <dirent.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#include <errno.h>
#include <fcntl.h>
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

util_file_t *util_file_load(int dirfd, char *path) {
  if (NULL == path) {
    errno = EINVAL;
    return NULL;
  }

  util_file_t *file = NULL;
  struct stat  buf;
  int          fd = -1;

  if (NULL == (file = malloc(sizeof(util_file_t))))
    goto end; // errno set by malloc

  if ((fd = openat(dirfd, path, O_RDONLY)) < 0)
    goto end; // errno set by open

  if (fstat(fd, &buf) < 0)
    goto end; // errno set by fstat

  if (NULL == (file->content = mmap(0, (file->size = buf.st_size), PROT_READ, MAP_PRIVATE, fd, 0)))
    goto end; // errno set by mmap

end:
  if (fd != -1)
    close(fd);

  if (NULL == file->content) {
    free(file);
    return NULL;
  }

  return file;
}

void util_file_free(util_file_t *file) {
  if (NULL == file)
    return;

  munmap(file->content, file->size);
  free(file);
}
