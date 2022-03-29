package main

/*
#include <string.h>

#ifdef __APPLE__
  #include <sys/ptrace.h>
#elif __linux__
  #include <sys/prctl.h>
#endif
int disable_ptrace() {
#ifdef __APPLE__
  return ptrace(PT_DENY_ATTACH, 0, 0, 0);
#elif __linux__
  return prctl(PR_SET_DUMPABLE, 0);
#endif
  return 1;
}

char *string_from_argv(int i, char **argv) {
  return strdup(argv[i]);
}
*/
import "C"

func disablePtrace() bool {
	return C.disable_ptrace() == C.int(0)
}


