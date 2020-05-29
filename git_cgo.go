package git

/*
#cgo windows CFLAGS: -I${SRCDIR}/windows/install/include/
#cgo windows LDFLAGS: -L${SRCDIR}/windows/install/lib -lgit2
#cgo darwin CFLAGS: -I${SRCDIR}/darwin/install/include/
#cgo darwin LDFLAGS: -L${SRCDIR}/darwin/install/lib/ -lgit2 -framework CoreFoundation -framework Security -lz -L/usr/lib -liconv
#cgo linux CFLAGS: -I${SRCDIR}/linux/install/include/
#cgo linux LDFLAGS: -L${SRCDIR}/linux/install/lib/ -lgit2 -lrt -lpthread -lz
#include <git2.h>

#if LIBGIT2_VER_MAJOR != 1 || LIBGIT2_VER_MINOR != 0
# error "Invalid libgit2 version; this git2go supports libgit2 v1.0"
#endif

*/
import "C"
