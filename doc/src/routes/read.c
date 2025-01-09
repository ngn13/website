#include <linux/limits.h>
#include <ctorm/all.h>
#include <sys/stat.h>
#include <sys/mman.h>

#include <dirent.h>
#include <string.h>
#include <unistd.h>
#include <stdlib.h>

#include <stdio.h>
#include <errno.h>
#include <cmark.h>
#include <fcntl.h>

#include "routes.h"

void GET_read(req_t *req, res_t *res) {
  const char *lang = REQ_QUERY("lang");
  const char *name = REQ_QUERY("name");

  char           fp[PATH_MAX + 1], md[NAME_MAX + 1];
  struct dirent *dirent   = NULL;
  int32_t        name_len = 0, ent_len = 0;
  DIR           *dir = NULL;

  bzero(fp, sizeof(fp));
  bzero(md, sizeof(md));

  if (NULL == name)
    return RES_REDIRECT("/");

  name_len = strlen(name);

  if (NULL == lang)
    lang = "en";

  if (NULL == (dir = opendir("docs"))) {
    error("failed to open the docs dir: %s", strerror(errno));
    RES_SENDFILE("html/internal.html");
    res->code = 500;
    return;
  }

  while ((dirent = readdir(dir)) != NULL) {
    if (strncmp(dirent->d_name, lang, PATH_MAX) == 0)
      break;
  }

  closedir(dir);

  if (NULL == dirent) {
    RES_SENDFILE("html/notfound.html");
    return;
  }

  snprintf(fp, sizeof(fp), "docs/%s", lang);

  if (NULL == (dir = opendir(fp))) {
    error("failed to open the language dir: %s", strerror(errno));
    RES_SENDFILE("html/internal.html");
    res->code = 500;
    return;
  }

  while ((dirent = readdir(dir)) != NULL) {
    if ((ent_len = strlen(dirent->d_name) - 3) != name_len)
      continue;

    if (strncmp(dirent->d_name, name, name_len) == 0) {
      memcpy(md, dirent->d_name, ent_len + 3);
      break;
    }
  }

  closedir(dir);

  if (NULL == dirent) {
    RES_SENDFILE("html/notfound.html");
    return;
  }

  char       *md_content = NULL, md_fp[PATH_MAX + 1];
  struct stat md_st;
  int         md_fd = 0;

  snprintf(md_fp, sizeof(fp), "%s/%s", fp, md);

  if ((md_fd = open(md_fp, O_RDONLY)) < 0) {
    error("failed to open %s: %s", fp, strerror(errno));
    goto err_internal_close;
  }

  if (fstat(md_fd, &md_st) < 0) {
    error("failed to fstat %s: %s", fp, strerror(errno));
    goto err_internal_close;
  }

  if ((md_content = mmap(0, md_st.st_size, PROT_READ, MAP_PRIVATE, md_fd, 0)) == NULL) {
    error("failed to mmap %s: %s", fp, strerror(errno));
    goto err_internal_close;
  }

  char *parsed = cmark_markdown_to_html(md_content, md_st.st_size, CMARK_OPT_DEFAULT);
  RES_SEND(parsed);

  free(parsed);
  munmap(md_content, md_st.st_size);
  close(md_fd);

  return;

err_internal_close:
  close(md_fd);

  RES_SENDFILE("html/internal.html");
  return;
}
