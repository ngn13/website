#include <sys/mman.h>

#include <stdlib.h>
#include <unistd.h>

#include <stdio.h>
#include <fcntl.h>
#include <errno.h>

#include "file.h"

file_t *file_load(int dirfd, char *path) {
  if (NULL == path) {
    errno = EINVAL;
    return NULL;
  }

  file_t *file = NULL;
  int     fd   = -1;

  // open the taarget file
  if ((fd = openat(dirfd, path, O_RDONLY)) < 0)
    goto end; // errno set by open

  // allocate a new file structure
  if (NULL == (file = calloc(1, sizeof(file_t))))
    goto end; // errno set by malloc

  // calculate the file size
  if ((file->size = lseek(fd, 0, SEEK_END)) < 0)
    goto end;

  // memory map the file
  if (NULL ==
      (file->content = mmap(0, file->size, PROT_READ, MAP_PRIVATE, fd, 0)))
    goto end; // errno set by mmap

end:
  if (fd != -1)
    close(fd);

  if (NULL != file && NULL == file->content) {
    free(file);
    return NULL;
  }

  return file;
}

void file_free(file_t *file) {
  if (NULL == file)
    return;

  munmap(file->content, file->size);
  free(file);
}
