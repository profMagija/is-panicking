#include "is_panicking.h"

#include "stdio.h"

g_t *get_g()
{
    void *g;

#if defined(__aarch64__)
    __asm__("mov %0, x28\n"
            : "=r"(g));
#elif defined(__x86_64__)
    __asm__("mov %%r14, %0\n"
            : "=r"(g));
#else
#error "Unsupported architecture"
#endif

    return g;
}

int is_panicking()
{
    p_t *p = get_g()->panic;
    return p != 0 && !p->goexit && !p->recovered;
}