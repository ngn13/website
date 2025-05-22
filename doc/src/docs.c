#include <linux/limits.h>

#include <dirent.h>
#include <string.h>
#include <stdint.h>

#include "util.h"
#include "docs.h"

#define DOCS_LANG_CODE_LEN 2

bool _docs_parse_name(docs_t *docs, char *ext) {
  // check the extension
  uint64_t ext_pos = util_endswith(docs->name, ext);

  if (ext_pos == 0)
    return false;

  // example.en.json\0 => example.en\0json\0
  docs->name[ext_pos] = 0;

  // example.en\0json\0
  //        |
  //        `--- find this
  for (docs->lang = docs->name; *docs->lang != 0 && *docs->lang != '.';
      docs->lang++)
    continue;

  if (*docs->lang != '.')
    return false;

  // example.en\0json\0 => example.en\0json\0
  *docs->lang++ = 0;

  // example\0en\0json
  // |        |   |
  // |        |   `--- ext_pos
  // |        `-- lang
  // `-- name
  return strlen(docs->lang) == DOCS_LANG_CODE_LEN && *docs->name != 0;
}

void _docs_clean(docs_t *docs) {
  if (NULL == docs->file)
    return;

  file_free(docs->file);
  docs->file = NULL;
}

bool docs_init(docs_t *docs, char *dir) {
  if (NULL == docs || NULL == dir) {
    errno = EINVAL;
    return false;
  }

  memset(docs, 0, sizeof(*docs));
  return NULL != (docs->dir = opendir(dir));
}

char *docs_next(docs_t *docs, char *name, bool content) {
  if (NULL == docs) {
    errno = EINVAL;
    return false;
  }

  struct dirent *ent = NULL;
  _docs_clean(docs);

  while (NULL != (ent = readdir(docs->dir))) {
    if (*ent->d_name == '.')
      continue;

    strncpy(docs->name, ent->d_name, NAME_MAX);

    if (!_docs_parse_name(docs, content ? ".md" : ".json"))
      continue;

    if (NULL == name || strncmp(docs->name, name, NAME_MAX) == 0)
      break;
  }

  if (NULL == ent) {
    errno = ENOENT;
    return NULL;
  }

  if (NULL == (docs->file = file_load(dirfd(docs->dir), ent->d_name)))
    return NULL;

  return docs->file->content;
}

void docs_reset(docs_t *docs) {
  if (NULL != docs)
    rewinddir(docs->dir);
}

void docs_free(docs_t *docs) {
  if (NULL == docs)
    return;

  _docs_clean(docs);
  closedir(docs->dir);

  memset(docs, 0, sizeof(*docs));
}
