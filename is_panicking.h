typedef struct _panic p_t;

typedef struct
{
    void *stack_hi, *stack_lo;
    void *stackguard0, *stackguard1;
    p_t *panic;
    void *defer;
    void *m;
    // ...
} g_t;

struct _panic
{
    void *argp;
    void *arg0, *arg1;
    struct _panic *link;
    void *pc;
    void *sp;
    char recovered, aborted, goexit;
};

g_t *get_g();